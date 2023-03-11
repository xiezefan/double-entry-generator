package cmb

import (
	"github.com/deb-sig/double-entry-generator/pkg/config"
	"github.com/deb-sig/double-entry-generator/pkg/ir"
	"github.com/deb-sig/double-entry-generator/pkg/util"
	"log"
)

type CMB struct {
}

func (c CMB) IgnoreItem(o *ir.Order, cfg *config.Config) bool {
	if cfg.CMB == nil || (cfg.CMB.IgnoreItem == nil && cfg.CMB.IgnoreTxType == nil) {
		return false
	}

	// get seperator
	defaultSep := ","
	matchFunc := util.SplitFindContains
	if cfg.CMB.IgnoreItem != nil {
		if matchFunc(*cfg.CMB.IgnoreItem, o.Item, defaultSep, true) {
			return true
		}
	}
	if cfg.CMB.IgnoreTxType != nil {
		if matchFunc(*cfg.CMB.IgnoreTxType, o.TxTypeOriginal, defaultSep, true) {
			return true
		}
	}

	return false
}

// GetAllCandidateAccounts returns all accounts defined in config.
func (c CMB) GetAllCandidateAccounts(cfg *config.Config) map[string]bool {
	// uniqMap will be used to create the concepts.
	uniqMap := make(map[string]bool)

	if cfg.CMB == nil || len(cfg.CMB.Rules) == 0 {
		return uniqMap
	}

	for _, r := range cfg.CMB.Rules {
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


func (c CMB) GetAccountsAndTags(o *ir.Order, cfg *config.Config, _, _ string) (string, string, map[ir.Account]string, []string) {
	var resCommission string
	// check this tx whether has commission
	if o.Commission != 0 {
		if cfg.DefaultCommissionAccount == "" {
			log.Fatalf("Found a tx with commission, but not setting the `defaultCommissionAccount` in config file!")
		} else {
			resCommission = cfg.DefaultCommissionAccount
		}
	}

	if cfg.CMB == nil || len(cfg.CMB.Rules) == 0 {
		return cfg.DefaultMinusAccount, cfg.DefaultPlusAccount, map[ir.Account]string{
			ir.CommissionAccount: resCommission,
		}, []string{}
	}

	resMinus := cfg.DefaultMinusAccount
	resPlus := cfg.DefaultPlusAccount

	for _, r := range cfg.CMB.Rules {

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

		if r.Type != nil {
			match = matchFunc(*r.Type, o.TxTypeOriginal, sep, match)
		}
		if r.TxType != nil {
			match = matchFunc(*r.TxType, o.TypeOriginal, sep, match)
		}
		if r.Method != nil {
			match = matchFunc(*r.Method, o.Method, sep, match)
		}
		if r.Item != nil {
			match = matchFunc(*r.Item, o.Item, sep, match)
		}

		if match {
			// Support multiple matches, like one rule matches the minus accout, the other rule matches the plus account.
			if r.TargetAccount != nil {
				if o.Type == ir.TypeRecv {
					resMinus = *r.TargetAccount
				} else {
					resPlus = *r.TargetAccount
				}
			}
			if r.MethodAccount != nil {
				if o.Type == ir.TypeRecv {
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
	}, []string{}
}

