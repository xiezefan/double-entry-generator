package cmb_credit

// Config is the configuration for WeChat.
type Config struct {
	Rules          []Rule  `mapstructure:"rules,omitempty"`
	IgnoreItem     *string `mapstructure:"ignoreItem,omitempty"`
	IgnoreCategory *string `mapstructure:"ignoreCategory,omitempty"`
	IgnorePeer     *string `mapstructure:"ignorePeer,omitempty"`
}

// Rule is the type for match rules.
type Rule struct {
	Item              *string  `mapstructure:"item,omitempty"`
	Peer              *string  `mapstructure:"peer,omitempty"`
	Category          *string  `mapstructure:"category,omitempty"`
	TxType            *string  `mapstructure:"txType,omitempty"`
	Money             *float64 `mapstructure:"money,omitempty"`
	Seperator         *string  `mapstructure:"sep,omitempty"` // default: ,
	Method            *string  `mapstructure:"method,omitempty"`
	MethodAccount     *string  `mapstructure:"methodAccount,omitempty"`
	TargetAccount     *string  `mapstructure:"targetAccount,omitempty"`
	CommissionAccount *string  `mapstructure:"commissionAccount,omitempty"`
	FullMatch         bool     `mapstructure:"fullMatch,omitempty"`
}
