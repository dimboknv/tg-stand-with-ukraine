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

// MessagesSearchRequest represents TL type `messages.search#a0fda762`.
// Gets back found messages
//
// See https://core.telegram.org/method/messages.search for reference.
type MessagesSearchRequest struct {
	// Flags, see TL conditional fields¹
	//
	// Links:
	//  1) https://core.telegram.org/mtproto/TL-combinators#conditional-fields
	Flags bin.Fields
	// User or chat, histories with which are searched, or (inputPeerEmpty)¹ constructor for
	// global search
	//
	// Links:
	//  1) https://core.telegram.org/constructor/inputPeerEmpty
	Peer InputPeerClass
	// Text search request
	Q string
	// Only return messages sent by the specified user ID
	//
	// Use SetFromID and GetFromID helpers.
	FromID InputPeerClass
	// Thread ID¹
	//
	// Links:
	//  1) https://core.telegram.org/api/threads
	//
	// Use SetTopMsgID and GetTopMsgID helpers.
	TopMsgID int
	// Filter to return only specified message types
	Filter MessagesFilterClass
	// If a positive value was transferred, only messages with a sending date bigger than the
	// transferred one will be returned
	MinDate int
	// If a positive value was transferred, only messages with a sending date smaller than
	// the transferred one will be returned
	MaxDate int
	// Only return messages starting from the specified message ID
	OffsetID int
	// Additional offset¹
	//
	// Links:
	//  1) https://core.telegram.org/api/offsets
	AddOffset int
	// Number of results to return¹
	//
	// Links:
	//  1) https://core.telegram.org/api/offsets
	Limit int
	// Maximum message ID to return¹
	//
	// Links:
	//  1) https://core.telegram.org/api/offsets
	MaxID int
	// Minimum message ID to return¹
	//
	// Links:
	//  1) https://core.telegram.org/api/offsets
	MinID int
	// Hash¹
	//
	// Links:
	//  1) https://core.telegram.org/api/offsets
	Hash int64
}

// MessagesSearchRequestTypeID is TL type id of MessagesSearchRequest.
const MessagesSearchRequestTypeID = 0xa0fda762

// Ensuring interfaces in compile-time for MessagesSearchRequest.
var (
	_ bin.Encoder     = &MessagesSearchRequest{}
	_ bin.Decoder     = &MessagesSearchRequest{}
	_ bin.BareEncoder = &MessagesSearchRequest{}
	_ bin.BareDecoder = &MessagesSearchRequest{}
)

func (s *MessagesSearchRequest) Zero() bool {
	if s == nil {
		return true
	}
	if !(s.Flags.Zero()) {
		return false
	}
	if !(s.Peer == nil) {
		return false
	}
	if !(s.Q == "") {
		return false
	}
	if !(s.FromID == nil) {
		return false
	}
	if !(s.TopMsgID == 0) {
		return false
	}
	if !(s.Filter == nil) {
		return false
	}
	if !(s.MinDate == 0) {
		return false
	}
	if !(s.MaxDate == 0) {
		return false
	}
	if !(s.OffsetID == 0) {
		return false
	}
	if !(s.AddOffset == 0) {
		return false
	}
	if !(s.Limit == 0) {
		return false
	}
	if !(s.MaxID == 0) {
		return false
	}
	if !(s.MinID == 0) {
		return false
	}
	if !(s.Hash == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (s *MessagesSearchRequest) String() string {
	if s == nil {
		return "MessagesSearchRequest(nil)"
	}
	type Alias MessagesSearchRequest
	return fmt.Sprintf("MessagesSearchRequest%+v", Alias(*s))
}

// FillFrom fills MessagesSearchRequest from given interface.
func (s *MessagesSearchRequest) FillFrom(from interface {
	GetPeer() (value InputPeerClass)
	GetQ() (value string)
	GetFromID() (value InputPeerClass, ok bool)
	GetTopMsgID() (value int, ok bool)
	GetFilter() (value MessagesFilterClass)
	GetMinDate() (value int)
	GetMaxDate() (value int)
	GetOffsetID() (value int)
	GetAddOffset() (value int)
	GetLimit() (value int)
	GetMaxID() (value int)
	GetMinID() (value int)
	GetHash() (value int64)
}) {
	s.Peer = from.GetPeer()
	s.Q = from.GetQ()
	if val, ok := from.GetFromID(); ok {
		s.FromID = val
	}

	if val, ok := from.GetTopMsgID(); ok {
		s.TopMsgID = val
	}

	s.Filter = from.GetFilter()
	s.MinDate = from.GetMinDate()
	s.MaxDate = from.GetMaxDate()
	s.OffsetID = from.GetOffsetID()
	s.AddOffset = from.GetAddOffset()
	s.Limit = from.GetLimit()
	s.MaxID = from.GetMaxID()
	s.MinID = from.GetMinID()
	s.Hash = from.GetHash()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*MessagesSearchRequest) TypeID() uint32 {
	return MessagesSearchRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*MessagesSearchRequest) TypeName() string {
	return "messages.search"
}

// TypeInfo returns info about TL type.
func (s *MessagesSearchRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "messages.search",
		ID:   MessagesSearchRequestTypeID,
	}
	if s == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Peer",
			SchemaName: "peer",
		},
		{
			Name:       "Q",
			SchemaName: "q",
		},
		{
			Name:       "FromID",
			SchemaName: "from_id",
			Null:       !s.Flags.Has(0),
		},
		{
			Name:       "TopMsgID",
			SchemaName: "top_msg_id",
			Null:       !s.Flags.Has(1),
		},
		{
			Name:       "Filter",
			SchemaName: "filter",
		},
		{
			Name:       "MinDate",
			SchemaName: "min_date",
		},
		{
			Name:       "MaxDate",
			SchemaName: "max_date",
		},
		{
			Name:       "OffsetID",
			SchemaName: "offset_id",
		},
		{
			Name:       "AddOffset",
			SchemaName: "add_offset",
		},
		{
			Name:       "Limit",
			SchemaName: "limit",
		},
		{
			Name:       "MaxID",
			SchemaName: "max_id",
		},
		{
			Name:       "MinID",
			SchemaName: "min_id",
		},
		{
			Name:       "Hash",
			SchemaName: "hash",
		},
	}
	return typ
}

