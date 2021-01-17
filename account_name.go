package zengin

import (
	"errors"
	"fmt"
	"strings"

	"golang.org/x/text/unicode/norm"
	"golang.org/x/text/width"
)

type AccountName string

func (an AccountName) ZenginText() (string, error) {
	str := shortName(string(an))
	l := LengthSJIS(str)
	if l > 30 {
		return "", errors.New(fmt.Sprintf("account name length is invalid: %d %s", l, str))
	}
	s := str
	for i := 0; i < 30-l; i++ {
		s = s + " "
	}
	return s, nil
}

func shortName(str string) string {
	s := width.Narrow.String(norm.NFD.String(str))
	if strings.HasPrefix(s, "ﾕｳｹﾞﾝｶﾞｲｼｬ") {
		s = strings.TrimPrefix(s, "ﾕｳｹﾞﾝｶﾞｲｼｬ")
		s = "ﾕ)" + s
	}
	if strings.HasPrefix(s, "ｶﾌﾞｼｷｶﾞｲｼｬ") {
		s = strings.TrimPrefix(s, "ｶﾌﾞｼｷｶﾞｲｼｬ")
		s = "ｶ)" + s
	}
	if strings.HasSuffix(s, "ｶﾌﾞｼｷｶﾞｲｼｬ") {
		s = strings.TrimSuffix(s, "ｶﾌﾞｼｷｶﾞｲｼｬ")
		s = s + "(ｶ"
	}
	if LengthSJIS(s) > 30 {
		return string([]rune(s)[:30])
	}
	return s
}
