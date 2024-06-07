package services

import (
	"authentication/models"
	"authentication/repository"
	"authentication/token"
	"log"
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type Login struct {
	logger          *log.Logger
	flags           *models.Flags
	loginRepository *repository.Login
}

func NewLogin(l *log.Logger, f *models.Flags) *Login {
	return &Login{
		logger:          l,
		flags:           f,
		loginRepository: repository.Init(),
	}
}

func (l *Login) GetToken(model models.LoginRequest, origin string) (string, *models.ErrorDetail) {
	user, err := l.loginRepository.GetUserByName(model.UserName, model.Password)

	if err != nil {
		return "", err
	}

	var claims = &models.JwtClaims{
		ComapnyId: strconv.Itoa(user.Id),
		Username:  user.Name,
		Roles:     user.Roles,
		StandardClaims: jwt.StandardClaims{
			Audience: origin,
		},
	}

	var tokenCreationTime = time.Now().UTC()

	var expirationTime = tokenCreationTime.Add(time.Duration(2) * time.Hour)

	return token.GenerateToken(claims, expirationTime)

}

func (*Login) VerifyToken(tokenString, referer string) (bool, *models.JwtClaims) {
	return token.VerifyToken(tokenString, referer)
}
