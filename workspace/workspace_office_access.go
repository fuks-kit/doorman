package workspace

import (
	"strings"
)

// isGroupMember checks if the given email is a member of the given group.
func isGroupMember(group, email string) (isMember bool, _ error) {
	member, err := adminService.Members.HasMember(group, email).Do()
	if err != nil {
		return false, err
	}

	return member.IsMember, nil
}

// checkFuksPermission checks if the given email has permission to access the office.
func checkFuksPermission(email string) (access bool, _ error) {
	user, err := adminService.Users.Get(email).Do()
	if err != nil {
		return false, err
	}

	// TODO: Uncomment this
	// if user.IsAdmin {
	// 	 return true, nil
	// }

	if user.OrgUnitPath == "/aktive" {
		return true, nil
	}

	// Check group membership

	member, err := isGroupMember("aktive@fuks.org", email)
	if err != nil {
		return false, err
	}

	return member, nil
}

func checkVisitorPermission(uid, email string) (access bool, _ error) {
	return false, nil
}

func HasOfficeAccess(uid, email string) (access bool, _ error) {
	if strings.HasSuffix(email, "@fuks.org") {
		return checkFuksPermission(email)
	} else {
		return checkVisitorPermission(uid, email)
	}
}
