package zengin

import "fmt"

type AccountNumber int

func (an AccountNumber) String() string {
	return fmt.Sprintf("%07d", an)
}
