package yahooapi

import (
	"testing"
	)

func TestKeyphrase(t *testing.T) {
	_, err := KeyphraseExtraction("そばやんけ！")
	if err != nil {
		t.Error(err.Error())
	}
}
