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
	member, err := isGroupMember("aktive@fuks.org", email)
	if err != nil {
		return nil, err
	}

	return &OfficePermission{
		HasAccess:    member,
		IsFuksMember: true,
		IsActiveFuks: member,
	}, nil
}

func checkVisitorPermission(uid string) (permission *OfficePermission, _ error) {
	access, err := GetAuthUserFromSheetCache(uid)
	if err != nil {
		return nil, err
	}

	return &OfficePermission{
		HasAccess:    access,
		IsFuksMember: false,
		IsActiveFuks: false,
	}, nil
}

func HasOfficeAccess(uid, email string) (access *OfficePermission, _ error) {
	if strings.HasSuffix(email, "@fuks.org") {
		return checkFuksPermission(email)
	} else {
		return checkVisitorPermission(uid)
	}
}