// SetFlags sets flags for non-zero fields.
func (s *MessagesSearchRequest) SetFlags() {
	if !(s.FromID == nil) {
		s.Flags.Set(0)
	}
	if !(s.TopMsgID == 0) {
		s.Flags.Set(1)
	}
}

// Encode implements bin.Encoder.
func (s *MessagesSearchRequest) Encode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode messages.search#a0fda762 as nil")
	}
	b.PutID(MessagesSearchRequestTypeID)
	return s.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (s *MessagesSearchRequest) EncodeBare(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode messages.search#a0fda762 as nil")
	}
	s.SetFlags()
	if err := s.Flags.Encode(b); err != nil {
		return fmt.Errorf("unable to encode messages.search#a0fda762: field flags: %w", err)
	}
	if s.Peer == nil {
		return fmt.Errorf("unable to encode messages.search#a0fda762: field peer is nil")
	}
	if err := s.Peer.Encode(b); err != nil {
		return fmt.Errorf("unable to encode messages.search#a0fda762: field peer: %w", err)
	}
	b.PutString(s.Q)
	if s.Flags.Has(0) {
		if s.FromID == nil {
			return fmt.Errorf("unable to encode messages.search#a0fda762: field from_id is nil")
		}
		if err := s.FromID.Encode(b); err != nil {
			return fmt.Errorf("unable to encode messages.search#a0fda762: field from_id: %w", err)
		}
	}
	if s.Flags.Has(1) {
		b.PutInt(s.TopMsgID)
	}
	if s.Filter == nil {
		return fmt.Errorf("unable to encode messages.search#a0fda762: field filter is nil")
	}
	if err := s.Filter.Encode(b); err != nil {
		return fmt.Errorf("unable to encode messages.search#a0fda762: field filter: %w", err)
	}
	b.PutInt(s.MinDate)
	b.PutInt(s.MaxDate)
	b.PutInt(s.OffsetID)
	b.PutInt(s.AddOffset)
	b.PutInt(s.Limit)
	b.PutInt(s.MaxID)
	b.PutInt(s.MinID)
	b.PutLong(s.Hash)
	return nil
}

// Decode implements bin.Decoder.
func (s *MessagesSearchRequest) Decode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode messages.search#a0fda762 to nil")
	}
	if err := b.ConsumeID(MessagesSearchRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode messages.search#a0fda762: %w", err)
	}
	return s.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (s *MessagesSearchRequest) DecodeBare(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode messages.search#a0fda762 to nil")
	}
	{
		if err := s.Flags.Decode(b); err != nil {
			return fmt.Errorf("unable to decode messages.search#a0fda762: field flags: %w", err)
		}
	}
	{
		value, err := DecodeInputPeer(b)
		if err != nil {
			return fmt.Errorf("unable to decode messages.search#a0fda762: field peer: %w", err)
		}
		s.Peer = value
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode messages.search#a0fda762: field q: %w", err)
		}
		s.Q = value
	}
	if s.Flags.Has(0) {
		value, err := DecodeInputPeer(b)
		if err != nil {
			return fmt.Errorf("unable to decode messages.search#a0fda762: field from_id: %w", err)
		}
		s.FromID = value
	}
	if s.Flags.Has(1) {
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode messages.search#a0fda762: field top_msg_id: %w", err)
		}
		s.TopMsgID = value
	}
	{
		value, err := DecodeMessagesFilter(b)
		if err != nil {
			return fmt.Errorf("unable to decode messages.search#a0fda762: field filter: %w", err)
		}
		s.Filter = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode messages.search#a0fda762: field min_date: %w", err)
		}
		s.MinDate = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode messages.search#a0fda762: field max_date: %w", err)
		}
		s.MaxDate = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode messages.search#a0fda762: field offset_id: %w", err)
		}
		s.OffsetID = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode messages.search#a0fda762: field add_offset: %w", err)
		}
		s.AddOffset = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode messages.search#a0fda762: field limit: %w", err)
		}
		s.Limit = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode messages.search#a0fda762: field max_id: %w", err)
		}
		s.MaxID = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode messages.search#a0fda762: field min_id: %w", err)
		}
		s.MinID = value
	}
	{
		value, err := b.Long()
		if err != nil {
			return fmt.Errorf("unable to decode messages.search#a0fda762: field hash: %w", err)
		}
		s.Hash = value
	}
	return nil
}

