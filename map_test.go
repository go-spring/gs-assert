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

package assert_test

import (
	"testing"

	"github.com/go-spring/assert"
	"github.com/go-spring/assert/internal"
)

func TestMap_Length(t *testing.T) {
	m := new(internal.MockTestingT)
	testMap := map[string]int{"a": 1}
	assert.ThatMap(m, testMap).Length(1)
	assert.ThatString(t, m.String()).Equal("")

	m.Reset()
	assert.ThatMap(m, testMap).Length(0)
	assert.ThatString(t, m.String()).Equal(`error# Assertion failed: expected map to have length 0, but it has length 1
    got: {"a":1}`)

	m.Reset()
	assert.ThatMap(m, testMap).Must().Length(0, "index is 0")
	assert.ThatString(t, m.String()).Equal(`fatal# Assertion failed: expected map to have length 0, but it has length 1
    got: {"a":1}
message: "index is 0"`)
}

func TestMap_Equal(t *testing.T) {
	m := new(internal.MockTestingT)
	testMap := map[string]int{"a": 1}
	expectMap := map[string]int{"a": 1}
	assert.ThatMap(m, testMap).Equal(expectMap)
	assert.ThatString(t, m.String()).Equal("")

	m.Reset()
	assert.ThatMap(m, testMap).Equal(map[string]int{"b": 2})
	assert.ThatString(t, m.String()).Equal(`error# Assertion failed: expected maps to be equal, but key 'a' is missing
    got: {"a":1}
 expect: {"b":2}`)

	m.Reset()
	assert.ThatMap(m, testMap).Must().Equal(map[string]int{"b": 2}, "index is 0")
	assert.ThatString(t, m.String()).Equal(`fatal# Assertion failed: expected maps to be equal, but key 'a' is missing
    got: {"a":1}
 expect: {"b":2}
message: "index is 0"`)
}

func TestMap_NotEqual(t *testing.T) {
	m := new(internal.MockTestingT)
	testMap := map[string]int{"a": 1}
	assert.ThatMap(m, testMap).NotEqual(map[string]int{"b": 2})
	assert.ThatString(t, m.String()).Equal("")

	m.Reset()
	assert.ThatMap(m, testMap).NotEqual(testMap)
	assert.ThatString(t, m.String()).Equal(`error# Assertion failed: expected maps to be different, but they are equal
    got: {"a":1}`)

	m.Reset()
	assert.ThatMap(m, testMap).Must().NotEqual(testMap, "index is 0")
	assert.ThatString(t, m.String()).Equal(`fatal# Assertion failed: expected maps to be different, but they are equal
    got: {"a":1}
message: "index is 0"`)
}

func TestMap_Empty(t *testing.T) {
	m := new(internal.MockTestingT)
	var emptyMap map[string]int
	assert.ThatMap(m, emptyMap).Empty()
	assert.ThatString(t, m.String()).Equal("")

	m.Reset()
	assert.ThatMap(m, map[string]int{"a": 1}).Empty()
	assert.ThatString(t, m.String()).Equal(`error# Assertion failed: expected map to be empty, but it is not
    got: {"a":1}`)

	m.Reset()
	assert.ThatMap(m, map[string]int{"a": 1}).Must().Empty("index is 0")
	assert.ThatString(t, m.String()).Equal(`fatal# Assertion failed: expected map to be empty, but it is not
    got: {"a":1}
message: "index is 0"`)
}

func TestMap_NotEmpty(t *testing.T) {
	m := new(internal.MockTestingT)
	testMap := map[string]int{"a": 1}
	assert.ThatMap(m, testMap).NotEmpty()
	assert.ThatString(t, m.String()).Equal("")

	m.Reset()
	var emptyMap map[string]int
	assert.ThatMap(m, emptyMap).NotEmpty()
	assert.ThatString(t, m.String()).Equal(`error# Assertion failed: expected map to be non-empty, but it is empty
    got: null`)

	m.Reset()
	assert.ThatMap(m, emptyMap).Must().NotEmpty("index is 0")
	assert.ThatString(t, m.String()).Equal(`fatal# Assertion failed: expected map to be non-empty, but it is empty
    got: null
message: "index is 0"`)
}

