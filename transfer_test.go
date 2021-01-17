package zengin_test

import (
	"testing"

	"github.com/yagitatsu/zengin"
)

func TestTransfer_DataRecord(t *testing.T) {
	tr := &zengin.Transfer{
		BankCode:      zengin.BankCode(1),
		BranchCode:    zengin.BranchCode(1),
		AccountType:   zengin.AccountType(1),
		AccountNumber: zengin.AccountNumber(123456),
		AccountName:   zengin.AccountName("ヤマダタロウ"),
		Amount:        100,
	}
	str, err := tr.DataRecord()
	if err != nil {
		t.Errorf("failed to do DataRecord: err=%s", err)
	}
	if zengin.LengthSJIS(str) != zengin.RowLength {
		t.Errorf("row length is invalid: length=%d, data=%+v", zengin.LengthSJIS(str), tr)
	}
}
