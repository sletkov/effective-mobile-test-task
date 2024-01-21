package service

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/sletkov/effective-mobile-test-task/internal/service/model"
)

// Add age to user by data from 3rd-party api response
func Agify(ageResponse *http.Response, u *model.User) error {
	var ageInfo = struct {
		Count int    `json:"count"`
		Name  string `json:"name"`
		Age   int    `json:"age"`
	}{}

	data, err := io.ReadAll(ageResponse.Body)

	if err != nil {
		return err
	}

	if err := json.Unmarshal(data, &ageInfo); err != nil {
		return err
	}

	u.Age = ageInfo.Age

	return nil
}

// Add gender to user by data from 3rd-party api response
func Genderize(genderResponse *http.Response, u *model.User) error {
	var genderInfo = struct {
		Count       int     `json:"count"`
		Name        string  `json:"name"`
		Gender      string  `json:"gender"`
		Probability float32 `json:"probability"`
	}{}

	data, err := io.ReadAll(genderResponse.Body)

	if err != nil {
		return err
	}

	if err := json.Unmarshal(data, &genderInfo); err != nil {
		return err
	}

	u.Gender = genderInfo.Gender

	return nil
}

// Add nationality to user by data from 3rd-party api response
func Nationalize(genderResponse *http.Response, u *model.User) error {
	var nationalityInfo = struct {
		Count   int    `json:"count"`
		Name    string `json:"name"`
		Country []struct {
			CountryId   string  `json:"country_id"`
			Probability float32 `json:"probability"`
		}
	}{}

	data, err := io.ReadAll(genderResponse.Body)

	if err != nil {
		return err
	}

	if err := json.Unmarshal(data, &nationalityInfo); err != nil {
		return err
	}

	// The first element always has the most probability
	u.Nationality = nationalityInfo.Country[0].CountryId

	return nil
}
