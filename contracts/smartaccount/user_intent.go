package smartaccount

import (
	"encoding/hex"
	"github.com/TractionEye/storm-sdk-go/contracts/hw"
	schemas "github.com/TractionEye/storm-sdk-go/tlb"

	"github.com/xssnick/tonutils-go/address"
	"github.com/xssnick/tonutils-go/tlb"
)

type PaymentRequest struct {
	Amount       *tlb.Coins       `tlb:"." json:"amount"`
	VaultAddress *address.Address `tlb:"addr" json:"vault_address"`
}

// UserIntentPayload user_intent_payload#_ amm_address:MsgAddressInt sa_address:MsgAddressInt direction:Direction order:^UserOrder = UserIntentPayload;
type UserIntentPayload struct {
	VAmm         *address.Address `tlb:"addr" json:"v_amm"`
	SmartAccount *address.Address `tlb:"addr" json:"smart_account"`
	Direction    uint             `tlb:"## 1" json:"direction"`
	Order        *schemas.Order   `tlb:"^" json:"order"`
}

// UserIntent _ query_id:UserQueryId created_at:uint32 reference_query_id:(Maybe UserQueryId) public_key:bits256 intent:^UserIntentPayload = UserIntent;
type UserIntent struct {
	QueryId          *hw.QueryId        `tlb:"." json:"query_id"`
	CreatedAt        uint64             `tlb:"## 32" json:"created_at"`
	ReferenceQueryId *hw.QueryId        `tlb:"maybe ." json:"reference_query_id"`
	PublicKey        []byte             `tlb:"bits 256" json:"public_key"` // pk
	Intent           *UserIntentPayload `tlb:"^"`
}

func (msg *UserIntent) Hash() (string, error) {
	msgCell, err := tlb.ToCell(msg)
	if err != nil {
		return "", err
	}

	return hex.EncodeToString(msgCell.Hash()), nil
}
