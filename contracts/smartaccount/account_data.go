package smartaccount

import (
	"crypto/ed25519"
	"crypto/hmac"
	"crypto/sha512"
	"encoding/json"
	"strings"

	"github.com/pkg/errors"
	"github.com/xssnick/tonutils-go/address"
	"github.com/xssnick/tonutils-go/tvm/cell"
	"golang.org/x/crypto/pbkdf2"
)

type AccountData struct {
	Type      uint8            `tlb:"## 8" json:"type"`
	Factory   *address.Address `tlb:"addr" json:"factory"`
	Owner     *address.Address `tlb:"addr" json:"owner"`
	Balances  *BalanceList     `tlb:"." json:"balances"`
	Version   uint8            `tlb:"## 8" json:"version"`
	Keys      *Keys            `tlb:"^" json:"keys"`
	Positions *UserPositions   `tlb:"." json:"positions"`
	Highload  *Highload        `tlb:"^" json:"highload"`
}

type Keys struct {
	Hot            []byte          `tlb:"bits 256" json:"hot"`
	Cold           []byte          `tlb:"bits 256" json:"cold"`
	UserPublicKeys *UserPublicKeys `tlb:"." json:"user_public_keys"`
	KeysCount      int64           `tlb:"## 8" json:"keys_count"`
}

type Highload struct {
	OldQueries    *cell.Dictionary `tlb:"dict 13" json:"old_queries"`
	Queries       *cell.Dictionary `tlb:"dict 13" json:"queries"`
	LastCleanTime uint64           `tlb:"## 64" json:"last_clean_time"`
	Timeout       uint64           `tlb:"## 24" json:"timeout"`
}

type UserPublicKeys struct {
	Values []PublicKey `json:"values"`
}

type PublicKey ed25519.PublicKey

func (m *UserPublicKeys) ExtractFromSeed(mnemonic []string) error {
	mac := hmac.New(sha512.New, []byte(strings.Join(mnemonic, " ")))
	hash := mac.Sum(nil)
	k := pbkdf2.Key(hash, []byte("TON default seed"), 100000, 32, sha512.New)

	privateKey := ed25519.NewKeyFromSeed(k)
	publicKey := privateKey.Public().(ed25519.PublicKey)
	m.Values = make([]PublicKey, 0, 10)
	m.Values = append(m.Values, PublicKey(publicKey))
	return nil
}

func (m *UserPublicKeys) LoadFromCell(slice *cell.Slice) error {
	deviceDict, err := slice.LoadDict(256)
	if err != nil {
		return errors.Wrap(err, "load devices dict")
	}

	if deviceDict.IsEmpty() {
		m.Values = make([]PublicKey, 0)
		return nil
	}

	kv, err := deviceDict.LoadAll()
	if err != nil {
		return errors.Wrap(err, "load all dict devices")
	}

	pks := make([]PublicKey, 0)
	for _, v := range kv {
		pk, err := v.Key.LoadBigUInt(256)
		if err != nil {
			return errors.Wrap(err, "load all dict devices")
		}

		pks = append(pks, pk.Bytes())
	}

	m.Values = pks

	return nil
}

func (m *UserPublicKeys) ToCell() (*cell.Cell, error) {
	dict, err := m.ToDictionary()
	if err != nil {
		return nil, err
	}

	return cell.BeginCell().MustStoreDict(dict).EndCell(), nil
}

func (m *UserPublicKeys) ToDictionary() (*cell.Dictionary, error) {
	dict := cell.NewDict(256)

	for _, item := range m.Values {
		key := cell.BeginCell().MustStoreUInt(0, 256).EndCell()
		err := dict.Set(key, cell.BeginCell().MustStoreSlice(item, 256).EndCell())
		if err != nil {
			return nil, err
		}
	}

	return dict, nil
}

type BalanceList struct {
	Balances map[*address.Address]uint64 `json:"balances"`
}

type BalanceEntry struct {
	Addr  address.Address `json:"address"`
	Count uint64          `json:"amount"`
}

func (b *BalanceList) MarshalJSON() ([]byte, error) {
	entries := make([]BalanceEntry, 0, len(b.Balances))
	for addrPtr, count := range b.Balances {
		entries = append(entries, BalanceEntry{
			Addr:  *addrPtr,
			Count: count,
		})
	}
	return json.Marshal(entries)
}

func (b *BalanceList) LoadFromCell(slice *cell.Slice) error {
	balanceDict, err := slice.LoadDict(267)
	if err != nil {
		return errors.Wrap(err, "load balances dict")
	}

	if balanceDict.IsEmpty() {
		b.Balances = make(map[*address.Address]uint64)
		return nil
	}

	kv, err := balanceDict.LoadAll()
	if err != nil {
		return errors.Wrap(err, "load all dict balances")
	}

	balances := make(map[*address.Address]uint64)

	for _, v := range kv {
		addr := v.Key.MustLoadAddr()
		balance := v.Value.MustLoadCoins()
		balances[addr] = balance
	}

	b.Balances = balances

	return nil
}

func (b *BalanceList) ToDictionary() (*cell.Dictionary, error) {
	devicesDict := cell.NewDict(267)
	for k, v := range b.Balances {
		key := cell.BeginCell().MustStoreAddr(k).EndCell()
		value := cell.BeginCell().MustStoreCoins(v).EndCell()

		err := devicesDict.Set(key, value)
		if err != nil {
			return nil, err
		}
	}

	return devicesDict, nil
}

func (b *BalanceList) ToCell() (*cell.Cell, error) {
	dict, err := b.ToDictionary()
	if err != nil {
		return nil, errors.Wrap(err, "to dictionary")
	}

	return cell.BeginCell().MustStoreDict(dict).EndCell(), nil
}