func TestMap_ContainsKey(t *testing.T) {
	m := new(internal.MockTestingT)
	testMap := map[string]int{"a": 1}
	assert.ThatMap(m, testMap).ContainsKey("a")
	assert.ThatString(t, m.String()).Equal("")

	m.Reset()
	assert.ThatMap(m, testMap).ContainsKey("b")
	assert.ThatString(t, m.String()).Equal(`error# Assertion failed: expected map to contain key 'b', but it is missing
    got: {"a":1}`)

	m.Reset()
	assert.ThatMap(m, testMap).Must().ContainsKey("b", "index is 0")
	assert.ThatString(t, m.String()).Equal(`fatal# Assertion failed: expected map to contain key 'b', but it is missing
    got: {"a":1}
message: "index is 0"`)
}

func TestMap_NotContainsKey(t *testing.T) {
	m := new(internal.MockTestingT)
	testMap := map[string]int{"a": 1}
	assert.ThatMap(m, testMap).NotContainsKey("b")
	assert.ThatString(t, m.String()).Equal("")

	m.Reset()
	assert.ThatMap(m, testMap).NotContainsKey("a")
	assert.ThatString(t, m.String()).Equal(`error# Assertion failed: expected map not to contain key 'a', but it is found
    got: {"a":1}`)

	m.Reset()
	assert.ThatMap(m, testMap).Must().NotContainsKey("a", "index is 0")
	assert.ThatString(t, m.String()).Equal(`fatal# Assertion failed: expected map not to contain key 'a', but it is found
    got: {"a":1}
message: "index is 0"`)
}

func TestMap_ContainsValue(t *testing.T) {
	m := new(internal.MockTestingT)
	testMap := map[string]int{"a": 1}
	assert.ThatMap(m, testMap).ContainsValue(1)
	assert.ThatString(t, m.String()).Equal("")

	m.Reset()
	assert.ThatMap(m, testMap).ContainsValue(2)
	assert.ThatString(t, m.String()).Equal(`error# Assertion failed: expected map to contain value 2, but it is missing
    got: {"a":1}`)

	m.Reset()
	assert.ThatMap(m, testMap).Must().ContainsValue(2, "index is 0")
	assert.ThatString(t, m.String()).Equal(`fatal# Assertion failed: expected map to contain value 2, but it is missing
    got: {"a":1}
message: "index is 0"`)
}

func TestMap_NotContainsValue(t *testing.T) {
	m := new(internal.MockTestingT)
	testMap := map[string]int{"a": 1}
	assert.ThatMap(m, testMap).NotContainsValue(2)
	assert.ThatString(t, m.String()).Equal("")

	m.Reset()
	assert.ThatMap(m, testMap).NotContainsValue(1)
	assert.ThatString(t, m.String()).Equal(`error# Assertion failed: expected map not to contain value 1, but it is found
    got: {"a":1}`)

	m.Reset()
	assert.ThatMap(m, testMap).Must().NotContainsValue(1, "index is 0")
	assert.ThatString(t, m.String()).Equal(`fatal# Assertion failed: expected map not to contain value 1, but it is found
    got: {"a":1}
message: "index is 0"`)
}

func TestMap_ContainsKeyValue(t *testing.T) {
	m := new(internal.MockTestingT)
	testMap := map[string]int{"a": 1}
	assert.ThatMap(m, testMap).ContainsKeyValue("a", 1)
	assert.ThatString(t, m.String()).Equal("")

	m.Reset()
	assert.ThatMap(m, testMap).ContainsKeyValue("a", 2)
	assert.ThatString(t, m.String()).Equal(`error# Assertion failed: expected value 2 for key 'a', but got 1 instead
    got: {"a":1}`)

	m.Reset()
	assert.ThatMap(m, testMap).Must().ContainsKeyValue("a", 2, "index is 0")
	assert.ThatString(t, m.String()).Equal(`fatal# Assertion failed: expected value 2 for key 'a', but got 1 instead
    got: {"a":1}
message: "index is 0"`)
}

