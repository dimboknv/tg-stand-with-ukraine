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

// LangpackGetDifferenceRequest represents TL type `langpack.getDifference#cd984aa5`.
// Get new strings in languagepack
//
// See https://core.telegram.org/method/langpack.getDifference for reference.
type LangpackGetDifferenceRequest struct {
	// Language pack
	LangPack string
	// Language code
	LangCode string
	// Previous localization pack version
	FromVersion int
}

// LangpackGetDifferenceRequestTypeID is TL type id of LangpackGetDifferenceRequest.
const LangpackGetDifferenceRequestTypeID = 0xcd984aa5

// Ensuring interfaces in compile-time for LangpackGetDifferenceRequest.
var (
	_ bin.Encoder     = &LangpackGetDifferenceRequest{}
	_ bin.Decoder     = &LangpackGetDifferenceRequest{}
	_ bin.BareEncoder = &LangpackGetDifferenceRequest{}
	_ bin.BareDecoder = &LangpackGetDifferenceRequest{}
)

func (g *LangpackGetDifferenceRequest) Zero() bool {
	if g == nil {
		return true
	}
	if !(g.LangPack == "") {
		return false
	}
	if !(g.LangCode == "") {
		return false
	}
	if !(g.FromVersion == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (g *LangpackGetDifferenceRequest) String() string {
	if g == nil {
		return "LangpackGetDifferenceRequest(nil)"
	}
	type Alias LangpackGetDifferenceRequest
	return fmt.Sprintf("LangpackGetDifferenceRequest%+v", Alias(*g))
}

// FillFrom fills LangpackGetDifferenceRequest from given interface.
func (g *LangpackGetDifferenceRequest) FillFrom(from interface {
	GetLangPack() (value string)
	GetLangCode() (value string)
	GetFromVersion() (value int)
}) {
	g.LangPack = from.GetLangPack()
	g.LangCode = from.GetLangCode()
	g.FromVersion = from.GetFromVersion()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*LangpackGetDifferenceRequest) TypeID() uint32 {
	return LangpackGetDifferenceRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*LangpackGetDifferenceRequest) TypeName() string {
	return "langpack.getDifference"
}

// TypeInfo returns info about TL type.
func (g *LangpackGetDifferenceRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "langpack.getDifference",
		ID:   LangpackGetDifferenceRequestTypeID,
	}
	if g == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "LangPack",
			SchemaName: "lang_pack",
		},
		{
			Name:       "LangCode",
			SchemaName: "lang_code",
		},
		{
			Name:       "FromVersion",
			SchemaName: "from_version",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (g *LangpackGetDifferenceRequest) Encode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode langpack.getDifference#cd984aa5 as nil")
	}
	b.PutID(LangpackGetDifferenceRequestTypeID)
	return g.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (g *LangpackGetDifferenceRequest) EncodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode langpack.getDifference#cd984aa5 as nil")
	}
	b.PutString(g.LangPack)
	b.PutString(g.LangCode)
	b.PutInt(g.FromVersion)
	return nil
}

// Decode implements bin.Decoder.
func (g *LangpackGetDifferenceRequest) Decode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode langpack.getDifference#cd984aa5 to nil")
	}
	if err := b.ConsumeID(LangpackGetDifferenceRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode langpack.getDifference#cd984aa5: %w", err)
	}
	return g.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (g *LangpackGetDifferenceRequest) DecodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode langpack.getDifference#cd984aa5 to nil")
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode langpack.getDifference#cd984aa5: field lang_pack: %w", err)
		}
		g.LangPack = value
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode langpack.getDifference#cd984aa5: field lang_code: %w", err)
		}
		g.LangCode = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode langpack.getDifference#cd984aa5: field from_version: %w", err)
		}
		g.FromVersion = value
	}
	return nil
}

// GetLangPack returns value of LangPack field.
func (g *LangpackGetDifferenceRequest) GetLangPack() (value string) {
	if g == nil {
		return
	}
	return g.LangPack
}

// GetLangCode returns value of LangCode field.
func (g *LangpackGetDifferenceRequest) GetLangCode() (value string) {
	if g == nil {
		return
	}
	return g.LangCode
}

// GetFromVersion returns value of FromVersion field.
func (g *LangpackGetDifferenceRequest) GetFromVersion() (value int) {
	if g == nil {
		return
	}
	return g.FromVersion
}

// LangpackGetDifference invokes method langpack.getDifference#cd984aa5 returning error if any.
// Get new strings in languagepack
//
// Possible errors:
//  400 LANG_PACK_INVALID: The provided language pack is invalid.
//
// See https://core.telegram.org/method/langpack.getDifference for reference.
func (c *Client) LangpackGetDifference(ctx context.Context, request *LangpackGetDifferenceRequest) (*LangPackDifference, error) {
	var result LangPackDifference

	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
