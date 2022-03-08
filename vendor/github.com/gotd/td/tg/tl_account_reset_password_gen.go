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

// AccountResetPasswordRequest represents TL type `account.resetPassword#9308ce1b`.
// Initiate a 2FA password reset: can only be used if the user is already logged-in, see
// here for more info »¹
//
// Links:
//  1) https://core.telegram.org/api/srp#password-reset
//
// See https://core.telegram.org/method/account.resetPassword for reference.
type AccountResetPasswordRequest struct {
}

// AccountResetPasswordRequestTypeID is TL type id of AccountResetPasswordRequest.
const AccountResetPasswordRequestTypeID = 0x9308ce1b

// Ensuring interfaces in compile-time for AccountResetPasswordRequest.
var (
	_ bin.Encoder     = &AccountResetPasswordRequest{}
	_ bin.Decoder     = &AccountResetPasswordRequest{}
	_ bin.BareEncoder = &AccountResetPasswordRequest{}
	_ bin.BareDecoder = &AccountResetPasswordRequest{}
)

func (r *AccountResetPasswordRequest) Zero() bool {
	if r == nil {
		return true
	}

	return true
}

// String implements fmt.Stringer.
func (r *AccountResetPasswordRequest) String() string {
	if r == nil {
		return "AccountResetPasswordRequest(nil)"
	}
	type Alias AccountResetPasswordRequest
	return fmt.Sprintf("AccountResetPasswordRequest%+v", Alias(*r))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*AccountResetPasswordRequest) TypeID() uint32 {
	return AccountResetPasswordRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*AccountResetPasswordRequest) TypeName() string {
	return "account.resetPassword"
}

// TypeInfo returns info about TL type.
func (r *AccountResetPasswordRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "account.resetPassword",
		ID:   AccountResetPasswordRequestTypeID,
	}
	if r == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{}
	return typ
}

// Encode implements bin.Encoder.
func (r *AccountResetPasswordRequest) Encode(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't encode account.resetPassword#9308ce1b as nil")
	}
	b.PutID(AccountResetPasswordRequestTypeID)
	return r.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (r *AccountResetPasswordRequest) EncodeBare(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't encode account.resetPassword#9308ce1b as nil")
	}
	return nil
}

// Decode implements bin.Decoder.
func (r *AccountResetPasswordRequest) Decode(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't decode account.resetPassword#9308ce1b to nil")
	}
	if err := b.ConsumeID(AccountResetPasswordRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode account.resetPassword#9308ce1b: %w", err)
	}
	return r.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (r *AccountResetPasswordRequest) DecodeBare(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't decode account.resetPassword#9308ce1b to nil")
	}
	return nil
}

// AccountResetPassword invokes method account.resetPassword#9308ce1b returning error if any.
// Initiate a 2FA password reset: can only be used if the user is already logged-in, see
// here for more info »¹
//
// Links:
//  1) https://core.telegram.org/api/srp#password-reset
//
// See https://core.telegram.org/method/account.resetPassword for reference.
func (c *Client) AccountResetPassword(ctx context.Context) (AccountResetPasswordResultClass, error) {
	var result AccountResetPasswordResultBox

	request := &AccountResetPasswordRequest{}
	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return nil, err
	}
	return result.ResetPasswordResult, nil
}
