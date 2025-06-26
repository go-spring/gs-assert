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

func TestSlice_Nil(t *testing.T) {
	m := new(internal.MockTestingT)
	var s []int = nil
	assert.ThatSlice(m, s).Nil()
	assert.ThatString(t, m.String()).Equal("")

	m.Reset()
	s = []int{1, 2}
	assert.ThatSlice(m, s).OutputValueAsStd().Nil()
	assert.ThatString(t, m.String()).Equal(`error# Assertion failed: expected slice to be nil, but it is not
    got: [1 2]`)

	m.Reset()
	assert.ThatSlice(m, s).OutputValueAsJSON().Must().Nil("index is 0")
	assert.ThatString(t, m.String()).Equal(`fatal# Assertion failed: expected slice to be nil, but it is not
    got: [1,2]
message: "index is 0"`)
}

func TestSlice_NotNil(t *testing.T) {
	m := new(internal.MockTestingT)
	var s = []int{1, 2}
	assert.ThatSlice(m, s).NotNil()
	assert.ThatString(t, m.String()).Equal("")

	m.Reset()
	s = nil
	assert.ThatSlice(m, s).OutputValueAsStd().NotNil()
	assert.ThatString(t, m.String()).Equal(`error# Assertion failed: expected slice not to be nil, but it is
    got: []`)

	m.Reset()
	assert.ThatSlice(m, s).OutputValueAsJSON().Must().NotNil("index is 0")
	assert.ThatString(t, m.String()).Equal(`fatal# Assertion failed: expected slice not to be nil, but it is
    got: null
message: "index is 0"`)
}

func TestSlice_Empty(t *testing.T) {
	m := new(internal.MockTestingT)
	assert.ThatSlice(m, []int{}).Empty()
	assert.ThatString(t, m.String()).Equal("")

	m.Reset()
	assert.ThatSlice(m, []int{1, 2}).OutputValueAsStd().Empty()
	assert.ThatString(t, m.String()).Equal(`error# Assertion failed: expected slice to be empty, but it is not
    got: [1 2]`)

	m.Reset()
	assert.ThatSlice(m, []int{1, 2}).OutputValueAsJSON().Must().Empty("index is 0")
	assert.ThatString(t, m.String()).Equal(`fatal# Assertion failed: expected slice to be empty, but it is not
    got: [1,2]
message: "index is 0"`)
}

func TestSlice_NotEmpty(t *testing.T) {
	m := new(internal.MockTestingT)
	assert.ThatSlice(m, []string{"hello"}).NotEmpty()
	assert.ThatString(t, m.String()).Equal("")

	m.Reset()
	assert.ThatSlice(m, []string(nil)).OutputValueAsStd().NotEmpty()
	assert.ThatString(t, m.String()).Equal(`error# Assertion failed: expected slice not to be empty, but it is
    got: []`)

	m.Reset()
	assert.ThatSlice(m, []string(nil)).OutputValueAsJSON().Must().NotEmpty("index is 0")
	assert.ThatString(t, m.String()).Equal(`fatal# Assertion failed: expected slice not to be empty, but it is
    got: null
message: "index is 0"`)
}

func TestSlice_Length(t *testing.T) {
	m := new(internal.MockTestingT)
	assert.ThatSlice(m, []float64{1.1, 2.2, 3.3}).Length(3)
	assert.ThatString(t, m.String()).Equal("")

	m.Reset()
	assert.ThatSlice(m, []float64{1.1, 2.2}).OutputValueAsStd().Length(3)
	assert.ThatString(t, m.String()).Equal(`error# Assertion failed: expected slice to have length 3, but it has length 2
    got: [1.1 2.2]`)

	m.Reset()
	assert.ThatSlice(m, []float64{1.1}).OutputValueAsJSON().Must().Length(0, "index is 0")
	assert.ThatString(t, m.String()).Equal(`fatal# Assertion failed: expected slice to have length 0, but it has length 1
    got: [1.1]
message: "index is 0"`)
}

