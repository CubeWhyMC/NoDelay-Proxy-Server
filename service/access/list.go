package access

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/layou233/ZBProxy/common/set"
	"github.com/layou233/ZBProxy/config"
)

func GetTargetList(listName string) (set.StringSet, error) {
	list, ok := config.Config.Lists[listName]
	if ok {
		return list, nil
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