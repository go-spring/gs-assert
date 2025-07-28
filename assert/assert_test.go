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
	"bytes"
	"errors"
	"fmt"
	"io"
	"slices"
	"testing"

	"github.com/go-spring/gs-assert/assert"
	"github.com/go-spring/gs-assert/internal"
)

func TestPanic(t *testing.T) {
	m := new(internal.MockTestingT)

	m.Reset()
	assert.Panic(m, func() { panic("this is an error") }, "an error")
	assert.ThatString(t, m.String()).Equal("")

	m.Reset()
	assert.Panic(m, func() {}, "an error")
	assert.ThatString(t, m.String()).Equal("error# Assertion failed: did not panic")

	m.Reset()
	assert.Panic(m, func() { panic("this is an error") }, `an error \`)
	assert.ThatString(t, m.String()).Equal("error# Assertion failed: invalid pattern")

	m.Reset()
	assert.Panic(m, func() { panic("there's no error") }, "an error")
	assert.ThatString(t, m.String()).Equal(`error# Assertion failed: got "there's no error" which does not match "an error"`)

	m.Reset()
	assert.Panic(m, func() { panic("there's no error") }, "an error", "index is 0")
	assert.ThatString(t, m.String()).Equal(`error# Assertion failed: got "there's no error" which does not match "an error"
 message: "index is 0"`)

	m.Reset()
	assert.Panic(m, func() { panic(errors.New("there's no error")) }, "an error")
	assert.ThatString(t, m.String()).Equal(`error# Assertion failed: got "there's no error" which does not match "an error"`)

	m.Reset()
	assert.Panic(m, func() { panic(bytes.NewBufferString("there's no error")) }, "an error")
	assert.ThatString(t, m.String()).Equal(`error# Assertion failed: got "there's no error" which does not match "an error"`)

	m.Reset()
	assert.Panic(m, func() { panic([]string{"there's no error"}) }, "an error")
	assert.ThatString(t, m.String()).Equal(`error# Assertion failed: got "[there's no error]" which does not match "an error"`)
}

func TestThat_True(t *testing.T) {
	m := new(internal.MockTestingT)

	m.Reset()
	assert.That(m, true).True()
	assert.ThatString(t, m.String()).Equal("")

	m.Reset()
	assert.That(m, false).True()
	assert.ThatString(t, m.String()).Equal("error# Assertion failed: expected value to be true, but it is false")

	m.Reset()
	assert.That(m, false).Require().True("index is 0")
	assert.ThatString(t, m.String()).Equal(`fatal# Assertion failed: expected value to be true, but it is false
 message: "index is 0"`)
}

func TestThat_False(t *testing.T) {
	m := new(internal.MockTestingT)

	m.Reset()
	assert.That(m, false).False()
	assert.ThatString(t, m.String()).Equal("")

	m.Reset()
	assert.That(m, true).False()
	assert.ThatString(t, m.String()).Equal("error# Assertion failed: expected value to be false, but it is true")

	m.Reset()
	assert.That(m, true).Require().False("index is 0")
	assert.ThatString(t, m.String()).Equal(`fatal# Assertion failed: expected value to be false, but it is true
 message: "index is 0"`)
}

func TestThat_Nil(t *testing.T) {
	m := new(internal.MockTestingT)

	m.Reset()
	assert.That(m, nil).Nil()
	assert.ThatString(t, m.String()).Equal("")

	m.Reset()
	var a []string
	assert.That(m, a).Nil()
	assert.ThatString(t, m.String()).Equal("")

	m.Reset()
	var s map[string]string
	assert.That(m, s).Nil()
	assert.ThatString(t, m.String()).Equal("")

	m.Reset()
	assert.That(m, 3).Nil()
	assert.ThatString(t, m.String()).Equal(`error# Assertion failed: expected value to be nil, but it is not
  actual: (int) 3`)

	m.Reset()
	assert.That(m, 3).Require().Nil("index is 0")
	assert.ThatString(t, m.String()).Equal(`fatal# Assertion failed: expected value to be nil, but it is not
  actual: (int) 3
 message: "index is 0"`)
}

func TestThat_NotNil(t *testing.T) {
	m := new(internal.MockTestingT)

	m.Reset()
	assert.That(m, 3).NotNil()
	assert.ThatString(t, m.String()).Equal("")

	m.Reset()
	assert.That(m, make([]string, 0)).NotNil()
	assert.ThatString(t, m.String()).Equal("")

	m.Reset()
	assert.That(m, make(map[string]string)).NotNil()
	assert.ThatString(t, m.String()).Equal("")

	m.Reset()
	assert.That(m, nil).NotNil()
	assert.ThatString(t, m.String()).Equal(`error# Assertion failed: expected value to be non-nil, but it is nil`)

	m.Reset()
	assert.That(m, nil).Require().NotNil("index is 0")
	assert.ThatString(t, m.String()).Equal(`fatal# Assertion failed: expected value to be non-nil, but it is nil
 message: "index is 0"`)
}

func TestThat_Equal(t *testing.T) {
	m := new(internal.MockTestingT)

	m.Reset()
	assert.That(m, 0).Equal(0)
	assert.ThatString(t, m.String()).Equal("")

	m.Reset()
	assert.That(m, []string{"a"}).Equal([]string{"a"})
	assert.ThatString(t, m.String()).Equal("")

	type SimpleText struct {
		text string
	}

	type AnotherSimpleText struct {
		text string
	}

	type SimpleMessage struct {
		message string
	}

	m.Reset()
	assert.That(m, SimpleText{text: "a"}).Equal(SimpleText{text: "a"})
	assert.ThatString(t, m.String()).Equal("")

	m.Reset()
	assert.That(m, SimpleText{text: "a"}).Equal(AnotherSimpleText{text: "a"})
	assert.ThatString(t, m.String()).Equal(`error# Assertion failed: expected values to be equal, but they are different
  actual: (assert_test.SimpleText) assert_test.SimpleText{text:"a"}
expected: (assert_test.AnotherSimpleText) assert_test.AnotherSimpleText{text:"a"}`)

	m.Reset()
	assert.That(m, SimpleText{text: "a"}).Equal(SimpleMessage{message: "a"})
	assert.ThatString(t, m.String()).Equal(`error# Assertion failed: expected values to be equal, but they are different
  actual: (assert_test.SimpleText) assert_test.SimpleText{text:"a"}
expected: (assert_test.SimpleMessage) assert_test.SimpleMessage{message:"a"}`)

	m.Reset()
	assert.That(m, 0).Require().Equal("0")
	assert.ThatString(t, m.String()).Equal(`fatal# Assertion failed: expected values to be equal, but they are different
  actual: (int) 0
expected: (string) "0"`)

	m.Reset()
	assert.That(m, 0).Require().Equal("0", "index is 0")
	assert.ThatString(t, m.String()).Equal(`fatal# Assertion failed: expected values to be equal, but they are different
  actual: (int) 0
expected: (string) "0"
 message: "index is 0"`)
}

func TestThat_NotEqual(t *testing.T) {
	m := new(internal.MockTestingT)

	m.Reset()
	assert.That(m, "0").NotEqual(0)
	assert.ThatString(t, m.String()).Equal("")

	m.Reset()
	assert.That(m, "0").NotEqual("0")
	assert.ThatString(t, m.String()).Equal(`error# Assertion failed: expected values to be different, but they are equal
  actual: (string) "0"`)

	m.Reset()
	assert.That(m, "0").Require().NotEqual("0", "index is 0")
	assert.ThatString(t, m.String()).Equal(`fatal# Assertion failed: expected values to be different, but they are equal
  actual: (string) "0"
 message: "index is 0"`)
}

func TestThat_Same(t *testing.T) {
	m := new(internal.MockTestingT)

	m.Reset()
	assert.That(m, "0").Same("0")
	assert.ThatString(t, m.String()).Equal("")

	m.Reset()
	assert.That(m, 0).Same("0")
	assert.ThatString(t, m.String()).Equal(`error# Assertion failed: expected values to be same, but they are different
  actual: (int) 0
expected: (string) "0"`)

	m.Reset()
	assert.That(m, 0).Require().Same("0", "index is 0")
	assert.ThatString(t, m.String()).Equal(`fatal# Assertion failed: expected values to be same, but they are different
  actual: (int) 0
expected: (string) "0"
 message: "index is 0"`)
}

func TestThat_NotSame(t *testing.T) {
	m := new(internal.MockTestingT)

	m.Reset()
	assert.That(m, "0").NotSame(0)
	assert.ThatString(t, m.String()).Equal("")

	m.Reset()
	assert.That(m, "0").NotSame("0")
	assert.ThatString(t, m.String()).Equal(`error# Assertion failed: expected values to be different, but they are same
  actual: (string) "0"`)

	m.Reset()
	assert.That(m, "0").Require().NotSame("0", "index is 0")
	assert.ThatString(t, m.String()).Equal(`fatal# Assertion failed: expected values to be different, but they are same
  actual: (string) "0"
 message: "index is 0"`)
}

func TestThat_TypeOf(t *testing.T) {
	m := new(internal.MockTestingT)

	m.Reset()
	assert.That(m, new(int)).TypeOf((*int)(nil))
	assert.ThatString(t, m.String()).Equal("")

	m.Reset()
	assert.That(m, "string").TypeOf((*fmt.Stringer)(nil))
	assert.ThatString(t, m.String()).Equal(`error# Assertion failed: expected type to be assignable to target, but it does not
  actual: string
expected: fmt.Stringer`)

	m.Reset()
	assert.That(m, "string").Require().TypeOf((*fmt.Stringer)(nil))
	assert.ThatString(t, m.String()).Equal(`fatal# Assertion failed: expected type to be assignable to target, but it does not
  actual: string
expected: fmt.Stringer`)
}

func TestThat_Implements(t *testing.T) {
	m := new(internal.MockTestingT)

	m.Reset()
	assert.That(m, errors.New("error")).Implements((*error)(nil))
	assert.ThatString(t, m.String()).Equal("")

	m.Reset()
	assert.That(m, new(int)).Implements((*int)(nil))
	assert.ThatString(t, m.String()).Equal("error# Assertion failed: expected target to implement should be interface")

	m.Reset()
	assert.That(m, new(int)).Require().Implements((*io.Reader)(nil))
	assert.ThatString(t, m.String()).Equal(`fatal# Assertion failed: expected type to implement target interface, but it does not
  actual: *int
expected: io.Reader`)
}

type Node struct{}

func (t *Node) Has(key string) (bool, error) {
	return false, nil
}

func (t *Node) Contains(key string) (bool, error) {
	return false, nil
}

type Tree struct {
	Keys []string
}

func (t *Tree) Has(key string) bool {
	return slices.Contains(t.Keys, key)
}

func (t *Tree) Contains(key string) bool {
	return slices.Contains(t.Keys, key)
}

func TestThat_Has(t *testing.T) {
	m := new(internal.MockTestingT)

	m.Reset()
	assert.That(m, &Tree{Keys: []string{"1"}}).Has("1")
	assert.ThatString(t, m.String()).Equal("")

	m.Reset()
	assert.That(m, 1).Has("1")
	assert.ThatString(t, m.String()).Equal("error# Assertion failed: method 'Has' not found on type int")

	m.Reset()
	assert.That(m, &Node{}).Has("2")
	assert.ThatString(t, m.String()).Equal("error# Assertion failed: method 'Has' on type *assert_test.Node should return only a bool, but it does not")

	m.Reset()
	assert.That(m, &Tree{}).Require().Has("2")
	assert.ThatString(t, m.String()).Equal("fatal# Assertion failed: method 'Has' on type *assert_test.Tree should return true when using param \"2\", but it does not")
}

func TestThat_Contains(t *testing.T) {
	m := new(internal.MockTestingT)

	m.Reset()
	assert.That(m, &Tree{Keys: []string{"1"}}).Contains("1")
	assert.ThatString(t, m.String()).Equal("")

	m.Reset()
	assert.That(m, 1).Contains("1")
	assert.ThatString(t, m.String()).Equal("error# Assertion failed: method 'Contains' not found on type int")

	m.Reset()
	assert.That(m, &Node{}).Contains("2")
	assert.ThatString(t, m.String()).Equal("error# Assertion failed: method 'Contains' on type *assert_test.Node should return only a bool, but it does not")

	m.Reset()
	assert.That(m, &Tree{}).Require().Contains("2")
	assert.ThatString(t, m.String()).Equal("fatal# Assertion failed: method 'Contains' on type *assert_test.Tree should return true when using param \"2\", but it does not")
}
