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

func TestSlice_Length(t *testing.T) {
	m := new(MockTestingT)
	assert.ThatSlice(m, []float64{1.1, 2.2, 3.3}).Length(3)
	assert.ThatString(t, m.String()).Equal("")

	m.Reset()
	assert.ThatSlice(m, []float64{1.1, 2.2}).Length(3)
	assert.ThatString(t, m.String()).Equal(`Assertion failed: length mismatch:
    got: length 2 ([]float64) [1.1 2.2]
 expect: length 3`)

	m.Reset()
	assert.ThatSlice(m, []float64{1.1}).Length(0, "index is 0")
	assert.ThatString(t, m.String()).Equal(`Assertion failed: length mismatch:
    got: length 1 ([]float64) [1.1]
 expect: length 0
message: index is 0`)
}

func TestSlice_Equal(t *testing.T) {
	m := new(MockTestingT)
	assert.ThatSlice(m, []int{1, 2, 3}).Equal([]int{1, 2, 3})
	assert.ThatString(t, m.String()).Equal("")

	m.Reset()
	assert.ThatSlice(m, []int{1, 2, 3}).Equal([]int{4, 5})
	assert.ThatString(t, m.String()).Equal(`Assertion failed: slices not equal:
    got: ([]int) [1 2 3]
 expect: ([]int) [4 5]`)

	m.Reset()
	assert.ThatSlice(m, []int{1, 2, 3}).Equal([]int{4, 5}, "index is 0")
	assert.ThatString(t, m.String()).Equal(`Assertion failed: slices not equal:
    got: ([]int) [1 2 3]
 expect: ([]int) [4 5]
message: index is 0`)
}

func TestSlice_NotEqual(t *testing.T) {
	m := new(MockTestingT)
	assert.ThatSlice(m, []string{"a", "b"}).NotEqual([]string{"c", "d"})
	assert.ThatString(t, m.String()).Equal("")

	m.Reset()
	assert.ThatSlice(m, []string{"a", "b"}).NotEqual([]string{"a", "b"})
	assert.ThatString(t, m.String()).Equal(`Assertion failed: slices are equal:
    got: ([]string) [a b]
 expect: not equal to [a b]`)

	m.Reset()
	assert.ThatSlice(m, []string{"a", "b"}).NotEqual([]string{"a", "b"}, "index is 0")
	assert.ThatString(t, m.String()).Equal(`Assertion failed: slices are equal:
    got: ([]string) [a b]
 expect: not equal to [a b]
message: index is 0`)
}

func TestSlice_IsNil(t *testing.T) {
	m := new(MockTestingT)
	var s []int = nil
	assert.ThatSlice(m, s).IsNil()
	assert.ThatString(t, m.String()).Equal("")

	m.Reset()
	s = []int{1, 2}
	assert.ThatSlice(m, s).IsNil()
	assert.ThatString(t, m.String()).Equal(`Assertion failed: got [1 2] is not nil:
    got: ([]int) [1 2]
 expect: nil slice`)

	m.Reset()
	assert.ThatSlice(m, s).IsNil("index is 0")
	assert.ThatString(t, m.String()).Equal(`Assertion failed: got [1 2] is not nil:
    got: ([]int) [1 2]
 expect: nil slice
message: index is 0`)
}

func TestSlice_IsNotNil(t *testing.T) {
	m := new(MockTestingT)
	var s = []int{1, 2}
	assert.ThatSlice(m, s).IsNotNil()
	assert.ThatString(t, m.String()).Equal("")

	m.Reset()
	s = nil
	assert.ThatSlice(m, s).IsNotNil()
	assert.ThatString(t, m.String()).Equal(`Assertion failed: got [] is nil:
    got: ([]int) []
 expect: non-nil slice`)

	m.Reset()
	assert.ThatSlice(m, s).IsNotNil("index is 0")
	assert.ThatString(t, m.String()).Equal(`Assertion failed: got [] is nil:
    got: ([]int) []
 expect: non-nil slice
message: index is 0`)
}

