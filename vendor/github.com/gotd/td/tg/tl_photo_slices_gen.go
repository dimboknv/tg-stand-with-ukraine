//go:build !no_gotd_slices
// +build !no_gotd_slices

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

// PhotoClassArray is adapter for slice of PhotoClass.
type PhotoClassArray []PhotoClass

// Sort sorts slice of PhotoClass.
func (s PhotoClassArray) Sort(less func(a, b PhotoClass) bool) PhotoClassArray {
	sort.Slice(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// SortStable sorts slice of PhotoClass.
func (s PhotoClassArray) SortStable(less func(a, b PhotoClass) bool) PhotoClassArray {
	sort.SliceStable(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// Retain filters in-place slice of PhotoClass.
func (s PhotoClassArray) Retain(keep func(x PhotoClass) bool) PhotoClassArray {
	n := 0
	for _, x := range s {
		if keep(x) {
			s[n] = x
			n++
		}
	}
	s = s[:n]

	return s
}

// First returns first element of slice (if exists).
func (s PhotoClassArray) First() (v PhotoClass, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[0], true
}

// Last returns last element of slice (if exists).
func (s PhotoClassArray) Last() (v PhotoClass, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[len(s)-1], true
}

// PopFirst returns first element of slice (if exists) and deletes it.
func (s *PhotoClassArray) PopFirst() (v PhotoClass, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[0]

	// Delete by index from SliceTricks.
	copy(a[0:], a[1:])
	var zero PhotoClass
	a[len(a)-1] = zero
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// Pop returns last element of slice (if exists) and deletes it.
func (s *PhotoClassArray) Pop() (v PhotoClass, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[len(a)-1]
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// SortByID sorts slice of PhotoClass by ID.
func (s PhotoClassArray) SortByID() PhotoClassArray {
	return s.Sort(func(a, b PhotoClass) bool {
		return a.GetID() < b.GetID()
	})
}

// SortStableByID sorts slice of PhotoClass by ID.
func (s PhotoClassArray) SortStableByID() PhotoClassArray {
	return s.SortStable(func(a, b PhotoClass) bool {
		return a.GetID() < b.GetID()
	})
}

// FillPhotoEmptyMap fills only PhotoEmpty constructors to given map.
func (s PhotoClassArray) FillPhotoEmptyMap(to map[int64]*PhotoEmpty) {
	for _, elem := range s {
		value, ok := elem.(*PhotoEmpty)
		if !ok {
			continue
		}
		to[value.GetID()] = value
	}
}

// PhotoEmptyToMap collects only PhotoEmpty constructors to map.
func (s PhotoClassArray) PhotoEmptyToMap() map[int64]*PhotoEmpty {
	r := make(map[int64]*PhotoEmpty, len(s))
	s.FillPhotoEmptyMap(r)
	return r
}

// AsPhotoEmpty returns copy with only PhotoEmpty constructors.
func (s PhotoClassArray) AsPhotoEmpty() (to PhotoEmptyArray) {
	for _, elem := range s {
		value, ok := elem.(*PhotoEmpty)
		if !ok {
			continue
		}
		to = append(to, *value)
	}

	return to
}

// FillPhotoMap fills only Photo constructors to given map.
func (s PhotoClassArray) FillPhotoMap(to map[int64]*Photo) {
	for _, elem := range s {
		value, ok := elem.(*Photo)
		if !ok {
			continue
		}
		to[value.GetID()] = value
	}
}

// PhotoToMap collects only Photo constructors to map.
func (s PhotoClassArray) PhotoToMap() map[int64]*Photo {
	r := make(map[int64]*Photo, len(s))
	s.FillPhotoMap(r)
	return r
}

// AsPhoto returns copy with only Photo constructors.
func (s PhotoClassArray) AsPhoto() (to PhotoArray) {
	for _, elem := range s {
		value, ok := elem.(*Photo)
		if !ok {
			continue
		}
		to = append(to, *value)
	}

	return to
}

// FillNotEmptyMap fills only NotEmpty constructors to given map.
func (s PhotoClassArray) FillNotEmptyMap(to map[int64]*Photo) {
	for _, elem := range s {
		value, ok := elem.AsNotEmpty()
		if !ok {
			continue
		}
		to[value.GetID()] = value
	}
}

// NotEmptyToMap collects only NotEmpty constructors to map.
func (s PhotoClassArray) NotEmptyToMap() map[int64]*Photo {
	r := make(map[int64]*Photo, len(s))
	s.FillNotEmptyMap(r)
	return r
}

// AppendOnlyNotEmpty appends only NotEmpty constructors to
// given slice.
func (s PhotoClassArray) AppendOnlyNotEmpty(to []*Photo) []*Photo {
	for _, elem := range s {
		value, ok := elem.AsNotEmpty()
		if !ok {
			continue
		}
		to = append(to, value)
	}

	return to
}

// AsNotEmpty returns copy with only NotEmpty constructors.
func (s PhotoClassArray) AsNotEmpty() (to []*Photo) {
	return s.AppendOnlyNotEmpty(to)
}

// FirstAsNotEmpty returns first element of slice (if exists).
func (s PhotoClassArray) FirstAsNotEmpty() (v *Photo, ok bool) {
	value, ok := s.First()
	if !ok {
		return
	}
	return value.AsNotEmpty()
}

// LastAsNotEmpty returns last element of slice (if exists).
func (s PhotoClassArray) LastAsNotEmpty() (v *Photo, ok bool) {
	value, ok := s.Last()
	if !ok {
		return
	}
	return value.AsNotEmpty()
}

// PopFirstAsNotEmpty returns element of slice (if exists).
func (s *PhotoClassArray) PopFirstAsNotEmpty() (v *Photo, ok bool) {
	value, ok := s.PopFirst()
	if !ok {
		return
	}
	return value.AsNotEmpty()
}

// PopAsNotEmpty returns element of slice (if exists).
func (s *PhotoClassArray) PopAsNotEmpty() (v *Photo, ok bool) {
	value, ok := s.Pop()
	if !ok {
		return
	}
	return value.AsNotEmpty()
}

// PhotoEmptyArray is adapter for slice of PhotoEmpty.
type PhotoEmptyArray []PhotoEmpty

// Sort sorts slice of PhotoEmpty.
func (s PhotoEmptyArray) Sort(less func(a, b PhotoEmpty) bool) PhotoEmptyArray {
	sort.Slice(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// SortStable sorts slice of PhotoEmpty.
func (s PhotoEmptyArray) SortStable(less func(a, b PhotoEmpty) bool) PhotoEmptyArray {
	sort.SliceStable(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// Retain filters in-place slice of PhotoEmpty.
func (s PhotoEmptyArray) Retain(keep func(x PhotoEmpty) bool) PhotoEmptyArray {
	n := 0
	for _, x := range s {
		if keep(x) {
			s[n] = x
			n++
		}
	}
	s = s[:n]

	return s
}

// First returns first element of slice (if exists).
func (s PhotoEmptyArray) First() (v PhotoEmpty, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[0], true
}

// Last returns last element of slice (if exists).
func (s PhotoEmptyArray) Last() (v PhotoEmpty, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[len(s)-1], true
}

// PopFirst returns first element of slice (if exists) and deletes it.
func (s *PhotoEmptyArray) PopFirst() (v PhotoEmpty, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[0]

	// Delete by index from SliceTricks.
	copy(a[0:], a[1:])
	var zero PhotoEmpty
	a[len(a)-1] = zero
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// Pop returns last element of slice (if exists) and deletes it.
func (s *PhotoEmptyArray) Pop() (v PhotoEmpty, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[len(a)-1]
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// SortByID sorts slice of PhotoEmpty by ID.
func (s PhotoEmptyArray) SortByID() PhotoEmptyArray {
	return s.Sort(func(a, b PhotoEmpty) bool {
		return a.GetID() < b.GetID()
	})
}

// SortStableByID sorts slice of PhotoEmpty by ID.
func (s PhotoEmptyArray) SortStableByID() PhotoEmptyArray {
	return s.SortStable(func(a, b PhotoEmpty) bool {
		return a.GetID() < b.GetID()
	})
}

// FillMap fills constructors to given map.
func (s PhotoEmptyArray) FillMap(to map[int64]PhotoEmpty) {
	for _, value := range s {
		to[value.GetID()] = value
	}
}

// ToMap collects constructors to map.
func (s PhotoEmptyArray) ToMap() map[int64]PhotoEmpty {
	r := make(map[int64]PhotoEmpty, len(s))
	s.FillMap(r)
	return r
}

// PhotoArray is adapter for slice of Photo.
type PhotoArray []Photo

// Sort sorts slice of Photo.
func (s PhotoArray) Sort(less func(a, b Photo) bool) PhotoArray {
	sort.Slice(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// SortStable sorts slice of Photo.
func (s PhotoArray) SortStable(less func(a, b Photo) bool) PhotoArray {
	sort.SliceStable(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// Retain filters in-place slice of Photo.
func (s PhotoArray) Retain(keep func(x Photo) bool) PhotoArray {
	n := 0
	for _, x := range s {
		if keep(x) {
			s[n] = x
			n++
		}
	}
	s = s[:n]

	return s
}

// First returns first element of slice (if exists).
func (s PhotoArray) First() (v Photo, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[0], true
}

// Last returns last element of slice (if exists).
func (s PhotoArray) Last() (v Photo, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[len(s)-1], true
}

// PopFirst returns first element of slice (if exists) and deletes it.
func (s *PhotoArray) PopFirst() (v Photo, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[0]

	// Delete by index from SliceTricks.
	copy(a[0:], a[1:])
	var zero Photo
	a[len(a)-1] = zero
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// Pop returns last element of slice (if exists) and deletes it.
func (s *PhotoArray) Pop() (v Photo, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[len(a)-1]
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// SortByID sorts slice of Photo by ID.
func (s PhotoArray) SortByID() PhotoArray {
	return s.Sort(func(a, b Photo) bool {
		return a.GetID() < b.GetID()
	})
}

// SortStableByID sorts slice of Photo by ID.
func (s PhotoArray) SortStableByID() PhotoArray {
	return s.SortStable(func(a, b Photo) bool {
		return a.GetID() < b.GetID()
	})
}

// SortByDate sorts slice of Photo by Date.
func (s PhotoArray) SortByDate() PhotoArray {
	return s.Sort(func(a, b Photo) bool {
		return a.GetDate() < b.GetDate()
	})
}

// SortStableByDate sorts slice of Photo by Date.
func (s PhotoArray) SortStableByDate() PhotoArray {
	return s.SortStable(func(a, b Photo) bool {
		return a.GetDate() < b.GetDate()
	})
}

// FillMap fills constructors to given map.
func (s PhotoArray) FillMap(to map[int64]Photo) {
	for _, value := range s {
		to[value.GetID()] = value
	}
}

// ToMap collects constructors to map.
func (s PhotoArray) ToMap() map[int64]Photo {
	r := make(map[int64]Photo, len(s))
	s.FillMap(r)
	return r
}