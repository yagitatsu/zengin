package zengin_test

import (
	"testing"

	"github.com/yagitatsu/zengin"
)

func TestBranchCode_String(t *testing.T) {
	code := zengin.BranchCode(1)
	expected := "001"

	if code.String() != expected {
		t.Errorf("%s does not equal expected %s", code, expected)
	}
}
