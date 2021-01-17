package zengin_test

import (
	"testing"

	"github.com/yagitatsu/zengin"
)

func TestBankCode_String(t *testing.T) {
	code := zengin.BankCode(1)
	expected := "0001"

	if code.String() != expected {
		t.Errorf("%s does not equal expected %s", code, expected)
	}
}
