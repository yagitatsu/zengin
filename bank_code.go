package zengin

import "fmt"

type BankCode int

func (bc BankCode) String() string {
	return fmt.Sprintf("%04d", bc)
}
