package context

import (
	"github.com/labstack/echo/v4"
	"github.com/onsonr/sonr/internal/gateway/views"
	"golang.org/x/exp/rand"
)

func GetCreateProfileData(c echo.Context) views.CreateProfileData {
	_, err := Get(c)
	if err != nil {
		return views.CreateProfileData{}
	}
	return views.CreateProfileData{
		FirstNumber: rand.Intn(5) + 1,
		LastNumber:  rand.Intn(4) + 1,
	}
}

func GetPasskeyCreateData(c echo.Context) views.CreatePasskeyData {
	sess, err := Get(c)
	if err != nil {
		return views.CreatePasskeyData{}
	}
	profile, err := GetProfile(c)
	if err != nil {
		return views.CreatePasskeyData{}
	}
	return views.CreatePasskeyData{
		Address:       profile.Address,
		Handle:        profile.Handle,
		Name:          profile.Name,
		Challenge:     sess.Session().Challenge,
		CreationBlock: "00001",
	}
}
