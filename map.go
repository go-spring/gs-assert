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
	t TestingT
	v map[K]V
}

// ThatMap returns a MapAssertion for the given testing object and map value.
func ThatMap[K comparable, V comparable](t TestingT, v map[K]V) *MapAssertion[K, V] {
	return &MapAssertion[K, V]{
		t: t,
		v: v,
	}
}

// Length asserts that the map has the expected length.
func (a *MapAssertion[K, V]) Length(length int, msg ...string) {
	a.t.Helper()
	if len(a.v) != length {
		str := fmt.Sprintf(`length mismatch:
    got: length %d (%T) %v
 expect: length %d`, len(a.v), a.v, a.v, length)
		fail(a.t, str, msg...)
	}
}

// Equal asserts that the map is equal to the expected map.
func (a *MapAssertion[K, V]) Equal(expect map[K]V, msg ...string) {
	a.t.Helper()
	if len(a.v) != len(expect) {
		str := fmt.Sprintf(`map length mismatch:
    got: length %d (%T) %v
 expect: length %d`, len(a.v), a.v, a.v, len(expect))
		fail(a.t, str, msg...)
		return
	}
	for k, v := range a.v {
		if expectV, ok := expect[k]; !ok || v != expectV {
			str := fmt.Sprintf(`map content mismatch:
    got: key %v value (%T) %v
 expect: key %v value (%T) %v`, k, v, v, k, expectV, expectV)
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
			str := fmt.Sprintf(`maps are equal:
    got: (%T) %v
 expect: not equal to %v`, a.v, a.v, expect)
			fail(a.t, str, msg...)
		}
	}
}

// IsEmpty asserts that the map is empty.
func (a *MapAssertion[K, V]) IsEmpty(msg ...string) {
	a.t.Helper()
	if len(a.v) != 0 {
		str := fmt.Sprintf(`map is not empty:
    got: (%T) %v
 expect: empty map`, a.v, a.v)
		fail(a.t, str, msg...)
	}
}

// IsNotEmpty asserts that the map is not empty.
func (a *MapAssertion[K, V]) IsNotEmpty(msg ...string) {
	a.t.Helper()
	if len(a.v) == 0 {
		str := fmt.Sprintf(`map is empty:
    got: (%T) %v
 expect: non-empty map`, a.v, a.v)
		fail(a.t, str, msg...)
	}
}

// Contains asserts that the map contains the expected key.
func (a *MapAssertion[K, V]) Contains(key K, msg ...string) {
	a.t.Helper()
	if _, ok := a.v[key]; !ok {
		str := fmt.Sprintf(`map does not contain the key:
    got: (%T) %v
 expect: map containing key %v`, a.v, a.v, key)
		fail(a.t, str, msg...)
	}
}

