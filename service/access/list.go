package access

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/CubeWhyMC/NoDelay-Proxy-Server/common/set"
	"github.com/CubeWhyMC/NoDelay-Proxy-Server/config"
)

func GetTargetList(listName string) (set.StringSet, error) {
	set, ok := config.Config.Lists[listName]
	if ok {
		return set, nil
	}
	return nil, fmt.Errorf("list %q not found", listName)
}

func IsWhitelist(playerName string) (bool, error) {
	resp, err := http.Get(config.Config.PrivateConfig.ListAPI + "?playerName=" + playerName)
	if err != nil {
		return false, fmt.Errorf("failed to make HTTP request: %w", err)
	}
	defer resp.Body.Close()

	bytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return false, fmt.Errorf("failed to read HTTP response body: %w", err)
	}

	return playerName == string(bytes), nil
}

func IsFirstTime(playerName string) bool {
	var accessedPlayers = make(map[string]bool)
	_, exists := accessedPlayers[playerName]
	if !exists {
		accessedPlayers[playerName] = true
		return true
	}
	return false
}