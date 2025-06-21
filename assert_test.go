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
)

type MockTestingT struct {
	buf bytes.Buffer
}

func (m *MockTestingT) Reset() {
	m.buf.Reset()
}

func (m *MockTestingT) Helper() {}

func (m *MockTestingT) Error(args ...any) {
	for _, arg := range args {
		m.buf.WriteString(fmt.Sprint(arg))
	}
}

func (m *MockTestingT) String() string {
	return m.buf.String()
}

func TestTrue(t *testing.T) {
	m := new(MockTestingT)
	assert.That(m, true).IsTrue()
	assert.ThatString(t, m.String()).Equal("")

	m.Reset()
	assert.That(m, false).IsTrue()
	assert.ThatString(t, m.String()).Equal("got false but expect true")

	m.Reset()
	assert.That(m, false).IsTrue("param (index=0)")
	assert.ThatString(t, m.String()).Equal("got false but expect true\nmessage: param (index=0)")
}

func TestFalse(t *testing.T) {
	m := new(MockTestingT)
	assert.That(m, false).IsFalse()
	assert.ThatString(t, m.String()).Equal("")

	m.Reset()
	assert.That(m, true).IsFalse()
	assert.ThatString(t, m.String()).Equal("got true but expect false")

	m.Reset()
	assert.That(m, true).IsFalse("param (index=0)")
	assert.ThatString(t, m.String()).Equal(`got true but expect false
message: param (index=0)`)
}

func TestNil(t *testing.T) {
	m := new(MockTestingT)
	assert.That(m, nil).IsNil()
	assert.ThatString(t, m.String()).Equal("")

	m.Reset()
	var a []string
	assert.That(m, a).IsNil()
	assert.ThatString(t, m.String()).Equal("")

	m.Reset()
	var s map[string]string
	assert.That(m, s).IsNil()
	assert.ThatString(t, m.String()).Equal("")

	m.Reset()
	assert.That(m, 3).IsNil()
	assert.ThatString(t, m.String()).Equal("got (int) 3 but expect nil")

	m.Reset()
	assert.That(m, 3).IsNil("param (index=0)")
	assert.ThatString(t, m.String()).Equal(`got (int) 3 but expect nil
message: param (index=0)`)
}

func TestNotNil(t *testing.T) {
	m := new(MockTestingT)
	assert.That(m, 3).IsNotNil()
	assert.ThatString(t, m.String()).Equal("")

	m.Reset()
	assert.That(m, make([]string, 0)).IsNotNil()
	assert.ThatString(t, m.String()).Equal("")

	m.Reset()
	assert.That(m, make(map[string]string)).IsNotNil()
	assert.ThatString(t, m.String()).Equal("")

	m.Reset()
	assert.That(m, nil).IsNotNil()
	assert.ThatString(t, m.String()).Equal("got nil but expect not nil")

	m.Reset()
	assert.That(m, nil).IsNotNil("param (index=0)")
	assert.ThatString(t, m.String()).Equal(`got nil but expect not nil
message: param (index=0)`)
}

