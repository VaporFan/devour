/*
 * mtStats Devour - Web API Package
 */

package webapi

import (
	"github.com/mtstats/devour/utils/config"
)

var (
	lastSeqNum = 0
)

type WebApi struct {
	urls     config.SteamAPI
	key      string
	language string
}

func LoadApi(config config.Configuration) (api WebApi) {

	api.key = config.SteamAPI.Key
	api.urls = config.SteamAPI

	return api

}
