package zengin

import "fmt"

type BranchCode int

func (bc BranchCode) String() string {
	return fmt.Sprintf("%03d", bc)
}
