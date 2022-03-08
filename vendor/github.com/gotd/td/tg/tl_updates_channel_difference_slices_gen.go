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

// UpdatesChannelDifferenceClassArray is adapter for slice of UpdatesChannelDifferenceClass.
type UpdatesChannelDifferenceClassArray []UpdatesChannelDifferenceClass

// Sort sorts slice of UpdatesChannelDifferenceClass.
func (s UpdatesChannelDifferenceClassArray) Sort(less func(a, b UpdatesChannelDifferenceClass) bool) UpdatesChannelDifferenceClassArray {
	sort.Slice(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// SortStable sorts slice of UpdatesChannelDifferenceClass.
func (s UpdatesChannelDifferenceClassArray) SortStable(less func(a, b UpdatesChannelDifferenceClass) bool) UpdatesChannelDifferenceClassArray {
	sort.SliceStable(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// Retain filters in-place slice of UpdatesChannelDifferenceClass.
func (s UpdatesChannelDifferenceClassArray) Retain(keep func(x UpdatesChannelDifferenceClass) bool) UpdatesChannelDifferenceClassArray {
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
func (s UpdatesChannelDifferenceClassArray) First() (v UpdatesChannelDifferenceClass, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[0], true
}

// Last returns last element of slice (if exists).
func (s UpdatesChannelDifferenceClassArray) Last() (v UpdatesChannelDifferenceClass, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[len(s)-1], true
}

// PopFirst returns first element of slice (if exists) and deletes it.
func (s *UpdatesChannelDifferenceClassArray) PopFirst() (v UpdatesChannelDifferenceClass, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[0]

	// Delete by index from SliceTricks.
	copy(a[0:], a[1:])
	var zero UpdatesChannelDifferenceClass
	a[len(a)-1] = zero
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// Pop returns last element of slice (if exists) and deletes it.
func (s *UpdatesChannelDifferenceClassArray) Pop() (v UpdatesChannelDifferenceClass, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[len(a)-1]
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// AsUpdatesChannelDifferenceEmpty returns copy with only UpdatesChannelDifferenceEmpty constructors.
func (s UpdatesChannelDifferenceClassArray) AsUpdatesChannelDifferenceEmpty() (to UpdatesChannelDifferenceEmptyArray) {
	for _, elem := range s {
		value, ok := elem.(*UpdatesChannelDifferenceEmpty)
		if !ok {
			continue
		}
		to = append(to, *value)
	}

	return to
}

// AsUpdatesChannelDifferenceTooLong returns copy with only UpdatesChannelDifferenceTooLong constructors.
func (s UpdatesChannelDifferenceClassArray) AsUpdatesChannelDifferenceTooLong() (to UpdatesChannelDifferenceTooLongArray) {
	for _, elem := range s {
		value, ok := elem.(*UpdatesChannelDifferenceTooLong)
		if !ok {
			continue
		}
		to = append(to, *value)
	}

	return to
}

// AsUpdatesChannelDifference returns copy with only UpdatesChannelDifference constructors.
func (s UpdatesChannelDifferenceClassArray) AsUpdatesChannelDifference() (to UpdatesChannelDifferenceArray) {
	for _, elem := range s {
		value, ok := elem.(*UpdatesChannelDifference)
		if !ok {
			continue
		}
		to = append(to, *value)
	}

	return to
}

// AppendOnlyNotEmpty appends only NotEmpty constructors to
// given slice.
func (s UpdatesChannelDifferenceClassArray) AppendOnlyNotEmpty(to []NotEmptyUpdatesChannelDifference) []NotEmptyUpdatesChannelDifference {
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
func (s UpdatesChannelDifferenceClassArray) AsNotEmpty() (to []NotEmptyUpdatesChannelDifference) {
	return s.AppendOnlyNotEmpty(to)
}

// FirstAsNotEmpty returns first element of slice (if exists).
func (s UpdatesChannelDifferenceClassArray) FirstAsNotEmpty() (v NotEmptyUpdatesChannelDifference, ok bool) {
	value, ok := s.First()
	if !ok {
		return
	}
	return value.AsNotEmpty()
}

// LastAsNotEmpty returns last element of slice (if exists).
func (s UpdatesChannelDifferenceClassArray) LastAsNotEmpty() (v NotEmptyUpdatesChannelDifference, ok bool) {
	value, ok := s.Last()
	if !ok {
		return
	}
	return value.AsNotEmpty()
}

// PopFirstAsNotEmpty returns element of slice (if exists).
func (s *UpdatesChannelDifferenceClassArray) PopFirstAsNotEmpty() (v NotEmptyUpdatesChannelDifference, ok bool) {
	value, ok := s.PopFirst()
	if !ok {
		return
	}
	return value.AsNotEmpty()
}

// PopAsNotEmpty returns element of slice (if exists).
func (s *UpdatesChannelDifferenceClassArray) PopAsNotEmpty() (v NotEmptyUpdatesChannelDifference, ok bool) {
	value, ok := s.Pop()
	if !ok {
		return
	}
	return value.AsNotEmpty()
}

// UpdatesChannelDifferenceEmptyArray is adapter for slice of UpdatesChannelDifferenceEmpty.
type UpdatesChannelDifferenceEmptyArray []UpdatesChannelDifferenceEmpty

// Sort sorts slice of UpdatesChannelDifferenceEmpty.
func (s UpdatesChannelDifferenceEmptyArray) Sort(less func(a, b UpdatesChannelDifferenceEmpty) bool) UpdatesChannelDifferenceEmptyArray {
	sort.Slice(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// SortStable sorts slice of UpdatesChannelDifferenceEmpty.
func (s UpdatesChannelDifferenceEmptyArray) SortStable(less func(a, b UpdatesChannelDifferenceEmpty) bool) UpdatesChannelDifferenceEmptyArray {
	sort.SliceStable(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// Retain filters in-place slice of UpdatesChannelDifferenceEmpty.
func (s UpdatesChannelDifferenceEmptyArray) Retain(keep func(x UpdatesChannelDifferenceEmpty) bool) UpdatesChannelDifferenceEmptyArray {
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
func (s UpdatesChannelDifferenceEmptyArray) First() (v UpdatesChannelDifferenceEmpty, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[0], true
}

// Last returns last element of slice (if exists).
func (s UpdatesChannelDifferenceEmptyArray) Last() (v UpdatesChannelDifferenceEmpty, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[len(s)-1], true
}

// PopFirst returns first element of slice (if exists) and deletes it.
func (s *UpdatesChannelDifferenceEmptyArray) PopFirst() (v UpdatesChannelDifferenceEmpty, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[0]

	// Delete by index from SliceTricks.
	copy(a[0:], a[1:])
	var zero UpdatesChannelDifferenceEmpty
	a[len(a)-1] = zero
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// Pop returns last element of slice (if exists) and deletes it.
func (s *UpdatesChannelDifferenceEmptyArray) Pop() (v UpdatesChannelDifferenceEmpty, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[len(a)-1]
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// UpdatesChannelDifferenceTooLongArray is adapter for slice of UpdatesChannelDifferenceTooLong.
type UpdatesChannelDifferenceTooLongArray []UpdatesChannelDifferenceTooLong

// Sort sorts slice of UpdatesChannelDifferenceTooLong.
func (s UpdatesChannelDifferenceTooLongArray) Sort(less func(a, b UpdatesChannelDifferenceTooLong) bool) UpdatesChannelDifferenceTooLongArray {
	sort.Slice(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// SortStable sorts slice of UpdatesChannelDifferenceTooLong.
func (s UpdatesChannelDifferenceTooLongArray) SortStable(less func(a, b UpdatesChannelDifferenceTooLong) bool) UpdatesChannelDifferenceTooLongArray {
	sort.SliceStable(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// Retain filters in-place slice of UpdatesChannelDifferenceTooLong.
func (s UpdatesChannelDifferenceTooLongArray) Retain(keep func(x UpdatesChannelDifferenceTooLong) bool) UpdatesChannelDifferenceTooLongArray {
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
func (s UpdatesChannelDifferenceTooLongArray) First() (v UpdatesChannelDifferenceTooLong, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[0], true
}

// Last returns last element of slice (if exists).
func (s UpdatesChannelDifferenceTooLongArray) Last() (v UpdatesChannelDifferenceTooLong, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[len(s)-1], true
}

// PopFirst returns first element of slice (if exists) and deletes it.
func (s *UpdatesChannelDifferenceTooLongArray) PopFirst() (v UpdatesChannelDifferenceTooLong, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[0]

	// Delete by index from SliceTricks.
	copy(a[0:], a[1:])
	var zero UpdatesChannelDifferenceTooLong
	a[len(a)-1] = zero
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// Pop returns last element of slice (if exists) and deletes it.
func (s *UpdatesChannelDifferenceTooLongArray) Pop() (v UpdatesChannelDifferenceTooLong, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[len(a)-1]
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// UpdatesChannelDifferenceArray is adapter for slice of UpdatesChannelDifference.
type UpdatesChannelDifferenceArray []UpdatesChannelDifference

// Sort sorts slice of UpdatesChannelDifference.
func (s UpdatesChannelDifferenceArray) Sort(less func(a, b UpdatesChannelDifference) bool) UpdatesChannelDifferenceArray {
	sort.Slice(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// SortStable sorts slice of UpdatesChannelDifference.
func (s UpdatesChannelDifferenceArray) SortStable(less func(a, b UpdatesChannelDifference) bool) UpdatesChannelDifferenceArray {
	sort.SliceStable(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// Retain filters in-place slice of UpdatesChannelDifference.
func (s UpdatesChannelDifferenceArray) Retain(keep func(x UpdatesChannelDifference) bool) UpdatesChannelDifferenceArray {
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
func (s UpdatesChannelDifferenceArray) First() (v UpdatesChannelDifference, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[0], true
}

// Last returns last element of slice (if exists).
func (s UpdatesChannelDifferenceArray) Last() (v UpdatesChannelDifference, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[len(s)-1], true
}

// PopFirst returns first element of slice (if exists) and deletes it.
func (s *UpdatesChannelDifferenceArray) PopFirst() (v UpdatesChannelDifference, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[0]

	// Delete by index from SliceTricks.
	copy(a[0:], a[1:])
	var zero UpdatesChannelDifference
	a[len(a)-1] = zero
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// Pop returns last element of slice (if exists) and deletes it.
func (s *UpdatesChannelDifferenceArray) Pop() (v UpdatesChannelDifference, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[len(a)-1]
	a = a[:len(a)-1]
	*s = a

	return v, true
}
