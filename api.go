package mojang

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"time"
)

//NameAndUUID store player name and UUID
type NameAndUUID struct {
	Name string `json:"name"`
	UUID string `json:"id"`
}

// GetUUID gets UUID by user name at a time.
// when time is nil, the current time is used.
func GetUUID(UserName string, time *time.Time) (nu NameAndUUID, err error) {
	//构造URL
	u := url.URL{
		Scheme: "https",
		Host:   "api.mojang.com",
		Path:   "users/profiles/minecraft/" + UserName,
	}
	if time != nil {
		u.RawQuery = fmt.Sprintf("at=%d", time.Unix())
	}
	//构造请求
	var reque *http.Request
	reque, err = http.NewRequest(http.MethodGet, u.String(), nil)
	if err != nil {
		return
	}
	reque.Header.Set("User-Agent", "miaoscraft")
	//执行请求
	var resp *http.Response
	resp, err = http.DefaultClient.Do(reque)
	if err != nil {
		return
	}
	if resp.StatusCode != 200 {
		err = fmt.Errorf("status code: %d", resp.StatusCode)
		return
	}
	defer resp.Body.Close()
	//解析结果
	d := json.NewDecoder(resp.Body)
	err = d.Decode(&nu)
	return
}
