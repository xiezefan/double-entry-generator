package cmb

import (
	"time"
)

const (
	// localTimeFmt set time format to utc+8
	localTimeFmt = "20060102 15:04:05 +0800 CST"
)

type Statistics struct {
	TxDate		string	`json:"tx_date,omitempty"`
	TxTime		string	`json:"tx_time,omitempty"`
	Income		string	`json:"income,omitempty"`
	Expense		string	`json:"expense,omitempty"`
	Balance		string	`json:"balance,omitempty"`
	TxType		string	`json:"tx_type,omitempty"`
	Item		string	`json:"item,omitempty"`
}

type Order struct {
	PayTime			time.Time // 付款时间
	Type			OrderType // 收/支
	TxType			string  // 交易类型
	Item           	string  // 交易备注
	Money			float64 // 金额
	Balance			float64 // 金额
	Account			string // 交易账号
}

type OrderType string

const (
	OrderTypeSend    OrderType = "支出"
	OrderTypeRecv              = "收入"
)