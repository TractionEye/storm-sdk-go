package smartaccount

import (
	"context"
	"math/big"
	"time"

	"github.com/TractionEye/storm-sdk-go/contracts/smartaccount"
	"github.com/pkg/errors"
	"github.com/xssnick/tonutils-go/address"
	"github.com/xssnick/tonutils-go/tlb"
	"github.com/xssnick/tonutils-go/ton"
	"github.com/xssnick/tonutils-go/ton/wallet"
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

func (f *Factory) DeploySmartAccount(ctx context.Context, mnemonic []string) (*tlb.Transaction, error) {
	queryId := uint64(time.Now().Unix())
	v := &smartaccount.DeployOrdinarySA{
		QueryID:    queryId,
		PublicKeys: nil,
	}
	pks := smartaccount.UserPublicKeys{}
	pks.ExtractFromSeed(mnemonic)

	dict, err := pks.ToDictionary()
	if err != nil {
		return nil, err
	}

	v.PublicKeys = dict

	payload, err := tlb.ToCell(v)
	if err != nil {
		return nil, err
	}

	walletInstance, err := wallet.FromSeed(f.API, mnemonic, wallet.V4R2)
	if err != nil {
		return nil, err
	}
	msg := wallet.SimpleMessage(f.Addr, tlb.MustFromTON("0.2"), payload)
	tx, _, err := walletInstance.SendWaitTransaction(context.Background(), msg)

	return tx, err
}


