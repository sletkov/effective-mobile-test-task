package model

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUserValidate(t *testing.T) {
	testCases := []struct {
		name    string
		u       User
		isValid bool
	}{
		{
			name: "valid",
			u: User{
				Name:        "Ivan",
				Surname:     "Ivanov",
				Patronymic:  "Ivanovich",
				Age:         20,
				Gender:      "male",
				Nationality: "RU",
			},
			isValid: true,
		},

		{
			name: "empty patronymic",
			u: User{
				Name:        "Ivan",
				Surname:     "Ivanov",
				Age:         20,
				Gender:      "male",
				Nationality: "RU",
			},
			isValid: true,
		},

		{
			name: "empty name",
			u: User{
				Surname:     "Ivanov",
				Patronymic:  "Ivanovich",
				Age:         20,
				Gender:      "male",
				Nationality: "RU",
			},
			isValid: false,
		},

		{
			name: "empty surname",
			u: User{
				Name:        "Ivan",
				Patronymic:  "Ivanovich",
				Age:         20,
				Gender:      "male",
				Nationality: "RU",
			},
			isValid: false,
		},

		{
			name: "empty age",
			u: User{
				Name:        "Ivan",
				Surname:     "Ivanov",
				Patronymic:  "Ivanovich",
				Gender:      "male",
				Nationality: "RU",
			},
			isValid: false,
		},

		{
			name: "invalid age",
			u: User{
				Name:        "Ivan",
				Surname:     "Ivanov",
				Age:         -1,
				Patronymic:  "Ivanovich",
				Gender:      "male",
				Nationality: "RU",
			},
			isValid: false,
		},

		{
			name: "empty gender",
			u: User{
				Name:        "Ivan",
				Surname:     "Ivanov",
				Patronymic:  "Ivanovich",
				Age:         20,
				Nationality: "RU",
			},
			isValid: false,
		},

		{
			name: "empty nationality",
			u: User{
				Name:       "Ivan",
				Surname:    "Ivanov",
				Patronymic: "Ivanovich",
				Age:        20,
				Gender:     "male",
			},
			isValid: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if tc.isValid {
				assert.NoError(t, tc.u.Validate())
			} else {
				assert.Error(t, tc.u.Validate())
			}
		})
	}
}

func TestCreateUserUserValidate(t *testing.T) {
	testCases := []struct {
		name    string
		u       CreateUser
		isValid bool
	}{
		{
			name: "valid",
			u: CreateUser{
				Name:       "Ivan",
				Surname:    "Ivanov",
				Patronymic: "Ivanovich",
			},
			isValid: true,
		},

		{
			name: "empty patronymic",
			u: CreateUser{
				Name:    "Ivan",
				Surname: "Ivanov",
			},
			isValid: true,
		},

		{
			name: "empty name",
			u: CreateUser{
				Surname:    "Ivanov",
				Patronymic: "Ivanovich",
			},
			isValid: false,
		},

		{
			name: "empty surname",
			u: CreateUser{
				Name:       "Ivan",
				Patronymic: "Ivanovich",
			},
			isValid: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if tc.isValid {
				assert.NoError(t, tc.u.Validate())
			} else {
				assert.Error(t, tc.u.Validate())
			}
		})
	}
}

func TestUpdateUserValidate(t *testing.T) {
	testCases := []struct {
		name    string
		u       UpdateUser
		isValid bool
	}{
		{
			name: "invalid name",
			u: UpdateUser{
				Name: "Иван",
			},
			isValid: false,
		},

		{
			name: "invalid surname",
			u: UpdateUser{
				Surname: "Иванов",
			},
			isValid: false,
		},

		{
			name: "invalid patronymic",
			u: UpdateUser{
				Name: "Иванович",
			},
			isValid: false,
		},

		{
			name: "invalid age",
			u: UpdateUser{
				Age: -1,
			},
			isValid: false,
		},

		{
			name: "invalid gender",
			u: UpdateUser{
				Gender: "email",
			},
			isValid: false,
		},

		{
			name: "invalid nationality",
			u: UpdateUser{
				Nationality: "Russia",
			},
			isValid: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if tc.isValid {
				assert.NoError(t, tc.u.Validate())
			} else {
				assert.Error(t, tc.u.Validate())
			}
		})
	}
}

func TestUserFilterValidate(t *testing.T) {
	testCases := []struct {
		name    string
		u       UserFilter
		isValid bool
	}{
		{
			name: "invalid name",
			u: UserFilter{
				Name: "Иван",
			},
			isValid: false,
		},

		{
			name: "invalid surname",
			u: UserFilter{
				Surname: "Иванов",
			},
			isValid: false,
		},

		{
			name: "invalid patronymic",
			u: UserFilter{
				Name: "Иванович",
			},
			isValid: false,
		},

		{
			name: "invalid ageFrom",
			u: UserFilter{
				AgeFrom: 101,
			},
			isValid: false,
		},

		{
			name: "invalid ageTo",
			u: UserFilter{
				AgeFrom: -1,
			},
			isValid: false,
		},

		{
			name: "invalid gender",
			u: UserFilter{
				Gender: "email",
			},
			isValid: false,
		},

		{
			name: "invalid nationality",
			u: UserFilter{
				Nationality: "Russia",
			},
			isValid: false,
		},

		{
			name: "invalid limit",
			u: UserFilter{
				Limit: -1,
			},
			isValid: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if tc.isValid {
				assert.NoError(t, tc.u.Validate())
			} else {
				assert.Error(t, tc.u.Validate())
			}
		})
	}
}
