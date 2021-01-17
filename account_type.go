package zengin

type AccountType int

const (
	AccountType_Ordinary AccountType = 1 // 普通
	AccountType_Checking AccountType = 2 // 当座
	AccountType_Deposit  AccountType = 4 // 貯蓄
)

func (at AccountType) String() string {
	switch at {
	case AccountType_Ordinary:
		return "普通"
	case AccountType_Checking:
		return "当座"
	case AccountType_Deposit:
		return "貯蓄"
	}
	return ""
}
