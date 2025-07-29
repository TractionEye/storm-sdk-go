package opcode

type Opcode = uint

var (
	OpCompleteOrder    = Opcode(0xcf90d618)
	OpUpdatePosition   = Opcode(0x60dfc677)
	OpRevertGasOnError = Opcode(0x2089a59f)

	OpRequestCreateOrder = Opcode(0xe0db7753)
	OpCreateOrder        = Opcode(0xa39843f4)
	OpJettonTransfer     = Opcode(0x0f8a7ea5)
	ProvidePosition      = Opcode(0x8865b402)
	LiquidatePosition    = Opcode(0x13076670)
	AddExecutorAmount    = Opcode(0x5dd66579)
	CompleteOrder        = Opcode(0xcf90d618)
	OrderCreated_V1      = Opcode(0x3a943ce6)
	OrderCreated_V2      = Opcode(0x80f4c55b)

	OpError                      = Opcode(0xffffffff)
	VaultRequestWithdrawPosition = Opcode(0x0226df66)

	WalletSignedExternalW5R1 = Opcode(0x7369676e)

	PayFunding      = Opcode(0xb652c441)
	SyncOraclePrice = Opcode(0xc4ca405d)
)

var HighloadW3InternalTransfer = Opcode(0xae42e5a4)

var PositionOpcodes = []Opcode{OpCompleteOrder, OpUpdatePosition, OpRevertGasOnError, OpError}
