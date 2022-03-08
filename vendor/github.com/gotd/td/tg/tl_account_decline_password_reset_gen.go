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

// AccountDeclinePasswordResetRequest represents TL type `account.declinePasswordReset#4c9409f6`.
// Abort a pending 2FA password reset, see here for more info »¹
//
// Links:
//  1) https://core.telegram.org/api/srp#password-reset
//
// See https://core.telegram.org/method/account.declinePasswordReset for reference.
type AccountDeclinePasswordResetRequest struct {
}

// AccountDeclinePasswordResetRequestTypeID is TL type id of AccountDeclinePasswordResetRequest.
const AccountDeclinePasswordResetRequestTypeID = 0x4c9409f6

// Ensuring interfaces in compile-time for AccountDeclinePasswordResetRequest.
var (
	_ bin.Encoder     = &AccountDeclinePasswordResetRequest{}
	_ bin.Decoder     = &AccountDeclinePasswordResetRequest{}
	_ bin.BareEncoder = &AccountDeclinePasswordResetRequest{}
	_ bin.BareDecoder = &AccountDeclinePasswordResetRequest{}
)

func (d *AccountDeclinePasswordResetRequest) Zero() bool {
	if d == nil {
		return true
	}

	return true
}

// String implements fmt.Stringer.
func (d *AccountDeclinePasswordResetRequest) String() string {
	if d == nil {
		return "AccountDeclinePasswordResetRequest(nil)"
	}
	type Alias AccountDeclinePasswordResetRequest
	return fmt.Sprintf("AccountDeclinePasswordResetRequest%+v", Alias(*d))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*AccountDeclinePasswordResetRequest) TypeID() uint32 {
	return AccountDeclinePasswordResetRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*AccountDeclinePasswordResetRequest) TypeName() string {
	return "account.declinePasswordReset"
}

// TypeInfo returns info about TL type.
func (d *AccountDeclinePasswordResetRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "account.declinePasswordReset",
		ID:   AccountDeclinePasswordResetRequestTypeID,
	}
	if d == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{}
	return typ
}

// Encode implements bin.Encoder.
func (d *AccountDeclinePasswordResetRequest) Encode(b *bin.Buffer) error {
	if d == nil {
		return fmt.Errorf("can't encode account.declinePasswordReset#4c9409f6 as nil")
	}
	b.PutID(AccountDeclinePasswordResetRequestTypeID)
	return d.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (d *AccountDeclinePasswordResetRequest) EncodeBare(b *bin.Buffer) error {
	if d == nil {
		return fmt.Errorf("can't encode account.declinePasswordReset#4c9409f6 as nil")
	}
	return nil
}

// Decode implements bin.Decoder.
func (d *AccountDeclinePasswordResetRequest) Decode(b *bin.Buffer) error {
	if d == nil {
		return fmt.Errorf("can't decode account.declinePasswordReset#4c9409f6 to nil")
	}
	if err := b.ConsumeID(AccountDeclinePasswordResetRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode account.declinePasswordReset#4c9409f6: %w", err)
	}
	return d.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (d *AccountDeclinePasswordResetRequest) DecodeBare(b *bin.Buffer) error {
	if d == nil {
		return fmt.Errorf("can't decode account.declinePasswordReset#4c9409f6 to nil")
	}
	return nil
}

// AccountDeclinePasswordReset invokes method account.declinePasswordReset#4c9409f6 returning error if any.
// Abort a pending 2FA password reset, see here for more info »¹
//
// Links:
//  1) https://core.telegram.org/api/srp#password-reset
//
// Possible errors:
//  400 RESET_REQUEST_MISSING: No password reset is in progress.
//
// See https://core.telegram.org/method/account.declinePasswordReset for reference.
func (c *Client) AccountDeclinePasswordReset(ctx context.Context) (bool, error) {
	var result BoolBox

	request := &AccountDeclinePasswordResetRequest{}
	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return false, err
	}
	_, ok := result.Bool.(*BoolTrue)
	return ok, nil
}