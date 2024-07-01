package sqlobjects

import (
	"errors"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name     string
	Email    string
	Password string
	Token    UserToken `gorm:"foreignKey:UserID;references:ID"`
}

type UserToken struct {
	gorm.Model
	UserID     uint
	Token      string
	Expiration time.Time `gorm:"type:datetime"`
}

func GetUserByEmail(email string, database *gorm.DB) (User, error) {

	user := User{Email: email}
	result := database.Where(user).First(&user)

	if result.Error != nil {
		return User{}, result.Error

	} else {
		return user, nil

	}

}

func CreateSessionToken(user User, database *gorm.DB) (string, error) {

	token, err := GenerateToken()

	if err != nil {
		return "", err
	}

	expiration := time.Now().Add(5 * 24 * time.Hour)
	userToken := UserToken{
		UserID:     user.ID,
		Token:      token,
		Expiration: expiration,
	}
	fmt.Println("Before creating token")

	err = database.Create(&userToken).Error
	if err != nil {
		return "", err
	}
	fmt.Println("After creating token")

	return token, nil

}

func GetUserByToken(token string, database *gorm.DB) (User, error) {
	var user User

	err := database.Joins("JOIN user_tokens ON user_tokens.user_id = users.id").
		Where("user_tokens.token = ?", token).
		Preload("Token").
		First(&user).Error

	if err != nil {
		return User{}, errors.New("could not find user with token")
	}

	if user.Token.Expiration.Before(time.Now()) {

		message := fmt.Sprint(user.Token.Expiration.UTC())
		return User{}, errors.New(message)
	}

	return user, nil
}

func GetUserByID(id int, database *gorm.DB) (User, error) {

	user := User{}
	result := database.First(&user, id)

	if result.Error != nil {
		return User{}, result.Error

	} else {
		return user, nil

	}
}

func IsPasswordValid(password string, user User) bool {

	isValid, err := comparePasswordAndHash(password, user.Password)
	if err != nil {
		return false
	}

	return isValid
}

func CreateUser(email string, password string, database *gorm.DB) (User, error) {

	hashedPassword, err := generateFromPassword(password, &HashingParams)

	if err != nil {
		return User{}, nil
	}

	newUser := User{Email: email, Password: hashedPassword, Name: email}
	err = database.Create(&newUser).Error
	if err != nil {
		return User{}, err
	}

	return newUser, nil

}