func TestPanic(t *testing.T) {
	m := new(MockTestingT)
	assert.Panic(m, func() { panic("this is an error") }, "an error")
	assert.ThatString(t, m.String()).Equal("")

	m.Reset()
	assert.Panic(m, func() {}, "an error")
	assert.ThatString(t, m.String()).Equal("did not panic")

	m.Reset()
	assert.Panic(m, func() { panic("this is an error") }, `an error \`)
	assert.ThatString(t, m.String()).Equal("invalid pattern")

	m.Reset()
	assert.Panic(m, func() { panic("there's no error") }, "an error")
	assert.ThatString(t, m.String()).Equal(`got "there's no error" which does not match "an error"`)

	m.Reset()
	assert.Panic(m, func() { panic("there's no error") }, "an error", "param (index=0)")
	assert.ThatString(t, m.String()).Equal(`got "there's no error" which does not match "an error"
message: param (index=0)`)

	m.Reset()
	assert.Panic(m, func() { panic(errors.New("there's no error")) }, "an error")
	assert.ThatString(t, m.String()).Equal(`got "there's no error" which does not match "an error"`)

	m.Reset()
	assert.Panic(m, func() { panic(bytes.NewBufferString("there's no error")) }, "an error")
	assert.ThatString(t, m.String()).Equal(`got "there's no error" which does not match "an error"`)

	m.Reset()
	assert.Panic(m, func() { panic([]string{"there's no error"}) }, "an error")
	assert.ThatString(t, m.String()).Equal(`got "[there's no error]" which does not match "an error"`)
}

func TestThat_Equal(t *testing.T) {
	m := new(MockTestingT)
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
	assert.ThatString(t, m.String()).Equal(`got (struct { Text string }) {a} but expect (struct { Text string "json:\"text\"" }) {a}`)

	m.Reset()
	assert.That(m, struct {
		text string
	}{text: "a"}).Equal(struct {
		msg string
	}{msg: "a"})
	assert.ThatString(t, m.String()).Equal("got (struct { text string }) {a} but expect (struct { msg string }) {a}")

	m.Reset()
	assert.That(m, 0).Equal("0")
	assert.ThatString(t, m.String()).Equal("got (int) 0 but expect (string) 0")

	m.Reset()
	assert.That(m, 0).Equal("0", "param (index=0)")
	assert.ThatString(t, m.String()).Equal(`got (int) 0 but expect (string) 0
message: param (index=0)`)
}

func TestThat_NotEqual(t *testing.T) {
	m := new(MockTestingT)
	assert.That(m, "0").NotEqual(0)

	m.Reset()
	assert.That(m, []string{"a"}).NotEqual([]string{"a"})
	assert.ThatString(t, m.String()).Equal("got ([]string) [a] but expect not ([]string) [a]")

	m.Reset()
	assert.That(m, "0").NotEqual("0")
	assert.ThatString(t, m.String()).Equal("got (string) 0 but expect not (string) 0")

	m.Reset()
	assert.That(m, "0").NotEqual("0", "param (index=0)")
	assert.ThatString(t, m.String()).Equal(`got (string) 0 but expect not (string) 0
message: param (index=0)`)
}

func TestThat_Same(t *testing.T) {
	m := new(MockTestingT)
	assert.That(m, "0").IsSame("0")

	m.Reset()
	assert.That(m, 0).IsSame("0")
	assert.ThatString(t, m.String()).Equal("got (int) 0 but expect (string) 0")

	m.Reset()
	assert.That(m, 0).IsSame("0", "param (index=0)")
	assert.ThatString(t, m.String()).Equal(`got (int) 0 but expect (string) 0
message: param (index=0)`)
}

func TestThat_NotSame(t *testing.T) {
	m := new(MockTestingT)
	assert.That(m, "0").IsNotSame(0)

	m.Reset()
	assert.That(m, "0").IsNotSame("0")
	assert.ThatString(t, m.String()).Equal("expect not (string) 0")

	m.Reset()
	assert.That(m, "0").IsNotSame("0", "param (index=0)")
	assert.ThatString(t, m.String()).Equal(`expect not (string) 0
message: param (index=0)`)
}

func TestThat_TypeOf(t *testing.T) {
	m := new(MockTestingT)
	assert.That(m, new(int)).IsTypeOf((*int)(nil))

	m.Reset()
	assert.That(m, "string").IsTypeOf((*fmt.Stringer)(nil))
	assert.ThatString(t, m.String()).Equal("got type (string) but expect type (fmt.Stringer)")
}

func TestThat_Implements(t *testing.T) {
	m := new(MockTestingT)
	assert.That(m, errors.New("error")).Implements((*error)(nil))
	assert.ThatString(t, m.String()).Equal("")

	m.Reset()
	assert.That(m, new(int)).Implements((*int)(nil))
	assert.ThatString(t, m.String()).Equal("expect should be interface")

	m.Reset()
	assert.That(m, new(int)).Implements((*io.Reader)(nil))
	assert.ThatString(t, m.String()).Equal("got type (*int) but expect type (io.Reader)")
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
	m := new(MockTestingT)
	assert.That(m, 1).Has("1")
	assert.ThatString(t, m.String()).Equal("method 'Has' not found on type int")

	m.Reset()
	assert.That(m, &Node{}).Has("2")
	assert.ThatString(t, m.String()).Equal("method 'Has' must return only a bool")

	m.Reset()
	assert.That(m, &Tree{}).Has("2")
	assert.ThatString(t, m.String()).Equal("got (*assert_test.Tree) &{[]} not has (string) 2")

	m.Reset()
	assert.That(m, &Tree{Keys: []string{"1"}}).Has("1")
}

func TestThat_Contains(t *testing.T) {
	m := new(MockTestingT)
	assert.That(m, 1).Contains("1")
	assert.ThatString(t, m.String()).Equal("method 'Contains' not found on type int")

	m.Reset()
	assert.That(m, &Node{}).Contains("2")
	assert.ThatString(t, m.String()).Equal("method 'Contains' must return only a bool")

	m.Reset()
	assert.That(m, &Tree{}).Contains("2")
	assert.ThatString(t, m.String()).Equal("got (*assert_test.Tree) &{[]} not contains (string) 2")

	m.Reset()
	assert.That(m, &Tree{Keys: []string{"1"}}).Contains("1")
}

func TestThat_InSlice(t *testing.T) {
	m := new(MockTestingT)
	assert.That(m, 1).InSlice("1")
	assert.ThatString(t, m.String()).Equal("unsupported expect value (string) 1")

	m.Reset()
	assert.That(m, 1).InSlice([]string{"1"})
	assert.ThatString(t, m.String()).Equal("got (int) 1 is not in ([]string) [1]")

	m.Reset()
	assert.That(m, int64(1)).InSlice([]int64{3, 2})
	assert.ThatString(t, m.String()).Equal("got (int64) 1 is not in ([]int64) [3 2]")

	m.Reset()
	assert.That(m, int64(1)).InSlice([]int64{3, 2, 1})
	assert.That(m, "1").InSlice([]string{"3", "2", "1"})
}

func TestThat_NotInSlice(t *testing.T) {
	m := new(MockTestingT)
	assert.That(m, 1).NotInSlice("1")
	assert.ThatString(t, m.String()).Equal("unsupported expect value (string) 1")

	m.Reset()
	assert.That(m, 1).NotInSlice([]string{"1"})
	assert.ThatString(t, m.String()).Equal("got type (int) doesn't match expect type ([]string)")

	m.Reset()
	assert.That(m, "1").NotInSlice([]string{"3", "2", "1"})
	assert.ThatString(t, m.String()).Equal("got (string) 1 is in ([]string) [3 2 1]")

	m.Reset()
	assert.That(m, int64(1)).NotInSlice([]int64{3, 2})
}

func TestThat_InMapKeys(t *testing.T) {
	m := new(MockTestingT)
	assert.That(m, 1).InMapKeys("1")
	assert.ThatString(t, m.String()).Equal("unsupported expect value (string) 1")

	m.Reset()
	assert.That(m, 1).InMapKeys(map[string]string{"1": "1"})
	assert.ThatString(t, m.String()).Equal("got (int) 1 is not in keys of (map[string]string) map[1:1]")

	m.Reset()
	assert.That(m, int64(1)).InMapKeys(map[int64]int64{3: 1, 2: 2, 1: 3})
	assert.That(m, "1").InMapKeys(map[string]string{"3": "1", "2": "2", "1": "3"})
}

func TestThat_InMapValues(t *testing.T) {
	m := new(MockTestingT)
	assert.That(m, 1).InMapValues("1")
	assert.ThatString(t, m.String()).Equal("unsupported expect value (string) 1")

	m.Reset()
	assert.That(m, 1).InMapValues(map[string]string{"1": "1"})
	assert.ThatString(t, m.String()).Equal("got (int) 1 is not in values of (map[string]string) map[1:1]")

	m.Reset()
	assert.That(m, int64(1)).InMapValues(map[int64]int64{3: 1, 2: 2, 1: 3})
	assert.That(m, "1").InMapValues(map[string]string{"3": "1", "2": "2", "1": "3"})
}
