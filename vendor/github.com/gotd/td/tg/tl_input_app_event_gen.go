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

// InputAppEvent represents TL type `inputAppEvent#1d1b1245`.
// Event that occured in the application.
//
// See https://core.telegram.org/constructor/inputAppEvent for reference.
type InputAppEvent struct {
	// Client's exact timestamp for the event
	Time float64
	// Type of event
	Type string
	// Arbitrary numeric value for more convenient selection of certain event types, or
	// events referring to a certain object
	Peer int64
	// Details of the event
	Data JSONValueClass
}

// InputAppEventTypeID is TL type id of InputAppEvent.
const InputAppEventTypeID = 0x1d1b1245

// Ensuring interfaces in compile-time for InputAppEvent.
var (
	_ bin.Encoder     = &InputAppEvent{}
	_ bin.Decoder     = &InputAppEvent{}
	_ bin.BareEncoder = &InputAppEvent{}
	_ bin.BareDecoder = &InputAppEvent{}
)

func (i *InputAppEvent) Zero() bool {
	if i == nil {
		return true
	}
	if !(i.Time == 0) {
		return false
	}
	if !(i.Type == "") {
		return false
	}
	if !(i.Peer == 0) {
		return false
	}
	if !(i.Data == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (i *InputAppEvent) String() string {
	if i == nil {
		return "InputAppEvent(nil)"
	}
	type Alias InputAppEvent
	return fmt.Sprintf("InputAppEvent%+v", Alias(*i))
}

// FillFrom fills InputAppEvent from given interface.
func (i *InputAppEvent) FillFrom(from interface {
	GetTime() (value float64)
	GetType() (value string)
	GetPeer() (value int64)
	GetData() (value JSONValueClass)
}) {
	i.Time = from.GetTime()
	i.Type = from.GetType()
	i.Peer = from.GetPeer()
	i.Data = from.GetData()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*InputAppEvent) TypeID() uint32 {
	return InputAppEventTypeID
}

// TypeName returns name of type in TL schema.
func (*InputAppEvent) TypeName() string {
	return "inputAppEvent"
}

// TypeInfo returns info about TL type.
func (i *InputAppEvent) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "inputAppEvent",
		ID:   InputAppEventTypeID,
	}
	if i == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Time",
			SchemaName: "time",
		},
		{
			Name:       "Type",
			SchemaName: "type",
		},
		{
			Name:       "Peer",
			SchemaName: "peer",
		},
		{
			Name:       "Data",
			SchemaName: "data",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (i *InputAppEvent) Encode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't encode inputAppEvent#1d1b1245 as nil")
	}
	b.PutID(InputAppEventTypeID)
	return i.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (i *InputAppEvent) EncodeBare(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't encode inputAppEvent#1d1b1245 as nil")
	}
	b.PutDouble(i.Time)
	b.PutString(i.Type)
	b.PutLong(i.Peer)
	if i.Data == nil {
		return fmt.Errorf("unable to encode inputAppEvent#1d1b1245: field data is nil")
	}
	if err := i.Data.Encode(b); err != nil {
		return fmt.Errorf("unable to encode inputAppEvent#1d1b1245: field data: %w", err)
	}
	return nil
}

// Decode implements bin.Decoder.
func (i *InputAppEvent) Decode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't decode inputAppEvent#1d1b1245 to nil")
	}
	if err := b.ConsumeID(InputAppEventTypeID); err != nil {
		return fmt.Errorf("unable to decode inputAppEvent#1d1b1245: %w", err)
	}
	return i.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (i *InputAppEvent) DecodeBare(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't decode inputAppEvent#1d1b1245 to nil")
	}
	{
		value, err := b.Double()
		if err != nil {
			return fmt.Errorf("unable to decode inputAppEvent#1d1b1245: field time: %w", err)
		}
		i.Time = value
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode inputAppEvent#1d1b1245: field type: %w", err)
		}
		i.Type = value
	}
	{
		value, err := b.Long()
		if err != nil {
			return fmt.Errorf("unable to decode inputAppEvent#1d1b1245: field peer: %w", err)
		}
		i.Peer = value
	}
	{
		value, err := DecodeJSONValue(b)
		if err != nil {
			return fmt.Errorf("unable to decode inputAppEvent#1d1b1245: field data: %w", err)
		}
		i.Data = value
	}
	return nil
}

// GetTime returns value of Time field.
func (i *InputAppEvent) GetTime() (value float64) {
	if i == nil {
		return
	}
	return i.Time
}

// GetType returns value of Type field.
func (i *InputAppEvent) GetType() (value string) {
	if i == nil {
		return
	}
	return i.Type
}

// GetPeer returns value of Peer field.
func (i *InputAppEvent) GetPeer() (value int64) {
	if i == nil {
		return
	}
	return i.Peer
}

// GetData returns value of Data field.
func (i *InputAppEvent) GetData() (value JSONValueClass) {
	if i == nil {
		return
	}
	return i.Data
}