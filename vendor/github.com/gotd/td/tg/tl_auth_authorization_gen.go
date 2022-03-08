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

// AuthAuthorization represents TL type `auth.authorization#33fb7bb8`.
// Contains user authorization info.
//
// See https://core.telegram.org/constructor/auth.authorization for reference.
type AuthAuthorization struct {
	// Flags, see TL conditional fields¹
	//
	// Links:
	//  1) https://core.telegram.org/mtproto/TL-combinators#conditional-fields
	Flags bin.Fields
	// SetupPasswordRequired field of AuthAuthorization.
	SetupPasswordRequired bool
	// OtherwiseReloginDays field of AuthAuthorization.
	//
	// Use SetOtherwiseReloginDays and GetOtherwiseReloginDays helpers.
	OtherwiseReloginDays int
	// Temporary passport¹ sessions
	//
	// Links:
	//  1) https://core.telegram.org/passport
	//
	// Use SetTmpSessions and GetTmpSessions helpers.
	TmpSessions int
	// Info on authorized user
	User UserClass
}

// AuthAuthorizationTypeID is TL type id of AuthAuthorization.
const AuthAuthorizationTypeID = 0x33fb7bb8

// construct implements constructor of AuthAuthorizationClass.
func (a AuthAuthorization) construct() AuthAuthorizationClass { return &a }

// Ensuring interfaces in compile-time for AuthAuthorization.
var (
	_ bin.Encoder     = &AuthAuthorization{}
	_ bin.Decoder     = &AuthAuthorization{}
	_ bin.BareEncoder = &AuthAuthorization{}
	_ bin.BareDecoder = &AuthAuthorization{}

	_ AuthAuthorizationClass = &AuthAuthorization{}
)

func (a *AuthAuthorization) Zero() bool {
	if a == nil {
		return true
	}
	if !(a.Flags.Zero()) {
		return false
	}
	if !(a.SetupPasswordRequired == false) {
		return false
	}
	if !(a.OtherwiseReloginDays == 0) {
		return false
	}
	if !(a.TmpSessions == 0) {
		return false
	}
	if !(a.User == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (a *AuthAuthorization) String() string {
	if a == nil {
		return "AuthAuthorization(nil)"
	}
	type Alias AuthAuthorization
	return fmt.Sprintf("AuthAuthorization%+v", Alias(*a))
}

// FillFrom fills AuthAuthorization from given interface.
func (a *AuthAuthorization) FillFrom(from interface {
	GetSetupPasswordRequired() (value bool)
	GetOtherwiseReloginDays() (value int, ok bool)
	GetTmpSessions() (value int, ok bool)
	GetUser() (value UserClass)
}) {
	a.SetupPasswordRequired = from.GetSetupPasswordRequired()
	if val, ok := from.GetOtherwiseReloginDays(); ok {
		a.OtherwiseReloginDays = val
	}

	if val, ok := from.GetTmpSessions(); ok {
		a.TmpSessions = val
	}

	a.User = from.GetUser()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*AuthAuthorization) TypeID() uint32 {
	return AuthAuthorizationTypeID
}

// TypeName returns name of type in TL schema.
func (*AuthAuthorization) TypeName() string {
	return "auth.authorization"
}

// TypeInfo returns info about TL type.
func (a *AuthAuthorization) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "auth.authorization",
		ID:   AuthAuthorizationTypeID,
	}
	if a == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "SetupPasswordRequired",
			SchemaName: "setup_password_required",
			Null:       !a.Flags.Has(1),
		},
		{
			Name:       "OtherwiseReloginDays",
			SchemaName: "otherwise_relogin_days",
			Null:       !a.Flags.Has(1),
		},
		{
			Name:       "TmpSessions",
			SchemaName: "tmp_sessions",
			Null:       !a.Flags.Has(0),
		},
		{
			Name:       "User",
			SchemaName: "user",
		},
	}
	return typ
}

// SetFlags sets flags for non-zero fields.
func (a *AuthAuthorization) SetFlags() {
	if !(a.SetupPasswordRequired == false) {
		a.Flags.Set(1)
	}
	if !(a.OtherwiseReloginDays == 0) {
		a.Flags.Set(1)
	}
	if !(a.TmpSessions == 0) {
		a.Flags.Set(0)
	}
}

