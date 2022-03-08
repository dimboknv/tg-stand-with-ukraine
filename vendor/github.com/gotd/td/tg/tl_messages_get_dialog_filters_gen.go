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

// MessagesGetDialogFiltersRequest represents TL type `messages.getDialogFilters#f19ed96d`.
// Get folders¹
//
// Links:
//  1) https://core.telegram.org/api/folders
//
// See https://core.telegram.org/method/messages.getDialogFilters for reference.
type MessagesGetDialogFiltersRequest struct {
}

// MessagesGetDialogFiltersRequestTypeID is TL type id of MessagesGetDialogFiltersRequest.
const MessagesGetDialogFiltersRequestTypeID = 0xf19ed96d

// Ensuring interfaces in compile-time for MessagesGetDialogFiltersRequest.
var (
	_ bin.Encoder     = &MessagesGetDialogFiltersRequest{}
	_ bin.Decoder     = &MessagesGetDialogFiltersRequest{}
	_ bin.BareEncoder = &MessagesGetDialogFiltersRequest{}
	_ bin.BareDecoder = &MessagesGetDialogFiltersRequest{}
)

func (g *MessagesGetDialogFiltersRequest) Zero() bool {
	if g == nil {
		return true
	}

	return true
}

// String implements fmt.Stringer.
func (g *MessagesGetDialogFiltersRequest) String() string {
	if g == nil {
		return "MessagesGetDialogFiltersRequest(nil)"
	}
	type Alias MessagesGetDialogFiltersRequest
	return fmt.Sprintf("MessagesGetDialogFiltersRequest%+v", Alias(*g))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*MessagesGetDialogFiltersRequest) TypeID() uint32 {
	return MessagesGetDialogFiltersRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*MessagesGetDialogFiltersRequest) TypeName() string {
	return "messages.getDialogFilters"
}

// TypeInfo returns info about TL type.
func (g *MessagesGetDialogFiltersRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "messages.getDialogFilters",
		ID:   MessagesGetDialogFiltersRequestTypeID,
	}
	if g == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{}
	return typ
}

// Encode implements bin.Encoder.
func (g *MessagesGetDialogFiltersRequest) Encode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode messages.getDialogFilters#f19ed96d as nil")
	}
	b.PutID(MessagesGetDialogFiltersRequestTypeID)
	return g.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (g *MessagesGetDialogFiltersRequest) EncodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode messages.getDialogFilters#f19ed96d as nil")
	}
	return nil
}

// Decode implements bin.Decoder.
func (g *MessagesGetDialogFiltersRequest) Decode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode messages.getDialogFilters#f19ed96d to nil")
	}
	if err := b.ConsumeID(MessagesGetDialogFiltersRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode messages.getDialogFilters#f19ed96d: %w", err)
	}
	return g.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (g *MessagesGetDialogFiltersRequest) DecodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode messages.getDialogFilters#f19ed96d to nil")
	}
	return nil
}

// MessagesGetDialogFilters invokes method messages.getDialogFilters#f19ed96d returning error if any.
// Get folders¹
//
// Links:
//  1) https://core.telegram.org/api/folders
//
// See https://core.telegram.org/method/messages.getDialogFilters for reference.
func (c *Client) MessagesGetDialogFilters(ctx context.Context) ([]DialogFilter, error) {
	var result DialogFilterVector

	request := &MessagesGetDialogFiltersRequest{}
	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return nil, err
	}
	return []DialogFilter(result.Elems), nil
}