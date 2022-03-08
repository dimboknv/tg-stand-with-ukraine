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

// HelpSupportName represents TL type `help.supportName#8c05f1c9`.
// Localized name for telegram support
//
// See https://core.telegram.org/constructor/help.supportName for reference.
type HelpSupportName struct {
	// Localized name
	Name string
}

// HelpSupportNameTypeID is TL type id of HelpSupportName.
const HelpSupportNameTypeID = 0x8c05f1c9

// Ensuring interfaces in compile-time for HelpSupportName.
var (
	_ bin.Encoder     = &HelpSupportName{}
	_ bin.Decoder     = &HelpSupportName{}
	_ bin.BareEncoder = &HelpSupportName{}
	_ bin.BareDecoder = &HelpSupportName{}
)

func (s *HelpSupportName) Zero() bool {
	if s == nil {
		return true
	}
	if !(s.Name == "") {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (s *HelpSupportName) String() string {
	if s == nil {
		return "HelpSupportName(nil)"
	}
	type Alias HelpSupportName
	return fmt.Sprintf("HelpSupportName%+v", Alias(*s))
}

// FillFrom fills HelpSupportName from given interface.
func (s *HelpSupportName) FillFrom(from interface {
	GetName() (value string)
}) {
	s.Name = from.GetName()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*HelpSupportName) TypeID() uint32 {
	return HelpSupportNameTypeID
}

// TypeName returns name of type in TL schema.
func (*HelpSupportName) TypeName() string {
	return "help.supportName"
}

// TypeInfo returns info about TL type.
func (s *HelpSupportName) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "help.supportName",
		ID:   HelpSupportNameTypeID,
	}
	if s == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Name",
			SchemaName: "name",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (s *HelpSupportName) Encode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode help.supportName#8c05f1c9 as nil")
	}
	b.PutID(HelpSupportNameTypeID)
	return s.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (s *HelpSupportName) EncodeBare(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode help.supportName#8c05f1c9 as nil")
	}
	b.PutString(s.Name)
	return nil
}

// Decode implements bin.Decoder.
func (s *HelpSupportName) Decode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode help.supportName#8c05f1c9 to nil")
	}
	if err := b.ConsumeID(HelpSupportNameTypeID); err != nil {
		return fmt.Errorf("unable to decode help.supportName#8c05f1c9: %w", err)
	}
	return s.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (s *HelpSupportName) DecodeBare(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode help.supportName#8c05f1c9 to nil")
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode help.supportName#8c05f1c9: field name: %w", err)
		}
		s.Name = value
	}
	return nil
}

// GetName returns value of Name field.
func (s *HelpSupportName) GetName() (value string) {
	if s == nil {
		return
	}
	return s.Name
}
