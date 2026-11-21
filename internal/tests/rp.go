package tests

import (
	"github.com/descope/virtualwebauthn"
	"github.com/HolyShitDutch/kashanka/config"
)

func RelyingParty() virtualwebauthn.RelyingParty {
	return virtualwebauthn.RelyingParty{
		ID:     config.RPID,
		Name:   config.RPDisplayName,
		Origin: config.RPOrigins[0],
	}
}