// Encode implements bin.Encoder.
func (a *AuthAuthorization) Encode(b *bin.Buffer) error {
	if a == nil {
		return fmt.Errorf("can't encode auth.authorization#33fb7bb8 as nil")
	}
	b.PutID(AuthAuthorizationTypeID)
	return a.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (a *AuthAuthorization) EncodeBare(b *bin.Buffer) error {
	if a == nil {
		return fmt.Errorf("can't encode auth.authorization#33fb7bb8 as nil")
	}
	a.SetFlags()
	if err := a.Flags.Encode(b); err != nil {
		return fmt.Errorf("unable to encode auth.authorization#33fb7bb8: field flags: %w", err)
	}
	if a.Flags.Has(1) {
		b.PutInt(a.OtherwiseReloginDays)
	}
	if a.Flags.Has(0) {
		b.PutInt(a.TmpSessions)
	}
	if a.User == nil {
		return fmt.Errorf("unable to encode auth.authorization#33fb7bb8: field user is nil")
	}
	if err := a.User.Encode(b); err != nil {
		return fmt.Errorf("unable to encode auth.authorization#33fb7bb8: field user: %w", err)
	}
	return nil
}

// Decode implements bin.Decoder.
func (a *AuthAuthorization) Decode(b *bin.Buffer) error {
	if a == nil {
		return fmt.Errorf("can't decode auth.authorization#33fb7bb8 to nil")
	}
	if err := b.ConsumeID(AuthAuthorizationTypeID); err != nil {
		return fmt.Errorf("unable to decode auth.authorization#33fb7bb8: %w", err)
	}
	return a.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (a *AuthAuthorization) DecodeBare(b *bin.Buffer) error {
	if a == nil {
		return fmt.Errorf("can't decode auth.authorization#33fb7bb8 to nil")
	}
	{
		if err := a.Flags.Decode(b); err != nil {
			return fmt.Errorf("unable to decode auth.authorization#33fb7bb8: field flags: %w", err)
		}
	}
	a.SetupPasswordRequired = a.Flags.Has(1)
	if a.Flags.Has(1) {
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode auth.authorization#33fb7bb8: field otherwise_relogin_days: %w", err)
		}
		a.OtherwiseReloginDays = value
	}
	if a.Flags.Has(0) {
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode auth.authorization#33fb7bb8: field tmp_sessions: %w", err)
		}
		a.TmpSessions = value
	}
	{
		value, err := DecodeUser(b)
		if err != nil {
			return fmt.Errorf("unable to decode auth.authorization#33fb7bb8: field user: %w", err)
		}
		a.User = value
	}
	return nil
}

// SetSetupPasswordRequired sets value of SetupPasswordRequired conditional field.
func (a *AuthAuthorization) SetSetupPasswordRequired(value bool) {
	if value {
		a.Flags.Set(1)
		a.SetupPasswordRequired = true
	} else {
		a.Flags.Unset(1)
		a.SetupPasswordRequired = false
	}
}

// GetSetupPasswordRequired returns value of SetupPasswordRequired conditional field.
func (a *AuthAuthorization) GetSetupPasswordRequired() (value bool) {
	if a == nil {
		return
	}
	return a.Flags.Has(1)
}

// SetOtherwiseReloginDays sets value of OtherwiseReloginDays conditional field.
func (a *AuthAuthorization) SetOtherwiseReloginDays(value int) {
	a.Flags.Set(1)
	a.OtherwiseReloginDays = value
}

// GetOtherwiseReloginDays returns value of OtherwiseReloginDays conditional field and
// boolean which is true if field was set.
func (a *AuthAuthorization) GetOtherwiseReloginDays() (value int, ok bool) {
	if a == nil {
		return
	}
	if !a.Flags.Has(1) {
		return value, false
	}
	return a.OtherwiseReloginDays, true
}

// SetTmpSessions sets value of TmpSessions conditional field.
func (a *AuthAuthorization) SetTmpSessions(value int) {
	a.Flags.Set(0)
	a.TmpSessions = value
}

// GetTmpSessions returns value of TmpSessions conditional field and
// boolean which is true if field was set.
func (a *AuthAuthorization) GetTmpSessions() (value int, ok bool) {
	if a == nil {
		return
	}
	if !a.Flags.Has(0) {
		return value, false
	}
	return a.TmpSessions, true
}

// GetUser returns value of User field.
func (a *AuthAuthorization) GetUser() (value UserClass) {
	if a == nil {
		return
	}
	return a.User
}

