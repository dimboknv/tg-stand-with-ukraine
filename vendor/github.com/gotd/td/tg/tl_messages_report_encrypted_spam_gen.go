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

// MessagesReportEncryptedSpamRequest represents TL type `messages.reportEncryptedSpam#4b0c8c0f`.
// Report a secret chat for spam
//
// See https://core.telegram.org/method/messages.reportEncryptedSpam for reference.
type MessagesReportEncryptedSpamRequest struct {
	// The secret chat to report
	Peer InputEncryptedChat
}

// MessagesReportEncryptedSpamRequestTypeID is TL type id of MessagesReportEncryptedSpamRequest.
const MessagesReportEncryptedSpamRequestTypeID = 0x4b0c8c0f

// Ensuring interfaces in compile-time for MessagesReportEncryptedSpamRequest.
var (
	_ bin.Encoder     = &MessagesReportEncryptedSpamRequest{}
	_ bin.Decoder     = &MessagesReportEncryptedSpamRequest{}
	_ bin.BareEncoder = &MessagesReportEncryptedSpamRequest{}
	_ bin.BareDecoder = &MessagesReportEncryptedSpamRequest{}
)

func (r *MessagesReportEncryptedSpamRequest) Zero() bool {
	if r == nil {
		return true
	}
	if !(r.Peer.Zero()) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (r *MessagesReportEncryptedSpamRequest) String() string {
	if r == nil {
		return "MessagesReportEncryptedSpamRequest(nil)"
	}
	type Alias MessagesReportEncryptedSpamRequest
	return fmt.Sprintf("MessagesReportEncryptedSpamRequest%+v", Alias(*r))
}

// FillFrom fills MessagesReportEncryptedSpamRequest from given interface.
func (r *MessagesReportEncryptedSpamRequest) FillFrom(from interface {
	GetPeer() (value InputEncryptedChat)
}) {
	r.Peer = from.GetPeer()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*MessagesReportEncryptedSpamRequest) TypeID() uint32 {
	return MessagesReportEncryptedSpamRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*MessagesReportEncryptedSpamRequest) TypeName() string {
	return "messages.reportEncryptedSpam"
}

// TypeInfo returns info about TL type.
func (r *MessagesReportEncryptedSpamRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "messages.reportEncryptedSpam",
		ID:   MessagesReportEncryptedSpamRequestTypeID,
	}
	if r == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Peer",
			SchemaName: "peer",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (r *MessagesReportEncryptedSpamRequest) Encode(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't encode messages.reportEncryptedSpam#4b0c8c0f as nil")
	}
	b.PutID(MessagesReportEncryptedSpamRequestTypeID)
	return r.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (r *MessagesReportEncryptedSpamRequest) EncodeBare(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't encode messages.reportEncryptedSpam#4b0c8c0f as nil")
	}
	if err := r.Peer.Encode(b); err != nil {
		return fmt.Errorf("unable to encode messages.reportEncryptedSpam#4b0c8c0f: field peer: %w", err)
	}
	return nil
}

// Decode implements bin.Decoder.
func (r *MessagesReportEncryptedSpamRequest) Decode(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't decode messages.reportEncryptedSpam#4b0c8c0f to nil")
	}
	if err := b.ConsumeID(MessagesReportEncryptedSpamRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode messages.reportEncryptedSpam#4b0c8c0f: %w", err)
	}
	return r.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (r *MessagesReportEncryptedSpamRequest) DecodeBare(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't decode messages.reportEncryptedSpam#4b0c8c0f to nil")
	}
	{
		if err := r.Peer.Decode(b); err != nil {
			return fmt.Errorf("unable to decode messages.reportEncryptedSpam#4b0c8c0f: field peer: %w", err)
		}
	}
	return nil
}

// GetPeer returns value of Peer field.
func (r *MessagesReportEncryptedSpamRequest) GetPeer() (value InputEncryptedChat) {
	if r == nil {
		return
	}
	return r.Peer
}

// MessagesReportEncryptedSpam invokes method messages.reportEncryptedSpam#4b0c8c0f returning error if any.
// Report a secret chat for spam
//
// Possible errors:
//  400 CHAT_ID_INVALID: The provided chat id is invalid.
//
// See https://core.telegram.org/method/messages.reportEncryptedSpam for reference.
func (c *Client) MessagesReportEncryptedSpam(ctx context.Context, peer InputEncryptedChat) (bool, error) {
	var result BoolBox

	request := &MessagesReportEncryptedSpamRequest{
		Peer: peer,
	}
	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return false, err
	}
	_, ok := result.Bool.(*BoolTrue)
	return ok, nil
}
