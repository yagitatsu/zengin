package zengin_test

import (
	"testing"

	"github.com/yagitatsu/zengin"
)

func TestAccountNumber_String(t *testing.T) {
	an := zengin.AccountNumber(123456)
	expected := "0123456"

	if an.String() != expected {
		t.Errorf("%s does not equal expected %s", an, expected)
	}
}
