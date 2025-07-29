package schemas

import "github.com/xssnick/tonutils-go/tlb"

type AddMarginOrder struct {
	_       tlb.Magic      `tlb:"$0100"`
	Payload LimitOrderData `tlb:"." json:"payload"`
}

type RemoveMarginOrder struct {
	_       tlb.Magic     `tlb:"$0101"`
	Payload StopOrderData `tlb:"." json:"payload"`
}

func (s AddMarginOrder) AsStopOrder() *StopOrderData {
	return nil
}

func (s AddMarginOrder) AsLimitOrder() *LimitOrderData {
	return &s.Payload
}

func (s RemoveMarginOrder) AsStopOrder() *StopOrderData {
	return &s.Payload
}

func (s RemoveMarginOrder) AsLimitOrder() *LimitOrderData {
	return nil
}
