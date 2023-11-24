package workspace

import "time"

const cacheDuration = time.Minute * 15

var authUsers map[string]*SheetAppUser
var lastFetchTime time.Time

func GetAuthUserFromSheetCache(userId string) (*SheetAppUser, bool, error) {
	if time.Since(lastFetchTime) < cacheDuration {
		user, found := authUsers[userId]
		return user, found, nil
	}

	users, err := GetAuthUserFromSheet()
	if err != nil {
		return nil, false, err
	}

	authUsers = make(map[string]*SheetAppUser)
	for _, user := range users {
		authUsers[user.UserId] = user
	}
	lastFetchTime = time.Now()

	user, found := authUsers[userId]
	return user, found, nil
}
