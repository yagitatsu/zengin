# zengin
Package zengin implements encoding of Zengin File.

# Usage

```go
	r := &zengin.Requester{
		Classification: zengin.RequesterClassification_EFT,
		RequesterCode:  1234567890,
	}
	tr := &zengin.Transfer{
		BankCode:      zengin.BankCode(1),
		BranchCode:    zengin.BranchCode(1),
		AccountType:   zengin.AccountType_Ordinary,
		AccountNumber: zengin.AccountNumber(123456),
		AccountName:   zengin.AccountName("ヤマダタロウ"),
		Amount:        100,
	}
	str, err := zengin.Marshal(r, []*zengin.Transfer{tr})
	if err != nil {
		t.Errorf("failed to marshal: err=%s", err)
	}

	fmt.Println(str)
```

```
12101234567890                                        00000000               000               00000000                 
20001               001               000010123456?????۳                       00000001000                    0        
8                                                                                                                       
9                                                                                                                       

```