func TestSlice_Equal(t *testing.T) {
	m := new(internal.MockTestingT)
	assert.ThatSlice(m, []int{1, 2, 3}).Equal([]int{1, 2, 3})
	assert.ThatString(t, m.String()).Equal("")

	m.Reset()
	assert.ThatSlice(m, []int{1, 2, 3}).OutputValueAsStd().Equal([]int{4, 5})
	assert.ThatString(t, m.String()).Equal(`error# Assertion failed: expected slices to be equal, but their lengths are different
    got: [1 2 3]
 expect: [4 5]`)

	m.Reset()
	assert.ThatSlice(m, []int{1, 2, 3}).OutputValueAsJSON().Must().Equal([]int{1, 2, 4}, "index is 0")
	assert.ThatString(t, m.String()).Equal(`fatal# Assertion failed: expected slices to be equal, but values at index 2 are different
    got: [1,2,3]
 expect: [1,2,4]
message: "index is 0"`)
}

func TestSlice_NotEqual(t *testing.T) {
	m := new(internal.MockTestingT)
	assert.ThatSlice(m, []string{"a", "b"}).NotEqual([]string{"c", "d"})
	assert.ThatString(t, m.String()).Equal("")

	m.Reset()
	assert.ThatSlice(m, []string{"a", "b"}).OutputValueAsStd().NotEqual([]string{"a", "b"})
	assert.ThatString(t, m.String()).Equal(`error# Assertion failed: expected slices to be different, but they are equal
    got: [a b]`)

	m.Reset()
	assert.ThatSlice(m, []string{"a", "b"}).OutputValueAsJSON().Must().NotEqual([]string{"a", "b"}, "index is 0")
	assert.ThatString(t, m.String()).Equal(`fatal# Assertion failed: expected slices to be different, but they are equal
    got: ["a","b"]
message: "index is 0"`)
}

func TestSlice_Contains(t *testing.T) {
	m := new(internal.MockTestingT)
	assert.ThatSlice(m, []int{1, 2, 3}).Contains(2)
	assert.ThatString(t, m.String()).Equal("")

	m.Reset()
	assert.ThatSlice(m, []int{1, 2, 3}).OutputValueAsStd().Contains(4)
	assert.ThatString(t, m.String()).Equal(`error# Assertion failed: expected slice to contain element 4, but it is missing
    got: [1 2 3]`)

	m.Reset()
	assert.ThatSlice(m, []int{1, 2, 3}).OutputValueAsJSON().Must().Contains(4, "index is 0")
	assert.ThatString(t, m.String()).Equal(`fatal# Assertion failed: expected slice to contain element 4, but it is missing
    got: [1,2,3]
message: "index is 0"`)
}

func TestSlice_NotContains(t *testing.T) {
	m := new(internal.MockTestingT)
	assert.ThatSlice(m, []int{1, 2, 3}).NotContains(4)
	assert.ThatString(t, m.String()).Equal("")

	m.Reset()
	assert.ThatSlice(m, []int{1, 2, 3}).OutputValueAsStd().NotContains(2)
	assert.ThatString(t, m.String()).Equal(`error# Assertion failed: expected slice not to contain element 2, but it is found
    got: [1 2 3]`)

	m.Reset()
	assert.ThatSlice(m, []int{1, 2, 3}).OutputValueAsJSON().Must().NotContains(2, "index is 0")
	assert.ThatString(t, m.String()).Equal(`fatal# Assertion failed: expected slice not to contain element 2, but it is found
    got: [1,2,3]
message: "index is 0"`)
}

