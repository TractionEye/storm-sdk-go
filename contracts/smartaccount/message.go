package smartaccount

import (
	"github.com/xssnick/tonutils-go/address"
	"github.com/xssnick/tonutils-go/tlb"
	"github.com/xssnick/tonutils-go/tvm/cell"
)

type DeployOrdinarySA struct {
	_          tlb.Magic        `tlb:"#764019e5" json:"_"`
	QueryID    uint64           `tlb:"## 64" json:"query_id"`
	PublicKeys *cell.Dictionary `tlb:"maybe dict 256" json:"public_keys"`
}

type DepositNativePayload struct {
	_               tlb.Magic        `tlb:"#29bb3721" json:"_"`
	QueryID         uint64           `tlb:"## 64" json:"query_id"`
	Amount          *tlb.Coins       `tlb:"." json:"amount"`
	ReceiverAddress *address.Address `tlb:"addr" json:"receiver_address"`
	Init            bool             `tlb:"bool" json:"init"`
	KeyInit         *cell.Dictionary `tlb:"maybe dict 256" json:"key_init"`
}

type DepositJettonPayload struct {
	_               tlb.Magic        `tlb:"#76840119" json:"_"`
	QueryID         uint64           `tlb:"## 64" json:"query_id"`
	ReceiverAddress *address.Address `tlb:"addr" json:"receiver_address"`
	Init            bool             `tlb:"bool" json:"init"`
	KeyInit         *cell.Dictionary `tlb:"maybe dict 256" json:"key_init"`
}

type WithdrawPayload struct {
	_            tlb.Magic        `tlb:"#6eec039d" json:"_"`
	QueryID      uint64           `tlb:"## 64" json:"query_id"`
	VaultAddress *address.Address `tlb:"addr" json:"vault_address"`
	Amount       *tlb.Coins       `tlb:"." json:"amount"`
}

type AddPublicKeyPayload struct {
	_         tlb.Magic `tlb:"#220c4c19" json:"_"`
	QueryID   uint64    `tlb:"## 64" json:"query_id"`
	PublicKey PublicKey `tlb:"bits 256" json:"public_key"`
}

type RemovePublicKeyPayload struct {
	_         tlb.Magic `tlb:"#7427ce1f" json:"_"`
	QueryID   uint64    `tlb:"## 64" json:"query_id"`
	PublicKey PublicKey `tlb:"bits 256" json:"public_key"`
}

type RemoveAllExceptCurrentPublicKeyPayload struct {
	_         tlb.Magic `tlb:"#5f9d0940" json:"_"`
	QueryID   uint64    `tlb:"## 64" json:"query_id"`
	PublicKey PublicKey `tlb:"bits 256" json:"public_key"`
}
