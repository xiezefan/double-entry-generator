package cmb_credit

import (
	"time"
)

const (
	// localTimeFmt set time format to utc+8
	billTimeFmt = "20060102150405 +0800 CST"
)

type Statistics struct {
	Data Data `json:"data,omitempty"`
}

type Data struct {
	Detail []BillItem `json:"detail,omitempty"`
}

type BillItem struct {
	BillId            string  `json:"billId,omitempty"`
	BillType          string  `json:"billType,omitempty"`
	BillDate          string  `json:"billDate,omitempty"`
	BillMonth         string  `json:"billMonth,omitempty"`
	Org               string  `json:"org,omitempty"`
	TransactionAmount string `transactionAmount:"org,omitempty"`
	Amount            string `json:"amount,omitempty"`
	Description       string  `json:"description,omitempty"`
	PostingDate       string  `json:"postingDate,omitempty"`
	Location          string  `json:"location,omitempty"`
	TotalStages       string  `json:"totalStages,omitempty"`
	CurrentStages     string  `json:"currentStages,omitempty"`
	RemainingStages   string  `json:"remainingStages,omitempty"`
	TransactionType   string  `json:"transactionType,omitempty"`
	CardNo            string  `json:"cardNo,omitempty"`
}

type Order struct {
	PayTime          time.Time // 付款时间
	PostingDate      string    // 入账日期
	TxType           OrderType // 收/支
	BillMonth        string    // 账单日期
	TransactionType  string    // 交易类型 -> TransactionType
	Item             string    // 交易备注 -> Description
	Money            float64   // 金额
	TransactionMoney float64   // 原交易金额
	CardNo           string    // 交易账号
	BillInstallment  string    // 账单分期信息
}

type OrderType string

const (
	OrderTypeSend OrderType = "支出"
	OrderTypeRecv           = "收入"
)