// AuthAuthorizationSignUpRequired represents TL type `auth.authorizationSignUpRequired#44747e9a`.
// An account with this phone number doesn't exist on telegram: the user has to enter
// basic information and sign up¹
//
// Links:
//  1) https://core.telegram.org/api/auth
//
// See https://core.telegram.org/constructor/auth.authorizationSignUpRequired for reference.
type AuthAuthorizationSignUpRequired struct {
	// Flags, see TL conditional fields¹
	//
	// Links:
	//  1) https://core.telegram.org/mtproto/TL-combinators#conditional-fields
	Flags bin.Fields
	// Telegram's terms of service: the user must read and accept the terms of service before
	// signing up to telegram
	//
	// Use SetTermsOfService and GetTermsOfService helpers.
	TermsOfService HelpTermsOfService
}

// AuthAuthorizationSignUpRequiredTypeID is TL type id of AuthAuthorizationSignUpRequired.
const AuthAuthorizationSignUpRequiredTypeID = 0x44747e9a

// construct implements constructor of AuthAuthorizationClass.
func (a AuthAuthorizationSignUpRequired) construct() AuthAuthorizationClass { return &a }

// Ensuring interfaces in compile-time for AuthAuthorizationSignUpRequired.
var (
	_ bin.Encoder     = &AuthAuthorizationSignUpRequired{}
	_ bin.Decoder     = &AuthAuthorizationSignUpRequired{}
	_ bin.BareEncoder = &AuthAuthorizationSignUpRequired{}
	_ bin.BareDecoder = &AuthAuthorizationSignUpRequired{}

	_ AuthAuthorizationClass = &AuthAuthorizationSignUpRequired{}
)

func (a *AuthAuthorizationSignUpRequired) Zero() bool {
	if a == nil {
		return true
	}
	if !(a.Flags.Zero()) {
		return false
	}
	if !(a.TermsOfService.Zero()) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (a *AuthAuthorizationSignUpRequired) String() string {
	if a == nil {
		return "AuthAuthorizationSignUpRequired(nil)"
	}
	type Alias AuthAuthorizationSignUpRequired
	return fmt.Sprintf("AuthAuthorizationSignUpRequired%+v", Alias(*a))
}

// FillFrom fills AuthAuthorizationSignUpRequired from given interface.
func (a *AuthAuthorizationSignUpRequired) FillFrom(from interface {
	GetTermsOfService() (value HelpTermsOfService, ok bool)
}) {
	if val, ok := from.GetTermsOfService(); ok {
		a.TermsOfService = val
	}

}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*AuthAuthorizationSignUpRequired) TypeID() uint32 {
	return AuthAuthorizationSignUpRequiredTypeID
}

// TypeName returns name of type in TL schema.
func (*AuthAuthorizationSignUpRequired) TypeName() string {
	return "auth.authorizationSignUpRequired"
}

// TypeInfo returns info about TL type.
func (a *AuthAuthorizationSignUpRequired) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "auth.authorizationSignUpRequired",
		ID:   AuthAuthorizationSignUpRequiredTypeID,
	}
	if a == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "TermsOfService",
			SchemaName: "terms_of_service",
			Null:       !a.Flags.Has(0),
		},
	}
	return typ
}

// SetFlags sets flags for non-zero fields.
func (a *AuthAuthorizationSignUpRequired) SetFlags() {
	if !(a.TermsOfService.Zero()) {
		a.Flags.Set(0)
	}
}

