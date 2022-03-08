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

// MessagesSendMultiMediaRequest represents TL type `messages.sendMultiMedia#f803138f`.
// Send an album or grouped media¹
//
// Links:
//  1) https://core.telegram.org/api/files#albums-grouped-media
//
// See https://core.telegram.org/method/messages.sendMultiMedia for reference.
type MessagesSendMultiMediaRequest struct {
	// Flags, see TL conditional fields¹
	//
	// Links:
	//  1) https://core.telegram.org/mtproto/TL-combinators#conditional-fields
	Flags bin.Fields
	// Whether to send the album silently (no notification triggered)
	Silent bool
	// Send in background?
	Background bool
	// Whether to clear drafts¹
	//
	// Links:
	//  1) https://core.telegram.org/api/drafts
	ClearDraft bool
	// Noforwards field of MessagesSendMultiMediaRequest.
	Noforwards bool
	// The destination chat
	Peer InputPeerClass
	// The message to reply to
	//
	// Use SetReplyToMsgID and GetReplyToMsgID helpers.
	ReplyToMsgID int
	// The medias to send
	MultiMedia []InputSingleMedia
	// Scheduled message date for scheduled messages
	//
	// Use SetScheduleDate and GetScheduleDate helpers.
	ScheduleDate int
	// SendAs field of MessagesSendMultiMediaRequest.
	//
	// Use SetSendAs and GetSendAs helpers.
	SendAs InputPeerClass
}

// MessagesSendMultiMediaRequestTypeID is TL type id of MessagesSendMultiMediaRequest.
const MessagesSendMultiMediaRequestTypeID = 0xf803138f

// Ensuring interfaces in compile-time for MessagesSendMultiMediaRequest.
var (
	_ bin.Encoder     = &MessagesSendMultiMediaRequest{}
	_ bin.Decoder     = &MessagesSendMultiMediaRequest{}
	_ bin.BareEncoder = &MessagesSendMultiMediaRequest{}
	_ bin.BareDecoder = &MessagesSendMultiMediaRequest{}
)

func (s *MessagesSendMultiMediaRequest) Zero() bool {
	if s == nil {
		return true
	}
	if !(s.Flags.Zero()) {
		return false
	}
	if !(s.Silent == false) {
		return false
	}
	if !(s.Background == false) {
		return false
	}
	if !(s.ClearDraft == false) {
		return false
	}
	if !(s.Noforwards == false) {
		return false
	}
	if !(s.Peer == nil) {
		return false
	}
	if !(s.ReplyToMsgID == 0) {
		return false
	}
	if !(s.MultiMedia == nil) {
		return false
	}
	if !(s.ScheduleDate == 0) {
		return false
	}
	if !(s.SendAs == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (s *MessagesSendMultiMediaRequest) String() string {
	if s == nil {
		return "MessagesSendMultiMediaRequest(nil)"
	}
	type Alias MessagesSendMultiMediaRequest
	return fmt.Sprintf("MessagesSendMultiMediaRequest%+v", Alias(*s))
}

// FillFrom fills MessagesSendMultiMediaRequest from given interface.
func (s *MessagesSendMultiMediaRequest) FillFrom(from interface {
	GetSilent() (value bool)
	GetBackground() (value bool)
	GetClearDraft() (value bool)
	GetNoforwards() (value bool)
	GetPeer() (value InputPeerClass)
	GetReplyToMsgID() (value int, ok bool)
	GetMultiMedia() (value []InputSingleMedia)
	GetScheduleDate() (value int, ok bool)
	GetSendAs() (value InputPeerClass, ok bool)
}) {
	s.Silent = from.GetSilent()
	s.Background = from.GetBackground()
	s.ClearDraft = from.GetClearDraft()
	s.Noforwards = from.GetNoforwards()
	s.Peer = from.GetPeer()
	if val, ok := from.GetReplyToMsgID(); ok {
		s.ReplyToMsgID = val
	}

	s.MultiMedia = from.GetMultiMedia()
	if val, ok := from.GetScheduleDate(); ok {
		s.ScheduleDate = val
	}

	if val, ok := from.GetSendAs(); ok {
		s.SendAs = val
	}

}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*MessagesSendMultiMediaRequest) TypeID() uint32 {
	return MessagesSendMultiMediaRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*MessagesSendMultiMediaRequest) TypeName() string {
	return "messages.sendMultiMedia"
}

// TypeInfo returns info about TL type.
func (s *MessagesSendMultiMediaRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "messages.sendMultiMedia",
		ID:   MessagesSendMultiMediaRequestTypeID,
	}
	if s == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Silent",
			SchemaName: "silent",
			Null:       !s.Flags.Has(5),
		},
		{
			Name:       "Background",
			SchemaName: "background",
			Null:       !s.Flags.Has(6),
		},
		{
			Name:       "ClearDraft",
			SchemaName: "clear_draft",
			Null:       !s.Flags.Has(7),
		},
		{
			Name:       "Noforwards",
			SchemaName: "noforwards",
			Null:       !s.Flags.Has(14),
		},
		{
			Name:       "Peer",
			SchemaName: "peer",
		},
		{
			Name:       "ReplyToMsgID",
			SchemaName: "reply_to_msg_id",
			Null:       !s.Flags.Has(0),
		},
		{
			Name:       "MultiMedia",
			SchemaName: "multi_media",
		},
		{
			Name:       "ScheduleDate",
			SchemaName: "schedule_date",
			Null:       !s.Flags.Has(10),
		},
		{
			Name:       "SendAs",
			SchemaName: "send_as",
			Null:       !s.Flags.Has(13),
		},
	}
	return typ
}

