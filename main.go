package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

type User struct {
	sessionID string
}

type SetInfo struct {
	roomID    string
	setID     int
	day       int
	month     int
	startTime string
	endTime   string
}

var roomIDMap = map[string]string{
	"232": "c710ea4b-91ae-4a9d-a3f2-ca7e31460b68",
}

func request(url string, sessionID string) (err error) {
	method := "GET"

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		fmt.Println(err)
		return
	}
	req.Header.Add("Host", "bgweixin.sspu.edu.cn")
	req.Header.Add("Accept", "*/*")
	req.Header.Add("User-Agent", "Mozilla/5.0 (Linux; Android 13; V2241A Build/TP1A.220624.014; wv) AppleWebKit/537.36 (KHTML, like Gecko) Version/4.0 Chrome/107.0.5304.141 Mobile Safari/537.36 XWEB/5015 MMWEBSDK/20230202 MMWEBID/5056 MicroMessenger/8.0.33.2320(0x28002151) WeChat/arm64 Weixin NetType/WIFI Language/zh_CN ABI/arm64")
	req.Header.Add("X-Requested-With", "XMLHttpRequest")
	req.Header.Add("Sec-Fetch-Site", "same-origin")
	req.Header.Add("Sec-Fetch-Mode", "cors")
	req.Header.Add("Sec-Fetch-Dest", "empty")
	req.Header.Add("Referer", "https://bgweixin.sspu.edu.cn/app/readroom/index.do")
	req.Header.Add("Accept-Language", "zh-CN,zh;q=0.9,en-US;q=0.8,en;q=0.7")
	req.Header.Add("Cookie", sessionID)

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(body))

	return nil
}

func reserve(user *User, sets []*SetInfo) (err error) {
	for _, set := range sets {
		url := fmt.Sprintf("https://bgweixin.sspu.edu.cn/app/readroom/ylsyySave.do?gsid=%v&zwid=%v&day=%v&month=%v&starttime=%v&endtime=%v", roomIDMap[set.roomID], set.setID, set.day, set.month, set.startTime, set.endTime)
		err = request(url, user.sessionID)
		if err == nil {
			break
		}
	}
	return
}

func main() {
	user := &User{sessionID: ""}
	sets := []*SetInfo{
		{
			roomID:    "232",
			setID:     28,
			day:       21,
			month:     3,
			startTime: "9%3A00",
			endTime:   "22%3A00",
		},
	}
	err := reserve(user, sets)
	if err != nil {
		fmt.Println(err)
	}
}
