package yahooapi

import (
	"encoding/xml"
	"net/http"
	"net/url"
	"strings"
	"os"
	"fmt"
	"io/ioutil"
	)

type Result struct {
	Keyphrase string `xml:"Keyphrase"`
	Score int `xml:"Score"`
}

func (r Result) String() string {
	return fmt.Sprintf("%s (%d)", r.Keyphrase, r.Score)
}

type Keyphrase struct {
	Results []Result `xml:"Result"`
}

func (k *Keyphrase) String() string {
	keyphrasese := make([]byte,0)

	for _, result := range k.Results {
		keyphrasese = append(keyphrasese, []byte(result.String() + "\n")...)
	}

	return string(keyphrasese)
}

func KeyphraseExtraction (text string) (*Keyphrase, error) {
	yahooId := os.Getenv("YAHOO_ID")

	client := &http.Client{}

	data := url.Values{
		"sentence" : {text},
	}

	req, err := http.NewRequest(
		"POST",
		"https://jlp.yahooapis.jp/KeyphraseService/V1/extract",
		strings.NewReader(data.Encode()),
		)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("Host", "jlp.yahooapis.jp")
	req.Header.Set("User-Agent", "Yahoo AppID:"+yahooId)

	resp, err := client.Do(req)
	if err != nil {
		return nil,err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil,err
	}
	defer resp.Body.Close()

	xmlData := new(Keyphrase)
	if err := xml.Unmarshal([]byte(string(body)), xmlData); err != nil {
		return nil,err
	}

	return xmlData, nil
}
