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

// DialogPeer represents TL type `dialogPeer#e56dbf05`.
// Peer
//
// See https://core.telegram.org/constructor/dialogPeer for reference.
type DialogPeer struct {
	// Peer
	Peer PeerClass
}

// DialogPeerTypeID is TL type id of DialogPeer.
const DialogPeerTypeID = 0xe56dbf05

// construct implements constructor of DialogPeerClass.
func (d DialogPeer) construct() DialogPeerClass { return &d }

// Ensuring interfaces in compile-time for DialogPeer.
var (
	_ bin.Encoder     = &DialogPeer{}
	_ bin.Decoder     = &DialogPeer{}
	_ bin.BareEncoder = &DialogPeer{}
	_ bin.BareDecoder = &DialogPeer{}

	_ DialogPeerClass = &DialogPeer{}
)

func (d *DialogPeer) Zero() bool {
	if d == nil {
		return true
	}
	if !(d.Peer == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (d *DialogPeer) String() string {
	if d == nil {
		return "DialogPeer(nil)"
	}
	type Alias DialogPeer
	return fmt.Sprintf("DialogPeer%+v", Alias(*d))
}

// FillFrom fills DialogPeer from given interface.
func (d *DialogPeer) FillFrom(from interface {
	GetPeer() (value PeerClass)
}) {
	d.Peer = from.GetPeer()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*DialogPeer) TypeID() uint32 {
	return DialogPeerTypeID
}

// TypeName returns name of type in TL schema.
func (*DialogPeer) TypeName() string {
	return "dialogPeer"
}

// TypeInfo returns info about TL type.
func (d *DialogPeer) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "dialogPeer",
		ID:   DialogPeerTypeID,
	}
	if d == nil {
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
func (d *DialogPeer) Encode(b *bin.Buffer) error {
	if d == nil {
		return fmt.Errorf("can't encode dialogPeer#e56dbf05 as nil")
	}
	b.PutID(DialogPeerTypeID)
	return d.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (d *DialogPeer) EncodeBare(b *bin.Buffer) error {
	if d == nil {
		return fmt.Errorf("can't encode dialogPeer#e56dbf05 as nil")
	}
	if d.Peer == nil {
		return fmt.Errorf("unable to encode dialogPeer#e56dbf05: field peer is nil")
	}
	if err := d.Peer.Encode(b); err != nil {
		return fmt.Errorf("unable to encode dialogPeer#e56dbf05: field peer: %w", err)
	}
	return nil
}

// Decode implements bin.Decoder.
func (d *DialogPeer) Decode(b *bin.Buffer) error {
	if d == nil {
		return fmt.Errorf("can't decode dialogPeer#e56dbf05 to nil")
	}
	if err := b.ConsumeID(DialogPeerTypeID); err != nil {
		return fmt.Errorf("unable to decode dialogPeer#e56dbf05: %w", err)
	}
	return d.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (d *DialogPeer) DecodeBare(b *bin.Buffer) error {
	if d == nil {
		return fmt.Errorf("can't decode dialogPeer#e56dbf05 to nil")
	}
	{
		value, err := DecodePeer(b)
		if err != nil {
			return fmt.Errorf("unable to decode dialogPeer#e56dbf05: field peer: %w", err)
		}
		d.Peer = value
	}
	return nil
}

// GetPeer returns value of Peer field.
func (d *DialogPeer) GetPeer() (value PeerClass) {
	if d == nil {
		return
	}
	return d.Peer
}

// DialogPeerFolder represents TL type `dialogPeerFolder#514519e2`.
// Peer folder¹
//
// Links:
//  1) https://core.telegram.org/api/folders#peer-folders
//
// See https://core.telegram.org/constructor/dialogPeerFolder for reference.
type DialogPeerFolder struct {
	// Peer folder ID, for more info click here¹
	//
	// Links:
	//  1) https://core.telegram.org/api/folders#peer-folders
	FolderID int
}

// DialogPeerFolderTypeID is TL type id of DialogPeerFolder.
const DialogPeerFolderTypeID = 0x514519e2

// construct implements constructor of DialogPeerClass.
func (d DialogPeerFolder) construct() DialogPeerClass { return &d }

// Ensuring interfaces in compile-time for DialogPeerFolder.
var (
	_ bin.Encoder     = &DialogPeerFolder{}
	_ bin.Decoder     = &DialogPeerFolder{}
	_ bin.BareEncoder = &DialogPeerFolder{}
	_ bin.BareDecoder = &DialogPeerFolder{}

	_ DialogPeerClass = &DialogPeerFolder{}
)

func (d *DialogPeerFolder) Zero() bool {
	if d == nil {
		return true
	}
	if !(d.FolderID == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (d *DialogPeerFolder) String() string {
	if d == nil {
		return "DialogPeerFolder(nil)"
	}
	type Alias DialogPeerFolder
	return fmt.Sprintf("DialogPeerFolder%+v", Alias(*d))
}

// FillFrom fills DialogPeerFolder from given interface.
func (d *DialogPeerFolder) FillFrom(from interface {
	GetFolderID() (value int)
}) {
	d.FolderID = from.GetFolderID()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*DialogPeerFolder) TypeID() uint32 {
	return DialogPeerFolderTypeID
}

// TypeName returns name of type in TL schema.
func (*DialogPeerFolder) TypeName() string {
	return "dialogPeerFolder"
}

// TypeInfo returns info about TL type.
func (d *DialogPeerFolder) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "dialogPeerFolder",
		ID:   DialogPeerFolderTypeID,
	}
	if d == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "FolderID",
			SchemaName: "folder_id",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (d *DialogPeerFolder) Encode(b *bin.Buffer) error {
	if d == nil {
		return fmt.Errorf("can't encode dialogPeerFolder#514519e2 as nil")
	}
	b.PutID(DialogPeerFolderTypeID)
	return d.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (d *DialogPeerFolder) EncodeBare(b *bin.Buffer) error {
	if d == nil {
		return fmt.Errorf("can't encode dialogPeerFolder#514519e2 as nil")
	}
	b.PutInt(d.FolderID)
	return nil
}

// Decode implements bin.Decoder.
func (d *DialogPeerFolder) Decode(b *bin.Buffer) error {
	if d == nil {
		return fmt.Errorf("can't decode dialogPeerFolder#514519e2 to nil")
	}
	if err := b.ConsumeID(DialogPeerFolderTypeID); err != nil {
		return fmt.Errorf("unable to decode dialogPeerFolder#514519e2: %w", err)
	}
	return d.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (d *DialogPeerFolder) DecodeBare(b *bin.Buffer) error {
	if d == nil {
		return fmt.Errorf("can't decode dialogPeerFolder#514519e2 to nil")
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode dialogPeerFolder#514519e2: field folder_id: %w", err)
		}
		d.FolderID = value
	}
	return nil
}

// GetFolderID returns value of FolderID field.
func (d *DialogPeerFolder) GetFolderID() (value int) {
	if d == nil {
		return
	}
	return d.FolderID
}

// DialogPeerClassName is schema name of DialogPeerClass.
const DialogPeerClassName = "DialogPeer"

// DialogPeerClass represents DialogPeer generic type.
//
// See https://core.telegram.org/type/DialogPeer for reference.
//
// Example:
//  g, err := tg.DecodeDialogPeer(buf)
//  if err != nil {
//      panic(err)
//  }
//  switch v := g.(type) {
//  case *tg.DialogPeer: // dialogPeer#e56dbf05
//  case *tg.DialogPeerFolder: // dialogPeerFolder#514519e2
//  default: panic(v)
//  }
type DialogPeerClass interface {
	bin.Encoder
	bin.Decoder
	bin.BareEncoder
	bin.BareDecoder
	construct() DialogPeerClass

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
}

// AsInput tries to map DialogPeerFolder to InputDialogPeerFolder.
func (d *DialogPeerFolder) AsInput() *InputDialogPeerFolder {
	value := new(InputDialogPeerFolder)
	value.FolderID = d.GetFolderID()

	return value
}

// DecodeDialogPeer implements binary de-serialization for DialogPeerClass.
func DecodeDialogPeer(buf *bin.Buffer) (DialogPeerClass, error) {
	id, err := buf.PeekID()
	if err != nil {
		return nil, err
	}
	switch id {
	case DialogPeerTypeID:
		// Decoding dialogPeer#e56dbf05.
		v := DialogPeer{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode DialogPeerClass: %w", err)
		}
		return &v, nil
	case DialogPeerFolderTypeID:
		// Decoding dialogPeerFolder#514519e2.
		v := DialogPeerFolder{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode DialogPeerClass: %w", err)
		}
		return &v, nil
	default:
		return nil, fmt.Errorf("unable to decode DialogPeerClass: %w", bin.NewUnexpectedID(id))
	}
}

// DialogPeer boxes the DialogPeerClass providing a helper.
type DialogPeerBox struct {
	DialogPeer DialogPeerClass
}

// Decode implements bin.Decoder for DialogPeerBox.
func (b *DialogPeerBox) Decode(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("unable to decode DialogPeerBox to nil")
	}
	v, err := DecodeDialogPeer(buf)
	if err != nil {
		return fmt.Errorf("unable to decode boxed value: %w", err)
	}
	b.DialogPeer = v
	return nil
}

// Encode implements bin.Encode for DialogPeerBox.
func (b *DialogPeerBox) Encode(buf *bin.Buffer) error {
	if b == nil || b.DialogPeer == nil {
		return fmt.Errorf("unable to encode DialogPeerClass as nil")
	}
	return b.DialogPeer.Encode(buf)
}