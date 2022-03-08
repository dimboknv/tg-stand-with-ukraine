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

// HelpUserInfoEmpty represents TL type `help.userInfoEmpty#f3ae2eed`.
// Internal use
//
// See https://core.telegram.org/constructor/help.userInfoEmpty for reference.
type HelpUserInfoEmpty struct {
}

// HelpUserInfoEmptyTypeID is TL type id of HelpUserInfoEmpty.
const HelpUserInfoEmptyTypeID = 0xf3ae2eed

// construct implements constructor of HelpUserInfoClass.
func (u HelpUserInfoEmpty) construct() HelpUserInfoClass { return &u }

// Ensuring interfaces in compile-time for HelpUserInfoEmpty.
var (
	_ bin.Encoder     = &HelpUserInfoEmpty{}
	_ bin.Decoder     = &HelpUserInfoEmpty{}
	_ bin.BareEncoder = &HelpUserInfoEmpty{}
	_ bin.BareDecoder = &HelpUserInfoEmpty{}

	_ HelpUserInfoClass = &HelpUserInfoEmpty{}
)

func (u *HelpUserInfoEmpty) Zero() bool {
	if u == nil {
		return true
	}

	return true
}

// String implements fmt.Stringer.
func (u *HelpUserInfoEmpty) String() string {
	if u == nil {
		return "HelpUserInfoEmpty(nil)"
	}
	type Alias HelpUserInfoEmpty
	return fmt.Sprintf("HelpUserInfoEmpty%+v", Alias(*u))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*HelpUserInfoEmpty) TypeID() uint32 {
	return HelpUserInfoEmptyTypeID
}

// TypeName returns name of type in TL schema.
func (*HelpUserInfoEmpty) TypeName() string {
	return "help.userInfoEmpty"
}

// TypeInfo returns info about TL type.
func (u *HelpUserInfoEmpty) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "help.userInfoEmpty",
		ID:   HelpUserInfoEmptyTypeID,
	}
	if u == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{}
	return typ
}

