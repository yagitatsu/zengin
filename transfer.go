package zengin

import (
	"fmt"

	"github.com/pkg/errors"
)

// Transfer is required data to transfer
type Transfer struct {
	BankCode      BankCode
	BranchCode    BranchCode
	AccountType   AccountType
	AccountNumber AccountNumber
	AccountName   AccountName
	Amount        int
}

func (t *Transfer) DataRecord() (string, error) {
	an, err := t.AccountName.ZenginText()
	if err != nil {
		return "", errors.Wrap(err, "failed to encode account name")
	}
	return fmt.Sprintf("2%s               %s               0000%d%s%s%010d0                    0        \r\n", t.BankCode, t.BranchCode, t.AccountType, t.AccountNumber, an, t.Amount), nil
}
