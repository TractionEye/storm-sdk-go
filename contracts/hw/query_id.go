package hw

type QueryId struct {
	Shift     uint16 `tlb:"## 10" json:"shift"`
	BitNumber uint16 `tlb:"## 13" json:"bit_number"`
}

func (i QueryId) Seqno() uint64 {
	return uint64(i.BitNumber + i.Shift*1023)
}

func FromSeqno(i uint64) *QueryId {
	shift := i / 1023
	bitNumber := i % 1023

	return &QueryId{
		Shift:     uint16(shift),
		BitNumber: uint16(bitNumber),
	}
}
