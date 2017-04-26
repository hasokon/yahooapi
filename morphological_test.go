package yahooapi

import (
	"testing"
	)

func TestMorphological(t *testing.T) {
	_, err := MorphologicalAnalysys("そばやんけ！")
	if err != nil {
		t.Error(err.Error())
	}
}
