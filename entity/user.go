package entity

import (
	"project3/infra/config"
	"project3/pkg/errs"
	"strings"
	"time"

	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
)

var invalidTokenErr = errs.NewUnauthenticatedError("invalid token")

type User struct {
	Id         int
	Full_name  string
	Email      string
	Password   string
	Role       string
	Created_at time.Time
	Updated_at time.Time
}

func (u *User) parseToken(tokenString string) (*jwt.Token, errs.MessageErr) {

	token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, invalidTokenErr
		}
		secretKey := config.GetAppConfig().JWTSecretKey
		return []byte(secretKey), nil
	})

	if err != nil {
		return nil, invalidTokenErr
	}

	return token, nil
}

func (u *User) bindTokenToUserEntity(claim jwt.MapClaims) errs.MessageErr {
	if id, ok := claim["id"].(float64); !ok {
		return invalidTokenErr
	} else {
		u.Id = int(id)
	}

	if email, ok := claim["email"].(string); !ok {
		return invalidTokenErr
	} else {
		u.Email = email
	}
	if role, ok := claim["role"].(string); ok {
		u.Role = role
	} else {
		return invalidTokenErr
	}
	return nil
}

func (u *User) ValidateToken(bearerToken string) errs.MessageErr {
	isBearer := strings.HasPrefix(bearerToken, "Bearer")

	if !isBearer {
		return invalidTokenErr
	}

	splitToken := strings.Split(bearerToken, " ")

	if len(splitToken) != 2 {
		return invalidTokenErr
	}

	tokenString := splitToken[1]

	token, err := u.parseToken(tokenString)

	if err != nil {
		return err
	}

	var mapClaims jwt.MapClaims

	if claims, ok := token.Claims.(jwt.MapClaims); !ok || !token.Valid {
		return invalidTokenErr
	} else {
		mapClaims = claims
	}

	err = u.bindTokenToUserEntity(mapClaims)

	return err
}

func (u *User) tokenClaim() jwt.MapClaims {
	return jwt.MapClaims{
		"id":    u.Id,
		"email": u.Email,
		"role":  u.Role,
	}
}

func (u *User) signToken(claims jwt.MapClaims) string {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	secretKey := config.GetAppConfig().JWTSecretKey

	tokenString, _ := token.SignedString([]byte(secretKey))

	return tokenString
}

func (u *User) GenerateToken() string {
	claims := u.tokenClaim()

	return u.signToken(claims)
}

func (u *User) HashPassword() errs.MessageErr {
	salt := 8

	bs, err := bcrypt.GenerateFromPassword([]byte(u.Password), salt)

	if err != nil {
		return errs.NewInternalServerError("something went wrong")
	}

	u.Password = string(bs)

	return nil
}

func (u *User) ComparePassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
	return err == nil
}
