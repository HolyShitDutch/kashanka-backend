package tests

import (
	"github.com/HolyShitDutch/kashanka/config"
	"github.com/descope/virtualwebauthn"
)

func RelyingParty() virtualwebauthn.RelyingParty {
	return virtualwebauthn.RelyingParty{
		ID:     config.RPID,
		Name:   config.RPDisplayName,
		Origin: config.RPOrigins[0],
	}
}
