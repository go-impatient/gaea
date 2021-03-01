package hooks

import (
	"moocss.com/gaea/pkg/log"
	"moocss.com/gaea/pkg/twirp"
)

// Init .
func Init(logger log.Logger) *twirp.ServerHooks {
	return twirp.ChainHooks(
		// 	NewHeaders(),
		NewRequestID(),
		NewLog(logger),
		// 	NewAuth(),
	)
}
