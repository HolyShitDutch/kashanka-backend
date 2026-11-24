package types

import "github.com/HolyShitDutch/kashanka/db/queries"

type UserTokenDTO struct {
	*queries.UserToken
	PlaintextToken []byte
}
