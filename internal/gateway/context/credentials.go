package context

import (
	"github.com/labstack/echo/v4"
	"github.com/onsonr/sonr/internal/gateway/models"
	"golang.org/x/exp/rand"
)

func GetCreateProfileData(c echo.Context) models.CreateProfileData {
	_, err := Get(c)
	if err != nil {
		return models.CreateProfileData{}
	}
	return models.CreateProfileData{
		FirstNumber: rand.Intn(5) + 1,
		LastNumber:  rand.Intn(4) + 1,
	}
}

func GetPasskeyCreateData(c echo.Context) models.CreatePasskeyData {
	sess, err := Get(c)
	if err != nil {
		return models.CreatePasskeyData{}
	}
	profile, err := GetProfile(c)
	if err != nil {
		return models.CreatePasskeyData{}
	}
	return models.CreatePasskeyData{
		Address:       profile.Address,
		Handle:        profile.Handle,
		Name:          profile.Name,
		Challenge:     sess.Session().Challenge,
		CreationBlock: "00001",
	}
}
