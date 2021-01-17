package zengin

import (
	"bufio"
	"bytes"
	"fmt"

	"github.com/pkg/errors"
	"golang.org/x/text/encoding/japanese"
	"golang.org/x/text/transform"
	"golang.org/x/text/unicode/norm"
	"golang.org/x/text/width"
)

const (
	RowLength      = 122
	ZENGIN_TRAILER = "8                                                                                                                       \r\n"
	//スペース1文字消えるので1つ多めに入れる
	ZENGIN_END = "9                                                                                                                       \r\n "
)

func Marshal(r *Requester, ts []*Transfer) (string, error) {
	buf := new(bytes.Buffer)
	writer := bufio.NewWriter(transform.NewWriter(buf, transform.Chain(
		norm.NFD,
		width.Narrow,
		japanese.ShiftJIS.NewEncoder(),
	)))
	fmt.Fprint(writer, r.HeaderRecord())

	for _, t := range ts {
		dr, err := t.DataRecord()
		if err != nil {
			return "", errors.Wrap(err, "failed to create data record")
		}
		fmt.Fprint(writer, dr)
	}
	fmt.Fprint(writer, ZENGIN_TRAILER)
	fmt.Fprint(writer, ZENGIN_END)
	writer.Flush()
	return buf.String(), nil
}
