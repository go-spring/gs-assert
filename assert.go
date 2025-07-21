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

// Package assert provides assertion helpers for testing,
// offering both functional and fluent assertion styles.
package assert

import (
	"encoding/json"
	"fmt"
	"reflect"
	"unsafe"

	"github.com/go-spring/assert/internal"
)

// Panic asserts that fn panics and the panic message matches expr.
// It reports an error if fn does not panic or if the recovered message does not satisfy expr.
func Panic(t internal.TestingT, fn func(), expr string, msg ...string) {
	t.Helper()
	internal.Panic(t, false, fn, expr, msg...)
}

type AssertionBase[T any] struct {
	fatalOnFailure bool
}

// Must 不要调用此方法，仅为 require 包提供的.
func (c *AssertionBase[T]) Must() T {
	c.fatalOnFailure = true
	return *(*T)(unsafe.Pointer(&c))
}

// toJsonString converts the given value to a JSON string.
func (c *AssertionBase[T]) toJsonString(v interface{}) string {
	b, err := json.Marshal(v)
	if err != nil {
		return err.Error()
	}
	return string(b)
}

// Assertion wraps a test context and a value for fluent assertions.
type Assertion struct {
	AssertionBase[*Assertion]
	t internal.TestingT
	v interface{}
}

// That creates an Assertion for the given value v and test context t.
func That(t internal.TestingT, v interface{}) *Assertion {
	return &Assertion{
		t: t,
		v: v,
	}
}

// True asserts that got is true. It reports an error if the value is false.
func (a *Assertion) True(msg ...string) {
	a.t.Helper()
	if !a.v.(bool) {
		str := fmt.Sprintf(`expected value to be true, but it is false`)
		internal.Fail(a.t, a.fatalOnFailure, str, msg...)
	}
}

// False asserts that got is false. It reports an error if the value is true.
func (a *Assertion) False(msg ...string) {
	a.t.Helper()
	if a.v.(bool) {
		str := fmt.Sprintf(`expected value to be false, but it is true`)
		internal.Fail(a.t, a.fatalOnFailure, str, msg...)
	}
}

func isNil(v reflect.Value) bool {
	switch v.Kind() {
	case reflect.Chan,
		reflect.Func,
		reflect.Interface,
		reflect.Map,
		reflect.Ptr,
		reflect.Slice,
		reflect.UnsafePointer:
		return v.IsNil()
	default:
		return !v.IsValid()
	}
}

// Nil asserts that got is nil. It reports an error if the value is not nil.
func (a *Assertion) Nil(msg ...string) {
	a.t.Helper()
	// Why can't we use got==nil to judge？Because if
	// a := (*int)(nil)        // %T == *int
	// b := (interface{})(nil) // %T == <nil>
	// then a==b is false, because they are different types.
	if !isNil(reflect.ValueOf(a.v)) {
		str := fmt.Sprintf(`expected value to be nil, but it is not
  actual: %v`, a.v)
		internal.Fail(a.t, a.fatalOnFailure, str, msg...)
	}
}

// NotNil asserts that got is not nil. It reports an error if the value is nil.
func (a *Assertion) NotNil(msg ...string) {
	a.t.Helper()
	if isNil(reflect.ValueOf(a.v)) {
		str := fmt.Sprintf(`expected value to be non-nil, but it is nil`)
		internal.Fail(a.t, a.fatalOnFailure, str, msg...)
	}
}

// Equal asserts that the wrapped value v is deeply equal to expect.
// It reports an error if the values are not deeply equal.
func (a *Assertion) Equal(expect interface{}, msg ...string) {
	a.t.Helper()
	if !reflect.DeepEqual(a.v, expect) {
		str := fmt.Sprintf(`expected values to be equal, but they are different
  actual: %v
expected: %v`, a.toJsonString(a.v), a.toJsonString(expect))
		internal.Fail(a.t, a.fatalOnFailure, str, msg...)
	}
}

// NotEqual asserts that the wrapped value v is not deeply equal to expect.
// It reports an error if the values are deeply equal.
func (a *Assertion) NotEqual(expect interface{}, msg ...string) {
	a.t.Helper()
	if reflect.DeepEqual(a.v, expect) {
		str := fmt.Sprintf(`expected values to be different, but they are equal
  actual: %v`, a.toJsonString(a.v))
		internal.Fail(a.t, a.fatalOnFailure, str, msg...)
	}
}

