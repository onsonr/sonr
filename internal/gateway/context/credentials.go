package context

import (
	"github.com/labstack/echo/v4"
	"github.com/onsonr/sonr/internal/gateway/models"
)

func GetCreateProfileData(c echo.Context) (models.CreateProfileData, error) {
	sess, err := Get(c)
	if err != nil {
		return models.CreateProfileData{}, err
	}
	return models.CreateProfileData{
		FirstNumber: sess.Session().IsHumanFirst,
		LastNumber:  sess.Session().IsHumanLast,
	}, nil
}

func GetPasskeyCreateData(c echo.Context) (models.CreatePasskeyData, error) {
	sess, err := Get(c)
	if err != nil {
		return models.CreatePasskeyData{}, err
	}
	profile, err := GetProfile(c)
	if err != nil {
		return models.CreatePasskeyData{}, err
	}
	return models.CreatePasskeyData{
		Address:       profile.Address,
		Handle:        profile.Handle,
		Name:          profile.Name,
		Challenge:     sess.Session().Challenge,
		CreationBlock: "00001",
	}, nil
}
