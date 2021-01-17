package zengin

import "fmt"

type RequesterClassification int

const (
	RequesterClassification_EFT RequesterClassification = 21
)

func (rc RequesterClassification) String() string {
	return fmt.Sprintf("%02d", rc)
}
