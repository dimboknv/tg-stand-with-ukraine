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

// HelpGetUserInfoRequest represents TL type `help.getUserInfo#38a08d3`.
// Internal use
//
// See https://core.telegram.org/method/help.getUserInfo for reference.
type HelpGetUserInfoRequest struct {
	// User ID
	UserID InputUserClass
}

// HelpGetUserInfoRequestTypeID is TL type id of HelpGetUserInfoRequest.
const HelpGetUserInfoRequestTypeID = 0x38a08d3

// Ensuring interfaces in compile-time for HelpGetUserInfoRequest.
var (
	_ bin.Encoder     = &HelpGetUserInfoRequest{}
	_ bin.Decoder     = &HelpGetUserInfoRequest{}
	_ bin.BareEncoder = &HelpGetUserInfoRequest{}
	_ bin.BareDecoder = &HelpGetUserInfoRequest{}
)

func (g *HelpGetUserInfoRequest) Zero() bool {
	if g == nil {
		return true
	}
	if !(g.UserID == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (g *HelpGetUserInfoRequest) String() string {
	if g == nil {
		return "HelpGetUserInfoRequest(nil)"
	}
	type Alias HelpGetUserInfoRequest
	return fmt.Sprintf("HelpGetUserInfoRequest%+v", Alias(*g))
}

// FillFrom fills HelpGetUserInfoRequest from given interface.
func (g *HelpGetUserInfoRequest) FillFrom(from interface {
	GetUserID() (value InputUserClass)
}) {
	g.UserID = from.GetUserID()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*HelpGetUserInfoRequest) TypeID() uint32 {
	return HelpGetUserInfoRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*HelpGetUserInfoRequest) TypeName() string {
	return "help.getUserInfo"
}

// TypeInfo returns info about TL type.
func (g *HelpGetUserInfoRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "help.getUserInfo",
		ID:   HelpGetUserInfoRequestTypeID,
	}
	if g == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "UserID",
			SchemaName: "user_id",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (g *HelpGetUserInfoRequest) Encode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode help.getUserInfo#38a08d3 as nil")
	}
	b.PutID(HelpGetUserInfoRequestTypeID)
	return g.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (g *HelpGetUserInfoRequest) EncodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode help.getUserInfo#38a08d3 as nil")
	}
	if g.UserID == nil {
		return fmt.Errorf("unable to encode help.getUserInfo#38a08d3: field user_id is nil")
	}
	if err := g.UserID.Encode(b); err != nil {
		return fmt.Errorf("unable to encode help.getUserInfo#38a08d3: field user_id: %w", err)
	}
	return nil
}

// Decode implements bin.Decoder.
func (g *HelpGetUserInfoRequest) Decode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode help.getUserInfo#38a08d3 to nil")
	}
	if err := b.ConsumeID(HelpGetUserInfoRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode help.getUserInfo#38a08d3: %w", err)
	}
	return g.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (g *HelpGetUserInfoRequest) DecodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode help.getUserInfo#38a08d3 to nil")
	}
	{
		value, err := DecodeInputUser(b)
		if err != nil {
			return fmt.Errorf("unable to decode help.getUserInfo#38a08d3: field user_id: %w", err)
		}
		g.UserID = value
	}
	return nil
}

// GetUserID returns value of UserID field.
func (g *HelpGetUserInfoRequest) GetUserID() (value InputUserClass) {
	if g == nil {
		return
	}
	return g.UserID
}

// HelpGetUserInfo invokes method help.getUserInfo#38a08d3 returning error if any.
// Internal use
//
// Possible errors:
//  403 USER_INVALID: Invalid user provided.
//
// See https://core.telegram.org/method/help.getUserInfo for reference.
func (c *Client) HelpGetUserInfo(ctx context.Context, userid InputUserClass) (HelpUserInfoClass, error) {
	var result HelpUserInfoBox

	request := &HelpGetUserInfoRequest{
		UserID: userid,
	}
	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return nil, err
	}
	return result.UserInfo, nil
}