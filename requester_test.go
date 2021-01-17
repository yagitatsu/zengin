package zengin_test

import (
	"testing"

	"github.com/yagitatsu/zengin"
)

func TestRequester_HeaderRecord(t *testing.T) {
	r := &zengin.Requester{
		Classification: zengin.RequesterClassification_EFT,
		RequesterCode:  1234567890,
	}

	if zengin.LengthSJIS(r.HeaderRecord()) != zengin.RowLength {
		t.Errorf("row length is invalid: length=%d, data=%+v", zengin.LengthSJIS(r.HeaderRecord()), r)
	}
}
