package cmb_credit

import (
	"fmt"
	"strconv"
	"time"
)

func (c *CmbCredit) translateToOrders(line BillItem) error {

	var bill = Order {}

	var err error

	bill.PayTime, err = time.Parse(billTimeFmt, line.BillDate + " +0800 CST")
	if err != nil {
		return fmt.Errorf("parse create time %s error: %v", line.BillDate, err)
	}

	bill.PostingDate = line.PostingDate
	bill.BillMonth = line.BillMonth
	bill.TransactionType = line.TransactionType
	bill.Money, err = strconv.ParseFloat(line.Amount, 64)
	if err != nil {
		return fmt.Errorf("parse money %s error: %v", line.Amount, err)
	}

	bill.TransactionMoney, err = strconv.ParseFloat(line.TransactionAmount, 64)
	if err != nil {
		return fmt.Errorf("parse transaction money %s error: %v", line.TransactionAmount, err)
	}
	if bill.Money > 0 {
		bill.TxType = OrderTypeSend
	} else {
		bill.TxType = OrderTypeRecv
		bill.Money = -bill.Money
		bill.TransactionMoney = -bill.TransactionMoney
	}
	bill.Item = line.Description
	bill.CardNo = line.CardNo
	if line.TotalStages != "" {
		bill.BillInstallment = line.CurrentStages + "/" + line.TotalStages
	}

	c.Orders = append(c.Orders, bill)
	return nil
}

