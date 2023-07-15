package workspace

import (
	"log"
	"strings"
)

func checkFuksAccess(email string) (access bool, _ error) {
	user, err := adminService.Users.Get(email).Do()
	if err != nil {
		return false, err
	}

	if user.IsAdmin {
		return true, nil
	}

	if user.OrgUnitPath == "/aktive" {
		return true, nil
	}

	// Check group membership
	group, err := adminService.Members.Get("aktive@fuks.org", email).Do()
	if err != nil {
		return false, err
	}

	// TODO: Check if user is member of "aktive" group
	log.Printf("########### group=%v", group)

	return false, nil
}

func checkExternalAccess(uid, email string) (access bool, _ error) {
	return false, nil
}

func HasOfficeAccess(uid, email string) (access bool, _ error) {
	if strings.HasSuffix(email, "@fuks.org") {
		return checkFuksAccess(email)
	} else {
		return checkExternalAccess(uid, email)
	}
}
