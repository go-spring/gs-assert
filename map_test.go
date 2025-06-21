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
)

func TestMap_Length(t *testing.T) {
	m := new(MockTestingT)
	testMap := map[string]int{"a": 1}
	assert.ThatMap(m, testMap).Length(1)
	assert.ThatString(t, m.String()).Equal("")

	m.Reset()
	assert.ThatMap(m, testMap).Length(0)
	assert.ThatString(t, m.String()).Equal(`length mismatch:
    got: length 1 (map[string]int) map[a:1]
 expect: length 0`)

	m.Reset()
	assert.ThatMap(m, testMap).Length(0, "param (index=0)")
	assert.ThatString(t, m.String()).Equal(`length mismatch:
    got: length 1 (map[string]int) map[a:1]
 expect: length 0
message: param (index=0)`)
}

func TestMap_Equal(t *testing.T) {
	m := new(MockTestingT)
	testMap := map[string]int{"a": 1}
	expectMap := map[string]int{"a": 1}
	assert.ThatMap(m, testMap).Equal(expectMap)
	assert.ThatString(t, m.String()).Equal("")

	m.Reset()
	assert.ThatMap(m, testMap).Equal(map[string]int{"b": 2})
	assert.ThatString(t, m.String()).Equal(`map content mismatch:
    got: key a value (int) 1
 expect: key a value (int) 0`)

	m.Reset()
	assert.ThatMap(m, testMap).Equal(map[string]int{"b": 2}, "param (index=0)")
	assert.ThatString(t, m.String()).Equal(`map content mismatch:
    got: key a value (int) 1
 expect: key a value (int) 0
message: param (index=0)`)
}

func TestMap_NotEqual(t *testing.T) {
	m := new(MockTestingT)
	testMap := map[string]int{"a": 1}
	assert.ThatMap(m, testMap).NotEqual(map[string]int{"b": 2})
	assert.ThatString(t, m.String()).Equal("")

	m.Reset()
	assert.ThatMap(m, testMap).NotEqual(testMap)
	assert.ThatString(t, m.String()).Equal(`maps are equal:
    got: (map[string]int) map[a:1]
 expect: not equal to map[a:1]`)

	m.Reset()
	assert.ThatMap(m, testMap).NotEqual(testMap, "param (index=0)")
	assert.ThatString(t, m.String()).Equal(`maps are equal:
    got: (map[string]int) map[a:1]
 expect: not equal to map[a:1]
message: param (index=0)`)
}

func TestMap_IsEmpty(t *testing.T) {
	m := new(MockTestingT)
	var emptyMap map[string]int
	assert.ThatMap(m, emptyMap).IsEmpty()
	assert.ThatString(t, m.String()).Equal("")

	m.Reset()
	assert.ThatMap(m, map[string]int{"a": 1}).IsEmpty()
	assert.ThatString(t, m.String()).Equal(`map is not empty:
    got: (map[string]int) map[a:1]
 expect: empty map`)

	m.Reset()
	assert.ThatMap(m, map[string]int{"a": 1}).IsEmpty("param (index=0)")
	assert.ThatString(t, m.String()).Equal(`map is not empty:
    got: (map[string]int) map[a:1]
 expect: empty map
message: param (index=0)`)
}

func TestMap_IsNotEmpty(t *testing.T) {
	m := new(MockTestingT)
	testMap := map[string]int{"a": 1}
	assert.ThatMap(m, testMap).IsNotEmpty()
	assert.ThatString(t, m.String()).Equal("")

	m.Reset()
	var emptyMap map[string]int
	assert.ThatMap(m, emptyMap).IsNotEmpty()
	assert.ThatString(t, m.String()).Equal(`map is empty:
    got: (map[string]int) map[]
 expect: non-empty map`)

	m.Reset()
	assert.ThatMap(m, emptyMap).IsNotEmpty("param (index=0)")
	assert.ThatString(t, m.String()).Equal(`map is empty:
    got: (map[string]int) map[]
 expect: non-empty map
message: param (index=0)`)
}

func TestMap_Contains(t *testing.T) {
	m := new(MockTestingT)
	testMap := map[string]int{"a": 1}
	assert.ThatMap(m, testMap).Contains("a")
	assert.ThatString(t, m.String()).Equal("")

	m.Reset()
	assert.ThatMap(m, testMap).Contains("b")
	assert.ThatString(t, m.String()).Equal(`map does not contain the key:
    got: (map[string]int) map[a:1]
 expect: map containing key b`)

	m.Reset()
	assert.ThatMap(m, testMap).Contains("b", "param (index=0)")
	assert.ThatString(t, m.String()).Equal(`map does not contain the key:
    got: (map[string]int) map[a:1]
 expect: map containing key b
message: param (index=0)`)
}

