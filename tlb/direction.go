package tlb

type Direction string

const (
	DirectionLong  Direction = "long"
	DirectionShort Direction = "short"
)

const (
	ContractDirectionLong  int = 0
	ContractDirectionShort int = 1
)

func (dir Direction) GetInt() int {
	switch {
	case dir == DirectionLong:
		return ContractDirectionLong
	case dir == DirectionShort:
		return ContractDirectionShort
	default:
		panic("invalid direction")
	}
}

func DirectionFromInt(dir int) Direction {
	switch dir {
	case ContractDirectionLong:
		return DirectionLong
	case ContractDirectionShort:
		return DirectionShort
	default:
		panic("invalid direction")
	}
}
