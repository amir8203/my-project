package services

import (
	"my-project/src/api/dto"
	"my-project/src/config"
	"time"

	service_errors "my-project/src/services/service_errors"

	"github.com/golang-jwt/jwt"
)

type TokenService struct {
	cfg *config.Config
}

type TokenDto struct {
	UserId    int
	Username  string
	MobileNumber string
}

func NewTokenService(cfg *config.Config) *TokenService {
	return &TokenService{
		cfg: cfg,
	}
}

func (s *TokenService) GenerateToken(token *TokenDto) (*dto.TokenDetail, error) {
	// tokenDetail
	td := &dto.TokenDetail{}
	td.AccessTokenExpireTime = int(time.Now().Add(s.cfg.JWT.AccessTokenExpireDuration * time.Minute).Unix())
	td.RefreshTokenExpireTime = int(time.Now().Add(s.cfg.JWT.RefreshTokenExpireDuration * time.Minute).Unix())

	// accessTokenClaims
	atc := jwt.MapClaims{}
	atc["user_id"] = token.UserId
	atc["username"] = token.Username
	atc["mobile_number"] = token.MobileNumber
	atc["exp"] = td.AccessTokenExpireTime

	// accessToken
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atc)

	var err error
	td.AccessToken, err = at.SignedString([]byte(s.cfg.JWT.Secret))

	if err != nil {
		return nil, err
	}

	// refreshTokenClaim
	rtc := jwt.MapClaims{}

	rtc["user_id"] = token.UserId
	rtc["exp"] = td.RefreshTokenExpireTime

	rt := jwt.NewWithClaims(jwt.SigningMethodHS256, rtc)

	td.RefreshToken, err = rt.SignedString([]byte(s.cfg.JWT.RefreshSecret))

	if err != nil {
		return nil, err
	}

	return td, nil
}

func (s *TokenService) VerifyToken(token string) (*jwt.Token, error) {
	// accessToken
	at , err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, &service_errors.ServiceError{EndUserMessage: service_errors.UnExpectedError}
		}
		return []byte(s.cfg.JWT.Secret), nil
	})
	if err != nil {
		return nil, err
	}
	return at, nil
}

func (s *TokenService) GetClaims(token string) (claimMap map[string]interface{},err error) {
	claimMap = map[string]interface{}{}

	verifyToken, err := s.VerifyToken(token)
	if err != nil {
		return nil, err
	}

	claims, ok := verifyToken.Claims.(jwt.MapClaims)
	if ok && verifyToken.Valid{
		for k, v := range claims {
			claimMap[k] = v
		}
		return claimMap, nil
	}
	return nil, &service_errors.ServiceError{EndUserMessage: service_errors.ClaimsNotFound}

}