package yahooapi

import (
	"testing"
	"fmt"
	)

func TestKana2Kanji(t *testing.T) {
	result, err := ChangeKana2Kanji("そばやんけ！")
	if err != nil {
		t.Error(err.Error())
	}
	fmt.Print(result)
}