// GetPeer returns value of Peer field.
func (s *MessagesSearchRequest) GetPeer() (value InputPeerClass) {
	if s == nil {
		return
	}
	return s.Peer
}

// GetQ returns value of Q field.
func (s *MessagesSearchRequest) GetQ() (value string) {
	if s == nil {
		return
	}
	return s.Q
}

// SetFromID sets value of FromID conditional field.
func (s *MessagesSearchRequest) SetFromID(value InputPeerClass) {
	s.Flags.Set(0)
	s.FromID = value
}

// GetFromID returns value of FromID conditional field and
// boolean which is true if field was set.
func (s *MessagesSearchRequest) GetFromID() (value InputPeerClass, ok bool) {
	if s == nil {
		return
	}
	if !s.Flags.Has(0) {
		return value, false
	}
	return s.FromID, true
}

// SetTopMsgID sets value of TopMsgID conditional field.
func (s *MessagesSearchRequest) SetTopMsgID(value int) {
	s.Flags.Set(1)
	s.TopMsgID = value
}

// GetTopMsgID returns value of TopMsgID conditional field and
// boolean which is true if field was set.
func (s *MessagesSearchRequest) GetTopMsgID() (value int, ok bool) {
	if s == nil {
		return
	}
	if !s.Flags.Has(1) {
		return value, false
	}
	return s.TopMsgID, true
}

// GetFilter returns value of Filter field.
func (s *MessagesSearchRequest) GetFilter() (value MessagesFilterClass) {
	if s == nil {
		return
	}
	return s.Filter
}

// GetMinDate returns value of MinDate field.
func (s *MessagesSearchRequest) GetMinDate() (value int) {
	if s == nil {
		return
	}
	return s.MinDate
}

// GetMaxDate returns value of MaxDate field.
func (s *MessagesSearchRequest) GetMaxDate() (value int) {
	if s == nil {
		return
	}
	return s.MaxDate
}

// GetOffsetID returns value of OffsetID field.
func (s *MessagesSearchRequest) GetOffsetID() (value int) {
	if s == nil {
		return
	}
	return s.OffsetID
}

// GetAddOffset returns value of AddOffset field.
func (s *MessagesSearchRequest) GetAddOffset() (value int) {
	if s == nil {
		return
	}
	return s.AddOffset
}

// GetLimit returns value of Limit field.
func (s *MessagesSearchRequest) GetLimit() (value int) {
	if s == nil {
		return
	}
	return s.Limit
}

// GetMaxID returns value of MaxID field.
func (s *MessagesSearchRequest) GetMaxID() (value int) {
	if s == nil {
		return
	}
	return s.MaxID
}

// GetMinID returns value of MinID field.
func (s *MessagesSearchRequest) GetMinID() (value int) {
	if s == nil {
		return
	}
	return s.MinID
}

// GetHash returns value of Hash field.
func (s *MessagesSearchRequest) GetHash() (value int64) {
	if s == nil {
		return
	}
	return s.Hash
}

// MessagesSearch invokes method messages.search#a0fda762 returning error if any.
// Gets back found messages
//
// Possible errors:
//  400 CHANNEL_INVALID: The provided channel is invalid.
//  400 CHANNEL_PRIVATE: You haven't joined this channel/supergroup.
//  400 CHAT_ADMIN_REQUIRED: You must be an admin in this chat to do this.
//  400 FROM_PEER_INVALID: The specified from_id is invalid.
//  400 INPUT_FILTER_INVALID: The specified filter is invalid.
//  400 INPUT_USER_DEACTIVATED: The specified user was deleted.
//  400 MSG_ID_INVALID: Invalid message ID provided.
//  400 PEER_ID_INVALID: The provided peer id is invalid.
//  400 PEER_ID_NOT_SUPPORTED: The provided peer ID is not supported.
//  400 SEARCH_QUERY_EMPTY: The search query is empty.
//  400 USER_ID_INVALID: The provided user ID is invalid.
//
// See https://core.telegram.org/method/messages.search for reference.
func (c *Client) MessagesSearch(ctx context.Context, request *MessagesSearchRequest) (MessagesMessagesClass, error) {
	var result MessagesMessagesBox

	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return nil, err
	}
	return result.Messages, nil
}