func TestMap_NotContains(t *testing.T) {
	m := new(MockTestingT)
	testMap := map[string]int{"a": 1}
	assert.ThatMap(m, testMap).NotContains("b")
	assert.ThatString(t, m.String()).Equal("")

	m.Reset()
	assert.ThatMap(m, testMap).NotContains("a")
	assert.ThatString(t, m.String()).Equal(`map contains the key:
    got: (map[string]int) map[a:1]
 expect: map not containing key a`)

	m.Reset()
	assert.ThatMap(m, testMap).NotContains("a", "param (index=0)")
	assert.ThatString(t, m.String()).Equal(`map contains the key:
    got: (map[string]int) map[a:1]
 expect: map not containing key a
message: param (index=0)`)
}

func TestMap_ContainsValue(t *testing.T) {
	m := new(MockTestingT)
	testMap := map[string]int{"a": 1}
	assert.ThatMap(m, testMap).ContainsValue(1)
	assert.ThatString(t, m.String()).Equal("")

	m.Reset()
	assert.ThatMap(m, testMap).ContainsValue(2)
	assert.ThatString(t, m.String()).Equal(`map does not contain the value:
    got: (map[string]int) map[a:1]
 expect: map containing value 2`)

	m.Reset()
	assert.ThatMap(m, testMap).ContainsValue(2, "param (index=0)")
	assert.ThatString(t, m.String()).Equal(`map does not contain the value:
    got: (map[string]int) map[a:1]
 expect: map containing value 2
message: param (index=0)`)
}

func TestMap_NotContainsValue(t *testing.T) {
	m := new(MockTestingT)
	testMap := map[string]int{"a": 1}
	assert.ThatMap(m, testMap).NotContainsValue(2)
	assert.ThatString(t, m.String()).Equal("")

	m.Reset()
	assert.ThatMap(m, testMap).NotContainsValue(1)
	assert.ThatString(t, m.String()).Equal(`map contains the value:
    got: (map[string]int) map[a:1]
 expect: map not containing value 1`)

	m.Reset()
	assert.ThatMap(m, testMap).NotContainsValue(1, "param (index=0)")
	assert.ThatString(t, m.String()).Equal(`map contains the value:
    got: (map[string]int) map[a:1]
 expect: map not containing value 1
message: param (index=0)`)
}

func TestMap_HasKeyValue(t *testing.T) {
	m := new(MockTestingT)
	testMap := map[string]int{"a": 1}
	assert.ThatMap(m, testMap).HasKeyValue("a", 1)
	assert.ThatString(t, m.String()).Equal("")

	m.Reset()
	assert.ThatMap(m, testMap).HasKeyValue("a", 2)
	assert.ThatString(t, m.String()).Equal(`key-value pair mismatch:
    got: key a value (int) 1
 expect: key a value (int) 2`)

	m.Reset()
	assert.ThatMap(m, testMap).HasKeyValue("a", 2, "param (index=0)")
	assert.ThatString(t, m.String()).Equal(`key-value pair mismatch:
    got: key a value (int) 1
 expect: key a value (int) 2
message: param (index=0)`)
}

func TestMap_ContainsKeys(t *testing.T) {
	m := new(MockTestingT)
	testMap := map[string]int{"a": 1, "b": 2}
	assert.ThatMap(m, testMap).ContainsKeys([]string{"a", "b"})
	assert.ThatString(t, m.String()).Equal("")

	m.Reset()
	assert.ThatMap(m, testMap).ContainsKeys([]string{"c"})
	assert.ThatString(t, m.String()).Equal(`map does not contain all keys:
    got: (map[string]int) map[a:1 b:2]
 expect: map containing key c`)

	m.Reset()
	assert.ThatMap(m, testMap).ContainsKeys([]string{"c"}, "param (index=0)")
	assert.ThatString(t, m.String()).Equal(`map does not contain all keys:
    got: (map[string]int) map[a:1 b:2]
 expect: map containing key c
message: param (index=0)`)
}

func TestMap_NotContainsKeys(t *testing.T) {
	m := new(MockTestingT)
	testMap := map[string]int{"a": 1, "b": 2}
	assert.ThatMap(m, testMap).NotContainsKeys([]string{"c"})
	assert.ThatString(t, m.String()).Equal("")

	m.Reset()
	assert.ThatMap(m, testMap).NotContainsKeys([]string{"a"})
	assert.ThatString(t, m.String()).Equal(`map contains unexpected key:
    got: (map[string]int) map[a:1 b:2]
 expect: map not containing key a`)

	m.Reset()
	assert.ThatMap(m, testMap).NotContainsKeys([]string{"a"}, "param (index=0)")
	assert.ThatString(t, m.String()).Equal(`map contains unexpected key:
    got: (map[string]int) map[a:1 b:2]
 expect: map not containing key a
message: param (index=0)`)
}