func TestMap_ContainsKeys(t *testing.T) {
	m := new(internal.MockTestingT)
	testMap := map[string]int{"a": 1, "b": 2}
	assert.ThatMap(m, testMap).ContainsKeys([]string{"a", "b"})
	assert.ThatString(t, m.String()).Equal("")

	m.Reset()
	assert.ThatMap(m, testMap).ContainsKeys([]string{"c"})
	assert.ThatString(t, m.String()).Equal(`error# Assertion failed: expected map to contain key 'c', but it is missing
    got: {"a":1,"b":2}`)

	m.Reset()
	assert.ThatMap(m, testMap).Must().ContainsKeys([]string{"c"}, "index is 0")
	assert.ThatString(t, m.String()).Equal(`fatal# Assertion failed: expected map to contain key 'c', but it is missing
    got: {"a":1,"b":2}
message: "index is 0"`)
}

func TestMap_NotContainsKeys(t *testing.T) {
	m := new(internal.MockTestingT)
	testMap := map[string]int{"a": 1, "b": 2}
	assert.ThatMap(m, testMap).NotContainsKeys([]string{"c"})
	assert.ThatString(t, m.String()).Equal("")

	m.Reset()
	assert.ThatMap(m, testMap).NotContainsKeys([]string{"a"})
	assert.ThatString(t, m.String()).Equal(`error# Assertion failed: expected map not to contain key 'a', but it is found
    got: {"a":1,"b":2}`)

	m.Reset()
	assert.ThatMap(m, testMap).Must().NotContainsKeys([]string{"a"}, "index is 0")
	assert.ThatString(t, m.String()).Equal(`fatal# Assertion failed: expected map not to contain key 'a', but it is found
    got: {"a":1,"b":2}
message: "index is 0"`)
}

func TestMap_ContainsValues(t *testing.T) {
	m := new(internal.MockTestingT)
	testMap := map[string]int{"a": 1, "b": 2}
	assert.ThatMap(m, testMap).ContainsValues([]int{1, 2})
	assert.ThatString(t, m.String()).Equal("")

	m.Reset()
	assert.ThatMap(m, testMap).ContainsValues([]int{3})
	assert.ThatString(t, m.String()).Equal(`error# Assertion failed: expected map to contain value 3, but it is missing
    got: {"a":1,"b":2}`)

	m.Reset()
	assert.ThatMap(m, testMap).Must().ContainsValues([]int{3}, "index is 0")
	assert.ThatString(t, m.String()).Equal(`fatal# Assertion failed: expected map to contain value 3, but it is missing
    got: {"a":1,"b":2}
message: "index is 0"`)
}

func TestMap_NotContainsValues(t *testing.T) {
	m := new(internal.MockTestingT)
	testMap := map[string]int{"a": 1, "b": 2}
	assert.ThatMap(m, testMap).NotContainsValues([]int{3})
	assert.ThatString(t, m.String()).Equal("")

	m.Reset()
	assert.ThatMap(m, testMap).NotContainsValues([]int{1})
	assert.ThatString(t, m.String()).Equal(`error# Assertion failed: expected map not to contain value 1, but it is found
    got: {"a":1,"b":2}`)

	m.Reset()
	assert.ThatMap(m, testMap).Must().NotContainsValues([]int{1}, "index is 0")
	assert.ThatString(t, m.String()).Equal(`fatal# Assertion failed: expected map not to contain value 1, but it is found
    got: {"a":1,"b":2}
message: "index is 0"`)
}