func TestSlice_ContainsSlice(t *testing.T) {
	m := new(internal.MockTestingT)
	assert.ThatSlice(m, []int{1, 2, 3, 4}).ContainsSlice([]int{2, 3})
	assert.ThatString(t, m.String()).Equal("")

	m.Reset()
	assert.ThatSlice(m, []int{1, 2, 3, 4}).OutputValueAsStd().ContainsSlice([]int{2, 4})
	assert.ThatString(t, m.String()).Equal(`error# Assertion failed: expected slice to contain sub-slice, but it is not
    got: [1 2 3 4]
    sub: [2 4]`)

	m.Reset()
	assert.ThatSlice(m, []int{1, 2, 3, 4}).OutputValueAsJSON().Must().ContainsSlice([]int{2, 4}, "index is 0")
	assert.ThatString(t, m.String()).Equal(`fatal# Assertion failed: expected slice to contain sub-slice, but it is not
    got: [1,2,3,4]
    sub: [2,4]
message: "index is 0"`)
}

func TestSlice_NotContainsSlice(t *testing.T) {
	m := new(internal.MockTestingT)
	assert.ThatSlice(m, []int{1, 2, 3, 4}).NotContainsSlice([]int{2, 4})
	assert.ThatString(t, m.String()).Equal("")

	m.Reset()
	assert.ThatSlice(m, []int{1, 2, 3, 4}).OutputValueAsStd().NotContainsSlice([]int{2, 3})
	assert.ThatString(t, m.String()).Equal(`error# Assertion failed: expected slice not to contain sub-slice, but it is
    got: [1 2 3 4]
    sub: [2 3]`)

	m.Reset()
	assert.ThatSlice(m, []int{1, 2, 3, 4}).OutputValueAsJSON().Must().NotContainsSlice([]int{2, 3}, "index is 0")
	assert.ThatString(t, m.String()).Equal(`fatal# Assertion failed: expected slice not to contain sub-slice, but it is
    got: [1,2,3,4]
    sub: [2,3]
message: "index is 0"`)
}

func TestSlice_HasPrefix(t *testing.T) {
	m := new(internal.MockTestingT)
	assert.ThatSlice(m, []int{1, 2, 3}).HasPrefix([]int{1, 2})
	assert.ThatString(t, m.String()).Equal("")

	m.Reset()
	assert.ThatSlice(m, []int{1, 2, 3}).OutputValueAsStd().HasPrefix([]int{1, 2, 3, 4})
	assert.ThatString(t, m.String()).Equal(`error# Assertion failed: expected slice to start with prefix, but it is not
    got: [1 2 3]
 prefix: [1 2 3 4]`)

	m.Reset()
	assert.ThatSlice(m, []int{1, 2, 3}).OutputValueAsJSON().Must().HasPrefix([]int{2, 3}, "index is 0")
	assert.ThatString(t, m.String()).Equal(`fatal# Assertion failed: expected slice to start with prefix, but it is not
    got: [1,2,3]
 prefix: [2,3]
message: "index is 0"`)
}

func TestSlice_HasSuffix(t *testing.T) {
	m := new(internal.MockTestingT)
	assert.ThatSlice(m, []int{1, 2, 3}).HasSuffix([]int{2, 3})
	assert.ThatString(t, m.String()).Equal("")

	m.Reset()
	assert.ThatSlice(m, []int{1, 2, 3}).OutputValueAsStd().HasSuffix([]int{1, 2, 3, 4})
	assert.ThatString(t, m.String()).Equal(`error# Assertion failed: expected slice to end with suffix, but it is not
    got: [1 2 3]
 suffix: [1 2 3 4]`)

	m.Reset()
	assert.ThatSlice(m, []int{1, 2, 3}).OutputValueAsJSON().Must().HasSuffix([]int{1, 2}, "index is 0")
	assert.ThatString(t, m.String()).Equal(`fatal# Assertion failed: expected slice to end with suffix, but it is not
    got: [1,2,3]
 suffix: [1,2]
message: "index is 0"`)
}