// Encode implements bin.Encoder.
func (a *AuthAuthorizationSignUpRequired) Encode(b *bin.Buffer) error {
	if a == nil {
		return fmt.Errorf("can't encode auth.authorizationSignUpRequired#44747e9a as nil")
	}
	b.PutID(AuthAuthorizationSignUpRequiredTypeID)
	return a.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (a *AuthAuthorizationSignUpRequired) EncodeBare(b *bin.Buffer) error {
	if a == nil {
		return fmt.Errorf("can't encode auth.authorizationSignUpRequired#44747e9a as nil")
	}
	a.SetFlags()
	if err := a.Flags.Encode(b); err != nil {
		return fmt.Errorf("unable to encode auth.authorizationSignUpRequired#44747e9a: field flags: %w", err)
	}
	if a.Flags.Has(0) {
		if err := a.TermsOfService.Encode(b); err != nil {
			return fmt.Errorf("unable to encode auth.authorizationSignUpRequired#44747e9a: field terms_of_service: %w", err)
		}
	}
	return nil
}

// Decode implements bin.Decoder.
func (a *AuthAuthorizationSignUpRequired) Decode(b *bin.Buffer) error {
	if a == nil {
		return fmt.Errorf("can't decode auth.authorizationSignUpRequired#44747e9a to nil")
	}
	if err := b.ConsumeID(AuthAuthorizationSignUpRequiredTypeID); err != nil {
		return fmt.Errorf("unable to decode auth.authorizationSignUpRequired#44747e9a: %w", err)
	}
	return a.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (a *AuthAuthorizationSignUpRequired) DecodeBare(b *bin.Buffer) error {
	if a == nil {
		return fmt.Errorf("can't decode auth.authorizationSignUpRequired#44747e9a to nil")
	}
	{
		if err := a.Flags.Decode(b); err != nil {
			return fmt.Errorf("unable to decode auth.authorizationSignUpRequired#44747e9a: field flags: %w", err)
		}
	}
	if a.Flags.Has(0) {
		if err := a.TermsOfService.Decode(b); err != nil {
			return fmt.Errorf("unable to decode auth.authorizationSignUpRequired#44747e9a: field terms_of_service: %w", err)
		}
	}
	return nil
}

// SetTermsOfService sets value of TermsOfService conditional field.
func (a *AuthAuthorizationSignUpRequired) SetTermsOfService(value HelpTermsOfService) {
	a.Flags.Set(0)
	a.TermsOfService = value
}

// GetTermsOfService returns value of TermsOfService conditional field and
// boolean which is true if field was set.
func (a *AuthAuthorizationSignUpRequired) GetTermsOfService() (value HelpTermsOfService, ok bool) {
	if a == nil {
		return
	}
	if !a.Flags.Has(0) {
		return value, false
	}
	return a.TermsOfService, true
}

// AuthAuthorizationClassName is schema name of AuthAuthorizationClass.
const AuthAuthorizationClassName = "auth.Authorization"

// AuthAuthorizationClass represents auth.Authorization generic type.
//
// See https://core.telegram.org/type/auth.Authorization for reference.
//
// Example:
//  g, err := tg.DecodeAuthAuthorization(buf)
//  if err != nil {
//      panic(err)
//  }
//  switch v := g.(type) {
//  case *tg.AuthAuthorization: // auth.authorization#33fb7bb8
//  case *tg.AuthAuthorizationSignUpRequired: // auth.authorizationSignUpRequired#44747e9a
//  default: panic(v)
//  }
type AuthAuthorizationClass interface {
	bin.Encoder
	bin.Decoder
	bin.BareEncoder
	bin.BareDecoder
	construct() AuthAuthorizationClass

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

// DecodeAuthAuthorization implements binary de-serialization for AuthAuthorizationClass.
func DecodeAuthAuthorization(buf *bin.Buffer) (AuthAuthorizationClass, error) {
	id, err := buf.PeekID()
	if err != nil {
		return nil, err
	}
	switch id {
	case AuthAuthorizationTypeID:
		// Decoding auth.authorization#33fb7bb8.
		v := AuthAuthorization{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode AuthAuthorizationClass: %w", err)
		}
		return &v, nil
	case AuthAuthorizationSignUpRequiredTypeID:
		// Decoding auth.authorizationSignUpRequired#44747e9a.
		v := AuthAuthorizationSignUpRequired{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode AuthAuthorizationClass: %w", err)
		}
		return &v, nil
	default:
		return nil, fmt.Errorf("unable to decode AuthAuthorizationClass: %w", bin.NewUnexpectedID(id))
	}
}

// AuthAuthorization boxes the AuthAuthorizationClass providing a helper.
type AuthAuthorizationBox struct {
	Authorization AuthAuthorizationClass
}

// Decode implements bin.Decoder for AuthAuthorizationBox.
func (b *AuthAuthorizationBox) Decode(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("unable to decode AuthAuthorizationBox to nil")
	}
	v, err := DecodeAuthAuthorization(buf)
	if err != nil {
		return fmt.Errorf("unable to decode boxed value: %w", err)
	}
	b.Authorization = v
	return nil
}

// Encode implements bin.Encode for AuthAuthorizationBox.
func (b *AuthAuthorizationBox) Encode(buf *bin.Buffer) error {
	if b == nil || b.Authorization == nil {
		return fmt.Errorf("unable to encode AuthAuthorizationClass as nil")
	}
	return b.Authorization.Encode(buf)
}