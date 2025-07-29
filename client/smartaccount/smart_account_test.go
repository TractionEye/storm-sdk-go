package smartaccount_test

import (
	"context"
	"github.com/storm-trade/sdk-go/client/smartaccount"
	"github.com/stretchr/testify/require"
	"github.com/xssnick/tonutils-go/address"
	"testing"
)

func TestClient_GetStorageData(t *testing.T) {
	client := smartaccount.NewClient(API, address.MustParseAddr("kQC82GIBp5xKv8RPXg3aDUH2_6OYDPsOoW0fqlopKKqX1Oo5"))
	data, err := client.GetStorageData(context.TODO())

	require.Nil(t, err)
	require.Equal(t, data.Factory.StringRaw(), "0:e0a015de9ff8c25a62991bb07f0cfcb9e791b4b1a898bc9334aceab5e67dc74d")
}
