package autorizaton

import "strings"

type Auth struct {
	logins map[string]string
}

func New() (*Auth, error) {
	logins := map[string]string{
		"validuser": "password",
	}
	return &Auth{logins: logins}, nil
}

func (a Auth) Check(username string, password string) (bool, error) {

	lowCaseUsername := strings.ToLower(username)

	storedPassword, exist := a.logins[lowCaseUsername]

	if storedPassword != password || !exist {
		return false, nil
	}

	return true, nil
}
