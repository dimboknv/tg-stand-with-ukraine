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

// ChannelsReadMessageContentsRequest represents TL type `channels.readMessageContents#eab5dc38`.
// Mark channel/supergroup¹ message contents as read
//
// Links:
//  1) https://core.telegram.org/api/channel
//
// See https://core.telegram.org/method/channels.readMessageContents for reference.
type ChannelsReadMessageContentsRequest struct {
	// Channel/supergroup¹
	//
	// Links:
	//  1) https://core.telegram.org/api/channel
	Channel InputChannelClass
	// IDs of messages whose contents should be marked as read
	ID []int
}

// ChannelsReadMessageContentsRequestTypeID is TL type id of ChannelsReadMessageContentsRequest.
const ChannelsReadMessageContentsRequestTypeID = 0xeab5dc38

// Ensuring interfaces in compile-time for ChannelsReadMessageContentsRequest.
var (
	_ bin.Encoder     = &ChannelsReadMessageContentsRequest{}
	_ bin.Decoder     = &ChannelsReadMessageContentsRequest{}
	_ bin.BareEncoder = &ChannelsReadMessageContentsRequest{}
	_ bin.BareDecoder = &ChannelsReadMessageContentsRequest{}
)

func (r *ChannelsReadMessageContentsRequest) Zero() bool {
	if r == nil {
		return true
	}
	if !(r.Channel == nil) {
		return false
	}
	if !(r.ID == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (r *ChannelsReadMessageContentsRequest) String() string {
	if r == nil {
		return "ChannelsReadMessageContentsRequest(nil)"
	}
	type Alias ChannelsReadMessageContentsRequest
	return fmt.Sprintf("ChannelsReadMessageContentsRequest%+v", Alias(*r))
}

// FillFrom fills ChannelsReadMessageContentsRequest from given interface.
func (r *ChannelsReadMessageContentsRequest) FillFrom(from interface {
	GetChannel() (value InputChannelClass)
	GetID() (value []int)
}) {
	r.Channel = from.GetChannel()
	r.ID = from.GetID()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*ChannelsReadMessageContentsRequest) TypeID() uint32 {
	return ChannelsReadMessageContentsRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*ChannelsReadMessageContentsRequest) TypeName() string {
	return "channels.readMessageContents"
}

// TypeInfo returns info about TL type.
func (r *ChannelsReadMessageContentsRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "channels.readMessageContents",
		ID:   ChannelsReadMessageContentsRequestTypeID,
	}
	if r == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Channel",
			SchemaName: "channel",
		},
		{
			Name:       "ID",
			SchemaName: "id",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (r *ChannelsReadMessageContentsRequest) Encode(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't encode channels.readMessageContents#eab5dc38 as nil")
	}
	b.PutID(ChannelsReadMessageContentsRequestTypeID)
	return r.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (r *ChannelsReadMessageContentsRequest) EncodeBare(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't encode channels.readMessageContents#eab5dc38 as nil")
	}
	if r.Channel == nil {
		return fmt.Errorf("unable to encode channels.readMessageContents#eab5dc38: field channel is nil")
	}
	if err := r.Channel.Encode(b); err != nil {
		return fmt.Errorf("unable to encode channels.readMessageContents#eab5dc38: field channel: %w", err)
	}
	b.PutVectorHeader(len(r.ID))
	for _, v := range r.ID {
		b.PutInt(v)
	}
	return nil
}

// Decode implements bin.Decoder.
func (r *ChannelsReadMessageContentsRequest) Decode(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't decode channels.readMessageContents#eab5dc38 to nil")
	}
	if err := b.ConsumeID(ChannelsReadMessageContentsRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode channels.readMessageContents#eab5dc38: %w", err)
	}
	return r.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (r *ChannelsReadMessageContentsRequest) DecodeBare(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't decode channels.readMessageContents#eab5dc38 to nil")
	}
	{
		value, err := DecodeInputChannel(b)
		if err != nil {
			return fmt.Errorf("unable to decode channels.readMessageContents#eab5dc38: field channel: %w", err)
		}
		r.Channel = value
	}
	{
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode channels.readMessageContents#eab5dc38: field id: %w", err)
		}

		if headerLen > 0 {
			r.ID = make([]int, 0, headerLen%bin.PreallocateLimit)
		}
		for idx := 0; idx < headerLen; idx++ {
			value, err := b.Int()
			if err != nil {
				return fmt.Errorf("unable to decode channels.readMessageContents#eab5dc38: field id: %w", err)
			}
			r.ID = append(r.ID, value)
		}
	}
	return nil
}

// GetChannel returns value of Channel field.
func (r *ChannelsReadMessageContentsRequest) GetChannel() (value InputChannelClass) {
	if r == nil {
		return
	}
	return r.Channel
}

// GetID returns value of ID field.
func (r *ChannelsReadMessageContentsRequest) GetID() (value []int) {
	if r == nil {
		return
	}
	return r.ID
}

// GetChannelAsNotEmpty returns mapped value of Channel field.
func (r *ChannelsReadMessageContentsRequest) GetChannelAsNotEmpty() (NotEmptyInputChannel, bool) {
	return r.Channel.AsNotEmpty()
}

// ChannelsReadMessageContents invokes method channels.readMessageContents#eab5dc38 returning error if any.
// Mark channel/supergroup¹ message contents as read
//
// Links:
//  1) https://core.telegram.org/api/channel
//
// Possible errors:
//  400 CHANNEL_INVALID: The provided channel is invalid.
//  400 CHANNEL_PRIVATE: You haven't joined this channel/supergroup.
//  400 MSG_ID_INVALID: Invalid message ID provided.
//
// See https://core.telegram.org/method/channels.readMessageContents for reference.
func (c *Client) ChannelsReadMessageContents(ctx context.Context, request *ChannelsReadMessageContentsRequest) (bool, error) {
	var result BoolBox

	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return false, err
	}
	_, ok := result.Bool.(*BoolTrue)
	return ok, nil
}
