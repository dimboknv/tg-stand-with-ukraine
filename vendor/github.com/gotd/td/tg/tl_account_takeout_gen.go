// Code generated by gotdgen, DO NOT EDIT.

package tg

import (
	"context"
	"errors"
	"fmt"
	"sort"
	"strings"

	"go.uber.org/multierr"

	"github.com/gotd/td/bin"
	"github.com/gotd/td/tdjson"
	"github.com/gotd/td/tdp"
	"github.com/gotd/td/tgerr"
)

// No-op definition for keeping imports.
var (
	_ = bin.Buffer{}
	_ = context.Background()
	_ = fmt.Stringer(nil)
	_ = strings.Builder{}
	_ = errors.Is
	_ = multierr.AppendInto
	_ = sort.Ints
	_ = tdp.Format
	_ = tgerr.Error{}
	_ = tdjson.Encoder{}
)

// AccountTakeout represents TL type `account.takeout#4dba4501`.
// Takout info
//
// See https://core.telegram.org/constructor/account.takeout for reference.
type AccountTakeout struct {
	// Takeout ID
	ID int64
}

// AccountTakeoutTypeID is TL type id of AccountTakeout.
const AccountTakeoutTypeID = 0x4dba4501

// Ensuring interfaces in compile-time for AccountTakeout.
var (
	_ bin.Encoder     = &AccountTakeout{}
	_ bin.Decoder     = &AccountTakeout{}
	_ bin.BareEncoder = &AccountTakeout{}
	_ bin.BareDecoder = &AccountTakeout{}
)

func (t *AccountTakeout) Zero() bool {
	if t == nil {
		return true
	}
	if !(t.ID == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (t *AccountTakeout) String() string {
	if t == nil {
		return "AccountTakeout(nil)"
	}
	type Alias AccountTakeout
	return fmt.Sprintf("AccountTakeout%+v", Alias(*t))
}

// FillFrom fills AccountTakeout from given interface.
func (t *AccountTakeout) FillFrom(from interface {
	GetID() (value int64)
}) {
	t.ID = from.GetID()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*AccountTakeout) TypeID() uint32 {
	return AccountTakeoutTypeID
}

// TypeName returns name of type in TL schema.
func (*AccountTakeout) TypeName() string {
	return "account.takeout"
}

// TypeInfo returns info about TL type.
func (t *AccountTakeout) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "account.takeout",
		ID:   AccountTakeoutTypeID,
	}
	if t == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "ID",
			SchemaName: "id",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (t *AccountTakeout) Encode(b *bin.Buffer) error {
	if t == nil {
		return fmt.Errorf("can't encode account.takeout#4dba4501 as nil")
	}
	b.PutID(AccountTakeoutTypeID)
	return t.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (t *AccountTakeout) EncodeBare(b *bin.Buffer) error {
	if t == nil {
		return fmt.Errorf("can't encode account.takeout#4dba4501 as nil")
	}
	b.PutLong(t.ID)
	return nil
}

// Decode implements bin.Decoder.
func (t *AccountTakeout) Decode(b *bin.Buffer) error {
	if t == nil {
		return fmt.Errorf("can't decode account.takeout#4dba4501 to nil")
	}
	if err := b.ConsumeID(AccountTakeoutTypeID); err != nil {
		return fmt.Errorf("unable to decode account.takeout#4dba4501: %w", err)
	}
	return t.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (t *AccountTakeout) DecodeBare(b *bin.Buffer) error {
	if t == nil {
		return fmt.Errorf("can't decode account.takeout#4dba4501 to nil")
	}
	{
		value, err := b.Long()
		if err != nil {
			return fmt.Errorf("unable to decode account.takeout#4dba4501: field id: %w", err)
		}
		t.ID = value
	}
	return nil
}

// GetID returns value of ID field.
func (t *AccountTakeout) GetID() (value int64) {
	if t == nil {
		return
	}
	return t.ID
}