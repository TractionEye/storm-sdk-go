package smartaccount

import (
	"context"
	"github.com/pkg/errors"
	"github.com/xssnick/tonutils-go/address"
	"github.com/xssnick/tonutils-go/ton"
	"math/big"
)

type Factory struct {
	API  ton.APIClientWrapped
	Addr *address.Address
}

func NewFactory(api ton.APIClientWrapped, addr *address.Address) *Factory {
	return &Factory{API: api, Addr: addr}
}

func (f *Factory) GetSmartAccountAddress(ctx context.Context, owner *address.Address) (*address.Address, error) {
	ctx = f.API.Client().StickyContext(ctx)

	ext, err := f.API.CurrentMasterchainInfo(ctx)
	if err != nil {
		return nil, err
	}

	hash := big.NewInt(0).SetBytes(owner.Data())

	res, err := f.API.RunGetMethod(ctx, ext, f.Addr, "get_nft_address_by_index", hash)
	if err != nil {
		return nil, errors.Wrap(err, "get addr")
	}

	addrSlice, err := res.Slice(0)
	if err != nil {
		return nil, err
	}

	return addrSlice.LoadAddr()
}

func (f *Factory) GetSmartAccount(ctx context.Context, owner *address.Address) (*Client, error) {
	addr, err := f.GetSmartAccountAddress(ctx, owner)
	if err != nil {
		return nil, err
	}

	return NewClient(f.API, addr), nil
}
