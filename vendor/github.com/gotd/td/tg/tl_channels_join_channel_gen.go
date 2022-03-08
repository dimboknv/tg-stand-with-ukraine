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

// ChannelsJoinChannelRequest represents TL type `channels.joinChannel#24b524c5`.
// Join a channel/supergroup
//
// See https://core.telegram.org/method/channels.joinChannel for reference.
type ChannelsJoinChannelRequest struct {
	// Channel/supergroup to join
	Channel InputChannelClass
}

// ChannelsJoinChannelRequestTypeID is TL type id of ChannelsJoinChannelRequest.
const ChannelsJoinChannelRequestTypeID = 0x24b524c5

// Ensuring interfaces in compile-time for ChannelsJoinChannelRequest.
var (
	_ bin.Encoder     = &ChannelsJoinChannelRequest{}
	_ bin.Decoder     = &ChannelsJoinChannelRequest{}
	_ bin.BareEncoder = &ChannelsJoinChannelRequest{}
	_ bin.BareDecoder = &ChannelsJoinChannelRequest{}
)

func (j *ChannelsJoinChannelRequest) Zero() bool {
	if j == nil {
		return true
	}
	if !(j.Channel == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (j *ChannelsJoinChannelRequest) String() string {
	if j == nil {
		return "ChannelsJoinChannelRequest(nil)"
	}
	type Alias ChannelsJoinChannelRequest
	return fmt.Sprintf("ChannelsJoinChannelRequest%+v", Alias(*j))
}

// FillFrom fills ChannelsJoinChannelRequest from given interface.
func (j *ChannelsJoinChannelRequest) FillFrom(from interface {
	GetChannel() (value InputChannelClass)
}) {
	j.Channel = from.GetChannel()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*ChannelsJoinChannelRequest) TypeID() uint32 {
	return ChannelsJoinChannelRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*ChannelsJoinChannelRequest) TypeName() string {
	return "channels.joinChannel"
}

// TypeInfo returns info about TL type.
func (j *ChannelsJoinChannelRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "channels.joinChannel",
		ID:   ChannelsJoinChannelRequestTypeID,
	}
	if j == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Channel",
			SchemaName: "channel",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (j *ChannelsJoinChannelRequest) Encode(b *bin.Buffer) error {
	if j == nil {
		return fmt.Errorf("can't encode channels.joinChannel#24b524c5 as nil")
	}
	b.PutID(ChannelsJoinChannelRequestTypeID)
	return j.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (j *ChannelsJoinChannelRequest) EncodeBare(b *bin.Buffer) error {
	if j == nil {
		return fmt.Errorf("can't encode channels.joinChannel#24b524c5 as nil")
	}
	if j.Channel == nil {
		return fmt.Errorf("unable to encode channels.joinChannel#24b524c5: field channel is nil")
	}
	if err := j.Channel.Encode(b); err != nil {
		return fmt.Errorf("unable to encode channels.joinChannel#24b524c5: field channel: %w", err)
	}
	return nil
}

// Decode implements bin.Decoder.
func (j *ChannelsJoinChannelRequest) Decode(b *bin.Buffer) error {
	if j == nil {
		return fmt.Errorf("can't decode channels.joinChannel#24b524c5 to nil")
	}
	if err := b.ConsumeID(ChannelsJoinChannelRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode channels.joinChannel#24b524c5: %w", err)
	}
	return j.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (j *ChannelsJoinChannelRequest) DecodeBare(b *bin.Buffer) error {
	if j == nil {
		return fmt.Errorf("can't decode channels.joinChannel#24b524c5 to nil")
	}
	{
		value, err := DecodeInputChannel(b)
		if err != nil {
			return fmt.Errorf("unable to decode channels.joinChannel#24b524c5: field channel: %w", err)
		}
		j.Channel = value
	}
	return nil
}

// GetChannel returns value of Channel field.
func (j *ChannelsJoinChannelRequest) GetChannel() (value InputChannelClass) {
	if j == nil {
		return
	}
	return j.Channel
}

// GetChannelAsNotEmpty returns mapped value of Channel field.
func (j *ChannelsJoinChannelRequest) GetChannelAsNotEmpty() (NotEmptyInputChannel, bool) {
	return j.Channel.AsNotEmpty()
}

// ChannelsJoinChannel invokes method channels.joinChannel#24b524c5 returning error if any.
// Join a channel/supergroup
//
// Possible errors:
//  400 CHANNELS_TOO_MUCH: You have joined too many channels/supergroups.
//  400 CHANNEL_INVALID: The provided channel is invalid.
//  400 CHANNEL_PRIVATE: You haven't joined this channel/supergroup.
//  400 CHAT_INVALID: Invalid chat.
//  400 INVITE_HASH_EMPTY: The invite hash is empty.
//  400 INVITE_HASH_EXPIRED: The invite link has expired.
//  400 INVITE_HASH_INVALID: The invite hash is invalid.
//  400 MSG_ID_INVALID: Invalid message ID provided.
//  400 PEER_ID_INVALID: The provided peer id is invalid.
//  400 USERS_TOO_MUCH: The maximum number of users has been exceeded (to create a chat, for example).
//  400 USER_ALREADY_PARTICIPANT: The user is already in the group.
//  400 USER_CHANNELS_TOO_MUCH: One of the users you tried to add is already in too many channels/supergroups.
//
// See https://core.telegram.org/method/channels.joinChannel for reference.
func (c *Client) ChannelsJoinChannel(ctx context.Context, channel InputChannelClass) (UpdatesClass, error) {
	var result UpdatesBox

	request := &ChannelsJoinChannelRequest{
		Channel: channel,
	}
	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return nil, err
	}
	return result.Updates, nil
}