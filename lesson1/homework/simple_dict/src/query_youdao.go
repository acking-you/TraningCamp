package src

//发现问题：有道的请求头是通过动态加盐和sign的方式来请求每一次翻译，所以如果想直接翻译所有，则需要破解出sign的计算原理
import (
	"crypto/md5"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"
)

//TODO sign解析未完成
func getMD5(content string) string {
	data := []byte(content)
	has := md5.Sum(data)
	md5str := fmt.Sprintf("%x", has) //将[]byte转成16进制
	return md5str
}

//暂未破解成功，这个sign不清楚怎么获得
func getYoudaoHeader(text string) (header string) {
	salt := strconv.FormatInt(time.Now().UnixNano()/1e5, 10)
	sign_str := "fanyideskweb" + text + salt + "Ygy_4c=r#e#4EX^NUGUc5"
	sign := getMD5(sign_str)
	//fmt.Println(sign)
	//fmt.Println(salt)
	var requestText = `i=%s&from=AUTO&to=AUTO&smartresult=dict&client=fanyideskweb&salt=%s&sign=%s&lts=1651919604941&bv=c2777327e4e29b7c4728f13e47bde9a5&doctype=json&version=2.1&keyfrom=fanyi.web&action=FY_BY_REALTlME`
	header = fmt.Sprintf(requestText, salt, sign, text)
	return
}

func QueryYoudao(word string) {
	client := &http.Client{}
	var data = strings.NewReader(getYoudaoHeader(word))
	req, err := http.NewRequest("POST", "https://fanyi.youdao.com/translate_o?smartresult=dict&smartresult=rule", data)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("Connection", "keep-alive")
	req.Header.Set("sec-ch-ua", `" Not A;Brand";v="99", "Chromium";v="99", "Google Chrome";v="99"`)
	req.Header.Set("Accept", "application/json, text/javascript, */*; q=0.01")
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded; charset=UTF-8")
	req.Header.Set("X-Requested-With", "XMLHttpRequest")
	req.Header.Set("sec-ch-ua-mobile", "?0")
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/99.0.4844.51 Safari/537.36")
	req.Header.Set("sec-ch-ua-platform", `"Windows"`)
	req.Header.Set("Origin", "https://fanyi.youdao.com")
	req.Header.Set("Sec-Fetch-Site", "same-origin")
	req.Header.Set("Sec-Fetch-Mode", "cors")
	req.Header.Set("Sec-Fetch-Dest", "empty")
	req.Header.Set("Referer", "https://fanyi.youdao.com/")
	req.Header.Set("Accept-Language", "zh-CN,zh;q=0.9")
	req.Header.Set("Cookie", "JSESSIONID=abc90mh6hMYs74Yg2uDcy; OUTFOX_SEARCH_USER_ID=-1937588996@10.110.96.158; OUTFOX_SEARCH_USER_ID_NCOO=1100955033.7040293; fanyi-ad-id=305838; fanyi-ad-closed=0; ___rl__test__cookies=1651915512175")
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	bodyText, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	if resp.StatusCode != 200 { //防止返回错误
		log.Fatal("bad Status code:", resp.StatusCode, "body", string(bodyText))
	}

	var dictResponse DictResponseYoudao
	err = json.Unmarshal(bodyText, &dictResponse)
	if err != nil {
		log.Fatal(err)
	}

	//打印翻译答案
	fmt.Println("translate from Youdao:")
	for _, item := range dictResponse.TranslateResult {
		for _, text := range item {
			fmt.Println(text.Tgt)
		}
	}
	for _, item := range dictResponse.SmartResult.Entries {
		fmt.Print(item)
	}
}