func TestMap_ContainsValues(t *testing.T) {
	m := new(MockTestingT)
	testMap := map[string]int{"a": 1, "b": 2}
	assert.ThatMap(m, testMap).ContainsValues([]int{1, 2})
	assert.ThatString(t, m.String()).Equal("")

	m.Reset()
	assert.ThatMap(m, testMap).ContainsValues([]int{3})
	assert.ThatString(t, m.String()).Equal(`map does not contain all values:
    got: (map[string]int) map[a:1 b:2]
 expect: map containing value 3`)

	m.Reset()
	assert.ThatMap(m, testMap).ContainsValues([]int{3}, "param (index=0)")
	assert.ThatString(t, m.String()).Equal(`map does not contain all values:
    got: (map[string]int) map[a:1 b:2]
 expect: map containing value 3
message: param (index=0)`)
}

func TestMap_NotContainsValues(t *testing.T) {
	m := new(MockTestingT)
	testMap := map[string]int{"a": 1, "b": 2}
	assert.ThatMap(m, testMap).NotContainsValues([]int{3})
	assert.ThatString(t, m.String()).Equal("")

	m.Reset()
	assert.ThatMap(m, testMap).NotContainsValues([]int{1})
	assert.ThatString(t, m.String()).Equal(`map contains unexpected value:
    got: (map[string]int) map[a:1 b:2]
 expect: map not containing value 1`)

	m.Reset()
	assert.ThatMap(m, testMap).NotContainsValues([]int{1}, "param (index=0)")
	assert.ThatString(t, m.String()).Equal(`map contains unexpected value:
    got: (map[string]int) map[a:1 b:2]
 expect: map not containing value 1
message: param (index=0)`)
}

func TestMap_IsSubsetOf(t *testing.T) {
	m := new(MockTestingT)
	testMap := map[string]int{"a": 1}
	superMap := map[string]int{"a": 1, "b": 2}
	assert.ThatMap(m, testMap).IsSubsetOf(superMap)
	assert.ThatString(t, m.String()).Equal("")

	m.Reset()
	assert.ThatMap(m, superMap).IsSubsetOf(testMap)
	assert.ThatString(t, m.String()).Equal(`map is not a subset:
    got: key b value (int) 2
 expect: key b value (int) 0`)

	m.Reset()
	assert.ThatMap(m, superMap).IsSubsetOf(testMap, "param (index=0)")
	assert.ThatString(t, m.String()).Equal(`map is not a subset:
    got: key b value (int) 2
 expect: key b value (int) 0
message: param (index=0)`)
}

func TestMap_IsSupersetOf(t *testing.T) {
	m := new(MockTestingT)
	testMap := map[string]int{"a": 1, "b": 2}
	subMap := map[string]int{"a": 1}
	assert.ThatMap(m, testMap).IsSupersetOf(subMap)
	assert.ThatString(t, m.String()).Equal("")

	m.Reset()
	assert.ThatMap(m, subMap).IsSupersetOf(testMap)
	assert.ThatString(t, m.String()).Equal(`map is not a superset:
    got: key b value (int) 0
 expect: key b value (int) 2`)

	m.Reset()
	assert.ThatMap(m, subMap).IsSupersetOf(testMap, "param (index=0)")
	assert.ThatString(t, m.String()).Equal(`map is not a superset:
    got: key b value (int) 0
 expect: key b value (int) 2
message: param (index=0)`)
}

func TestMap_HasSameKeys(t *testing.T) {
	m := new(MockTestingT)
	map1 := map[string]int{"a": 1, "b": 2}
	map2 := map[string]int{"b": 3, "a": 4}
	assert.ThatMap(m, map1).HasSameKeys(map2)
	assert.ThatString(t, m.String()).Equal("")

	m.Reset()
	assert.ThatMap(m, map1).HasSameKeys(map[string]int{"c": 3})
	assert.ThatString(t, m.String()).Equal(`map key count mismatch:
    got: count 2 (map[string]int) map[a:1 b:2]
 expect: count 1`)

	m.Reset()
	assert.ThatMap(m, map1).HasSameKeys(map[string]int{"c": 3}, "param (index=0)")
	assert.ThatString(t, m.String()).Equal(`map key count mismatch:
    got: count 2 (map[string]int) map[a:1 b:2]
 expect: count 1
message: param (index=0)`)
}

func TestMap_HasSameValues(t *testing.T) {
	m := new(MockTestingT)
	map1 := map[string]int{"a": 1, "b": 2}
	map2 := map[string]int{"x": 1, "y": 2}
	assert.ThatMap(m, map1).HasSameValues(map2)
	assert.ThatString(t, m.String()).Equal("")

	m.Reset()
	assert.ThatMap(m, map1).HasSameValues(map[string]int{"c": 3})
	assert.ThatString(t, m.String()).Equal(`map value count mismatch:
    got: count 2 (map[string]int) map[a:1 b:2]
 expect: count 1`)

	m.Reset()
	assert.ThatMap(m, map1).HasSameValues(map[string]int{"c": 3}, "param (index=0)")
	assert.ThatString(t, m.String()).Equal(`map value count mismatch:
    got: count 2 (map[string]int) map[a:1 b:2]
 expect: count 1
message: param (index=0)`)
}
