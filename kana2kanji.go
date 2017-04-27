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

type CandidateList struct {
	Candidates []string `xml:"Candidate"`
}

func (c CandidateList) String() string {
	candidates := []byte("<")

	for _, candidate := range c.Candidates {
		candidates = append(candidates, []byte(candidate + " ")...)
	}

	candidates = append(candidates, []byte(">\n")...)

	return string(candidates)
}

type SegmentMold struct {
	Text string `xml:"SegmentText"`
	Candidates CandidateList `xml:"CandidateList"`
}

func (s SegmentMold) String() string {
	return fmt.Sprintf("<%s>\n%s", s.Text, s.Candidates.String())
}

type SegmentListMold struct {
	Segments []SegmentMold `xml:"Segment"`
}

func (s SegmentListMold) String() string {
	segments := []byte{}

	for _, segment := range s.Segments {
		segments = append(segments, []byte(segment.String())...)
	}

	return string(segments)
}

type ResultMold struct {
	SegmentList SegmentListMold `xml:"SegmentList"`
}

func (r ResultMold) String() string {
	return r.SegmentList.String()
}

type Kana2Kanji struct {
	Result ResultMold `xml:"Result"`
}

func (k *Kana2Kanji) String() string {
	return k.Result.String()
}

func ChangeKana2Kanji (text string) (*Kana2Kanji, error) {
	yahooId := os.Getenv("YAHOO_ID")

	client := &http.Client{}

	data := url.Values{
		"sentence" : {text},
	}

	req, err := http.NewRequest(
		"POST",
		"https://jlp.yahooapis.jp/JIMService/V1/conversion",
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

	xmlData := new(Kana2Kanji)
	if err := xml.Unmarshal([]byte(string(body)), xmlData); err != nil {
		return nil,err
	}

	return xmlData, nil
}
