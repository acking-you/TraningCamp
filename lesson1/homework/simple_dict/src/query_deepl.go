package src

//Deepl 是我见过翻译句子到英文最好的翻译引擎
import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"unicode"
)


func getDeeplHeader(text string) (header string) {
	var rowHeader = `{"jsonrpc":"2.0","method": "LMT_handle_jobs","params":{"jobs":[{"kind":"default","sentences":[{"text":"%s","id":0,"prefix":""}],"raw_en_context_before":[],"raw_en_context_after":[],"preferred_num_beams":4}],"lang":{"preference":{"weight":{"DE":0.28607,"EN":1.46694,"ES":0.25353,"FR":0.26007,"IT":0.23263,"JA":0.8852,"NL":0.21019,"PL":0.19671,"PT":0.18584,"RU":0.18408,"ZH":5.92388,"BG":0.13602,"CS":0.16022,"DA":0.16811,"EL":0.13271,"ET":0.15213,"FI":0.19951,"HU":0.15119,"LT":0.14065,"LV":0.11681,"RO":0.14776,"SK":0.14827,"SL":0.13934,"SV":0.18033},"default":"default"},"source_lang_user_selected":"EN","target_lang":"ZH"},"priority":1,"commonJobParams":{"browserType":1},"timestamp":1651918799405},"id":83800015}`
	f := false
	for _,ch :=range text{	// 如果是ZH=>EN则进行另一套请求头
		if unicode.Is(unicode.Han,ch){
			f = true
			break
		}
	}
	if f{
		rowHeader =  `{"jsonrpc":"2.0","method": "LMT_handle_jobs","params":{"jobs":[{"kind":"default","sentences":[{"text":"%s","id":0,"prefix":""}],"raw_en_context_before":[],"raw_en_context_after":[],"preferred_num_beams":4,"quality":"fast"}],"lang":{"preference":{"weight":{"DE":0.23796,"EN":0.40289,"ES":0.18361,"FR":0.20778,"IT":0.14221,"JA":1.02578,"NL":0.15049,"PL":0.13814,"PT":0.13626,"RU":0.14849,"ZH":6.45278,"BG":0.10472,"CS":0.10801,"DA":0.1121,"EL":0.10823,"ET":0.1065,"FI":0.12662,"HU":0.10144,"LT":0.10389,"LV":0.08013,"RO":0.10744,"SK":0.10464,"SL":0.10203,"SV":0.12185},"default":"default"},"source_lang_user_selected":"auto","target_lang":"EN"},"priority":-1,"commonJobParams":{"regionalVariant":"en-US","browserType":1,"formality":null},"timestamp":1651912606763},"id":15480013}`
	}
	header = fmt.Sprintf(rowHeader,text)
	return
}

func QueryDeepl(sentence string) {
	client := &http.Client{}
	var data = strings.NewReader(getDeeplHeader(sentence))
	req, err := http.NewRequest("POST", "https://www2.deepl.com/jsonrpc?method=LMT_handle_jobs", data)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("authority", "www2.deepl.com")
	req.Header.Set("sec-ch-ua", `" Not A;Brand";v="99", "Chromium";v="99", "Google Chrome";v="99"`)
	req.Header.Set("sec-ch-ua-mobile", "?0")
	req.Header.Set("user-agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/99.0.4844.51 Safari/537.36")
	req.Header.Set("sec-ch-ua-platform", `"Windows"`)
	req.Header.Set("content-type", "application/json")
	req.Header.Set("accept", "*/*")
	req.Header.Set("origin", "https://www.deepl.com")
	req.Header.Set("sec-fetch-site", "same-site")
	req.Header.Set("sec-fetch-mode", "cors")
	req.Header.Set("sec-fetch-dest", "empty")
	req.Header.Set("referer", "https://www.deepl.com/")
	req.Header.Set("accept-language", "zh-CN,zh;q=0.9")
	req.Header.Set("cookie", "dapUid=0b6c1411-718e-44f2-8c50-30be8a7ab62e; privacySettings=%7B%22v%22%3A%221%22%2C%22t%22%3A1650067200%2C%22m%22%3A%22LAX_AUTO%22%2C%22consent%22%3A%5B%22NECESSARY%22%2C%22PERFORMANCE%22%2C%22COMFORT%22%5D%7D; LMTBID=v2|320b47d1-70c4-4544-9360-a9a112a1f176|e6e06c8d3f0168f11e168a5cbdf0b453; __cf_bm=k1Twpi99FvJCMnDy2Qy71VzeRvvRnp0T5Xs8PXYwe8g-1651912045-0-AdI+N4bHJ5jVMKGhAa4Yg9b+B3KGgWtT6pwbmgHV/43KraYaeVunoUVVGnpf1gb0KJkk0FxpxXJnyrupHhcQUcE=; dapVn=7; dapSid=%7B%22sid%22%3A%226403381c-d78f-424b-baf2-8ade07d83c40%22%2C%22lastUpdate%22%3A1651912606%7D")
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

	var dictResponse DictResponseDeepl
	err = json.Unmarshal(bodyText, &dictResponse)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("translate from Deepl:")

	//print translation result
	for _, translation := range dictResponse.Result.Translations {
		for _,beam := range translation.Beams{
			for _,sentences := range beam.Sentences{
				fmt.Println(sentences.Text)
			}
		}
	}
}
