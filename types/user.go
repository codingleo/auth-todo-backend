package types

import (
	"log"
	"time"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	ID        uint64         `json:"id" gorm:"primaryKey;autoIncrement:true"`
	FirstName string         `json:"firstName" gorm:"not null"`
	LastName  string         `json:"lastName" gorm:"not null"`
	Email     string         `json:"email" gorm:"unique;not null"`
	Password  string         `json:"password,omitempty" gorm:"not null" sql:"-"`
	CreatedAt time.Time      `json:"-"`
	UpdatedAt time.Time      `json:"-"`
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index"`
}

func (u *User) ComparePassword(password string) bool {
	return bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password)) == nil
}

func (u *User) BeforeCreate(tx *gorm.DB) error {
	u.Password = hashAndSalt(string(u.Password))
	return nil
}

func (u *User) Validate() map[string]string {
	errs := make(map[string]string)

	if u.FirstName == "" {
		errs["firstName"] = "First name is required"
	}

	if u.LastName == "" {
		errs["lastName"] = "Last name is required"
	}

	if u.Email == "" {
		errs["email"] = "Email is required"
	}

	if u.Password == "" {
		errs["password"] = "Password is required"
	}

	return errs
}

func hashAndSalt(password string) string {
	pwd := []byte(password)
	hash, err := bcrypt.GenerateFromPassword(pwd, 14)
	if err != nil {
		log.Println(err)
	}
	return string(hash)
}
