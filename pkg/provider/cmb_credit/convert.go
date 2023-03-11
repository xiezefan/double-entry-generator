package cmb_credit

import (
	"github.com/deb-sig/double-entry-generator/pkg/ir"
	"strconv"
	"strings"
)

func (c *CmbCredit) convertToIR() *ir.IR {
	i := ir.New()
	for _, o := range c.Orders {
		peer := ""
		item := ""
		if strings.Index(o.Item, "-") > 0 {
			ret := strings.Split(o.Item, "-")
			peer = ret[0]
			item = ret[1]
		} else {
			peer = o.Item
		}
		peer = strings.ReplaceAll(peer, "         ", " ")
		peer = strings.ReplaceAll(peer, "        ", " ")

		irO := ir.Order{
			Category: o.TransactionType,
			Peer:     peer,
			Item:     item,
			PayTime:  o.PayTime,
			Money:    o.Money,
			Type:   convertType(o.TxType),
			Method:   o.CardNo,
		}

		irO.Metadata = convertMetadata(o)
		i.Orders = append(i.Orders, irO)
	}
	return i
}

func convertMetadata(o Order) map[string]string {
	result := map[string]string{
		"bill_month":        o.BillMonth,
		"pay_time":          o.PayTime.String(),
		"posting_date":      o.PostingDate,
		"transaction_type":  string(o.TxType),
		"transaction_money": strconv.FormatFloat(o.TransactionMoney, 'f', 2, 64),
		"card_no":           o.CardNo,
	}
	if o.BillInstallment != "" {
		result["bill_installment"] = o.BillInstallment
	}
	return result
}

func convertType(orderType OrderType) ir.Type {
	switch orderType {
	case OrderTypeSend:
		return ir.TypeSend
	case OrderTypeRecv:
		return ir.TypeRecv
	default:
		panic("unknown order type " + orderType)
	}
}