// SetFlags sets flags for non-zero fields.
func (s *MessagesSendMultiMediaRequest) SetFlags() {
	if !(s.Silent == false) {
		s.Flags.Set(5)
	}
	if !(s.Background == false) {
		s.Flags.Set(6)
	}
	if !(s.ClearDraft == false) {
		s.Flags.Set(7)
	}
	if !(s.Noforwards == false) {
		s.Flags.Set(14)
	}
	if !(s.ReplyToMsgID == 0) {
		s.Flags.Set(0)
	}
	if !(s.ScheduleDate == 0) {
		s.Flags.Set(10)
	}
	if !(s.SendAs == nil) {
		s.Flags.Set(13)
	}
}

// Encode implements bin.Encoder.
func (s *MessagesSendMultiMediaRequest) Encode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode messages.sendMultiMedia#f803138f as nil")
	}
	b.PutID(MessagesSendMultiMediaRequestTypeID)
	return s.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (s *MessagesSendMultiMediaRequest) EncodeBare(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode messages.sendMultiMedia#f803138f as nil")
	}
	s.SetFlags()
	if err := s.Flags.Encode(b); err != nil {
		return fmt.Errorf("unable to encode messages.sendMultiMedia#f803138f: field flags: %w", err)
	}
	if s.Peer == nil {
		return fmt.Errorf("unable to encode messages.sendMultiMedia#f803138f: field peer is nil")
	}
	if err := s.Peer.Encode(b); err != nil {
		return fmt.Errorf("unable to encode messages.sendMultiMedia#f803138f: field peer: %w", err)
	}
	if s.Flags.Has(0) {
		b.PutInt(s.ReplyToMsgID)
	}
	b.PutVectorHeader(len(s.MultiMedia))
	for idx, v := range s.MultiMedia {
		if err := v.Encode(b); err != nil {
			return fmt.Errorf("unable to encode messages.sendMultiMedia#f803138f: field multi_media element with index %d: %w", idx, err)
		}
	}
	if s.Flags.Has(10) {
		b.PutInt(s.ScheduleDate)
	}
	if s.Flags.Has(13) {
		if s.SendAs == nil {
			return fmt.Errorf("unable to encode messages.sendMultiMedia#f803138f: field send_as is nil")
		}
		if err := s.SendAs.Encode(b); err != nil {
			return fmt.Errorf("unable to encode messages.sendMultiMedia#f803138f: field send_as: %w", err)
		}
	}
	return nil
}