func TestSlice_IsEmpty(t *testing.T) {
	m := new(MockTestingT)
	assert.ThatSlice(m, []int{}).IsEmpty()
	assert.ThatString(t, m.String()).Equal("")

	m.Reset()
	assert.ThatSlice(m, []int{1, 2}).IsEmpty()
	assert.ThatString(t, m.String()).Equal(`Assertion failed: slice is not empty:
    got: ([]int) [1 2]
 expect: empty slice`)

	m.Reset()
	assert.ThatSlice(m, []int{1, 2}).IsEmpty("index is 0")
	assert.ThatString(t, m.String()).Equal(`Assertion failed: slice is not empty:
    got: ([]int) [1 2]
 expect: empty slice
message: index is 0`)
}

func TestSlice_IsNotEmpty(t *testing.T) {
	m := new(MockTestingT)
	assert.ThatSlice(m, []string{"hello"}).IsNotEmpty()
	assert.ThatString(t, m.String()).Equal("")

	m.Reset()
	assert.ThatSlice(m, []string{}).IsNotEmpty()
	assert.ThatString(t, m.String()).Equal(`Assertion failed: slice is empty:
    got: ([]string) []
 expect: non-empty slice`)

	m.Reset()
	assert.ThatSlice(m, []string{}).IsNotEmpty("index is 0")
	assert.ThatString(t, m.String()).Equal(`Assertion failed: slice is empty:
    got: ([]string) []
 expect: non-empty slice
message: index is 0`)
}

func TestSlice_Contains(t *testing.T) {
	m := new(MockTestingT)
	assert.ThatSlice(m, []int{1, 2, 3}).Contains(2)
	assert.ThatString(t, m.String()).Equal("")

	m.Reset()
	assert.ThatSlice(m, []int{1, 2, 3}).Contains(4)
	assert.ThatString(t, m.String()).Equal(`Assertion failed: slice does not contain the expected element:
    got: ([]int) [1 2 3]
 expect: to contain element 4`)

	m.Reset()
	assert.ThatSlice(m, []int{1, 2, 3}).Contains(4, "index is 0")
	assert.ThatString(t, m.String()).Equal(`Assertion failed: slice does not contain the expected element:
    got: ([]int) [1 2 3]
 expect: to contain element 4
message: index is 0`)
}

func TestSlice_NotContains(t *testing.T) {
	m := new(MockTestingT)
	assert.ThatSlice(m, []int{1, 2, 3}).NotContains(4)
	assert.ThatString(t, m.String()).Equal("")

	m.Reset()
	assert.ThatSlice(m, []int{1, 2, 3}).NotContains(2)
	assert.ThatString(t, m.String()).Equal(`Assertion failed: slice contains the unexpected element:
    got: ([]int) [1 2 3]
 expect: not to contain element 2`)

	m.Reset()
	assert.ThatSlice(m, []int{1, 2, 3}).NotContains(2, "index is 0")
	assert.ThatString(t, m.String()).Equal(`Assertion failed: slice contains the unexpected element:
    got: ([]int) [1 2 3]
 expect: not to contain element 2
message: index is 0`)
}

func TestSlice_SubSlice(t *testing.T) {
	m := new(MockTestingT)
	assert.ThatSlice(m, []int{1, 2, 3, 4}).SubSlice([]int{2, 3})
	assert.ThatString(t, m.String()).Equal("")

	m.Reset()
	assert.ThatSlice(m, []int{1, 2, 3, 4}).SubSlice([]int{2, 4})
	assert.ThatString(t, m.String()).Equal(`Assertion failed: slice does not contain sub-slice:
    got: ([]int) [1 2 3 4]
 expect: to contain sub-slice [2 4]`)

	m.Reset()
	assert.ThatSlice(m, []int{1, 2, 3, 4}).SubSlice([]int{2, 4}, "index is 0")
	assert.ThatString(t, m.String()).Equal(`Assertion failed: slice does not contain sub-slice:
    got: ([]int) [1 2 3 4]
 expect: to contain sub-slice [2 4]
message: index is 0`)
}

