package smartaccount_test

import (
	"crypto/ed25519"
	"crypto/rand"
	"encoding/hex"
	"github.com/cristalhq/base64"
	"github.com/TractionEye/storm-sdk-go/contracts/smartaccount"
	schemas "github.com/TractionEye/storm-sdk-go/tlb"
	"github.com/stretchr/testify/require"
	"github.com/xssnick/tonutils-go/address"
	"github.com/xssnick/tonutils-go/tlb"
	"github.com/xssnick/tonutils-go/tvm/cell"
	"math/big"
	"testing"
	"time"
)

func GenerateRandomAddress() *address.Address {
	pub, _, _ := ed25519.GenerateKey(rand.Reader)
	addr := address.NewAddress(0, 0, cell.BeginCell().MustStoreSlice(pub, 32).EndCell().Hash())

	return addr
}

func GenerateRandomKey() smartaccount.PublicKey {
	pub, _, _ := ed25519.GenerateKey(rand.Reader)

	return smartaccount.PublicKey(pub)
}

func GenerateRandomInt(max int64) *big.Int {
	n, _ := rand.Int(rand.Reader, big.NewInt(max))

	return n
}

func Base64ToHex(str string) string {
	v, err := base64.StdEncoding.DecodeString(str)
	if err != nil {
		panic(err)
	}

	return hex.EncodeToString(v)
}

func GenerateRandomPosition() *schemas.PositionState {
	Notional, _ := tlb.FromNano(GenerateRandomInt(1_000_000), 9)
	Margin, _ := tlb.FromNano(GenerateRandomInt(1_000_000), 9)

	return &schemas.PositionState{
		Size:                         GenerateRandomInt(1_000_000_000),
		Direction:                    0,
		Margin:                       &Margin,
		OpenNotional:                 &Notional,
		LastUpdatedCumulativePremium: GenerateRandomInt(1_000_000).Uint64(),
		Fee:                          GenerateRandomInt(1_000_000).Uint64(),
		Discount:                     GenerateRandomInt(1_000_000).Uint64(),
		Rebate:                       GenerateRandomInt(1_000_000).Uint64(),
		LastUpdatedTimestamp:         uint64(time.Now().Unix()),
	}
}

