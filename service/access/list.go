package access

import (
  "encoding/json"
  "fmt"
  "io/ioutil"
  "net/http"

  "github.com/layou233/ZBProxy/common/set"
)

func GetTargetList(listName string) (set.StringSet, error) {
  	// 获取网站API地址
  	resp, err := http.Get(config.Config.PrivateConfig.ListAPI)
  	if err != nil {
    	return nil, fmt.Errorf("error fetching list: %v", err)
  	}
  	defer resp.Body.Close()

  	// 读取网页响应
  	body, err := ioutil.ReadAll(resp.Body)
  	if err != nil {
    	return nil, fmt.Errorf("error reading response body: %v", err)
  	}

  	// map的键为列表名称，值为字符串数组
  	var lists map[string][]string
  	err = json.Unmarshal(body, &lists)
  	if err != nil {
    	return nil, fmt.Errorf("error parsing JSON: %v", err)
  	}

  	// 在返回的map中查找指定的列表名称
  	list, ok := lists[listName]
  	if ok {
    // 将slice转换为StringSet
    	stringSet := make(set.StringSet) // 创建一个新的 set.StringSet
    	for _, item := range list {
      		stringSet[item] = struct{}{} // 加入元素
    	}
    	return stringSet, nil
  	}
  	return nil, fmt.Errorf("list %q not found", listName)
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