func TestSlice_NotSubSlice(t *testing.T) {
	m := new(MockTestingT)
	assert.ThatSlice(m, []int{1, 2, 3, 4}).NotSubSlice([]int{2, 4})
	assert.ThatString(t, m.String()).Equal("")

	m.Reset()
	assert.ThatSlice(m, []int{1, 2, 3, 4}).NotSubSlice([]int{2, 3})
	assert.ThatString(t, m.String()).Equal(`Assertion failed: slice contains sub-slice:
    got: ([]int) [1 2 3 4]
 expect: not to contain sub-slice [2 3]`)

	m.Reset()
	assert.ThatSlice(m, []int{1, 2, 3, 4}).NotSubSlice([]int{2, 3}, "index is 0")
	assert.ThatString(t, m.String()).Equal(`Assertion failed: slice contains sub-slice:
    got: ([]int) [1 2 3 4]
 expect: not to contain sub-slice [2 3]
message: index is 0`)
}

func TestSlice_HasPrefix(t *testing.T) {
	m := new(MockTestingT)
	assert.ThatSlice(m, []int{1, 2, 3}).HasPrefix([]int{1, 2})
	assert.ThatString(t, m.String()).Equal("")

	m.Reset()
	assert.ThatSlice(m, []int{1, 2, 3}).HasPrefix([]int{2, 3})
	assert.ThatString(t, m.String()).Equal(`Assertion failed: slice does not start with the expected prefix:
    got: ([]int) [1 2 3]
 expect: to start with [2 3]`)

	m.Reset()
	assert.ThatSlice(m, []int{1, 2, 3}).HasPrefix([]int{2, 3}, "index is 0")
	assert.ThatString(t, m.String()).Equal(`Assertion failed: slice does not start with the expected prefix:
    got: ([]int) [1 2 3]
 expect: to start with [2 3]
message: index is 0`)
}

func TestSlice_HasSuffix(t *testing.T) {
	m := new(MockTestingT)
	assert.ThatSlice(m, []int{1, 2, 3}).HasSuffix([]int{2, 3})
	assert.ThatString(t, m.String()).Equal("")

	m.Reset()
	assert.ThatSlice(m, []int{1, 2, 3}).HasSuffix([]int{1, 2})
	assert.ThatString(t, m.String()).Equal(`Assertion failed: slice does not end with the expected suffix:
    got: ([]int) [1 2 3]
 expect: to end with [1 2]`)

	m.Reset()
	assert.ThatSlice(m, []int{1, 2, 3}).HasSuffix([]int{1, 2}, "index is 0")
	assert.ThatString(t, m.String()).Equal(`Assertion failed: slice does not end with the expected suffix:
    got: ([]int) [1 2 3]
 expect: to end with [1 2]
message: index is 0`)
}

func TestSlice_IsIncreasing(t *testing.T) {
	m := new(MockTestingT)
	assert.ThatSlice(m, []int{1, 2, 3, 4}).IsIncreasing()
	assert.ThatString(t, m.String()).Equal("")

	m.Reset()
	assert.ThatSlice(m, []int{1, 3, 2, 4}).IsIncreasing()
	assert.ThatString(t, m.String()).Equal(`Assertion failed: slice is not increasing at 2:
    got: ([]int) [1 3 2 4]
 expect: strictly increasing order`)

	m.Reset()
	assert.ThatSlice(m, []int{1, 3, 2, 4}).IsIncreasing("index is 0")
	assert.ThatString(t, m.String()).Equal(`Assertion failed: slice is not increasing at 2:
    got: ([]int) [1 3 2 4]
 expect: strictly increasing order
message: index is 0`)
}

func TestSlice_NonIncreasing(t *testing.T) {
	m := new(MockTestingT)
	assert.ThatSlice(m, []int{4, 3, 2, 1}).NonIncreasing()
	assert.ThatString(t, m.String()).Equal("")

	m.Reset()
	assert.ThatSlice(m, []int{1, 2, 3, 4}).NonIncreasing()
	assert.ThatString(t, m.String()).Equal(`Assertion failed: slice is increasing at 1:
    got: ([]int) [1 2 3 4]
 expect: not strictly increasing`)

	m.Reset()
	assert.ThatSlice(m, []int{1, 2, 3, 4}).NonIncreasing("index is 0")
	assert.ThatString(t, m.String()).Equal(`Assertion failed: slice is increasing at 1:
    got: ([]int) [1 2 3 4]
 expect: not strictly increasing
message: index is 0`)
}

