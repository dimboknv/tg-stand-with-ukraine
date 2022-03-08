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

// RestrictionReason represents TL type `restrictionReason#d072acb4`.
// Restriction reason.
// Contains the reason why access to a certain object must be restricted. Clients are
// supposed to deny access to the channel if the platform field is equal to all or to the
// current platform (ios, android, wp, etc.). Platforms can be concatenated (ios-android,
// ios-wp), unknown platforms are to be ignored. The text is the error message that
// should be shown to the user.
//
// See https://core.telegram.org/constructor/restrictionReason for reference.
type RestrictionReason struct {
	// Platform identifier (ios, android, wp, all, etc.), can be concatenated with a dash as
	// separator (android-ios, ios-wp, etc)
	Platform string
	// Restriction reason (porno, terms, etc.)
	Reason string
	// Error message to be shown to the user
	Text string
}

// RestrictionReasonTypeID is TL type id of RestrictionReason.
const RestrictionReasonTypeID = 0xd072acb4

// Ensuring interfaces in compile-time for RestrictionReason.
var (
	_ bin.Encoder     = &RestrictionReason{}
	_ bin.Decoder     = &RestrictionReason{}
	_ bin.BareEncoder = &RestrictionReason{}
	_ bin.BareDecoder = &RestrictionReason{}
)

func (r *RestrictionReason) Zero() bool {
	if r == nil {
		return true
	}
	if !(r.Platform == "") {
		return false
	}
	if !(r.Reason == "") {
		return false
	}
	if !(r.Text == "") {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (r *RestrictionReason) String() string {
	if r == nil {
		return "RestrictionReason(nil)"
	}
	type Alias RestrictionReason
	return fmt.Sprintf("RestrictionReason%+v", Alias(*r))
}

// FillFrom fills RestrictionReason from given interface.
func (r *RestrictionReason) FillFrom(from interface {
	GetPlatform() (value string)
	GetReason() (value string)
	GetText() (value string)
}) {
	r.Platform = from.GetPlatform()
	r.Reason = from.GetReason()
	r.Text = from.GetText()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*RestrictionReason) TypeID() uint32 {
	return RestrictionReasonTypeID
}

// TypeName returns name of type in TL schema.
func (*RestrictionReason) TypeName() string {
	return "restrictionReason"
}

// TypeInfo returns info about TL type.
func (r *RestrictionReason) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "restrictionReason",
		ID:   RestrictionReasonTypeID,
	}
	if r == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Platform",
			SchemaName: "platform",
		},
		{
			Name:       "Reason",
			SchemaName: "reason",
		},
		{
			Name:       "Text",
			SchemaName: "text",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (r *RestrictionReason) Encode(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't encode restrictionReason#d072acb4 as nil")
	}
	b.PutID(RestrictionReasonTypeID)
	return r.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (r *RestrictionReason) EncodeBare(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't encode restrictionReason#d072acb4 as nil")
	}
	b.PutString(r.Platform)
	b.PutString(r.Reason)
	b.PutString(r.Text)
	return nil
}

// Decode implements bin.Decoder.
func (r *RestrictionReason) Decode(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't decode restrictionReason#d072acb4 to nil")
	}
	if err := b.ConsumeID(RestrictionReasonTypeID); err != nil {
		return fmt.Errorf("unable to decode restrictionReason#d072acb4: %w", err)
	}
	return r.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (r *RestrictionReason) DecodeBare(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't decode restrictionReason#d072acb4 to nil")
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode restrictionReason#d072acb4: field platform: %w", err)
		}
		r.Platform = value
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode restrictionReason#d072acb4: field reason: %w", err)
		}
		r.Reason = value
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode restrictionReason#d072acb4: field text: %w", err)
		}
		r.Text = value
	}
	return nil
}

// GetPlatform returns value of Platform field.
func (r *RestrictionReason) GetPlatform() (value string) {
	if r == nil {
		return
	}
	return r.Platform
}

// GetReason returns value of Reason field.
func (r *RestrictionReason) GetReason() (value string) {
	if r == nil {
		return
	}
	return r.Reason
}

// GetText returns value of Text field.
func (r *RestrictionReason) GetText() (value string) {
	if r == nil {
		return
	}
	return r.Text
}
