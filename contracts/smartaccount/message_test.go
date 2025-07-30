package smartaccount_test

import (
	"github.com/TractionEye/sdk-go/contracts/smartaccount"
	"github.com/stretchr/testify/require"
	"github.com/xssnick/tonutils-go/tlb"
	"testing"
)

func Test_DepositNativeMessage(t *testing.T) {
	Amount, _ := tlb.FromNano(GenerateRandomInt(1_000_000), 9)

	msg := &smartaccount.DepositNative{
		QueryID:         1,
		Amount:          &Amount,
		ReceiverAddress: GenerateRandomAddress(),
		Init:            false,
		KeyInit:         nil,
	}

	msgCell, err := tlb.ToCell(msg)
	require.Nil(t, err)

	msgLoaded := new(smartaccount.DepositNative)
	err = tlb.LoadFromCell(msgLoaded, msgCell.BeginParse())

	require.Nil(t, err)

	require.Equal(t, msgLoaded.QueryID, msg.QueryID)
}
