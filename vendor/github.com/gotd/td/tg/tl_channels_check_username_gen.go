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

// ChannelsCheckUsernameRequest represents TL type `channels.checkUsername#10e6bd2c`.
// Check if a username is free and can be assigned to a channel/supergroup
//
// See https://core.telegram.org/method/channels.checkUsername for reference.
type ChannelsCheckUsernameRequest struct {
	// The channel/supergroup¹ that will assigned the specified username
	//
	// Links:
	//  1) https://core.telegram.org/api/channel
	Channel InputChannelClass
	// The username to check
	Username string
}

// ChannelsCheckUsernameRequestTypeID is TL type id of ChannelsCheckUsernameRequest.
const ChannelsCheckUsernameRequestTypeID = 0x10e6bd2c

// Ensuring interfaces in compile-time for ChannelsCheckUsernameRequest.
var (
	_ bin.Encoder     = &ChannelsCheckUsernameRequest{}
	_ bin.Decoder     = &ChannelsCheckUsernameRequest{}
	_ bin.BareEncoder = &ChannelsCheckUsernameRequest{}
	_ bin.BareDecoder = &ChannelsCheckUsernameRequest{}
)

func (c *ChannelsCheckUsernameRequest) Zero() bool {
	if c == nil {
		return true
	}
	if !(c.Channel == nil) {
		return false
	}
	if !(c.Username == "") {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (c *ChannelsCheckUsernameRequest) String() string {
	if c == nil {
		return "ChannelsCheckUsernameRequest(nil)"
	}
	type Alias ChannelsCheckUsernameRequest
	return fmt.Sprintf("ChannelsCheckUsernameRequest%+v", Alias(*c))
}

// FillFrom fills ChannelsCheckUsernameRequest from given interface.
func (c *ChannelsCheckUsernameRequest) FillFrom(from interface {
	GetChannel() (value InputChannelClass)
	GetUsername() (value string)
}) {
	c.Channel = from.GetChannel()
	c.Username = from.GetUsername()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*ChannelsCheckUsernameRequest) TypeID() uint32 {
	return ChannelsCheckUsernameRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*ChannelsCheckUsernameRequest) TypeName() string {
	return "channels.checkUsername"
}

// TypeInfo returns info about TL type.
func (c *ChannelsCheckUsernameRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "channels.checkUsername",
		ID:   ChannelsCheckUsernameRequestTypeID,
	}
	if c == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Channel",
			SchemaName: "channel",
		},
		{
			Name:       "Username",
			SchemaName: "username",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (c *ChannelsCheckUsernameRequest) Encode(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't encode channels.checkUsername#10e6bd2c as nil")
	}
	b.PutID(ChannelsCheckUsernameRequestTypeID)
	return c.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (c *ChannelsCheckUsernameRequest) EncodeBare(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't encode channels.checkUsername#10e6bd2c as nil")
	}
	if c.Channel == nil {
		return fmt.Errorf("unable to encode channels.checkUsername#10e6bd2c: field channel is nil")
	}
	if err := c.Channel.Encode(b); err != nil {
		return fmt.Errorf("unable to encode channels.checkUsername#10e6bd2c: field channel: %w", err)
	}
	b.PutString(c.Username)
	return nil
}

// Decode implements bin.Decoder.
func (c *ChannelsCheckUsernameRequest) Decode(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't decode channels.checkUsername#10e6bd2c to nil")
	}
	if err := b.ConsumeID(ChannelsCheckUsernameRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode channels.checkUsername#10e6bd2c: %w", err)
	}
	return c.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (c *ChannelsCheckUsernameRequest) DecodeBare(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't decode channels.checkUsername#10e6bd2c to nil")
	}
	{
		value, err := DecodeInputChannel(b)
		if err != nil {
			return fmt.Errorf("unable to decode channels.checkUsername#10e6bd2c: field channel: %w", err)
		}
		c.Channel = value
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode channels.checkUsername#10e6bd2c: field username: %w", err)
		}
		c.Username = value
	}
	return nil
}

// GetChannel returns value of Channel field.
func (c *ChannelsCheckUsernameRequest) GetChannel() (value InputChannelClass) {
	if c == nil {
		return
	}
	return c.Channel
}

// GetUsername returns value of Username field.
func (c *ChannelsCheckUsernameRequest) GetUsername() (value string) {
	if c == nil {
		return
	}
	return c.Username
}

// GetChannelAsNotEmpty returns mapped value of Channel field.
func (c *ChannelsCheckUsernameRequest) GetChannelAsNotEmpty() (NotEmptyInputChannel, bool) {
	return c.Channel.AsNotEmpty()
}

// ChannelsCheckUsername invokes method channels.checkUsername#10e6bd2c returning error if any.
// Check if a username is free and can be assigned to a channel/supergroup
//
// Possible errors:
//  400 CHANNELS_ADMIN_PUBLIC_TOO_MUCH: You're admin of too many public channels, make some channels private to change the username of this channel.
//  400 CHANNEL_INVALID: The provided channel is invalid.
//  400 CHAT_ID_INVALID: The provided chat id is invalid.
//  400 USERNAME_INVALID: The provided username is not valid.
//
// See https://core.telegram.org/method/channels.checkUsername for reference.
func (c *Client) ChannelsCheckUsername(ctx context.Context, request *ChannelsCheckUsernameRequest) (bool, error) {
	var result BoolBox

	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return false, err
	}
	_, ok := result.Bool.(*BoolTrue)
	return ok, nil
}