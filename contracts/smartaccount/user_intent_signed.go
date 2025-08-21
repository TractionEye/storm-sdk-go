package smartaccount

import (
	"crypto/ed25519"
	"encoding/hex"

	"github.com/xssnick/tonutils-go/address"
	"github.com/xssnick/tonutils-go/tlb"
	"github.com/xssnick/tonutils-go/tvm/cell"
)

type SignedMessage struct {
	Message   *UserIntent `tlb:"^" json:"message"`
	Signature []byte      `tlb:"bits 512" json:"signature"`
	PublicKey PublicKey   `tlb:"bits 256" json:"public_key"`
}

type SignedCancelMessage struct {
 _         tlb.Magic      `tlb:"#db3ee02c" json:"_"`
 Message   *CancelMessage `tlb:"^" json:"message"`
 PublicKey []byte         `tlb:"bits 256" json:"public_key"`
 Signature []byte         `tlb:"bits 512" json:"signature"`
}

func (msg *SignedMessage) Hash() (string, error) {
	return msg.Message.Hash()
}

func (msg *SignedMessage) Marshal(saAddr *address.Address) (string, error) {

	userIntentCell, err := tlb.ToCell(msg.Message)
	if err != nil {
		return "", err
	}

	toSend := cell.BeginCell()
	toSend.MustStoreUInt(0x588b3270, 32)
	toSend.MustStoreRef(userIntentCell)
	toSend.MustStoreSlice(msg.Signature, 512)
	toSend.MustStoreSlice(msg.Message.PublicKey, 256)
	toCell := toSend.EndCell()

	toSendExt := &tlb.ExternalMessage{
		DstAddr: saAddr,
		Body: toCell,
	}

	extCellOk, err := tlb.ToCell(toSendExt)
	if err != nil {
    return "", err
  }

	cellBytes := extCellOk.ToBOCWithFlags(false)
	cbString := hex.EncodeToString(cellBytes)
	return cbString, nil
}

func SignMessage(msg *UserIntent, sk ed25519.PrivateKey) (*SignedMessage, error) {
	c, err := tlb.ToCell(msg)
	if err != nil {
		return nil, err
	}

	sign := c.Sign(sk)
	pk := PublicKey(sk.Public().(ed25519.PublicKey))

	return &SignedMessage{Message: msg, Signature: sign, PublicKey: pk}, nil
}

func SignCancelMessage(msg *CancelMessage, sk ed25519.PrivateKey) (*SignedCancelMessage, error) {
	c, err := tlb.ToCell(msg)
	if err != nil {
		return nil, err
	}

	sign := c.Sign(sk)
	pk := PublicKey(sk.Public().(ed25519.PublicKey))

	return &SignedCancelMessage{Message: msg, Signature: sign, PublicKey: pk}, nil
}

func (msg *SignedCancelMessage) Marshal(saAddr *address.Address) (string, error) {

	message, err := tlb.ToCell(msg.Message)
	if err != nil {
		return "", err
	}

	toSend := cell.BeginCell()
	toSend.MustStoreUInt(0xdb3ee02c, 32)
	toSend.MustStoreRef(message)
	toSend.MustStoreSlice(msg.Signature, 512)
	toSend.MustStoreSlice(msg.PublicKey, 256)
	toCell := toSend.EndCell()

	toSendExt := &tlb.ExternalMessage{
		DstAddr: saAddr,
		Body: toCell,
	}

	extCellOk, err := tlb.ToCell(toSendExt)
	if err != nil {
    return "", err
  }

	cellBytes := extCellOk.ToBOCWithFlags(false)
	cbString := hex.EncodeToString(cellBytes)
	return cbString, nil
}
