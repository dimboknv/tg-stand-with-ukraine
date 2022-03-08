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

// InvokeAfterMsgsRequest represents TL type `invokeAfterMsgs#3dc4b4f0`.
// Invokes a query after a successfull completion of previous queries
//
// See https://core.telegram.org/constructor/invokeAfterMsgs for reference.
type InvokeAfterMsgsRequest struct {
	// List of messages on which a current query depends
	MsgIDs []int64
	// The query itself
	Query bin.Object
}

// InvokeAfterMsgsRequestTypeID is TL type id of InvokeAfterMsgsRequest.
const InvokeAfterMsgsRequestTypeID = 0x3dc4b4f0

// Ensuring interfaces in compile-time for InvokeAfterMsgsRequest.
var (
	_ bin.Encoder     = &InvokeAfterMsgsRequest{}
	_ bin.Decoder     = &InvokeAfterMsgsRequest{}
	_ bin.BareEncoder = &InvokeAfterMsgsRequest{}
	_ bin.BareDecoder = &InvokeAfterMsgsRequest{}
)

func (i *InvokeAfterMsgsRequest) Zero() bool {
	if i == nil {
		return true
	}
	if !(i.MsgIDs == nil) {
		return false
	}
	if !(i.Query == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (i *InvokeAfterMsgsRequest) String() string {
	if i == nil {
		return "InvokeAfterMsgsRequest(nil)"
	}
	type Alias InvokeAfterMsgsRequest
	return fmt.Sprintf("InvokeAfterMsgsRequest%+v", Alias(*i))
}

// FillFrom fills InvokeAfterMsgsRequest from given interface.
func (i *InvokeAfterMsgsRequest) FillFrom(from interface {
	GetMsgIDs() (value []int64)
	GetQuery() (value bin.Object)
}) {
	i.MsgIDs = from.GetMsgIDs()
	i.Query = from.GetQuery()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*InvokeAfterMsgsRequest) TypeID() uint32 {
	return InvokeAfterMsgsRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*InvokeAfterMsgsRequest) TypeName() string {
	return "invokeAfterMsgs"
}

// TypeInfo returns info about TL type.
func (i *InvokeAfterMsgsRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "invokeAfterMsgs",
		ID:   InvokeAfterMsgsRequestTypeID,
	}
	if i == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "MsgIDs",
			SchemaName: "msg_ids",
		},
		{
			Name:       "Query",
			SchemaName: "query",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (i *InvokeAfterMsgsRequest) Encode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't encode invokeAfterMsgs#3dc4b4f0 as nil")
	}
	b.PutID(InvokeAfterMsgsRequestTypeID)
	return i.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (i *InvokeAfterMsgsRequest) EncodeBare(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't encode invokeAfterMsgs#3dc4b4f0 as nil")
	}
	b.PutVectorHeader(len(i.MsgIDs))
	for _, v := range i.MsgIDs {
		b.PutLong(v)
	}
	if err := i.Query.Encode(b); err != nil {
		return fmt.Errorf("unable to encode invokeAfterMsgs#3dc4b4f0: field query: %w", err)
	}
	return nil
}

// Decode implements bin.Decoder.
func (i *InvokeAfterMsgsRequest) Decode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't decode invokeAfterMsgs#3dc4b4f0 to nil")
	}
	if err := b.ConsumeID(InvokeAfterMsgsRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode invokeAfterMsgs#3dc4b4f0: %w", err)
	}
	return i.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (i *InvokeAfterMsgsRequest) DecodeBare(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't decode invokeAfterMsgs#3dc4b4f0 to nil")
	}
	{
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode invokeAfterMsgs#3dc4b4f0: field msg_ids: %w", err)
		}

		if headerLen > 0 {
			i.MsgIDs = make([]int64, 0, headerLen%bin.PreallocateLimit)
		}
		for idx := 0; idx < headerLen; idx++ {
			value, err := b.Long()
			if err != nil {
				return fmt.Errorf("unable to decode invokeAfterMsgs#3dc4b4f0: field msg_ids: %w", err)
			}
			i.MsgIDs = append(i.MsgIDs, value)
		}
	}
	{
		if err := i.Query.Decode(b); err != nil {
			return fmt.Errorf("unable to decode invokeAfterMsgs#3dc4b4f0: field query: %w", err)
		}
	}
	return nil
}

// GetMsgIDs returns value of MsgIDs field.
func (i *InvokeAfterMsgsRequest) GetMsgIDs() (value []int64) {
	if i == nil {
		return
	}
	return i.MsgIDs
}

// GetQuery returns value of Query field.
func (i *InvokeAfterMsgsRequest) GetQuery() (value bin.Object) {
	if i == nil {
		return
	}
	return i.Query
}
