package params

// Default simulation operation weights for messages and gov proposals
const (
	DefaultWeightMsgSend                      int = 100
	DefaultWeightMsgMultiSend                 int = 10
	DefaultWeightMsgCreateValidator           int = 100
	DefaultWeightMsgEditValidator             int = 5
	DefaultWeightMsgDelegate                  int = 100
	DefaultWeightMsgUndelegate                int = 100
	DefaultWeightMsgBeginRedelegate           int = 100
	DefaultWeightMsgCancelUnbondingDelegation int = 100
)
