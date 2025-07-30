package smartaccount_test

import (
	"context"
	"github.com/pkg/errors"
	"github.com/TractionEye/sdk-go/client/smartaccount"
	"github.com/stretchr/testify/require"
	"github.com/xssnick/tonutils-go/address"
	"github.com/xssnick/tonutils-go/liteclient"
	"github.com/xssnick/tonutils-go/ton"
	"testing"
)

func Client(ctx context.Context, configUrl string) ton.APIClientWrapped {
	client := liteclient.NewConnectionPool()
	err := client.AddConnectionsFromConfigUrl(ctx, configUrl)
	if err != nil {
		panic(errors.Wrap(err, "ls add config"))
	}

	api := ton.NewAPIClient(client, ton.ProofCheckPolicyUnsafe).WithRetry(10)

	return api
}

var API = Client(context.Background(), "https://ton-blockchain.github.io/testnet-global.config.json")

func TestFactory_GetSmartAccountAddress(t *testing.T) {
	factoryClient := smartaccount.NewFactory(API, address.MustParseAddr("kQDgoBXen_jCWmKZG7B_DPy555G0saiYvJM0rOq15n3HTWIN"))
	addr, err := factoryClient.GetSmartAccountAddress(context.Background(), address.MustParseAddr("0QCogh0uaL1-p7nBx4eB8yIrH2Pf5qQ5ZUUeS4on5VzAkeeI"))

	require.Nil(t, err)
	require.Equal(t, addr.StringRaw(), "0:bcd86201a79c4abfc44f5e0dda0d41f6ffa3980cfb0ea16d1faa5a2928aa97d4")
}
