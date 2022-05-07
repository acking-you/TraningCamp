package src

//彩云翻译引擎对应的结构体类型
type DictRequestCaiyun struct {
	TransType string `json:"trans_type"`
	Source    string `json:"source"`
	UserID    string `json:"user_id"`
}

type DictResponseCaiyun struct {
	Rc   int `json:"rc"`
	Wiki struct {
		KnownInLaguages int `json:"known_in_laguages"`
		Description     struct {
			Source string      `json:"source"`
			Target interface{} `json:"target"`
		} `json:"description"`
		ID   string `json:"id"`
		Item struct {
			Source string `json:"source"`
			Target string `json:"target"`
		} `json:"item"`
		ImageURL  string `json:"image_url"`
		IsSubject string `json:"is_subject"`
		Sitelink  string `json:"sitelink"`
	} `json:"wiki"`
	Dictionary struct {
		Prons struct {
			EnUs string `json:"en-us"`
			En   string `json:"en"`
		} `json:"prons"`
		Explanations []string      `json:"explanations"`
		Synonym      []string      `json:"synonym"`
		Antonym      []string      `json:"antonym"`
		WqxExample   [][]string    `json:"wqx_example"`
		Entry        string        `json:"entry"`
		Type         string        `json:"type"`
		Related      []interface{} `json:"related"`
		Source       string        `json:"source"`
	} `json:"dictionary"`
}

//Deepl对应的结构体
type DictResponseDeepl struct {
	Jsonrpc string `json:"jsonrpc"`
	ID int `json:"id"`
	Result struct {
		Translations []struct {
			Beams []struct {
				Sentences []struct {
					Text string `json:"text"`
					Ids []int `json:"ids"`
				} `json:"sentences"`
				NumSymbols int `json:"num_symbols"`
			} `json:"beams"`
			Quality string `json:"quality"`
		} `json:"translations"`
		TargetLang string `json:"target_lang"`
		SourceLang string `json:"source_lang"`
		SourceLangIsConfident bool `json:"source_lang_is_confident"`
		DetectedLanguages struct {
			EN float64 `json:"EN"`
			DE float64 `json:"DE"`
			FR float64 `json:"FR"`
			ES float64 `json:"ES"`
			PT float64 `json:"PT"`
			IT float64 `json:"IT"`
			NL float64 `json:"NL"`
			PL float64 `json:"PL"`
			RU float64 `json:"RU"`
			ZH float64 `json:"ZH"`
			JA float64 `json:"JA"`
			BG float64 `json:"BG"`
			CS float64 `json:"CS"`
			DA float64 `json:"DA"`
			EL float64 `json:"EL"`
			ET float64 `json:"ET"`
			FI float64 `json:"FI"`
			HU float64 `json:"HU"`
			LT float64 `json:"LT"`
			LV float64 `json:"LV"`
			RO float64 `json:"RO"`
			SK float64 `json:"SK"`
			SL float64 `json:"SL"`
			SV float64 `json:"SV"`
			Unsupported float64 `json:"unsupported"`
		} `json:"detectedLanguages"`
	} `json:"result"`
}

// Youdao
type DictResponseYoudao struct {
	TranslateResult [][]struct {
		Tgt string `json:"tgt"`
		Src string `json:"src"`
	} `json:"translateResult"`
	ErrorCode int `json:"errorCode"`
	Type string `json:"type"`
	SmartResult struct {
		Entries []string `json:"entries"`
		Type int `json:"type"`
	} `json:"smartResult"`
}
