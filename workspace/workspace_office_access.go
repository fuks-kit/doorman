package workspace

import (
	"strings"
)

type OfficePermission struct {
	UID          string // The uid of the user
	EMail        string // The email of the user
	Name         string // The name of the user
	Organization string // The organization of the user
	HasAccess    bool   // Has access to the fuks office
	IsFuksMember bool   // Is a member of the fuks group
	IsActiveFuks bool   // Is active fuks member
}

// isGroupMember checks if the given email is a member of the given group.
func isGroupMember(group, email string) (isMember bool, _ error) {
	member, err := adminService.Members.HasMember(group, email).Do()
	if err != nil {
		return false, err
	}

	return member.IsMember, nil
}

func HasOfficeAccess(uid, email string) (access *OfficePermission, _ error) {
	access = &OfficePermission{
		UID:   uid,
		EMail: email,
	}

	//
	// Check if the user is a member of the fuks group
	//

	if strings.HasSuffix(email, "@fuks.org") {
		isActive, err := isGroupMember("aktive@fuks.org", email)
		if err != nil {
			return nil, err
		}

		access.IsFuksMember = true
		access.IsActiveFuks = isActive
		access.Organization = "fuks"

		if isActive {
			access.HasAccess = true
			return access, nil
		}
	}

	//
	// If the user-id is in the sheet, grant access
	//

	user, found, err := GetAuthUserFromSheetCache(uid)
	if err != nil {
		return nil, err
	}

	if found {
		access.HasAccess = true
		access.Name = user.Name
		access.Organization = user.Organization
	}

	return access, nil
}