func TestSlice_AllUnique(t *testing.T) {
	m := new(internal.MockTestingT)
	assert.ThatSlice(m, []int{1, 2, 3}).AllUnique()
	assert.ThatString(t, m.String()).Equal("")

	m.Reset()
	assert.ThatSlice(m, []int{1, 2, 1}).OutputValueAsStd().AllUnique()
	assert.ThatString(t, m.String()).Equal(`error# Assertion failed: expected all elements in the slice to be unique, but duplicate element 1 is found
    got: [1 2 1]`)

	m.Reset()
	assert.ThatSlice(m, []int{1, 2, 1}).OutputValueAsJSON().Must().AllUnique("index is 0")
	assert.ThatString(t, m.String()).Equal(`fatal# Assertion failed: expected all elements in the slice to be unique, but duplicate element 1 is found
    got: [1,2,1]
message: "index is 0"`)
}

func TestSlice_AllMatches(t *testing.T) {
	m := new(internal.MockTestingT)
	assert.ThatSlice(m, []int{2, 4, 6, 8}).AllMatches(func(n int) bool { return n%2 == 0 })
	assert.ThatString(t, m.String()).Equal("")

	m.Reset()
	assert.ThatSlice(m, []int{2, 3, 4, 6}).OutputValueAsStd().AllMatches(func(n int) bool { return n%2 == 0 })
	assert.ThatString(t, m.String()).Equal(`error# Assertion failed: expected all elements in the slice to satisfy the condition, but element 3 does not
    got: [2 3 4 6]`)

	m.Reset()
	assert.ThatSlice(m, []int{2, 3, 4, 6}).OutputValueAsJSON().Must().AllMatches(func(n int) bool { return n%2 == 0 }, "index is 0")
	assert.ThatString(t, m.String()).Equal(`fatal# Assertion failed: expected all elements in the slice to satisfy the condition, but element 3 does not
    got: [2,3,4,6]
message: "index is 0"`)
}

func TestSlice_AnyMatches(t *testing.T) {
	m := new(internal.MockTestingT)
	assert.ThatSlice(m, []int{1, 2, 3, 5}).AnyMatches(func(n int) bool { return n%2 == 0 })
	assert.ThatString(t, m.String()).Equal("")

	m.Reset()
	assert.ThatSlice(m, []int{1, 3, 5, 7}).OutputValueAsStd().AnyMatches(func(n int) bool { return n%2 == 0 })
	assert.ThatString(t, m.String()).Equal(`error# Assertion failed: expected at least one element in the slice to satisfy the condition, but none do
    got: [1 3 5 7]`)

	m.Reset()
	assert.ThatSlice(m, []int{1, 3, 5, 7}).OutputValueAsJSON().Must().AnyMatches(func(n int) bool { return n%2 == 0 }, "index is 0")
	assert.ThatString(t, m.String()).Equal(`fatal# Assertion failed: expected at least one element in the slice to satisfy the condition, but none do
    got: [1,3,5,7]
message: "index is 0"`)
}

func TestSlice_NoneMatches(t *testing.T) {
	m := new(internal.MockTestingT)
	assert.ThatSlice(m, []int{1, 3, 5, 7}).NoneMatches(func(n int) bool { return n%2 == 0 })
	assert.ThatString(t, m.String()).Equal("")

	m.Reset()
	assert.ThatSlice(m, []int{1, 2, 3, 5}).OutputValueAsStd().NoneMatches(func(n int) bool { return n%2 == 0 })
	assert.ThatString(t, m.String()).Equal(`error# Assertion failed: expected no element in the slice to satisfy the condition, but element 2 does
    got: [1 2 3 5]`)

	m.Reset()
	assert.ThatSlice(m, []int{1, 2, 3, 5}).OutputValueAsJSON().Must().NoneMatches(func(n int) bool { return n%2 == 0 }, "index is 0")
	assert.ThatString(t, m.String()).Equal(`fatal# Assertion failed: expected no element in the slice to satisfy the condition, but element 2 does
    got: [1,2,3,5]
message: "index is 0"`)
}
