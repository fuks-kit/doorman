package workspace

import (
	"strings"
)

type OfficePermission struct {
	HasAccess    bool // Has access to the fuks office
	IsFuksMember bool // Is a member of the fuks group
	IsActiveFuks bool // Is active fuks member
}

// isGroupMember checks if the given email is a member of the given group.
func isGroupMember(group, email string) (isMember bool, _ error) {
	member, err := adminService.Members.HasMember(group, email).Do()
	if err != nil {
		return false, err
	}

	return member.IsMember, nil
}

// checkFuksPermission checks if the given email has permission to access the office.
func checkFuksPermission(email string) (access *OfficePermission, _ error) {
	user, err := adminService.Users.Get(email).Do()
	if err != nil {
		return nil, err
	}

	// TODO: Uncomment this
	// if user.IsAdmin {
	// 	 return true, nil
	// }

	if user.OrgUnitPath == "/aktive" {
		return &OfficePermission{
			HasAccess:    true,
			IsFuksMember: true,
			IsActiveFuks: true,
		}, nil
	}

	member, err := isGroupMember("aktive@fuks.org", email)
	if err != nil {
		return nil, err
	}

	return &OfficePermission{
		HasAccess:    true,
		IsFuksMember: true,
		IsActiveFuks: member,
	}, nil
}

func checkVisitorPermission(uid, email string) (access *OfficePermission, _ error) {
	return &OfficePermission{}, nil
}

func HasOfficeAccess(uid, email string) (access *OfficePermission, _ error) {
	if strings.HasSuffix(email, "@fuks.org") {
		return checkFuksPermission(email)
	} else {
		return checkVisitorPermission(uid, email)
	}
}
