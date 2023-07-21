package workspace

import "time"

const cacheDuration = time.Minute * 15

var authUsers map[string]SheetAppUser
var lastFetchTime time.Time

func GetAuthUserFromSheetCache(userId string) (access bool, _ error) {
	if !lastFetchTime.IsZero() && time.Since(lastFetchTime) < cacheDuration {
		_, access = authUsers[userId]
		return
	}

	users, err := GetAuthUserFromSheet()
	if err != nil {
		return false, err
	}

	authUsers = make(map[string]SheetAppUser)
	for _, user := range users {
		authUsers[user.UserId] = user
	}
	lastFetchTime = time.Now()

	_, access = authUsers[userId]
	return
}