// Decode implements bin.Decoder.
func (s *MessagesSendMultiMediaRequest) Decode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode messages.sendMultiMedia#f803138f to nil")
	}
	if err := b.ConsumeID(MessagesSendMultiMediaRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode messages.sendMultiMedia#f803138f: %w", err)
	}
	return s.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (s *MessagesSendMultiMediaRequest) DecodeBare(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode messages.sendMultiMedia#f803138f to nil")
	}
	{
		if err := s.Flags.Decode(b); err != nil {
			return fmt.Errorf("unable to decode messages.sendMultiMedia#f803138f: field flags: %w", err)
		}
	}
	s.Silent = s.Flags.Has(5)
	s.Background = s.Flags.Has(6)
	s.ClearDraft = s.Flags.Has(7)
	s.Noforwards = s.Flags.Has(14)
	{
		value, err := DecodeInputPeer(b)
		if err != nil {
			return fmt.Errorf("unable to decode messages.sendMultiMedia#f803138f: field peer: %w", err)
		}
		s.Peer = value
	}
	if s.Flags.Has(0) {
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode messages.sendMultiMedia#f803138f: field reply_to_msg_id: %w", err)
		}
		s.ReplyToMsgID = value
	}
	{
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode messages.sendMultiMedia#f803138f: field multi_media: %w", err)
		}

		if headerLen > 0 {
			s.MultiMedia = make([]InputSingleMedia, 0, headerLen%bin.PreallocateLimit)
		}
		for idx := 0; idx < headerLen; idx++ {
			var value InputSingleMedia
			if err := value.Decode(b); err != nil {
				return fmt.Errorf("unable to decode messages.sendMultiMedia#f803138f: field multi_media: %w", err)
			}
			s.MultiMedia = append(s.MultiMedia, value)
		}
	}
	if s.Flags.Has(10) {
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode messages.sendMultiMedia#f803138f: field schedule_date: %w", err)
		}
		s.ScheduleDate = value
	}
	if s.Flags.Has(13) {
		value, err := DecodeInputPeer(b)
		if err != nil {
			return fmt.Errorf("unable to decode messages.sendMultiMedia#f803138f: field send_as: %w", err)
		}
		s.SendAs = value
	}
	return nil
}

// SetSilent sets value of Silent conditional field.
func (s *MessagesSendMultiMediaRequest) SetSilent(value bool) {
	if value {
		s.Flags.Set(5)
		s.Silent = true
	} else {
		s.Flags.Unset(5)
		s.Silent = false
	}
}

// GetSilent returns value of Silent conditional field.
func (s *MessagesSendMultiMediaRequest) GetSilent() (value bool) {
	if s == nil {
		return
	}
	return s.Flags.Has(5)
}

// SetBackground sets value of Background conditional field.
func (s *MessagesSendMultiMediaRequest) SetBackground(value bool) {
	if value {
		s.Flags.Set(6)
		s.Background = true
	} else {
		s.Flags.Unset(6)
		s.Background = false
	}
}

// GetBackground returns value of Background conditional field.
func (s *MessagesSendMultiMediaRequest) GetBackground() (value bool) {
	if s == nil {
		return
	}
	return s.Flags.Has(6)
}

// SetClearDraft sets value of ClearDraft conditional field.
func (s *MessagesSendMultiMediaRequest) SetClearDraft(value bool) {
	if value {
		s.Flags.Set(7)
		s.ClearDraft = true
	} else {
		s.Flags.Unset(7)
		s.ClearDraft = false
	}
}

// GetClearDraft returns value of ClearDraft conditional field.
func (s *MessagesSendMultiMediaRequest) GetClearDraft() (value bool) {
	if s == nil {
		return
	}
	return s.Flags.Has(7)
}

