package context

import (
	"github.com/labstack/echo/v4"
	"github.com/onsonr/sonr/internal/gateway/models"
)

func GetCreateProfileData(c echo.Context) models.CreateProfileData {
	sess, err := Get(c)
	if err != nil {
		return models.CreateProfileData{}
	}
	return models.CreateProfileData{
		FirstNumber: sess.Session().IsHumanFirst,
		LastNumber:  sess.Session().IsHumanLast,
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
