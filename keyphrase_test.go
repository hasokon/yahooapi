package yahooapi

import (
	"testing"
	"fmt"
	)

func TestKeyphrase(t *testing.T) {
	result, err := KeyphraseExtraction("そばやんけ！")
	if err != nil {
		t.Error(err.Error())
	}
	fmt.Print(result)
}
