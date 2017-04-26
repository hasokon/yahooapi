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

type Word struct {
	Surface string `xml:"surface"`
	Reading string `xml:"reading"`
	Pos string `xml:"pos"`
	Baseform string `xml:"baseform"`
}

func (w Word) String() string {
	return fmt.Sprintf("%s (%s)",w.Surface, w.Pos)
}

type WordList struct {
	Wordlist []Word `xml:"word"`
}

func (w WordList) String() string {
	words := []byte{}
	for _, word := range w.Wordlist {
		words = append(words, []byte(word.String()+"\n")...)
	}
	return string(words)
}

type MaResult struct {
	TotalCount int `xml:"total_count"`
	FilteredCount int `xml:"filtered_count"`
	Wordlist WordList `xml:"word_list"`
}

func (m MaResult) String() string {
	return fmt.Sprintf("Total Count: %d\nFilterCount: %d\n%s",m.TotalCount, m.FilteredCount,m.Wordlist)
}

type Morphological struct {
	Ma MaResult `xml:"ma_result"`
}

func (r *Morphological) String() string {
	return r.Ma.String()
}

func MorphologicalAnalysys(text string) (*Morphological, error) {
	yahooId := os.Getenv("YAHOO_ID")

	client := &http.Client{}

	data := url.Values{
		"sentence" : {text},
	}

	req, err := http.NewRequest(
		"POST",
		"https://jlp.yahooapis.jp/MAService/V1/parse",
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

	xmlData := new(Morphological)
	if err := xml.Unmarshal([]byte(string(body)), xmlData); err != nil {
		return nil,err
	}

	return xmlData, nil
}
