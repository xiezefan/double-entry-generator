package cmb_credit

import (
	"github.com/deb-sig/double-entry-generator/pkg/config"
	"github.com/deb-sig/double-entry-generator/pkg/ir"
	"github.com/deb-sig/double-entry-generator/pkg/util"
	"log"
)

type CmbCredit struct {}

func (c CmbCredit) GetAllCandidateAccounts(cfg *config.Config) map[string]bool {
	// uniqMap will be used to create the concepts.
	uniqMap := make(map[string]bool)

	if cfg.CmbCredit == nil || len(cfg.CmbCredit.Rules) == 0 {
		return uniqMap
	}

	for _, r := range cfg.CmbCredit.Rules {
		if r.MethodAccount != nil {
			uniqMap[*r.MethodAccount] = true
		}
		if r.TargetAccount != nil {
			uniqMap[*r.TargetAccount] = true
		}
		if r.CommissionAccount != nil {
			uniqMap[*r.CommissionAccount] = true
		}
	}
	uniqMap[cfg.DefaultPlusAccount] = true
	uniqMap[cfg.DefaultMinusAccount] = true
	return uniqMap
}

func (c CmbCredit) GetAccounts(o *ir.Order, cfg *config.Config, target, provider string) (string, string, map[ir.Account]string) {
	var resCommission string
	// check this tx whether has commission
	if o.Commission != 0 {
		if cfg.DefaultCommissionAccount == "" {
			log.Fatalf("Found a tx with commission, but not setting the `defaultCommissionAccount` in config file!")
		} else {
			resCommission = cfg.DefaultCommissionAccount
		}
	}

	if cfg.CmbCredit == nil || len(cfg.CmbCredit.Rules) == 0 {
		return cfg.DefaultMinusAccount, cfg.DefaultPlusAccount, map[ir.Account]string{
			ir.CommissionAccount: resCommission,
		}
	}


	// default TxType = Send
	resMinus := cfg.DefaultMinusAccount
	resPlus := cfg.DefaultPlusAccount

	for _, r := range cfg.CmbCredit.Rules {

		match := true
		// get seperator
		sep := ","
		if r.Seperator != nil {
			sep = *r.Seperator
		}

		matchFunc := util.SplitFindContains
		if r.FullMatch {
			matchFunc = util.SplitFindEquals
		}

		if r.Money != nil {
			match = match && *r.Money == o.Money
		}
		if r.TxType != nil {
			txType := ""
			if o.TxType == ir.TxTypeSend {
				txType = "支出"
			} else if o.TxType == ir.TxTypeRecv {
				txType = "收入"
			}
			match = matchFunc(*r.TxType, txType, sep, match)
		}
		if r.Method != nil {
			match = matchFunc(*r.Method, o.Method, sep, match)
		}
		if r.Item != nil {
			match = matchFunc(*r.Item, o.Item, sep, match)
		}
		if r.Category != nil {
			match = matchFunc(*r.Category, o.Category, sep, match)
		}
		if r.Peer != nil {
			match = matchFunc(*r.Peer, o.Peer, sep, match)
		}

		if match {
			// Support multiple matches, like one rule matches the minus accout, the other rule matches the plus account.
			if r.TargetAccount != nil {
				if o.TxType == ir.TxTypeRecv {
					resMinus = *r.TargetAccount
				} else {
					resPlus = *r.TargetAccount
				}
			}
			if r.MethodAccount != nil {
				if o.TxType == ir.TxTypeRecv {
					resPlus = *r.MethodAccount
				} else {
					resMinus = *r.MethodAccount
				}
			}
			if r.CommissionAccount != nil {
				resCommission = *r.CommissionAccount
			}
		}
	}

	return resMinus, resPlus, map[ir.Account]string{
		ir.CommissionAccount: resCommission,
	}
}

func (c CmbCredit) IgnoreItem(o *ir.Order, cfg *config.Config) bool {

	if cfg.CmbCredit == nil || (cfg.CmbCredit.IgnorePeer == nil && cfg.CmbCredit.IgnoreItem == nil && cfg.CmbCredit.IgnoreCategory == nil) {
		return false
	}

	// get seperator
	defaultSep := ","
	matchFunc := util.SplitFindContains

	if cfg.CmbCredit.IgnorePeer != nil {
		if matchFunc(*cfg.CmbCredit.IgnorePeer, o.Peer, defaultSep, true) {
			return true
		}
	}
	if cfg.CmbCredit.IgnoreItem != nil {
		if matchFunc(*cfg.CmbCredit.IgnoreItem, o.Item, defaultSep, true) {
			return true
		}
	}
	if cfg.CmbCredit.IgnoreCategory != nil {
		if matchFunc(*cfg.CmbCredit.IgnoreCategory, o.Category, defaultSep, true) {
			return true
		}
	}

	return false
}