func Test_LoadSmartAccountState(t *testing.T) {
	t.Run("should serialize account", func(t *testing.T) {
		accountData := smartaccount.AccountData{
			Type:    0,
			Factory: GenerateRandomAddress(),
			Owner:   GenerateRandomAddress(),
			Balances: &smartaccount.BalanceList{Balances: map[*address.Address]uint64{
				GenerateRandomAddress(): 228,
			}},
			Version: 0,
			Keys: &smartaccount.Keys{
				Hot:  GenerateRandomKey(),
				Cold: GenerateRandomKey(),
				UserPublicKeys: &smartaccount.UserPublicKeys{
					Values: []smartaccount.PublicKey{
						GenerateRandomKey(),
					},
				},
				KeysCount: 2,
			},
			Positions: &smartaccount.UserPositions{
				Values: map[smartaccount.UserPositionKey]*smartaccount.UserPosition{
					smartaccount.UserPositionKey{Market: GenerateRandomAddress(), Direction: schemas.DirectionLong}: {
						Locked: false,
						Data:   GenerateRandomPosition(),
					},
				},
			},
			Highload: &smartaccount.Highload{
				OldQueries:    cell.NewDict(13),
				Queries:       cell.NewDict(13),
				LastCleanTime: uint64(time.Now().Unix()),
				Timeout:       uint64(time.Now().Unix()) / 60 / 60 / 24,
			},
		}

		c, err := tlb.ToCell(accountData)
		require.Nil(t, err)

		loaded := new(smartaccount.AccountData)
		err = tlb.LoadFromCell(loaded, c.BeginParse())
		require.Nil(t, err)

		require.Equal(t, loaded.Factory.String(), accountData.Factory.String())
		require.Equal(t, loaded.Owner.String(), accountData.Owner.String())
		require.Equal(t, loaded.Version, accountData.Version)

		// balances
		values := map[string]uint64{}
		for k, v := range accountData.Balances.Balances {
			values[k.StringRaw()] = v
		}

		for k, v := range loaded.Balances.Balances {
			require.Equal(t, values[k.StringRaw()], v)
		}

		// keys
		require.Equal(t, accountData.Keys.KeysCount, loaded.Keys.KeysCount)
		require.Equal(t, hex.EncodeToString(accountData.Keys.Cold), hex.EncodeToString(loaded.Keys.Cold))
		require.Equal(t, hex.EncodeToString(accountData.Keys.Hot), hex.EncodeToString(loaded.Keys.Hot))

		pkValues := map[string]any{}
		for _, v := range accountData.Keys.UserPublicKeys.Values {
			pkValues[hex.EncodeToString(v)] = true
		}

		for _, v := range loaded.Keys.UserPublicKeys.Values {
			require.Equal(t, pkValues[hex.EncodeToString(v)], true)
		}

		// highload
		require.Equal(t, accountData.Highload.LastCleanTime, loaded.Highload.LastCleanTime)
		require.Equal(t, accountData.Highload.Timeout, loaded.Highload.Timeout)

		// positions
		posValues := map[string]*smartaccount.UserPosition{}
		for _, v := range accountData.Keys.UserPublicKeys.Values {
			pkValues[hex.EncodeToString(v)] = true
		}

		for k, v := range accountData.Positions.Values {
			posValues[k.String()] = v
		}

		for k, v := range loaded.Positions.Values {
			before, ok := posValues[k.String()]

			require.Equal(t, ok, true)

			require.Equal(t, v.Locked, before.Locked)

			require.Equal(t, v.Data.Direction, before.Data.Direction)
			require.Equal(t, v.Data.Fee, before.Data.Fee)
			require.Equal(t, v.Data.OpenNotional, before.Data.OpenNotional)
			require.Equal(t, v.Data.Size, before.Data.Size)
			require.Equal(t, v.Data.Margin, before.Data.Margin)
			require.Equal(t, v.Data.Discount, before.Data.Discount)
			require.Equal(t, v.Data.Rebate, before.Data.Rebate)
			require.Equal(t, v.Data.LastUpdatedTimestamp, before.Data.LastUpdatedTimestamp)
		}
	})

	t.Run("should deserialize account from tonviewer", func(t *testing.T) {
		accountData := new(smartaccount.AccountData)
		dataHex := Base64ToHex("te6cckECCQEAAUUAA4oAgBwUArvT/xhLTFMjdg/hn5c88jaWNRMXkmaVnVa8z7jpsAKiCHS5ovX6nucHHh4HzIisfY9/mpDllRR5LiiflXMCRgABAgMCBYFwAgQFAYMzULqTIHvNl2Xsv+RQorI76Nhr7m1JARYLrdU08ZkHzWFr5hR3F6kibGbK9mzWvg5Q5ZIJVuNEVhy/t7NC9CWRgUAGABcAAAAAGhdEGTtTgCAAUb/V7IcbbtBoVjN60XKzj6HvmM3b8kqQkUDtub8el+yMSTgL26VbzAAEAE2/yU5QX3xEDjWu99ds6pHxGgrBOIF7hJJrIeLL6ICozywoo+mrgAQCASAHCABDv/Kly62EOSqt5Vjj88Md/X+NUCre1apasBisub0/mtbxwABDv+XQ1rT31azmhyF+hPWPi3g9bGojEaGzwhX0rtfWxS5WwI8W324=")

		dataBoc, err := hex.DecodeString(dataHex)
		require.Nil(t, err)

		dataCell, err := cell.FromBOC(dataBoc)

		err = tlb.LoadFromCell(accountData, dataCell.BeginParse())
		require.Nil(t, err)

		cdata, err := tlb.ToCell(accountData)
		require.Nil(t, err)

		require.Equal(t, hex.EncodeToString(dataCell.Hash()), hex.EncodeToString(cdata.Hash()))
	})
}
