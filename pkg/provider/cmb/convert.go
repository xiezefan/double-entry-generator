package cmb

import (
	"github.com/deb-sig/double-entry-generator/pkg/ir"
)

// convertToIR convert CMB bills to IR.
func (w *CMB) convertToIR() *ir.IR {
	i := ir.New()
	for _, o := range w.Orders {
		irO := ir.Order{
			Peer: o.TxType,
			Item: o.Item,
			PayTime: o.PayTime,
			Money: o.Money,
			Type: convertType(o.Type),
			TxTypeOriginal: string(o.Type),
			TypeOriginal: o.TxType,
			Method: o.Account,
		}

		irO.Metadata = convertMetadata(o)
		i.Orders = append(i.Orders, irO)
	}
	return i
}

func convertMetadata(o Order) map[string]string {

	return map[string]string{
		"card_no": o.Account,
		"pay_time": o.PayTime.String(),
		"type": string(o.Type),
		"tx_type": o.TxType,
	}
}

func convertType(orderType OrderType) ir.Type {
	switch orderType {
	case OrderTypeSend:
		return ir.TypeSend
	case OrderTypeRecv:
		return ir.TypeRecv
	default:
		return ir.TypeUnknown
	}
}