// NotContains asserts that the map does not contain the expected key.
func (a *MapAssertion[K, V]) NotContains(key K, msg ...string) {
	a.t.Helper()
	if _, ok := a.v[key]; ok {
		str := fmt.Sprintf(`map contains the key:
    got: (%T) %v
 expect: map not containing key %v`, a.v, a.v, key)
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
	str := fmt.Sprintf(`map does not contain the value:
    got: (%T) %v
 expect: map containing value %v`, a.v, a.v, value)
	fail(a.t, str, msg...)
}

// NotContainsValue asserts that the map does not contain the expected value.
func (a *MapAssertion[K, V]) NotContainsValue(value V, msg ...string) {
	a.t.Helper()
	for _, v := range a.v {
		if v == value {
			str := fmt.Sprintf(`map contains the value:
    got: (%T) %v
 expect: map not containing value %v`, a.v, a.v, value)
			fail(a.t, str, msg...)
			return
		}
	}
}

// HasKeyValue asserts that the map contains the expected key-value pair.
func (a *MapAssertion[K, V]) HasKeyValue(key K, value V, msg ...string) {
	a.t.Helper()
	if v, ok := a.v[key]; !ok || v != value {
		str := fmt.Sprintf(`key-value pair mismatch:
    got: key %v value (%T) %v
 expect: key %v value (%T) %v`, key, v, v, key, value, value)
		fail(a.t, str, msg...)
	}
}

// ContainsKeys asserts that the map contains all the expected keys.
func (a *MapAssertion[K, V]) ContainsKeys(keys []K, msg ...string) {
	a.t.Helper()
	for _, key := range keys {
		if _, ok := a.v[key]; !ok {
			str := fmt.Sprintf(`map does not contain all keys:
    got: (%T) %v
 expect: map containing key %v`, a.v, a.v, key)
			fail(a.t, str, msg...)
			return
		}
	}
}

// NotContainsKeys asserts that the map does not contain any of the expected keys.
func (a *MapAssertion[K, V]) NotContainsKeys(keys []K, msg ...string) {
	a.t.Helper()
	for _, key := range keys {
		if _, ok := a.v[key]; ok {
			str := fmt.Sprintf(`map contains unexpected key:
    got: (%T) %v
 expect: map not containing key %v`, a.v, a.v, key)
			fail(a.t, str, msg...)
			return
		}
	}
}

// ContainsValues asserts that the map contains all the expected values.
func (a *MapAssertion[K, V]) ContainsValues(values []V, msg ...string) {
	a.t.Helper()
	for _, value := range values {
		found := false
		for _, v := range a.v {
			if v == value {
				found = true
				break
			}
		}
		if !found {
			str := fmt.Sprintf(`map does not contain all values:
    got: (%T) %v
 expect: map containing value %v`, a.v, a.v, value)
			fail(a.t, str, msg...)
			return
		}
	}
}

// NotContainsValues asserts that the map does not contain any of the expected values.
func (a *MapAssertion[K, V]) NotContainsValues(values []V, msg ...string) {
	a.t.Helper()
	for _, value := range values {
		for _, v := range a.v {
			if v == value {
				str := fmt.Sprintf(`map contains unexpected value:
    got: (%T) %v
 expect: map not containing value %v`, a.v, a.v, value)
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
		if expectV, ok := expect[k]; !ok || v != expectV {
			str := fmt.Sprintf(`map is not a subset:
    got: key %v value (%T) %v
 expect: key %v value (%T) %v`, k, v, v, k, expectV, expectV)
			fail(a.t, str, msg...)
			return
		}
	}
}

// IsSupersetOf asserts that the map is a superset of the expected map.
func (a *MapAssertion[K, V]) IsSupersetOf(expect map[K]V, msg ...string) {
	a.t.Helper()
	for k, v := range expect {
		if aV, ok := a.v[k]; !ok || aV != v {
			str := fmt.Sprintf(`map is not a superset:
    got: key %v value (%T) %v
 expect: key %v value (%T) %v`, k, aV, aV, k, v, v)
			fail(a.t, str, msg...)
			return
		}
	}
}

// HasSameKeys asserts that the map has the same keys as the expected map.
func (a *MapAssertion[K, V]) HasSameKeys(expect map[K]V, msg ...string) {
	a.t.Helper()
	if len(a.v) != len(expect) {
		str := fmt.Sprintf(`map key count mismatch:
    got: count %d (%T) %v
 expect: count %d`, len(a.v), a.v, a.v, len(expect))
		fail(a.t, str, msg...)
		return
	}
	for k := range a.v {
		if _, ok := expect[k]; !ok {
			str := fmt.Sprintf(`map keys do not match:
    got: (%T) %v
 expect: map with same keys`, a.v, a.v)
			fail(a.t, str, msg...)
			return
		}
	}
}

// HasSameValues asserts that the map has the same values as the expected map.
func (a *MapAssertion[K, V]) HasSameValues(expect map[K]V, msg ...string) {
	a.t.Helper()
	if len(a.v) != len(expect) {
		str := fmt.Sprintf(`map value count mismatch:
    got: count %d (%T) %v
 expect: count %d`, len(a.v), a.v, a.v, len(expect))
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
			str := fmt.Sprintf(`map values do not match:
    got: (%T) %v
 expect: map with same values`, a.v, a.v)
			fail(a.t, str, msg...)
			return
		}
	}
}
