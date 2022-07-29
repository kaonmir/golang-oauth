package models

import (
	"crypto/sha256"
	"errors"
	"log"

	"github.com/kaonmir/OAuth/db"
	"github.com/kaonmir/OAuth/forms"
	uuid "github.com/satori/go.uuid"
	"github.com/vmihailenco/msgpack"
)

var hash256 = sha256.New()

type User struct {
	Id       string `json:"id"`
	UserId   string `json:"user_id"`
	Name     string `json:"name"`
	Gender   string `json:"gender"`
	Password string `json:"password"`
}

func (h User) Signup(userPayload forms.UserSignup) (*User, error) {
	db := db.GetDB()
	id := uuid.NewV4()

	user := User{
		Id:       id.String(),
		UserId:   userPayload.Id,
		Name:     userPayload.Name,
		Gender:   userPayload.Gender,
		Password: hashPassword(userPayload.Password),
	}

	b, err := msgpack.Marshal(&user)
	if err != nil {
		return nil, err
	}

	err = db.Set(user.UserId, string(b), 0).Err()
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (h User) GetByID(id string, password string) (*User, error) {
	log.Printf("[DEBUG] id: %s, password: %s", id, password)

	db := db.GetDB()

	result, err := db.Get(id).Result()
	if err != nil {
		return nil, err
	}

	var user *User
	err = msgpack.Unmarshal([]byte(result), &user)
	if err != nil {
		return nil, err
	}

	if user.Password != hashPassword(password) {
		return nil, errors.New("password is wrong")
	}

	return user, nil
}

func hashPassword(password string) string {
	hash256.Reset()
	hash256.Write([]byte(password))
	return string(hash256.Sum(nil))
}