// SetNoforwards sets value of Noforwards conditional field.
func (s *MessagesSendMultiMediaRequest) SetNoforwards(value bool) {
	if value {
		s.Flags.Set(14)
		s.Noforwards = true
	} else {
		s.Flags.Unset(14)
		s.Noforwards = false
	}
}

// GetNoforwards returns value of Noforwards conditional field.
func (s *MessagesSendMultiMediaRequest) GetNoforwards() (value bool) {
	if s == nil {
		return
	}
	return s.Flags.Has(14)
}

// GetPeer returns value of Peer field.
func (s *MessagesSendMultiMediaRequest) GetPeer() (value InputPeerClass) {
	if s == nil {
		return
	}
	return s.Peer
}

// SetReplyToMsgID sets value of ReplyToMsgID conditional field.
func (s *MessagesSendMultiMediaRequest) SetReplyToMsgID(value int) {
	s.Flags.Set(0)
	s.ReplyToMsgID = value
}

// GetReplyToMsgID returns value of ReplyToMsgID conditional field and
// boolean which is true if field was set.
func (s *MessagesSendMultiMediaRequest) GetReplyToMsgID() (value int, ok bool) {
	if s == nil {
		return
	}
	if !s.Flags.Has(0) {
		return value, false
	}
	return s.ReplyToMsgID, true
}

// GetMultiMedia returns value of MultiMedia field.
func (s *MessagesSendMultiMediaRequest) GetMultiMedia() (value []InputSingleMedia) {
	if s == nil {
		return
	}
	return s.MultiMedia
}

// SetScheduleDate sets value of ScheduleDate conditional field.
func (s *MessagesSendMultiMediaRequest) SetScheduleDate(value int) {
	s.Flags.Set(10)
	s.ScheduleDate = value
}

// GetScheduleDate returns value of ScheduleDate conditional field and
// boolean which is true if field was set.
func (s *MessagesSendMultiMediaRequest) GetScheduleDate() (value int, ok bool) {
	if s == nil {
		return
	}
	if !s.Flags.Has(10) {
		return value, false
	}
	return s.ScheduleDate, true
}

// SetSendAs sets value of SendAs conditional field.
func (s *MessagesSendMultiMediaRequest) SetSendAs(value InputPeerClass) {
	s.Flags.Set(13)
	s.SendAs = value
}

// GetSendAs returns value of SendAs conditional field and
// boolean which is true if field was set.
func (s *MessagesSendMultiMediaRequest) GetSendAs() (value InputPeerClass, ok bool) {
	if s == nil {
		return
	}
	if !s.Flags.Has(13) {
		return value, false
	}
	return s.SendAs, true
}

// MessagesSendMultiMedia invokes method messages.sendMultiMedia#f803138f returning error if any.
// Send an album or grouped media¹
//
// Links:
//  1) https://core.telegram.org/api/files#albums-grouped-media
//
// Possible errors:
//  400 CHANNEL_PRIVATE: You haven't joined this channel/supergroup.
//  400 CHAT_ADMIN_REQUIRED: You must be an admin in this chat to do this.
//  400 MEDIA_CAPTION_TOO_LONG: The caption is too long.
//  400 MEDIA_EMPTY: The provided media object is invalid.
//  400 MEDIA_INVALID: Media invalid.
//  400 MULTI_MEDIA_TOO_LONG: Too many media files for album.
//  400 PEER_ID_INVALID: The provided peer id is invalid.
//  400 RANDOM_ID_EMPTY: Random ID empty.
//  400 SCHEDULE_DATE_TOO_LATE: You can't schedule a message this far in the future.
//  400 SCHEDULE_TOO_MUCH: There are too many scheduled messages.
//  420 SLOWMODE_WAIT_X: Slowmode is enabled in this chat: wait X seconds before sending another message to this chat.
//
// See https://core.telegram.org/method/messages.sendMultiMedia for reference.
// Can be used by bots.
func (c *Client) MessagesSendMultiMedia(ctx context.Context, request *MessagesSendMultiMediaRequest) (UpdatesClass, error) {
	var result UpdatesBox

	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return nil, err
	}
	return result.Updates, nil
}
