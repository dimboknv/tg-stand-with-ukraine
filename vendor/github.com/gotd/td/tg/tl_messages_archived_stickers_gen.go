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

// MessagesArchivedStickers represents TL type `messages.archivedStickers#4fcba9c8`.
// Archived stickersets
//
// See https://core.telegram.org/constructor/messages.archivedStickers for reference.
type MessagesArchivedStickers struct {
	// Number of archived stickers
	Count int
	// Archived stickersets
	Sets []StickerSetCoveredClass
}

// MessagesArchivedStickersTypeID is TL type id of MessagesArchivedStickers.
const MessagesArchivedStickersTypeID = 0x4fcba9c8

// Ensuring interfaces in compile-time for MessagesArchivedStickers.
var (
	_ bin.Encoder     = &MessagesArchivedStickers{}
	_ bin.Decoder     = &MessagesArchivedStickers{}
	_ bin.BareEncoder = &MessagesArchivedStickers{}
	_ bin.BareDecoder = &MessagesArchivedStickers{}
)

func (a *MessagesArchivedStickers) Zero() bool {
	if a == nil {
		return true
	}
	if !(a.Count == 0) {
		return false
	}
	if !(a.Sets == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (a *MessagesArchivedStickers) String() string {
	if a == nil {
		return "MessagesArchivedStickers(nil)"
	}
	type Alias MessagesArchivedStickers
	return fmt.Sprintf("MessagesArchivedStickers%+v", Alias(*a))
}

// FillFrom fills MessagesArchivedStickers from given interface.
func (a *MessagesArchivedStickers) FillFrom(from interface {
	GetCount() (value int)
	GetSets() (value []StickerSetCoveredClass)
}) {
	a.Count = from.GetCount()
	a.Sets = from.GetSets()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*MessagesArchivedStickers) TypeID() uint32 {
	return MessagesArchivedStickersTypeID
}

// TypeName returns name of type in TL schema.
func (*MessagesArchivedStickers) TypeName() string {
	return "messages.archivedStickers"
}

// TypeInfo returns info about TL type.
func (a *MessagesArchivedStickers) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "messages.archivedStickers",
		ID:   MessagesArchivedStickersTypeID,
	}
	if a == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Count",
			SchemaName: "count",
		},
		{
			Name:       "Sets",
			SchemaName: "sets",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (a *MessagesArchivedStickers) Encode(b *bin.Buffer) error {
	if a == nil {
		return fmt.Errorf("can't encode messages.archivedStickers#4fcba9c8 as nil")
	}
	b.PutID(MessagesArchivedStickersTypeID)
	return a.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (a *MessagesArchivedStickers) EncodeBare(b *bin.Buffer) error {
	if a == nil {
		return fmt.Errorf("can't encode messages.archivedStickers#4fcba9c8 as nil")
	}
	b.PutInt(a.Count)
	b.PutVectorHeader(len(a.Sets))
	for idx, v := range a.Sets {
		if v == nil {
			return fmt.Errorf("unable to encode messages.archivedStickers#4fcba9c8: field sets element with index %d is nil", idx)
		}
		if err := v.Encode(b); err != nil {
			return fmt.Errorf("unable to encode messages.archivedStickers#4fcba9c8: field sets element with index %d: %w", idx, err)
		}
	}
	return nil
}

// Decode implements bin.Decoder.
func (a *MessagesArchivedStickers) Decode(b *bin.Buffer) error {
	if a == nil {
		return fmt.Errorf("can't decode messages.archivedStickers#4fcba9c8 to nil")
	}
	if err := b.ConsumeID(MessagesArchivedStickersTypeID); err != nil {
		return fmt.Errorf("unable to decode messages.archivedStickers#4fcba9c8: %w", err)
	}
	return a.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (a *MessagesArchivedStickers) DecodeBare(b *bin.Buffer) error {
	if a == nil {
		return fmt.Errorf("can't decode messages.archivedStickers#4fcba9c8 to nil")
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode messages.archivedStickers#4fcba9c8: field count: %w", err)
		}
		a.Count = value
	}
	{
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode messages.archivedStickers#4fcba9c8: field sets: %w", err)
		}

		if headerLen > 0 {
			a.Sets = make([]StickerSetCoveredClass, 0, headerLen%bin.PreallocateLimit)
		}
		for idx := 0; idx < headerLen; idx++ {
			value, err := DecodeStickerSetCovered(b)
			if err != nil {
				return fmt.Errorf("unable to decode messages.archivedStickers#4fcba9c8: field sets: %w", err)
			}
			a.Sets = append(a.Sets, value)
		}
	}
	return nil
}

// GetCount returns value of Count field.
func (a *MessagesArchivedStickers) GetCount() (value int) {
	if a == nil {
		return
	}
	return a.Count
}

// GetSets returns value of Sets field.
func (a *MessagesArchivedStickers) GetSets() (value []StickerSetCoveredClass) {
	if a == nil {
		return
	}
	return a.Sets
}

// MapSets returns field Sets wrapped in StickerSetCoveredClassArray helper.
func (a *MessagesArchivedStickers) MapSets() (value StickerSetCoveredClassArray) {
	return StickerSetCoveredClassArray(a.Sets)
}
