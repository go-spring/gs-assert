/*
 * Copyright 2024 The Go-Spring Authors.
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

	"github.com/go-spring/assert"
	"github.com/go-spring/assert/internal"
)

func TestTrue(t *testing.T) {
	m := new(internal.MockTestingT)
	assert.That(m, true).True()
	assert.ThatString(t, m.String()).Equal("")

	m.Reset()
	assert.That(m, false).True()
	assert.ThatString(t, m.String()).Equal("error# Assertion failed: expected value to be true, but it is false")

	m.Reset()
	assert.That(m, false).Must().True("index is 0")
	assert.ThatString(t, m.String()).Equal(`fatal# Assertion failed: expected value to be true, but it is false
message: "index is 0"`)
}

func TestFalse(t *testing.T) {
	m := new(internal.MockTestingT)
	assert.That(m, false).False()
	assert.ThatString(t, m.String()).Equal("")

	m.Reset()
	assert.That(m, true).False()
	assert.ThatString(t, m.String()).Equal("error# Assertion failed: expected value to be false, but it is true")

	m.Reset()
	assert.That(m, true).Must().False("index is 0")
	assert.ThatString(t, m.String()).Equal(`fatal# Assertion failed: expected value to be false, but it is true
message: "index is 0"`)
}

func TestNil(t *testing.T) {
	m := new(internal.MockTestingT)
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
    got: 3`)

	m.Reset()
	assert.That(m, 3).Must().Nil("index is 0")
	assert.ThatString(t, m.String()).Equal(`fatal# Assertion failed: expected value to be nil, but it is not
    got: 3
message: "index is 0"`)
}

func TestNotNil(t *testing.T) {
	m := new(internal.MockTestingT)
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
	assert.That(m, nil).Must().NotNil("index is 0")
	assert.ThatString(t, m.String()).Equal(`fatal# Assertion failed: expected value to be non-nil, but it is nil
message: "index is 0"`)
}

func TestPanic(t *testing.T) {
	m := new(internal.MockTestingT)
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

func TestThat_Equal(t *testing.T) {
	m := new(internal.MockTestingT)
	assert.That(m, 0).Equal(0)
	assert.ThatString(t, m.String()).Equal("")

	m.Reset()
	assert.That(m, []string{"a"}).Equal([]string{"a"})
	assert.ThatString(t, m.String()).Equal("")

	m.Reset()
	assert.That(m, struct {
		text string
	}{text: "a"}).Equal(struct {
		text string
	}{text: "a"})
	assert.ThatString(t, m.String()).Equal("")

	m.Reset()
	assert.That(m, struct {
		Text string
	}{Text: "a"}).Equal(struct {
		Text string `json:"text"`
	}{Text: "a"})
	assert.ThatString(t, m.String()).Equal(`error# Assertion failed: expected values to be equal, but they are different
    got: {"Text":"a"}
 expect: {"text":"a"}`)

	m.Reset()
	assert.That(m, struct {
		text string
	}{text: "a"}).Equal(struct {
		msg string
	}{msg: "a"})
	assert.ThatString(t, m.String()).Equal(`error# Assertion failed: expected values to be equal, but they are different
    got: {}
 expect: {}`)

	m.Reset()
	assert.That(m, 0).Must().Equal("0")
	assert.ThatString(t, m.String()).Equal(`fatal# Assertion failed: expected values to be equal, but they are different
    got: 0
 expect: "0"`)

	m.Reset()
	assert.That(m, 0).Must().Equal("0", "index is 0")
	assert.ThatString(t, m.String()).Equal(`fatal# Assertion failed: expected values to be equal, but they are different
    got: 0
 expect: "0"
message: "index is 0"`)
}

func TestThat_NotEqual(t *testing.T) {
	m := new(internal.MockTestingT)
	assert.That(m, "0").NotEqual(0)
	assert.ThatString(t, m.String()).Equal("")

	m.Reset()
	assert.That(m, "0").NotEqual("0")
	assert.ThatString(t, m.String()).Equal(`error# Assertion failed: expected values to be different, but they are equal
    got: "0"`)

	m.Reset()
	assert.That(m, "0").Must().NotEqual("0", "index is 0")
	assert.ThatString(t, m.String()).Equal(`fatal# Assertion failed: expected values to be different, but they are equal
    got: "0"
message: "index is 0"`)
}

func TestThat_Same(t *testing.T) {
	m := new(internal.MockTestingT)
	assert.That(m, "0").Same("0")
	assert.ThatString(t, m.String()).Equal("")

	m.Reset()
	assert.That(m, 0).Same("0")
	assert.ThatString(t, m.String()).Equal(`error# Assertion failed: expected values to be same, but they are different
    got: 0
 expect: "0"`)

	m.Reset()
	assert.That(m, 0).Must().Same("0", "index is 0")
	assert.ThatString(t, m.String()).Equal(`fatal# Assertion failed: expected values to be same, but they are different
    got: 0
 expect: "0"
message: "index is 0"`)
}

func TestThat_NotSame(t *testing.T) {
	m := new(internal.MockTestingT)
	assert.That(m, "0").NotSame(0)
	assert.ThatString(t, m.String()).Equal("")

	m.Reset()
	assert.That(m, "0").NotSame("0")
	assert.ThatString(t, m.String()).Equal(`error# Assertion failed: expected values to be different, but they are same
    got: "0"`)

	m.Reset()
	assert.That(m, "0").Must().NotSame("0", "index is 0")
	assert.ThatString(t, m.String()).Equal(`fatal# Assertion failed: expected values to be different, but they are same
    got: "0"
message: "index is 0"`)
}

func TestThat_TypeOf(t *testing.T) {
	m := new(internal.MockTestingT)
	assert.That(m, new(int)).TypeOf((*int)(nil))
	assert.ThatString(t, m.String()).Equal("")

	m.Reset()
	assert.That(m, "string").TypeOf((*fmt.Stringer)(nil))
	assert.ThatString(t, m.String()).Equal(`error# Assertion failed: expected type to be assignable to target, but it is not
    got: string
 expect: fmt.Stringer`)

	m.Reset()
	assert.That(m, "string").Must().TypeOf((*fmt.Stringer)(nil))
	assert.ThatString(t, m.String()).Equal(`fatal# Assertion failed: expected type to be assignable to target, but it is not
    got: string
 expect: fmt.Stringer`)
}

func TestThat_Implements(t *testing.T) {
	m := new(internal.MockTestingT)

	assert.That(m, errors.New("error")).Implements((*error)(nil))
	assert.ThatString(t, m.String()).Equal("")

	m.Reset()
	assert.That(m, new(int)).Implements((*int)(nil))
	assert.ThatString(t, m.String()).Equal("error# Assertion failed: expect should be interface")

	m.Reset()
	assert.That(m, new(int)).Must().Implements((*io.Reader)(nil))
	assert.ThatString(t, m.String()).Equal(`fatal# Assertion failed: expected type to implement target interface, but it does not
    got: *int
 expect: io.Reader`)
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
	assert.That(m, &Tree{Keys: []string{"1"}}).Has("1")
	assert.ThatString(t, m.String()).Equal("")

	m.Reset()
	assert.That(m, 1).Has("1")
	assert.ThatString(t, m.String()).Equal("error# Assertion failed: method 'Has' not found on type int")

	m.Reset()
	assert.That(m, &Node{}).Has("2")
	assert.ThatString(t, m.String()).Equal("error# Assertion failed: method 'Has' on type *assert_test.Node should return only a bool")

	m.Reset()
	assert.That(m, &Tree{}).Must().Has("2")
	assert.ThatString(t, m.String()).Equal("fatal# Assertion failed: method 'Has' on type *assert_test.Tree should return true when using param \"2\", but it does not")
}

func TestThat_Contains(t *testing.T) {
	m := new(internal.MockTestingT)

	assert.That(m, &Tree{Keys: []string{"1"}}).Contains("1")
	assert.ThatString(t, m.String()).Equal("")

	m.Reset()
	assert.That(m, 1).Contains("1")
	assert.ThatString(t, m.String()).Equal("error# Assertion failed: method 'Contains' not found on type int")

	m.Reset()
	assert.That(m, &Node{}).Contains("2")
	assert.ThatString(t, m.String()).Equal("error# Assertion failed: method 'Contains' on type *assert_test.Node should return only a bool")

	m.Reset()
	assert.That(m, &Tree{}).Must().Contains("2")
	assert.ThatString(t, m.String()).Equal("fatal# Assertion failed: method 'Contains' on type *assert_test.Tree should return true when using param \"2\", but it does not")
}
