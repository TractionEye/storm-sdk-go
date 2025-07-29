package smartaccount

import (
	"context"
	"github.com/pkg/errors"
	"github.com/storm-trade/sdk-go/contracts/smartaccount"
	"github.com/xssnick/tonutils-go/address"
	"github.com/xssnick/tonutils-go/tlb"
	"github.com/xssnick/tonutils-go/ton"
	"github.com/xssnick/tonutils-go/ton/wallet"
	"github.com/xssnick/tonutils-go/tvm/cell"
	"time"
)

var (
	MinGas = tlb.MustFromTON("0.05")
)

type Client struct {
	addr *address.Address
	API  ton.APIClientWrapped
}

func NewClient(api ton.APIClientWrapped, addr *address.Address) *Client {
	return &Client{addr: addr, API: api}
}

func (c *Client) GetAddress() *address.Address {
	return c.addr
}

func (c *Client) GetStorageData(ctx context.Context) (*smartaccount.AccountData, error) {
	ctx = c.API.Client().StickyContext(ctx)

	ext, err := c.API.CurrentMasterchainInfo(ctx)
	if err != nil {
		return nil, err
	}

	account, err := c.API.GetAccount(ctx, ext, c.addr)
	if err != nil {
		return nil, err
	}

	if !account.IsActive {
		return nil, errors.New("account not active")
	}

	data := new(smartaccount.AccountData)
	if err := tlb.LoadFromCell(data, account.State.StateInit.Data.BeginParse()); err != nil {
		return nil, err
	}

	return data, nil
}

func (c *Client) DepositNative(from *wallet.Wallet, to *address.Address, amount *tlb.Coins) (*tlb.Transaction, error) {
	payload, err := c.BuildDepositNativePayload(from.WalletAddress(), amount)
	if err != nil {
		return nil, err
	}

	tonAmount, err := amount.Add(&MinGas)
	if err != nil {
		return nil, err
	}

	msg := wallet.SimpleMessage(to, *tonAmount, payload)
	tx, _, err := from.SendWaitTransaction(context.Background(), msg)

	return tx, err
}

func (c *Client) BuildDepositNativePayload(owner *address.Address, amount *tlb.Coins) (*cell.Cell, error) {
	queryId := uint64(time.Now().Unix())

	v := &smartaccount.DepositNative{
		QueryID:         queryId,
		Amount:          amount,
		ReceiverAddress: owner,
		Init:            false,
		KeyInit:         nil,
	}

	return tlb.ToCell(v)
}
