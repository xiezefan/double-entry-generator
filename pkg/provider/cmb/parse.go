package cmb

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

func (w *CMB) translateToOrders(array []string, account string) error {
	for idx, a := range array {
		a = strings.Trim(a, " ")
		a = strings.Trim(a, "\t")
		array[idx] = a
	}

	// EOF
	if array[0] == "" ||
		strings.HasPrefix(array[0], "#") {
		return  nil
	}

	var bill = Order {
		Account: account,
	}

	var err error

	// PayTime
	bill.PayTime, err = time.Parse(localTimeFmt, array[0] + " " + array[1] + " +0800 CST")
	if err != nil {
		return fmt.Errorf("parse create time %s error: %v", array[0], err)
	}

	// Type
	if array[2] != "" {
		bill.Type = OrderTypeRecv
	} else {
		bill.Type = OrderTypeSend
	}

	// TxType
	bill.TxType = array[5]

	// Item
	bill.Item = array[6]

	// Money
	var money string
	if bill.Type == OrderTypeRecv {
		money = array[2]
	} else {
		money = array[3]
	}
	bill.Money, err = strconv.ParseFloat(money, 64)
	if err != nil {
		return fmt.Errorf("parse money %s error: %v",  money, err)
	}

	// Balance
	bill.Balance, err = strconv.ParseFloat(array[4], 64)
	if err != nil {
		return fmt.Errorf("parse money %s error: %v", array[4], err)
	}

	w.Orders = append(w.Orders, bill)
	return nil
}