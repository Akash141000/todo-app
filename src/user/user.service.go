package user

import (
	"errors"
	"fmt"
	"time"
	"todoBackend/db"
	helperservice "todoBackend/src/helperService"
	"todoBackend/utils"

	"github.com/golang-jwt/jwt/v4"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
)

func CreateUser(user User) *mongo.InsertOneResult {
	hashedPassword := HashPassword(user)
	createdUser, _ := helperservice.InsertOne(db.UserModel, map[string]string{"email": user.Email, "password": hashedPassword})
	return createdUser
}

func LoginUser(user User) (interface{}, error) {

	//query filter
	filter := make(map[string]interface{})
	filter["email"] = user.Email

	userFound, err := helperservice.FindOne(db.UserModel, &filter)
	if err != nil {
		return nil, err
	}
	loginUser := &User{
		ID:       userFound["_id"].(primitive.ObjectID),
		Email:    userFound["email"].(string),
		Password: userFound["password"].(string),
	}
	hashedPassword := loginUser.Password
	isEqual := CompareHashedPassword(hashedPassword, user.Password)

	signingKey := []byte("jwtSigningKey")

	claims := &utils.Claims{
		Email: loginUser.Email,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: &jwt.NumericDate{
				Time: time.Now().Local().Add(time.Hour * time.Duration(24)),
			},
		},
	}
	refreshTokenClaims := claims

	if isEqual {
		accessToken, accessTokenError := jwt.NewWithClaims(jwt.SigningMethodHS256, *claims).SignedString(signingKey)

		// refresh token
		refreshTokenClaims.ExpiresAt.Time = time.Now().Local().Add(time.Hour * time.Duration(24) * time.Duration(30))
		refreshToken, refreshTokenError := jwt.NewWithClaims(jwt.SigningMethodHS256, *refreshTokenClaims).SignedString(signingKey)

		if accessTokenError != nil {
			fmt.Println("ACCESS TOKEN SIGNING ERROR>>", accessToken)
			return nil, errors.New("error logging in user")
		}
		if refreshTokenError != nil {
			fmt.Println("REFRESH TOKEN SIGNING ERROR>>", accessToken)
			return nil, errors.New("error logging in user")
		}

		return map[string]string{"accessToken": accessToken, "refreshToken": refreshToken}, nil
	}

	return nil, errors.New("invalid credentials")
}

func CompareHashedPassword(hashedPassword string, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	if err != nil {
		fmt.Println("Compare Password>>", err)
		return false
	}
	return true
}

func HashPassword(user User) string {
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(user.Password), 10)
	return string(hashedPassword)
}
