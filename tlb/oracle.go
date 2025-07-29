package schemas

import (
	"github.com/xssnick/tonutils-go/tlb"
	"github.com/xssnick/tonutils-go/tvm/cell"
)

func init() {
	tlb.Register(OraclePayloadSimple{})
	tlb.Register(OraclePayloadWithSettlement{})
	tlb.Register(OraclePayloadWithSettlementAndCreated{})
	tlb.Register(OraclePayloadSimpleWithCreated{})
}

type OracleRefType string

const (
	OracleRefSimple                    OracleRefType = "simple"
	OracleRefSimpleWithCreated         OracleRefType = "simpleWithCreated"
	OracleRefWithSettlementWithCreated OracleRefType = "withSettlementWithCreated"
	OracleRefWithSettlement            OracleRefType = "withSettlement"
)

type PriceRef struct {
	PriceRef      *cell.Cell `tlb:"^" json:"price_ref"`
	SignaturesRef *cell.Cell `tlb:"^" json:"signatures_ref"`
}

type DoublePriceRef struct {
	PriceRef                *cell.Cell `tlb:"^" json:"price_ref"`
	SignaturesRef           *cell.Cell `tlb:"^" json:"signatures_ref"`
	SettlementRef           *cell.Cell `tlb:"^" json:"settlement_ref"`
	SettlementSignaturesRef *cell.Cell `tlb:"^" json:"settlement_signatures_ref"`
}

type OraclePayloadSimpleWithCreated struct {
	_        tlb.Magic `tlb:"$0010"`
	PriceRef PriceRef  `tlb:"^"`

	CreatedAtRef struct {
		PriceRef      *cell.Cell `tlb:"^" json:"price_ref"`
		SignaturesRef *cell.Cell `tlb:"^" json:"signatures_ref"`
	} `tlb:"^"`
}

type OraclePayloadWithSettlementAndCreated struct {
	_            tlb.Magic      `tlb:"$0010"`
	PriceRef     DoublePriceRef `tlb:"^"`
	CreatedAtRef PriceRef       `tlb:"^"`
}

func (o OraclePayloadSimpleWithCreated) GetType() OracleRefType {
	return OracleRefSimpleWithCreated
}

func (o OraclePayloadWithSettlementAndCreated) GetType() OracleRefType {
	return OracleRefWithSettlementWithCreated
}

type OraclePayloadSimple struct {
	_             tlb.Magic  `tlb:"$0000"`
	PriceRef      *cell.Cell `tlb:"^" json:"price_ref"`
	SignaturesRef *cell.Cell `tlb:"^" json:"signatures_ref"`
}

func (o OraclePayloadSimple) GetType() OracleRefType {
	return OracleRefSimple
}

type OraclePayloadWithSettlement struct {
	_                       tlb.Magic  `tlb:"$0001"`
	PriceRef                *cell.Cell `tlb:"^" json:"price_ref"`
	SignaturesRef           *cell.Cell `tlb:"^" json:"signatures_ref"`
	SettlementRef           *cell.Cell `tlb:"^" json:"settlement_ref"`
	SettlementSignaturesRef *cell.Cell `tlb:"^" json:"settlement_signatures_ref"`
}

func (o OraclePayloadWithSettlement) GetType() OracleRefType {
	return OracleRefWithSettlement
}

type OracleRef interface {
	GetType() OracleRefType
}

type OraclePayload struct {
	Value OracleRef `json:"order"`
}

func BuildOracleRef(payload *OraclePayload) *cell.Cell {
	switch payload.Value.GetType() {
	case OracleRefWithSettlement:
		payload := payload.Value.(*OraclePayloadWithSettlement)

		return cell.BeginCell().
			MustStoreUInt(1, 8).
			MustStoreRef(payload.PriceRef).
			MustStoreRef(payload.SignaturesRef).
			MustStoreRef(payload.SettlementRef).
			MustStoreRef(payload.SettlementSignaturesRef).
			EndCell()

	case OracleRefSimpleWithCreated:
		payload := payload.Value.(*OraclePayloadSimpleWithCreated)

		priceRef := cell.BeginCell().
			MustStoreUInt(0, 8).
			MustStoreRef(payload.PriceRef.PriceRef).
			MustStoreRef(payload.PriceRef.SignaturesRef).
			EndCell()

		createdRef := cell.BeginCell().
			MustStoreUInt(0, 8).
			MustStoreRef(payload.CreatedAtRef.PriceRef).
			MustStoreRef(payload.CreatedAtRef.SignaturesRef).
			EndCell()

		return cell.BeginCell().
			MustStoreUInt(2, 8).
			MustStoreRef(priceRef).
			MustStoreRef(createdRef).
			EndCell()
	case OracleRefWithSettlementWithCreated:
		payload := payload.Value.(*OraclePayloadWithSettlementAndCreated)

		priceRef := cell.BeginCell().
			MustStoreUInt(1, 8).
			MustStoreRef(payload.PriceRef.PriceRef).
			MustStoreRef(payload.PriceRef.SignaturesRef).
			MustStoreRef(payload.PriceRef.SettlementRef).
			MustStoreRef(payload.PriceRef.SettlementSignaturesRef).
			EndCell()

		createdRef := cell.BeginCell().
			MustStoreUInt(0, 8).
			MustStoreRef(payload.CreatedAtRef.PriceRef).
			MustStoreRef(payload.CreatedAtRef.SignaturesRef).
			EndCell()

		return cell.BeginCell().
			MustStoreUInt(3, 8).
			MustStoreRef(priceRef).
			MustStoreRef(createdRef).
			EndCell()
	default:
		payload := payload.Value.(*OraclePayloadSimple)

		return cell.BeginCell().
			MustStoreUInt(0, 8).
			MustStoreRef(payload.PriceRef).
			MustStoreRef(payload.SignaturesRef).
			EndCell()
	}
}
