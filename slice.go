/*
 * Copyright 2025 The Go-Spring Authors.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *      https://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package assert

import (
	"cmp"
	"fmt"
)

type SliceAssertion[T cmp.Ordered] struct {
	t TestingT
	v []T
}

func ThatSlice[T cmp.Ordered](t TestingT, v []T) *SliceAssertion[T] {
	return &SliceAssertion[T]{
		t: t,
		v: v,
	}
}

// Length asserts that the slice has the expected length.
func (a *SliceAssertion[T]) Length(length int, msg ...string) *SliceAssertion[T] {
	a.t.Helper()
	if len(a.v) != length {
		str := fmt.Sprintf(`length mismatch:
    got: length %d (%T) %v
 expect: length %d`, len(a.v), a.v, a.v, length)
		fail(a.t, str, msg...)
	}
	return a
}

// Equal asserts that the slice is equal to the expected slice.
func (a *SliceAssertion[T]) Equal(expect []T, msg ...string) *SliceAssertion[T] {
	a.t.Helper()
	if len(a.v) != len(expect) {
		str := fmt.Sprintf(`slices not equal:
    got: (%T) %v
 expect: (%T) %v`, a.v, a.v, expect, expect)
		fail(a.t, str, msg...)
		return a
	}
	for i := range a.v {
		if a.v[i] != expect[i] {
			str := fmt.Sprintf(`slices not equal:
    got: (%T) %v
 expect: (%T) %v`, a.v, a.v, expect, expect)
			fail(a.t, str, msg...)
			return a
		}
	}
	return a
}

// NotEqual asserts that the slice is not equal to the expected slice.
func (a *SliceAssertion[T]) NotEqual(expect []T, msg ...string) *SliceAssertion[T] {
	a.t.Helper()
	if len(a.v) == len(expect) {
		equal := true
		for i := range a.v {
			if a.v[i] != expect[i] {
				equal = false
				break
			}
		}
		if equal {
			str := fmt.Sprintf(`slices are equal:
    got: (%T) %v
 expect: not equal to %v`, a.v, a.v, expect)
			fail(a.t, str, msg...)
		}
	}
	return a
}

// IsNil asserts that the slice is nil.
func (a *SliceAssertion[T]) IsNil(msg ...string) *SliceAssertion[T] {
	a.t.Helper()
	if a.v != nil {
		str := fmt.Sprintf(`got [1 2] is not nil:
    got: (%T) %v
 expect: nil slice`, a.v, a.v)
		fail(a.t, str, msg...)
	}
	return a
}

// IsNotNil asserts that the slice is not nil.
func (a *SliceAssertion[T]) IsNotNil(msg ...string) *SliceAssertion[T] {
	a.t.Helper()
	if a.v == nil {
		str := fmt.Sprintf(`got [] is nil:
    got: (%T) %v
 expect: non-nil slice`, a.v, a.v)
		fail(a.t, str, msg...)
	}
	return a
}

// IsEmpty asserts that the slice is empty.
func (a *SliceAssertion[T]) IsEmpty(msg ...string) *SliceAssertion[T] {
	a.t.Helper()
	if len(a.v) != 0 {
		str := fmt.Sprintf(`slice is not empty:
    got: (%T) %v
 expect: empty slice`, a.v, a.v)
		fail(a.t, str, msg...)
	}
	return a
}

// IsNotEmpty asserts that the slice is not empty.
func (a *SliceAssertion[T]) IsNotEmpty(msg ...string) *SliceAssertion[T] {
	a.t.Helper()
	if len(a.v) == 0 {
		str := fmt.Sprintf(`slice is empty:
    got: (%T) %v
 expect: non-empty slice`, a.v, a.v)
		fail(a.t, str, msg...)
	}
	return a
}

// Contains asserts that the slice contains the expected element.
func (a *SliceAssertion[T]) Contains(element T, msg ...string) *SliceAssertion[T] {
	a.t.Helper()
	for _, v := range a.v {
		if v == element {
			return a
		}
	}
	str := fmt.Sprintf(`slice does not contain the expected element:
    got: (%T) %v
 expect: to contain element %v`, a.v, a.v, element)
	fail(a.t, str, msg...)
	return a
}

// NotContains asserts that the slice does not contain the expected element.
func (a *SliceAssertion[T]) NotContains(element T, msg ...string) *SliceAssertion[T] {
	a.t.Helper()
	for _, v := range a.v {
		if v == element {
			str := fmt.Sprintf(`slice contains the unexpected element:
    got: (%T) %v
 expect: not to contain element %v`, a.v, a.v, element)
			fail(a.t, str, msg...)
			return a
		}
	}
	return a
}

// SubSlice asserts that the slice contains the expected sub-slice.
func (a *SliceAssertion[T]) SubSlice(sub []T, msg ...string) *SliceAssertion[T] {
	a.t.Helper()
	if len(sub) == 0 {
		return a
	}
	for i := 0; i <= len(a.v)-len(sub); i++ {
		match := true
		for j := 0; j < len(sub); j++ {
			if a.v[i+j] != sub[j] {
				match = false
				break
			}
		}
		if match {
			return a
		}
	}
	str := fmt.Sprintf(`slice does not contain sub-slice:
    got: (%T) %v
 expect: to contain sub-slice %v`, a.v, a.v, sub)
	fail(a.t, str, msg...)
	return a
}

// NotSubSlice asserts that the slice does not contain the expected sub-slice.
func (a *SliceAssertion[T]) NotSubSlice(sub []T, msg ...string) *SliceAssertion[T] {
	a.t.Helper()
	if len(sub) == 0 {
		return a
	}
	for i := 0; i <= len(a.v)-len(sub); i++ {
		match := true
		for j := 0; j < len(sub); j++ {
			if a.v[i+j] != sub[j] {
				match = false
				break
			}
		}
		if match {
			str := fmt.Sprintf(`slice contains sub-slice:
    got: (%T) %v
 expect: not to contain sub-slice %v`, a.v, a.v, sub)
			fail(a.t, str, msg...)
			return a
		}
	}
	return a
}

// HasPrefix asserts that the slice starts with the specified prefix.
func (a *SliceAssertion[T]) HasPrefix(prefix []T, msg ...string) *SliceAssertion[T] {
	a.t.Helper()
	if len(prefix) > len(a.v) {
		str := fmt.Sprintf(`slice does not start with the expected prefix:
    got: (%T) %v
 expect: to start with %v`, a.v, a.v, prefix)
		fail(a.t, str, msg...)
		return a
	}
	for i := range prefix {
		if a.v[i] != prefix[i] {
			str := fmt.Sprintf(`slice does not start with the expected prefix:
    got: (%T) %v
 expect: to start with %v`, a.v, a.v, prefix)
			fail(a.t, str, msg...)
			return a
		}
	}
	return a
}

// HasSuffix asserts that the slice ends with the specified suffix.
func (a *SliceAssertion[T]) HasSuffix(suffix []T, msg ...string) *SliceAssertion[T] {
	a.t.Helper()
	if len(suffix) > len(a.v) {
		str := fmt.Sprintf(`slice does not end with the expected suffix:
    got: (%T) %v
 expect: to end with %v`, a.v, a.v, suffix)
		fail(a.t, str, msg...)
		return a
	}
	offset := len(a.v) - len(suffix)
	for i := range suffix {
		if a.v[offset+i] != suffix[i] {
			str := fmt.Sprintf(`slice does not end with the expected suffix:
    got: (%T) %v
 expect: to end with %v`, a.v, a.v, suffix)
			fail(a.t, str, msg...)
			return a
		}
	}
	return a
}

// IsIncreasing asserts that the slice is strictly increasing.
func (a *SliceAssertion[T]) IsIncreasing(msg ...string) *SliceAssertion[T] {
	a.t.Helper()
	for i := 1; i < len(a.v); i++ {
		if a.v[i-1] >= a.v[i] {
			str := fmt.Sprintf(`slice is not increasing at %d:
    got: (%T) %v
 expect: strictly increasing order`, i, a.v, a.v)
			fail(a.t, str, msg...)
			return a
		}
	}
	return a
}

// NonIncreasing asserts that the slice is not strictly increasing.
func (a *SliceAssertion[T]) NonIncreasing(msg ...string) *SliceAssertion[T] {
	a.t.Helper()
	for i := 1; i < len(a.v); i++ {
		if a.v[i-1] < a.v[i] {
			str := fmt.Sprintf(`slice is increasing at %d:
    got: (%T) %v
 expect: not strictly increasing`, i, a.v, a.v)
			fail(a.t, str, msg...)
			return a
		}
	}
	return a
}

// IsDecreasing asserts that the slice is strictly decreasing.
func (a *SliceAssertion[T]) IsDecreasing(msg ...string) *SliceAssertion[T] {
	a.t.Helper()
	for i := 1; i < len(a.v); i++ {
		if a.v[i-1] <= a.v[i] {
			str := fmt.Sprintf(`slice is not decreasing at %d:
    got: (%T) %v
 expect: strictly decreasing order`, i, a.v, a.v)
			fail(a.t, str, msg...)
			return a
		}
	}
	return a
}

// NonDecreasing asserts that the slice is not strictly decreasing.
func (a *SliceAssertion[T]) NonDecreasing(msg ...string) *SliceAssertion[T] {
	a.t.Helper()
	for i := 1; i < len(a.v); i++ {
		if a.v[i-1] > a.v[i] {
			str := fmt.Sprintf(`slice is decreasing at %d:
    got: (%T) %v
 expect: not strictly decreasing`, i, a.v, a.v)
			fail(a.t, str, msg...)
			return a
		}
	}
	return a
}

// IsSorted asserts that the slice is sorted in ascending order.
func (a *SliceAssertion[T]) IsSorted(msg ...string) *SliceAssertion[T] {
	a.t.Helper()
	for i := 1; i < len(a.v); i++ {
		if a.v[i-1] > a.v[i] {
			str := fmt.Sprintf(`slice is not sorted in ascending order at %d:
    got: (%T) %v
 expect: sorted in ascending order`, i, a.v, a.v)
			fail(a.t, str, msg...)
			return a
		}
	}
	return a
}

// IsSortedDescending asserts that the slice is sorted in descending order.
func (a *SliceAssertion[T]) IsSortedDescending(msg ...string) *SliceAssertion[T] {
	a.t.Helper()
	for i := 1; i < len(a.v); i++ {
		if a.v[i-1] < a.v[i] {
			str := fmt.Sprintf(`slice is not sorted in descending order at %d:
    got: (%T) %v
 expect: sorted in descending order`, i, a.v, a.v)
			fail(a.t, str, msg...)
			return a
		}
	}
	return a
}

// IsUnique asserts that all elements in the slice are unique.
func (a *SliceAssertion[T]) IsUnique(msg ...string) *SliceAssertion[T] {
	a.t.Helper()
	seen := make(map[T]bool)
	for _, v := range a.v {
		if seen[v] {
			str := fmt.Sprintf(`duplicate element found at %v:
    got: (%T) %v
 expect: all elements to be unique`, v, a.v, a.v)
			fail(a.t, str, msg...)
			return a
		}
		seen[v] = true
	}
	return a
}

// IsUniqueBy asserts that all elements in the slice are unique based on a custom function.
func (a *SliceAssertion[T]) IsUniqueBy(fn func(T) interface{}, msg ...string) *SliceAssertion[T] {
	a.t.Helper()
	seen := make(map[interface{}]bool)
	for _, v := range a.v {
		key := fn(v)
		if seen[key] {
			str := fmt.Sprintf(`duplicate element based on key function:
    got: (%T) %v
 expect: all elements to be unique by length`, a.v, a.v)
			fail(a.t, str, msg...)
			return a
		}
		seen[key] = true
	}
	return a
}

// All asserts that all elements in the slice satisfy the given condition.
func (a *SliceAssertion[T]) All(fn func(T) bool, msg ...string) *SliceAssertion[T] {
	a.t.Helper()
	for _, v := range a.v {
		if !fn(v) {
			str := fmt.Sprintf(`element does not satisfy condition:
    got: (%T) %v
 expect: all elements should satisfy condition`, a.v, a.v)
			fail(a.t, str, msg...)
			return a
		}
	}
	return a
}

// Any asserts that at least one element in the slice satisfies the given condition.
func (a *SliceAssertion[T]) Any(fn func(T) bool, msg ...string) *SliceAssertion[T] {
	a.t.Helper()
	for _, v := range a.v {
		if fn(v) {
			return a
		}
	}
	str := fmt.Sprintf(`no element satisfies the condition:
    got: (%T) %v
 expect: any element should satisfy condition`, a.v, a.v)
	fail(a.t, str, msg...)
	return a
}

// None asserts that no element in the slice satisfies the given condition.
func (a *SliceAssertion[T]) None(fn func(T) bool, msg ...string) *SliceAssertion[T] {
	a.t.Helper()
	for _, v := range a.v {
		if fn(v) {
			str := fmt.Sprintf(`element satisfies the condition:
    got: (%T) %v
 expect: no element should satisfy condition`, a.v, a.v)
			fail(a.t, str, msg...)
			return a
		}
	}
	return a
}
