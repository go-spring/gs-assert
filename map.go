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
	"fmt"
)

// MapAssertion encapsulates a map value and a test handler for making assertions on the map.
type MapAssertion[K comparable, V comparable] struct {
	c AssertionConfig
	t TestingT
	v map[K]V
}

// ThatMap returns a MapAssertion for the given testing object and map value.
func ThatMap[K comparable, V comparable](t TestingT, v map[K]V, options ...AssertionOption) *MapAssertion[K, V] {
	c := AssertionConfig{
		outputValueAsJSON: true,
	}
	for _, opt := range options {
		opt(&c)
	}
	return &MapAssertion[K, V]{
		c: c,
		t: t,
		v: v,
	}
}

// Length asserts that the map has the expected length.
func (a *MapAssertion[K, V]) Length(length int, msg ...string) {
	a.t.Helper()
	if len(a.v) != length {
		str := fmt.Sprintf(`expected map to have length %v, but it has length %d
    got: %v`, length, len(a.v), a.c.OutputValue(a.v))
		fail(a.t, str, msg...)
	}
}

// Equal asserts that the map is equal to the expected map.
func (a *MapAssertion[K, V]) Equal(expect map[K]V, msg ...string) {
	a.t.Helper()
	if len(a.v) != len(expect) {
		str := fmt.Sprintf(`expected maps to be equal, but their lengths are different
    got: %v
 expect: %v`, a.c.OutputValue(a.v), a.c.OutputValue(expect))
		fail(a.t, str, msg...)
		return
	}
	for k, v := range a.v {
		if expectV, ok := expect[k]; !ok {
			str := fmt.Sprintf(`expected maps to be equal, but key '%v' is missing
    got: %v
 expect: %v`, k, a.c.OutputValue(a.v), a.c.OutputValue(expect))
			fail(a.t, str, msg...)
			return
		} else if v != expectV {
			str := fmt.Sprintf(`expected maps to be equal, but values for key '%v' are different
    got: %v
 expect: %v`, k, a.c.OutputValue(a.v), a.c.OutputValue(expect))
			fail(a.t, str, msg...)
			return
		}
	}
}

// NotEqual asserts that the map is not equal to the expected map.
func (a *MapAssertion[K, V]) NotEqual(expect map[K]V, msg ...string) {
	a.t.Helper()
	if len(a.v) == len(expect) {
		equal := true
		for k, v := range a.v {
			if expectV, ok := expect[k]; !ok || v != expectV {
				equal = false
				break
			}
		}
		if equal {
			str := fmt.Sprintf(`expected maps to be different, but they are equal
    got: %v`, a.c.OutputValue(a.v))
			fail(a.t, str, msg...)
		}
	}
}

// IsEmpty asserts that the map is empty.
func (a *MapAssertion[K, V]) IsEmpty(msg ...string) {
	a.t.Helper()
	if len(a.v) != 0 {
		str := fmt.Sprintf(`expected map to be empty, but it is not
    got: %v`, a.c.OutputValue(a.v))
		fail(a.t, str, msg...)
	}
}

// IsNotEmpty asserts that the map is not empty.
func (a *MapAssertion[K, V]) IsNotEmpty(msg ...string) {
	a.t.Helper()
	if len(a.v) == 0 {
		str := fmt.Sprintf(`expected map to be non-empty, but it is empty
    got: %v`, a.c.OutputValue(a.v))
		fail(a.t, str, msg...)
	}
}

// ContainsKey asserts that the map contains the expected key.
func (a *MapAssertion[K, V]) ContainsKey(key K, msg ...string) {
	a.t.Helper()
	if _, ok := a.v[key]; !ok {
		str := fmt.Sprintf(`expected map to contain key '%v', but it is missing
    got: %v`, key, a.c.OutputValue(a.v))
		fail(a.t, str, msg...)
	}
}

// NotContainsKey asserts that the map does not contain the expected key.
func (a *MapAssertion[K, V]) NotContainsKey(key K, msg ...string) {
	a.t.Helper()
	if _, ok := a.v[key]; ok {
		str := fmt.Sprintf(`expected map not to contain key '%v', but it is found
    got: %v`, key, a.c.OutputValue(a.v))
		fail(a.t, str, msg...)
	}
}

// ContainsValue asserts that the map contains the expected value.
func (a *MapAssertion[K, V]) ContainsValue(value V, msg ...string) {
	a.t.Helper()
	for _, v := range a.v {
		if v == value {
			return
		}
	}
	str := fmt.Sprintf(`expected map to contain value %#v, but it is missing
    got: %v`, value, a.c.OutputValue(a.v))
	fail(a.t, str, msg...)
}

// NotContainsValue asserts that the map does not contain the expected value.
func (a *MapAssertion[K, V]) NotContainsValue(value V, msg ...string) {
	a.t.Helper()
	for _, v := range a.v {
		if v == value {
			str := fmt.Sprintf(`expected map not to contain value %#v, but it is found
    got: %v`, value, a.c.OutputValue(a.v))
			fail(a.t, str, msg...)
			return
		}
	}
}

// HasKeyValue asserts that the map contains the expected key-value pair.
func (a *MapAssertion[K, V]) HasKeyValue(key K, value V, msg ...string) {
	a.t.Helper()
	if v, ok := a.v[key]; !ok {
		str := fmt.Sprintf(`expected map to contain key '%v', but it is missing
    got: %v`, key, a.c.OutputValue(a.v))
		fail(a.t, str, msg...)
	} else if v != value {
		str := fmt.Sprintf(`expected value %v for key '%v', but got %v instead
    got: %v`, key, value, v, a.c.OutputValue(a.v))
		fail(a.t, str, msg...)
	}
}

