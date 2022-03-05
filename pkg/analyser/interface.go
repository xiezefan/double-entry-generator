package analyser

import (
	"fmt"
	"github.com/deb-sig/double-entry-generator/pkg/analyser/alipay"
	"github.com/deb-sig/double-entry-generator/pkg/analyser/cmb"
	"github.com/deb-sig/double-entry-generator/pkg/analyser/htsec"
	"github.com/deb-sig/double-entry-generator/pkg/analyser/huobi"
	"github.com/deb-sig/double-entry-generator/pkg/analyser/wechat"
	"github.com/deb-sig/double-entry-generator/pkg/config"
	"github.com/deb-sig/double-entry-generator/pkg/consts"
	"github.com/deb-sig/double-entry-generator/pkg/ir"
)

// Interface is the interface of analyser.
type Interface interface {
	GetAllCandidateAccounts(cfg *config.Config) map[string]bool
	GetAccounts(o *ir.Order, cfg *config.Config, target, provider string) (string, string, map[ir.Account]string)
	IgnoreItem(o *ir.Order, cfg *config.Config) bool
}

// New creates a new analyser.
func New(providerName string) (Interface, error) {
	switch providerName {
	case consts.ProviderAlipay:
		return alipay.Alipay{}, nil
	case consts.ProviderWechat:
		return wechat.Wechat{}, nil
	case consts.ProviderHuobi:
		return huobi.Huobi{}, nil
	case consts.ProviderHtsec:
		return htsec.Htsec{}, nil
	case consts.ProviderCMB:
		return cmb.CMB{}, nil
	default:
		return nil, fmt.Errorf("fail to create the analyser for the given name %s", providerName)
	}
}
