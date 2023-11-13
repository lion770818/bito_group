package servers

import (
	"bito_group/internal/user"
)

type Apps struct {
	UserApp user.UserAppInterface
}

func NewApps(repos *Repos) *Apps {
	return &Apps{
		UserApp: user.NewUserApp(repos.UserRepo),
	}
}
