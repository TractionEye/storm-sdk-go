package smartaccount

import (
	"github.com/xssnick/tonutils-go/address"
	"github.com/xssnick/tonutils-go/tlb"
	"github.com/xssnick/tonutils-go/tvm/cell"
)

type DepositNative struct {
	_               tlb.Magic        `tlb:"#29bb3721" json:"_"`
	QueryID         uint64           `tlb:"## 64" json:"query_id"`
	Amount          *tlb.Coins       `tlb:"." json:"amount"`
	ReceiverAddress *address.Address `tlb:"addr" json:"receiver_address"`
	Init            bool             `tlb:"bool" json:"init"`
	KeyInit         *cell.Dictionary `tlb:"maybe dict 256" json:"key_init"`
}
