package zengin

import "fmt"

type Requester struct {
	Classification RequesterClassification
	RequesterCode  int
}

const HEADER_FORMAT = "1%s0%010d                                        00000000               000               00000000                 \r\n"

func (r *Requester) HeaderRecord() string {
	return fmt.Sprintf(HEADER_FORMAT, r.Classification, r.RequesterCode)
}