func TestMap_IsSubsetOf(t *testing.T) {
	m := new(internal.MockTestingT)
	testMap := map[string]int{"a": 1}
	superMap := map[string]int{"a": 1, "b": 2}
	assert.ThatMap(m, testMap).IsSubsetOf(superMap)
	assert.ThatString(t, m.String()).Equal("")

	m.Reset()
	assert.ThatMap(m, superMap).IsSubsetOf(testMap)
	assert.ThatString(t, m.String()).Equal(`error# Assertion failed: expected map to be a subset, but unexpected key 'b' is found
    got: {"a":1,"b":2}
 expect: {"a":1}`)

	m.Reset()
	assert.ThatMap(m, superMap).Must().IsSubsetOf(testMap, "index is 0")
	assert.ThatString(t, m.String()).Equal(`fatal# Assertion failed: expected map to be a subset, but unexpected key 'b' is found
    got: {"a":1,"b":2}
 expect: {"a":1}
message: "index is 0"`)
}

func TestMap_IsSupersetOf(t *testing.T) {
	m := new(internal.MockTestingT)
	testMap := map[string]int{"a": 1, "b": 2}
	subMap := map[string]int{"a": 1}
	assert.ThatMap(m, testMap).IsSupersetOf(subMap)
	assert.ThatString(t, m.String()).Equal("")

	m.Reset()
	assert.ThatMap(m, subMap).IsSupersetOf(testMap)
	assert.ThatString(t, m.String()).Equal(`error# Assertion failed: expected map to be a superset, but key 'b' is missing
    got: {"a":1}
 expect: {"a":1,"b":2}`)

	m.Reset()
	assert.ThatMap(m, subMap).Must().IsSupersetOf(testMap, "index is 0")
	assert.ThatString(t, m.String()).Equal(`fatal# Assertion failed: expected map to be a superset, but key 'b' is missing
    got: {"a":1}
 expect: {"a":1,"b":2}
message: "index is 0"`)
}

func TestMap_HasSameKeys(t *testing.T) {
	m := new(internal.MockTestingT)
	map1 := map[string]int{"a": 1, "b": 2}
	map2 := map[string]int{"b": 3, "a": 4}
	assert.ThatMap(m, map1).HasSameKeys(map2)
	assert.ThatString(t, m.String()).Equal("")

	m.Reset()
	assert.ThatMap(m, map1).HasSameKeys(map[string]int{"c": 3})
	assert.ThatString(t, m.String()).Equal(`error# Assertion failed: expected maps to have the same keys, but their lengths are different
    got: {"a":1,"b":2}
 expect: {"c":3}`)

	m.Reset()
	assert.ThatMap(m, map1).Must().HasSameKeys(map[string]int{"c": 3}, "index is 0")
	assert.ThatString(t, m.String()).Equal(`fatal# Assertion failed: expected maps to have the same keys, but their lengths are different
    got: {"a":1,"b":2}
 expect: {"c":3}
message: "index is 0"`)
}

func TestMap_HasSameValues(t *testing.T) {
	m := new(internal.MockTestingT)
	map1 := map[string]int{"a": 1, "b": 2}
	map2 := map[string]int{"x": 1, "y": 2}
	assert.ThatMap(m, map1).HasSameValues(map2)
	assert.ThatString(t, m.String()).Equal("")

	m.Reset()
	assert.ThatMap(m, map1).HasSameValues(map[string]int{"c": 3})
	assert.ThatString(t, m.String()).Equal(`error# Assertion failed: expected maps to have the same values, but their lengths are different
    got: {"a":1,"b":2}
 expect: {"c":3}`)

	m.Reset()
	assert.ThatMap(m, map1).Must().HasSameValues(map[string]int{"c": 3}, "index is 0")
	assert.ThatString(t, m.String()).Equal(`fatal# Assertion failed: expected maps to have the same values, but their lengths are different
    got: {"a":1,"b":2}
 expect: {"c":3}
message: "index is 0"`)
}
