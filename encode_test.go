package zengin_test

import (
	"fmt"
	"testing"

	"github.com/yagitatsu/zengin"
)

func TestMarshal(t *testing.T) {
	r := &zengin.Requester{
		Classification: zengin.RequesterClassification_EFT,
		RequesterCode:  1234567890,
	}
	tr := &zengin.Transfer{
		BankCode:      zengin.BankCode(1),
		BranchCode:    zengin.BranchCode(1),
		AccountType:   zengin.AccountType_Ordinary,
		AccountNumber: zengin.AccountNumber(123456),
		AccountName:   zengin.AccountName("ヤマダタロウ"),
		Amount:        100,
	}
	str, err := zengin.Marshal(r, []*zengin.Transfer{tr})
	if err != nil {
		t.Errorf("failed to marshal: err=%s", err)
	}

	fmt.Println(str)
}