// ContainsKeys asserts that the map contains all the expected keys.
func (a *MapAssertion[K, V]) ContainsKeys(expect []K, msg ...string) {
	a.t.Helper()
	for _, key := range expect {
		if _, ok := a.v[key]; !ok {
			str := fmt.Sprintf(`expected map to contain key '%v', but it is missing
    got: %v`, key, a.c.OutputValue(a.v))
			fail(a.t, str, msg...)
			return
		}
	}
}

// NotContainsKeys asserts that the map does not contain any of the expected keys.
func (a *MapAssertion[K, V]) NotContainsKeys(expect []K, msg ...string) {
	a.t.Helper()
	for _, key := range expect {
		if _, ok := a.v[key]; ok {
			str := fmt.Sprintf(`expected map not to contain key '%v', but it is found
    got: %v`, key, a.c.OutputValue(a.v))
			fail(a.t, str, msg...)
			return
		}
	}
}

// ContainsValues asserts that the map contains all the expected values.
func (a *MapAssertion[K, V]) ContainsValues(expect []V, msg ...string) {
	a.t.Helper()
	for _, value := range expect {
		found := false
		for _, v := range a.v {
			if v == value {
				found = true
				break
			}
		}
		if !found {
			str := fmt.Sprintf(`expected map to contain value %#v, but it is missing
    got: %v`, value, a.c.OutputValue(a.v))
			fail(a.t, str, msg...)
			return
		}
	}
}

// NotContainsValues asserts that the map does not contain any of the expected values.
func (a *MapAssertion[K, V]) NotContainsValues(expect []V, msg ...string) {
	a.t.Helper()
	for _, value := range expect {
		for _, v := range a.v {
			if v == value {
				str := fmt.Sprintf(`expected map not to contain value %#v, but it is found
    got: %v`, v, a.c.OutputValue(a.v))
				fail(a.t, str, msg...)
				return
			}
		}
	}
}

// IsSubsetOf asserts that the map is a subset of the expected map.
func (a *MapAssertion[K, V]) IsSubsetOf(expect map[K]V, msg ...string) {
	a.t.Helper()
	for k, v := range a.v {
		if expectV, ok := expect[k]; !ok {
			str := fmt.Sprintf(`expected map to be a subset, but unexpected key '%v' is found
    got: %v
 expect: %v`, k, a.c.OutputValue(a.v), a.c.OutputValue(expect))
			fail(a.t, str, msg...)
			return
		} else if v != expectV {
			str := fmt.Sprintf(`expected map to be a subset, but values for key '%v' are different 
    got: %v
 expect: %v`, k, a.c.OutputValue(a.v), a.c.OutputValue(expect))
			fail(a.t, str, msg...)
			return
		}
	}
}

// IsSupersetOf asserts that the map is a superset of the expected map.
func (a *MapAssertion[K, V]) IsSupersetOf(expect map[K]V, msg ...string) {
	a.t.Helper()
	for k, v := range expect {
		if aV, ok := a.v[k]; !ok {
			str := fmt.Sprintf(`expected map to be a superset, but key '%v' is missing
    got: %v
 expect: %v`, k, a.c.OutputValue(a.v), a.c.OutputValue(expect))
			fail(a.t, str, msg...)
			return
		} else if aV != v {
			str := fmt.Sprintf(`expected map to be a superset, but values for key '%v' are different
    got: %v
 expect: %v`, k, a.c.OutputValue(a.v), a.c.OutputValue(expect))
			fail(a.t, str, msg...)
			return
		}
	}
}

// HasSameKeys asserts that the map has the same keys as the expected map.
func (a *MapAssertion[K, V]) HasSameKeys(expect map[K]V, msg ...string) {
	a.t.Helper()
	if len(a.v) != len(expect) {
		str := fmt.Sprintf(`expected maps to have the same keys, but their lengths differ
    got: %v
 expect: %v`, a.c.OutputValue(a.v), a.c.OutputValue(expect))
		fail(a.t, str, msg...)
		return
	}
	for k := range a.v {
		if _, ok := expect[k]; !ok {
			str := fmt.Sprintf(`expected maps to have the same keys, but key '%v' is missing
    got: %v
 expect: %v`, k, a.c.OutputValue(a.v), a.c.OutputValue(expect))
			fail(a.t, str, msg...)
			return
		}
	}
}

// HasSameValues asserts that the map has the same values as the expected map.
func (a *MapAssertion[K, V]) HasSameValues(expect map[K]V, msg ...string) {
	a.t.Helper()
	if len(a.v) != len(expect) {
		str := fmt.Sprintf(`expected maps to have the same values, but their lengths differ
    got: %v
 expect: %v`, a.c.OutputValue(a.v), a.c.OutputValue(expect))
		fail(a.t, str, msg...)
		return
	}
	valueCount := make(map[V]int)
	for _, v := range a.v {
		valueCount[v]++
	}
	for _, v := range expect {
		valueCount[v]--
	}
	for _, count := range valueCount {
		if count != 0 {
			str := fmt.Sprintf(`expected maps to have the same values, but values are different
    got: %v
 expect: %v`, a.c.OutputValue(a.v), a.c.OutputValue(expect))
			fail(a.t, str, msg...)
			return
		}
	}
}
