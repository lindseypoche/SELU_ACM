package discord

import (
	"log"

	"github.com/cmd-ctrl-q/SELU_ACM/api/internal/listing"
)

type me struct {
	authorID  string
	stateID   string
	channelID string
	isBot     bool
	roles     *[]listing.Role

	// access level. 0-3
	accessLvl uint8

	// isAuthorized is mainly for validating the channel id
	isAuthorized bool
}

// Validate authors.
// returns (access level, is authorized)
// access level is 0-3 where 3 has full access and 0 has zero access.
// isAuthorized makes sure that the user is not a bot and that
// the channel is a valid channel to listen for data.
func Validate(v *me) (uint8, bool) {

	// validate user is not bot
	if v.isBot {
		v.isAuthorized = false
		return level0, v.isAuthorized
	}

	// validate channels
	// if channelID is empty, then channelID is not necessary and
	// channel validation should be skipped.
	if v.channelID != "" {
		for _, channel := range Config.Channels {
			if channel == v.channelID {
				v.isAuthorized = true
			}
		}
	}

	// check user's access level from roles
	al := accessLevel(v.roles)
	log.Println(al, " : ", v.isAuthorized)

	return al, v.isAuthorized
}

// Check user access level.
// Access level0 = no access
// Access level1 = acm member access and acm alumi.
// Access level2 = acm officer access.
// Access level3 = acm admin.
const (
	level0 = iota
	level1
	level2
	level3

	member  = "acm_member"  // level1
	alumni  = "acm_alumni"  // level1
	officer = "acm_officer" // level2
	admin   = "acm_admin"   // level3
)

// accessLevel returns the access level of a user for validation
// and before a user can access certain funtions.
func accessLevel(roles *[]listing.Role) uint8 {

	var accessLevel uint8

	if roles == nil || len(*roles) < 1 {
		return level0
	}

	var isMember, isOfficer, isAdmin bool

	// check for acm roles
	for _, role := range *roles {

		switch role.Name {
		case member:
			isMember = true
			break
		case officer:
			isOfficer = true
			break
		case admin:
			isAdmin = true
			break
		}
	}

	if isAdmin {
		accessLevel = level3
	} else if isOfficer {
		accessLevel = level2
	} else if isMember {
		accessLevel = level1
	} else {
		accessLevel = level0
	}

	return accessLevel
}