// Encode implements bin.Encoder.
func (u *HelpUserInfoEmpty) Encode(b *bin.Buffer) error {
	if u == nil {
		return fmt.Errorf("can't encode help.userInfoEmpty#f3ae2eed as nil")
	}
	b.PutID(HelpUserInfoEmptyTypeID)
	return u.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (u *HelpUserInfoEmpty) EncodeBare(b *bin.Buffer) error {
	if u == nil {
		return fmt.Errorf("can't encode help.userInfoEmpty#f3ae2eed as nil")
	}
	return nil
}

// Decode implements bin.Decoder.
func (u *HelpUserInfoEmpty) Decode(b *bin.Buffer) error {
	if u == nil {
		return fmt.Errorf("can't decode help.userInfoEmpty#f3ae2eed to nil")
	}
	if err := b.ConsumeID(HelpUserInfoEmptyTypeID); err != nil {
		return fmt.Errorf("unable to decode help.userInfoEmpty#f3ae2eed: %w", err)
	}
	return u.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (u *HelpUserInfoEmpty) DecodeBare(b *bin.Buffer) error {
	if u == nil {
		return fmt.Errorf("can't decode help.userInfoEmpty#f3ae2eed to nil")
	}
	return nil
}

// HelpUserInfo represents TL type `help.userInfo#1eb3758`.
// Internal use
//
// See https://core.telegram.org/constructor/help.userInfo for reference.
type HelpUserInfo struct {
	// Info
	Message string
	// Message entities for styled text¹
	//
	// Links:
	//  1) https://core.telegram.org/api/entities
	Entities []MessageEntityClass
	// Author
	Author string
	// Date
	Date int
}

// HelpUserInfoTypeID is TL type id of HelpUserInfo.
const HelpUserInfoTypeID = 0x1eb3758

// construct implements constructor of HelpUserInfoClass.
func (u HelpUserInfo) construct() HelpUserInfoClass { return &u }

// Ensuring interfaces in compile-time for HelpUserInfo.
var (
	_ bin.Encoder     = &HelpUserInfo{}
	_ bin.Decoder     = &HelpUserInfo{}
	_ bin.BareEncoder = &HelpUserInfo{}
	_ bin.BareDecoder = &HelpUserInfo{}

	_ HelpUserInfoClass = &HelpUserInfo{}
)

func (u *HelpUserInfo) Zero() bool {
	if u == nil {
		return true
	}
	if !(u.Message == "") {
		return false
	}
	if !(u.Entities == nil) {
		return false
	}
	if !(u.Author == "") {
		return false
	}
	if !(u.Date == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (u *HelpUserInfo) String() string {
	if u == nil {
		return "HelpUserInfo(nil)"
	}
	type Alias HelpUserInfo
	return fmt.Sprintf("HelpUserInfo%+v", Alias(*u))
}

// FillFrom fills HelpUserInfo from given interface.
func (u *HelpUserInfo) FillFrom(from interface {
	GetMessage() (value string)
	GetEntities() (value []MessageEntityClass)
	GetAuthor() (value string)
	GetDate() (value int)
}) {
	u.Message = from.GetMessage()
	u.Entities = from.GetEntities()
	u.Author = from.GetAuthor()
	u.Date = from.GetDate()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*HelpUserInfo) TypeID() uint32 {
	return HelpUserInfoTypeID
}

// TypeName returns name of type in TL schema.
func (*HelpUserInfo) TypeName() string {
	return "help.userInfo"
}

// TypeInfo returns info about TL type.
func (u *HelpUserInfo) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "help.userInfo",
		ID:   HelpUserInfoTypeID,
	}
	if u == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Message",
			SchemaName: "message",
		},
		{
			Name:       "Entities",
			SchemaName: "entities",
		},
		{
			Name:       "Author",
			SchemaName: "author",
		},
		{
			Name:       "Date",
			SchemaName: "date",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (u *HelpUserInfo) Encode(b *bin.Buffer) error {
	if u == nil {
		return fmt.Errorf("can't encode help.userInfo#1eb3758 as nil")
	}
	b.PutID(HelpUserInfoTypeID)
	return u.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (u *HelpUserInfo) EncodeBare(b *bin.Buffer) error {
	if u == nil {
		return fmt.Errorf("can't encode help.userInfo#1eb3758 as nil")
	}
	b.PutString(u.Message)
	b.PutVectorHeader(len(u.Entities))
	for idx, v := range u.Entities {
		if v == nil {
			return fmt.Errorf("unable to encode help.userInfo#1eb3758: field entities element with index %d is nil", idx)
		}
		if err := v.Encode(b); err != nil {
			return fmt.Errorf("unable to encode help.userInfo#1eb3758: field entities element with index %d: %w", idx, err)
		}
	}
	b.PutString(u.Author)
	b.PutInt(u.Date)
	return nil
}

// Decode implements bin.Decoder.
func (u *HelpUserInfo) Decode(b *bin.Buffer) error {
	if u == nil {
		return fmt.Errorf("can't decode help.userInfo#1eb3758 to nil")
	}
	if err := b.ConsumeID(HelpUserInfoTypeID); err != nil {
		return fmt.Errorf("unable to decode help.userInfo#1eb3758: %w", err)
	}
	return u.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (u *HelpUserInfo) DecodeBare(b *bin.Buffer) error {
	if u == nil {
		return fmt.Errorf("can't decode help.userInfo#1eb3758 to nil")
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode help.userInfo#1eb3758: field message: %w", err)
		}
		u.Message = value
	}
	{
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode help.userInfo#1eb3758: field entities: %w", err)
		}

		if headerLen > 0 {
			u.Entities = make([]MessageEntityClass, 0, headerLen%bin.PreallocateLimit)
		}
		for idx := 0; idx < headerLen; idx++ {
			value, err := DecodeMessageEntity(b)
			if err != nil {
				return fmt.Errorf("unable to decode help.userInfo#1eb3758: field entities: %w", err)
			}
			u.Entities = append(u.Entities, value)
		}
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode help.userInfo#1eb3758: field author: %w", err)
		}
		u.Author = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode help.userInfo#1eb3758: field date: %w", err)
		}
		u.Date = value
	}
	return nil
}

// GetMessage returns value of Message field.
func (u *HelpUserInfo) GetMessage() (value string) {
	if u == nil {
		return
	}
	return u.Message
}

// GetEntities returns value of Entities field.
func (u *HelpUserInfo) GetEntities() (value []MessageEntityClass) {
	if u == nil {
		return
	}
	return u.Entities
}

// GetAuthor returns value of Author field.
func (u *HelpUserInfo) GetAuthor() (value string) {
	if u == nil {
		return
	}
	return u.Author
}

// GetDate returns value of Date field.
func (u *HelpUserInfo) GetDate() (value int) {
	if u == nil {
		return
	}
	return u.Date
}

// MapEntities returns field Entities wrapped in MessageEntityClassArray helper.
func (u *HelpUserInfo) MapEntities() (value MessageEntityClassArray) {
	return MessageEntityClassArray(u.Entities)
}

// HelpUserInfoClassName is schema name of HelpUserInfoClass.
const HelpUserInfoClassName = "help.UserInfo"

// HelpUserInfoClass represents help.UserInfo generic type.
//
// See https://core.telegram.org/type/help.UserInfo for reference.
//
// Example:
//  g, err := tg.DecodeHelpUserInfo(buf)
//  if err != nil {
//      panic(err)
//  }
//  switch v := g.(type) {
//  case *tg.HelpUserInfoEmpty: // help.userInfoEmpty#f3ae2eed
//  case *tg.HelpUserInfo: // help.userInfo#1eb3758
//  default: panic(v)
//  }
type HelpUserInfoClass interface {
	bin.Encoder
	bin.Decoder
	bin.BareEncoder
	bin.BareDecoder
	construct() HelpUserInfoClass

	// TypeID returns type id in TL schema.
	//
	// See https://core.telegram.org/mtproto/TL-tl#remarks.
	TypeID() uint32
	// TypeName returns name of type in TL schema.
	TypeName() string
	// String implements fmt.Stringer.
	String() string
	// Zero returns true if current object has a zero value.
	Zero() bool

	// AsNotEmpty tries to map HelpUserInfoClass to HelpUserInfo.
	AsNotEmpty() (*HelpUserInfo, bool)
}

// AsNotEmpty tries to map HelpUserInfoEmpty to HelpUserInfo.
func (u *HelpUserInfoEmpty) AsNotEmpty() (*HelpUserInfo, bool) {
	return nil, false
}

// AsNotEmpty tries to map HelpUserInfo to HelpUserInfo.
func (u *HelpUserInfo) AsNotEmpty() (*HelpUserInfo, bool) {
	return u, true
}

// DecodeHelpUserInfo implements binary de-serialization for HelpUserInfoClass.
func DecodeHelpUserInfo(buf *bin.Buffer) (HelpUserInfoClass, error) {
	id, err := buf.PeekID()
	if err != nil {
		return nil, err
	}
	switch id {
	case HelpUserInfoEmptyTypeID:
		// Decoding help.userInfoEmpty#f3ae2eed.
		v := HelpUserInfoEmpty{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode HelpUserInfoClass: %w", err)
		}
		return &v, nil
	case HelpUserInfoTypeID:
		// Decoding help.userInfo#1eb3758.
		v := HelpUserInfo{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode HelpUserInfoClass: %w", err)
		}
		return &v, nil
	default:
		return nil, fmt.Errorf("unable to decode HelpUserInfoClass: %w", bin.NewUnexpectedID(id))
	}
}

// HelpUserInfo boxes the HelpUserInfoClass providing a helper.
type HelpUserInfoBox struct {
	UserInfo HelpUserInfoClass
}

// Decode implements bin.Decoder for HelpUserInfoBox.
func (b *HelpUserInfoBox) Decode(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("unable to decode HelpUserInfoBox to nil")
	}
	v, err := DecodeHelpUserInfo(buf)
	if err != nil {
		return fmt.Errorf("unable to decode boxed value: %w", err)
	}
	b.UserInfo = v
	return nil
}

// Encode implements bin.Encode for HelpUserInfoBox.
func (b *HelpUserInfoBox) Encode(buf *bin.Buffer) error {
	if b == nil || b.UserInfo == nil {
		return fmt.Errorf("unable to encode HelpUserInfoClass as nil")
	}
	return b.UserInfo.Encode(buf)
}
