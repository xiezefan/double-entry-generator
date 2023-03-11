package cmb


// Config is the configuration for WeChat.
type Config struct {
	Rules []Rule `mapstructure:"rules,omitempty"`
	IgnoreItem		  *string `mapstructure:"ignoreItem,omitempty"`
	IgnoreTxType	  *string `mapstructure:"ignoreTxType,omitempty"`
}

// Rule is the type for match rules.
type Rule struct {
	Item              *string `mapstructure:"item,omitempty"`
	Type              *string `mapstructure:"type,omitempty"`
	TxType            *string `mapstructure:"txType,omitempty"`
	Money             *float64 `mapstructure:"money,omitempty"`
	Seperator         *string `mapstructure:"sep,omitempty"` // default: ,
	Method            *string `mapstructure:"method,omitempty"`
	MethodAccount     *string `mapstructure:"methodAccount,omitempty"`
	TargetAccount     *string `mapstructure:"targetAccount,omitempty"`
	CommissionAccount *string `mapstructure:"commissionAccount,omitempty"`
	FullMatch         bool    `mapstructure:"fullMatch,omitempty"`
}