func TestSlice_IsDecreasing(t *testing.T) {
	m := new(MockTestingT)
	assert.ThatSlice(m, []int{4, 3, 2, 1}).IsDecreasing()
	assert.ThatString(t, m.String()).Equal("")

	m.Reset()
	assert.ThatSlice(m, []int{4, 5, 2, 1}).IsDecreasing()
	assert.ThatString(t, m.String()).Equal(`Assertion failed: slice is not decreasing at 1:
    got: ([]int) [4 5 2 1]
 expect: strictly decreasing order`)

	m.Reset()
	assert.ThatSlice(m, []int{4, 5, 2, 1}).IsDecreasing("index is 0")
	assert.ThatString(t, m.String()).Equal(`Assertion failed: slice is not decreasing at 1:
    got: ([]int) [4 5 2 1]
 expect: strictly decreasing order
message: index is 0`)
}

func TestSlice_NonDecreasing(t *testing.T) {
	m := new(MockTestingT)
	assert.ThatSlice(m, []int{4, 5, 6, 7}).NonDecreasing()
	assert.ThatString(t, m.String()).Equal("")

	m.Reset()
	assert.ThatSlice(m, []int{7, 6, 5, 4}).NonDecreasing()
	assert.ThatString(t, m.String()).Equal(`Assertion failed: slice is decreasing at 1:
    got: ([]int) [7 6 5 4]
 expect: not strictly decreasing`)

	m.Reset()
	assert.ThatSlice(m, []int{7, 6, 5, 4}).NonDecreasing("index is 0")
	assert.ThatString(t, m.String()).Equal(`Assertion failed: slice is decreasing at 1:
    got: ([]int) [7 6 5 4]
 expect: not strictly decreasing
message: index is 0`)
}

func TestSlice_IsSorted(t *testing.T) {
	m := new(MockTestingT)
	assert.ThatSlice(m, []int{1, 2, 3, 4}).IsSorted()
	assert.ThatString(t, m.String()).Equal("")

	m.Reset()
	assert.ThatSlice(m, []int{1, 3, 2, 4}).IsSorted()
	assert.ThatString(t, m.String()).Equal(`Assertion failed: slice is not sorted in ascending order at 2:
    got: ([]int) [1 3 2 4]
 expect: sorted in ascending order`)

	m.Reset()
	assert.ThatSlice(m, []int{1, 3, 2, 4}).IsSorted("index is 0")
	assert.ThatString(t, m.String()).Equal(`Assertion failed: slice is not sorted in ascending order at 2:
    got: ([]int) [1 3 2 4]
 expect: sorted in ascending order
message: index is 0`)
}

func TestSlice_IsSortedDescending(t *testing.T) {
	m := new(MockTestingT)
	assert.ThatSlice(m, []int{4, 3, 2, 1}).IsSortedDescending()
	assert.ThatString(t, m.String()).Equal("")

	m.Reset()
	assert.ThatSlice(m, []int{4, 5, 2, 1}).IsSortedDescending()
	assert.ThatString(t, m.String()).Equal(`Assertion failed: slice is not sorted in descending order at 1:
    got: ([]int) [4 5 2 1]
 expect: sorted in descending order`)

	m.Reset()
	assert.ThatSlice(m, []int{4, 5, 2, 1}).IsSortedDescending("index is 0")
	assert.ThatString(t, m.String()).Equal(`Assertion failed: slice is not sorted in descending order at 1:
    got: ([]int) [4 5 2 1]
 expect: sorted in descending order
message: index is 0`)
}

func TestSlice_IsUnique(t *testing.T) {
	m := new(MockTestingT)
	assert.ThatSlice(m, []int{1, 2, 3}).IsUnique()
	assert.ThatString(t, m.String()).Equal("")

	m.Reset()
	assert.ThatSlice(m, []int{1, 2, 1}).IsUnique()
	assert.ThatString(t, m.String()).Equal(`Assertion failed: duplicate element found at 1:
    got: ([]int) [1 2 1]
 expect: all elements to be unique`)

	m.Reset()
	assert.ThatSlice(m, []int{1, 2, 1}).IsUnique("index is 0")
	assert.ThatString(t, m.String()).Equal(`Assertion failed: duplicate element found at 1:
    got: ([]int) [1 2 1]
 expect: all elements to be unique
message: index is 0`)
}