// Same asserts that the wrapped value v and expect are the same (using Go ==).
// It reports an error if v != expect.
func (a *Assertion) Same(expect interface{}, msg ...string) {
	a.t.Helper()
	if a.v != expect {
		str := fmt.Sprintf(`expected values to be same, but they are different
  actual: %v
expected: %v`, a.toJsonString(a.v), a.toJsonString(expect))
		internal.Fail(a.t, a.fatalOnFailure, str, msg...)
	}
}

// NotSame asserts that the wrapped value v and expect are not the same (using Go !=).
// It reports an error if v == expect.
func (a *Assertion) NotSame(expect interface{}, msg ...string) {
	a.t.Helper()
	if a.v == expect {
		str := fmt.Sprintf(`expected values to be different, but they are same
  actual: %v`, a.toJsonString(a.v))
		internal.Fail(a.t, a.fatalOnFailure, str, msg...)
	}
}

// TypeOf asserts that the type of the wrapped value v is assignable to the type of expect.
// It supports pointer to interface types.
// It reports an error if the types are not assignable.
func (a *Assertion) TypeOf(expect interface{}, msg ...string) {
	a.t.Helper()

	e1 := reflect.TypeOf(a.v)
	e2 := reflect.TypeOf(expect)
	if e2.Kind() == reflect.Ptr && e2.Elem().Kind() == reflect.Interface {
		e2 = e2.Elem()
	}

	if !e1.AssignableTo(e2) {
		str := fmt.Sprintf(`expected type to be assignable to target, but it is not
  actual: %s
expected: %s`, e1.String(), e2.String())
		internal.Fail(a.t, a.fatalOnFailure, str, msg...)
	}
}

// Implements asserts that the type of the wrapped value v implements the interface type of expect.
// The expect parameter must be an interface or pointer to interface.
// It reports an error if v does not implement the interface.
func (a *Assertion) Implements(expect interface{}, msg ...string) {
	a.t.Helper()

	e1 := reflect.TypeOf(a.v)
	e2 := reflect.TypeOf(expect)
	if e2.Kind() == reflect.Ptr {
		if e2.Elem().Kind() == reflect.Interface {
			e2 = e2.Elem()
		} else {
			internal.Fail(a.t, a.fatalOnFailure, "expect should be interface", msg...)
			return
		}
	}

	if !e1.Implements(e2) {
		str := fmt.Sprintf(`expected type to implement target interface, but it does not
  actual: %s
expected: %s`, e1.String(), e2.String())
		internal.Fail(a.t, a.fatalOnFailure, str, msg...)
	}
}

// Has asserts that the wrapped value v has a method named 'Has' that returns true when passed expect.
// It reports an error if the method does not exist or returns false.
func (a *Assertion) Has(expect interface{}, msg ...string) {
	a.t.Helper()

	m := reflect.ValueOf(a.v).MethodByName("Has")
	if !m.IsValid() {
		str := fmt.Sprintf("method 'Has' not found on type %T", a.v)
		internal.Fail(a.t, a.fatalOnFailure, str, msg...)
		return
	}

	if m.Type().NumOut() != 1 || m.Type().Out(0).Kind() != reflect.Bool {
		str := fmt.Sprintf("method 'Has' on type %T should return only a bool", a.v)
		internal.Fail(a.t, a.fatalOnFailure, str, msg...)
		return
	}

	ret := m.Call([]reflect.Value{reflect.ValueOf(expect)})
	if !ret[0].Bool() {
		str := fmt.Sprintf(`method 'Has' on type %T should return true when using param %#v, but it does not`, a.v, expect)
		internal.Fail(a.t, a.fatalOnFailure, str, msg...)
	}
}

// Contains asserts that the wrapped value v has a method named 'Contains' that returns true when passed expect.
// It reports an error if the method does not exist or returns false.
func (a *Assertion) Contains(expect interface{}, msg ...string) {
	a.t.Helper()

	m := reflect.ValueOf(a.v).MethodByName("Contains")
	if !m.IsValid() {
		str := fmt.Sprintf("method 'Contains' not found on type %T", a.v)
		internal.Fail(a.t, a.fatalOnFailure, str, msg...)
		return
	}

	if m.Type().NumOut() != 1 || m.Type().Out(0).Kind() != reflect.Bool {
		str := fmt.Sprintf("method 'Contains' on type %T should return only a bool", a.v)
		internal.Fail(a.t, a.fatalOnFailure, str, msg...)
		return
	}

	ret := m.Call([]reflect.Value{reflect.ValueOf(expect)})
	if !ret[0].Bool() {
		str := fmt.Sprintf(`method 'Contains' on type %T should return true when using param %#v, but it does not`, a.v, expect)
		internal.Fail(a.t, a.fatalOnFailure, str, msg...)
	}
}
