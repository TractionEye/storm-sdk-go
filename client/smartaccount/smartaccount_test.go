package smartaccount_test

import (
	"context"
	"github.com/TractionEye/sdk-go/client/smartaccount"
	"github.com/stretchr/testify/require"
	"github.com/xssnick/tonutils-go/address"
	"github.com/xssnick/tonutils-go/tlb"
	"github.com/xssnick/tonutils-go/ton/wallet"
	"strings"
	"testing"
)

func TestClient_GetStorageData(t *testing.T) {
	client := smartaccount.NewClient(API, address.MustParseAddr("kQC82GIBp5xKv8RPXg3aDUH2_6OYDPsOoW0fqlopKKqX1Oo5"))
	data, err := client.GetStorageData(context.TODO())

	require.Nil(t, err)
	require.Equal(t, data.Factory.StringRaw(), "0:e0a015de9ff8c25a62991bb07f0cfcb9e791b4b1a898bc9334aceab5e67dc74d")
}

func TestClient_DepositNative(t *testing.T) {
	client := smartaccount.NewClient(API, address.MustParseAddr("kQC82GIBp5xKv8RPXg3aDUH2_6OYDPsOoW0fqlopKKqX1Oo5"))

	vaultAddress := address.MustParseAddr("kQCSnKC--Igca13vrtnVI-I0FYJxAvcJJNZDxZfRAVGeWJlq")
	amount := tlb.MustFromTON("1")

	sender, err := wallet.FromSeed(API, strings.Split("<seed here>", " "), wallet.V4R2)
	require.Nil(t, err)

	_, err = client.DepositNative(sender, sender.WalletAddress(), vaultAddress, &amount, false)
	require.Nil(t, err)
}

func TestClient_DepositJetton(t *testing.T) {
	client := smartaccount.NewClient(API, address.MustParseAddr("kQC82GIBp5xKv8RPXg3aDUH2_6OYDPsOoW0fqlopKKqX1Oo5"))

	vaultAddress := address.MustParseAddr("kQAr2Q423aDQrGb1ouVnH0PfMZu35JUhIoHbc349L9kYknJM")
	jettonAddress := address.MustParseAddr("kQAFM9tbxGoebAof8QKXKl3x7ez1cw-HNRET5wt58UUy-wR0")
	amount := tlb.MustFromTON("1")

	sender, err := wallet.FromSeed(API, strings.Split("<seed here>", " "), wallet.V4R2)
	require.Nil(t, err)

	_, err = client.DepositJetton(sender, sender.WalletAddress(), vaultAddress, jettonAddress, &amount, false)
	require.Nil(t, err)
}