func TestSlice_IsUniqueBy(t *testing.T) {
	m := new(MockTestingT)
	assert.ThatSlice(m, []string{"app", "banana", "strawberry"}).IsUniqueBy(func(s string) interface{} {
		return len(s)
	})
	assert.ThatString(t, m.String()).Equal("")

	m.Reset()
	assert.ThatSlice(m, []string{"apple", "grape", "orange"}).IsUniqueBy(func(s string) interface{} {
		return len(s)
	})
	assert.ThatString(t, m.String()).Equal(`Assertion failed: duplicate element based on key function:
    got: ([]string) [apple grape orange]
 expect: all elements to be unique by length`)

	m.Reset()
	assert.ThatSlice(m, []string{"apple", "grape", "orange"}).IsUniqueBy(func(s string) interface{} {
		return len(s)
	}, "index is 0")
	assert.ThatString(t, m.String()).Equal(`Assertion failed: duplicate element based on key function:
    got: ([]string) [apple grape orange]
 expect: all elements to be unique by length
message: index is 0`)
}

func TestSlice_All(t *testing.T) {
	m := new(MockTestingT)
	assert.ThatSlice(m, []int{2, 4, 6, 8}).All(func(n int) bool { return n%2 == 0 })
	assert.ThatString(t, m.String()).Equal("")

	m.Reset()
	assert.ThatSlice(m, []int{2, 3, 4, 6}).All(func(n int) bool { return n%2 == 0 })
	assert.ThatString(t, m.String()).Equal(`Assertion failed: element does not satisfy condition:
    got: ([]int) [2 3 4 6]
 expect: all elements should satisfy condition`)

	m.Reset()
	assert.ThatSlice(m, []int{2, 3, 4, 6}).All(func(n int) bool { return n%2 == 0 }, "index is 0")
	assert.ThatString(t, m.String()).Equal(`Assertion failed: element does not satisfy condition:
    got: ([]int) [2 3 4 6]
 expect: all elements should satisfy condition
message: index is 0`)
}

func TestSlice_Any(t *testing.T) {
	m := new(MockTestingT)
	assert.ThatSlice(m, []int{1, 3, 5, 7}).Any(func(n int) bool { return n%2 == 0 })
	assert.ThatString(t, m.String()).Equal(`Assertion failed: no element satisfies the condition:
    got: ([]int) [1 3 5 7]
 expect: any element should satisfy condition`)

	m.Reset()
	assert.ThatSlice(m, []int{1, 2, 3, 5}).Any(func(n int) bool { return n%2 == 0 })
	assert.ThatString(t, m.String()).Equal("")

	m.Reset()
	assert.ThatSlice(m, []int{1, 3, 5, 7}).Any(func(n int) bool { return n%2 == 0 }, "index is 0")
	assert.ThatString(t, m.String()).Equal(`Assertion failed: no element satisfies the condition:
    got: ([]int) [1 3 5 7]
 expect: any element should satisfy condition
message: index is 0`)
}

func TestSlice_None(t *testing.T) {
	m := new(MockTestingT)
	assert.ThatSlice(m, []int{1, 3, 5, 7}).None(func(n int) bool { return n%2 == 0 })
	assert.ThatString(t, m.String()).Equal("")

	m.Reset()
	assert.ThatSlice(m, []int{1, 2, 3, 5}).None(func(n int) bool { return n%2 == 0 })
	assert.ThatString(t, m.String()).Equal(`Assertion failed: element satisfies the condition:
    got: ([]int) [1 2 3 5]
 expect: no element should satisfy condition`)

	m.Reset()
	assert.ThatSlice(m, []int{1, 2, 3, 5}).None(func(n int) bool { return n%2 == 0 }, "index is 0")
	assert.ThatString(t, m.String()).Equal(`Assertion failed: element satisfies the condition:
    got: ([]int) [1 2 3 5]
 expect: no element should satisfy condition
message: index is 0`)
}
