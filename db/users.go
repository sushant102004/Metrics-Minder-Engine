package db

import (
	"errors"
	"metric-minder-engine/models"

	"github.com/rs/zerolog/log"
	"gorm.io/gorm"
)

type UsersDB struct{}

func NewUsersDB() UsersDB {
	return UsersDB{}
}

func (db *UsersDB) GetAllEmails() ([]string, error) {
	users := []models.GoogleUser{}

	tx := Conn.Find(&users)
	if tx.Error != nil {
		return nil, tx.Error
	}

	emails := make([]string, len(users))
	for _, u := range users {
		emails = append(emails, u.Email)
	}

	return emails, nil
}

func (db *UsersDB) GetAccessToken(email string) (string, error) {
	user := models.GoogleUser{}

	tx := Conn.First(&user, "email = "+email)
	if tx.Error != nil {
		return "", tx.Error
	}

	return user.AccessToken, nil
}

func (db *UsersDB) GetRefreshToken(email string) (string, error) {
	user := models.GoogleUser{}
	tx := Conn.Where("email = ?", email).First(&user)
	if tx.Error != nil {
		log.Error().Msg("error while fetching onboard status from db: " + tx.Error.Error())
		return "", tx.Error
	}

	return user.RefreshToken, nil
}

func (db UsersDB) GetPropertyID(email string) (string, error) {
	var property models.PropertyToMonitor

	tx := Conn.Where("email = ?", email).First(&property)
	if tx.Error != nil {
		if tx.Error == gorm.ErrRecordNotFound {
			return "", errors.New("property not found")
		}
		return "", tx.Error
	}

	return property.PropertyID, nil
}
