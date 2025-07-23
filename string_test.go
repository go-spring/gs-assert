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

func TestString_Length(t *testing.T) {
	m := new(internal.MockTestingT)

	m.Reset()
	assert.ThatString(m, "0").Length(1)
	assert.ThatString(t, m.String()).Equal("")

	m.Reset()
	assert.ThatString(m, "0").Length(0)
	assert.ThatString(t, m.String()).Equal(`error# Assertion failed: expected string to have length 0, but it has length 1
  actual: "0"`)

	m.Reset()
	assert.ThatString(m, "0").Must().Length(0, "index is 0")
	assert.ThatString(t, m.String()).Equal(`fatal# Assertion failed: expected string to have length 0, but it has length 1
  actual: "0"
 message: "index is 0"`)
}

func TestString_Blank(t *testing.T) {
	m := new(internal.MockTestingT)

	m.Reset()
	assert.ThatString(m, "   ").Blank()
	assert.ThatString(t, m.String()).Equal("")

	m.Reset()
	assert.ThatString(m, "hello").Blank()
	assert.ThatString(t, m.String()).Equal(`error# Assertion failed: expected string to contain only whitespace, but it does not
  actual: "hello"`)

	m.Reset()
	assert.ThatString(m, "hello").Must().Blank("index is 0")
	assert.ThatString(t, m.String()).Equal(`fatal# Assertion failed: expected string to contain only whitespace, but it does not
  actual: "hello"
 message: "index is 0"`)
}

func TestString_NotBlank(t *testing.T) {
	m := new(internal.MockTestingT)

	m.Reset()
	assert.ThatString(m, "hello").NotBlank()
	assert.ThatString(t, m.String()).Equal("")

	m.Reset()
	assert.ThatString(m, "   ").NotBlank()
	assert.ThatString(t, m.String()).Equal(`error# Assertion failed: expected string to be non-blank, but it is blank
  actual: "   "`)

	m.Reset()
	assert.ThatString(m, " \n  ").Must().NotBlank("index is 0")
	assert.ThatString(t, m.String()).Equal(`fatal# Assertion failed: expected string to be non-blank, but it is blank
  actual: " \n  "
 message: "index is 0"`)
}

func TestString_Equal(t *testing.T) {
	m := new(internal.MockTestingT)

	m.Reset()
	assert.ThatString(m, "0").Equal("0")
	assert.ThatString(t, m.String()).Equal("")

	m.Reset()
	assert.ThatString(m, "0").Equal("1")
	assert.ThatString(t, m.String()).Equal(`error# Assertion failed: expected strings to be equal, but they are not
  actual: "0"
expected: "1"`)

	m.Reset()
	assert.ThatString(m, "0").Must().Equal("1", "index is 0")
	assert.ThatString(t, m.String()).Equal(`fatal# Assertion failed: expected strings to be equal, but they are not
  actual: "0"
expected: "1"
 message: "index is 0"`)
}

func TestString_NotEqual(t *testing.T) {
	m := new(internal.MockTestingT)

	m.Reset()
	assert.ThatString(m, "0").NotEqual("1")
	assert.ThatString(t, m.String()).Equal("")

	m.Reset()
	assert.ThatString(m, "0").NotEqual("0")
	assert.ThatString(t, m.String()).Equal(`error# Assertion failed: expected strings to be different, but they are equal
  actual: "0"
expected: "0"`)

	m.Reset()
	assert.ThatString(m, "0").Must().NotEqual("0", "index is 0")
	assert.ThatString(t, m.String()).Equal(`fatal# Assertion failed: expected strings to be different, but they are equal
  actual: "0"
expected: "0"
 message: "index is 0"`)
}

func TestString_EqualFold(t *testing.T) {
	m := new(internal.MockTestingT)

	m.Reset()
	assert.ThatString(m, "hello, world!").EqualFold("Hello, World!")
	assert.ThatString(t, m.String()).Equal("")

	m.Reset()
	assert.ThatString(m, "hello, world!").EqualFold("Hello, Jimmy!")
	assert.ThatString(t, m.String()).Equal(`error# Assertion failed: expected strings to be equal (case-insensitive), but they are not
  actual: "hello, world!"
expected: "Hello, Jimmy!"`)

	m.Reset()
	assert.ThatString(m, "hello, world!").Must().EqualFold("Hello, Jimmy!", "index is 0")
	assert.ThatString(t, m.String()).Equal(`fatal# Assertion failed: expected strings to be equal (case-insensitive), but they are not
  actual: "hello, world!"
expected: "Hello, Jimmy!"
 message: "index is 0"`)
}

func TestString_JSONEqual(t *testing.T) {
	m := new(internal.MockTestingT)

	m.Reset()
	assert.ThatString(m, `{"a":0,"b":1}`).JSONEqual(`{"b":1,"a":0}`)
	assert.ThatString(t, m.String()).Equal("")

	m.Reset()
	assert.ThatString(m, `this is an error`).JSONEqual(`[{"b":1},{"a":0}]`)
	assert.ThatString(t, m.String()).Equal(`error# Assertion failed: expected strings to be JSON-equal, but failed to unmarshal actual value
  actual: "this is an error"
   error: "invalid character 'h' in literal true (expecting 'r')"`)

	m.Reset()
	assert.ThatString(m, `{"a":0,"b":1}`).Must().JSONEqual(`this is an error`)
	assert.ThatString(t, m.String()).Equal(`fatal# Assertion failed: expected strings to be JSON-equal, but failed to unmarshal expected value
expected: "this is an error"
   error: "invalid character 'h' in literal true (expecting 'r')"`)

	m.Reset()
	assert.ThatString(m, `{"a":0,"b":1}`).JSONEqual(`[{"b":1},{"a":0}]`)
	assert.ThatString(t, m.String()).Equal(`error# Assertion failed: expected strings to be JSON-equal, but they are not
  actual: "{\"a\":0,\"b\":1}"
expected: "[{\"b\":1},{\"a\":0}]"`)

	m.Reset()
	assert.ThatString(m, `{"a":0}`).Must().JSONEqual(`{"a":1}`, "index is 0")
	assert.ThatString(t, m.String()).Equal(`fatal# Assertion failed: expected strings to be JSON-equal, but they are not
  actual: "{\"a\":0}"
expected: "{\"a\":1}"
 message: "index is 0"`)
}

func TestString_Matches(t *testing.T) {
	m := new(internal.MockTestingT)

	m.Reset()
	assert.ThatString(m, "this is an error").Matches("this is an error")
	assert.ThatString(t, m.String()).Equal("")

	m.Reset()
	assert.ThatString(m, "this is an error").Matches("an error (")
	assert.ThatString(t, m.String()).Equal(`error# Assertion failed: expected string to match the pattern, but it does not
  actual: "this is an error"
 pattern: "an error ("
   error: "error parsing regexp: missing closing ): ` + "`an error (`\"")

	m.Reset()
	assert.ThatString(m, "there's no error").Matches("an error")
	assert.ThatString(t, m.String()).Equal(`error# Assertion failed: expected string to match the pattern, but it does not
  actual: "there's no error"
 pattern: "an error"`)

	m.Reset()
	assert.ThatString(m, "there's no error").Must().Matches("an error", "index is 0")
	assert.ThatString(t, m.String()).Equal(`fatal# Assertion failed: expected string to match the pattern, but it does not
  actual: "there's no error"
 pattern: "an error"
 message: "index is 0"`)
}

func TestString_HasPrefix(t *testing.T) {
	m := new(internal.MockTestingT)

	m.Reset()
	assert.ThatString(m, "hello, world!").HasPrefix("hello")
	assert.ThatString(t, m.String()).Equal("")

	m.Reset()
	assert.ThatString(m, "hello, world!").HasPrefix("Hello, Jimmy!")
	assert.ThatString(t, m.String()).Equal(`error# Assertion failed: expected string to start with the specified prefix, but it does not
  actual: "hello, world!"
  prefix: "Hello, Jimmy!"`)

	m.Reset()
	assert.ThatString(m, "hello, world!").Must().HasPrefix("Hello, Jimmy!", "index is 0")
	assert.ThatString(t, m.String()).Equal(`fatal# Assertion failed: expected string to start with the specified prefix, but it does not
  actual: "hello, world!"
  prefix: "Hello, Jimmy!"
 message: "index is 0"`)
}

func TestString_HasSuffix(t *testing.T) {
	m := new(internal.MockTestingT)

	m.Reset()
	assert.ThatString(m, "hello, world!").HasSuffix("world!")
	assert.ThatString(t, m.String()).Equal("")

	m.Reset()
	assert.ThatString(m, "hello, world!").HasSuffix("Hello, Jimmy!")
	assert.ThatString(t, m.String()).Equal(`error# Assertion failed: expected string to end with the specified suffix, but it does not
  actual: "hello, world!"
  suffix: "Hello, Jimmy!"`)

	m.Reset()
	assert.ThatString(m, "hello, world!").Must().HasSuffix("Hello, Jimmy!", "index is 0")
	assert.ThatString(t, m.String()).Equal(`fatal# Assertion failed: expected string to end with the specified suffix, but it does not
  actual: "hello, world!"
  suffix: "Hello, Jimmy!"
 message: "index is 0"`)
}

func TestString_Contains(t *testing.T) {
	m := new(internal.MockTestingT)

	m.Reset()
	assert.ThatString(m, "hello, world!").Contains("hello")
	assert.ThatString(t, m.String()).Equal("")

	m.Reset()
	assert.ThatString(m, "hello, world!").Contains("Hello, Jimmy!")
	assert.ThatString(t, m.String()).Equal(`error# Assertion failed: expected string to contain the specified substring, but it does not
  actual: "hello, world!"
     sub: "Hello, Jimmy!"`)

	m.Reset()
	assert.ThatString(m, "hello, world!").Must().Contains("Hello, Jimmy!", "index is 0")
	assert.ThatString(t, m.String()).Equal(`fatal# Assertion failed: expected string to contain the specified substring, but it does not
  actual: "hello, world!"
     sub: "Hello, Jimmy!"
 message: "index is 0"`)
}

func TestString_IsLowerCase(t *testing.T) {
	m := new(internal.MockTestingT)

	m.Reset()
	assert.ThatString(m, "hello").IsLowerCase()
	assert.ThatString(t, m.String()).Equal("")

	m.Reset()
	assert.ThatString(m, "Hello").IsLowerCase()
	assert.ThatString(t, m.String()).Equal(`error# Assertion failed: expected string to be all lowercase, but it is not
  actual: "Hello"`)

	m.Reset()
	assert.ThatString(m, "Hello").Must().IsLowerCase("index is 0")
	assert.ThatString(t, m.String()).Equal(`fatal# Assertion failed: expected string to be all lowercase, but it is not
  actual: "Hello"
 message: "index is 0"`)
}

func TestString_IsUpperCase(t *testing.T) {
	m := new(internal.MockTestingT)

	m.Reset()
	assert.ThatString(m, "HELLO").IsUpperCase()
	assert.ThatString(t, m.String()).Equal("")

	m.Reset()
	assert.ThatString(m, "Hello").IsUpperCase()
	assert.ThatString(t, m.String()).Equal(`error# Assertion failed: expected string to be all uppercase, but it is not
  actual: "Hello"`)

	m.Reset()
	assert.ThatString(m, "Hello").Must().IsUpperCase("index is 0")
	assert.ThatString(t, m.String()).Equal(`fatal# Assertion failed: expected string to be all uppercase, but it is not
  actual: "Hello"
 message: "index is 0"`)
}

func TestString_IsNumeric(t *testing.T) {
	m := new(internal.MockTestingT)

	m.Reset()
	assert.ThatString(m, "12345").IsNumeric()
	assert.ThatString(t, m.String()).Equal("")

	m.Reset()
	assert.ThatString(m, "123a45").IsNumeric()
	assert.ThatString(t, m.String()).Equal(`error# Assertion failed: expected string to contain only digits, but it does not
  actual: "123a45"`)

	m.Reset()
	assert.ThatString(m, "123a45").Must().IsNumeric("index is 0")
	assert.ThatString(t, m.String()).Equal(`fatal# Assertion failed: expected string to contain only digits, but it does not
  actual: "123a45"
 message: "index is 0"`)
}

func TestString_IsAlpha(t *testing.T) {
	m := new(internal.MockTestingT)

	m.Reset()
	assert.ThatString(m, "abcdef").IsAlpha()
	assert.ThatString(t, m.String()).Equal("")

	m.Reset()
	assert.ThatString(m, "abc123").IsAlpha()
	assert.ThatString(t, m.String()).Equal(`error# Assertion failed: expected string to contain only letters, but it does not
  actual: "abc123"`)

	m.Reset()
	assert.ThatString(m, "abc123").Must().IsAlpha("index is 0")
	assert.ThatString(t, m.String()).Equal(`fatal# Assertion failed: expected string to contain only letters, but it does not
  actual: "abc123"
 message: "index is 0"`)
}

func TestString_IsAlphaNumeric(t *testing.T) {
	m := new(internal.MockTestingT)

	m.Reset()
	assert.ThatString(m, "abc123").IsAlphaNumeric()
	assert.ThatString(t, m.String()).Equal("")

	m.Reset()
	assert.ThatString(m, "abc@123").IsAlphaNumeric()
	assert.ThatString(t, m.String()).Equal(`error# Assertion failed: expected string to contain only letters and digits, but it does not
  actual: "abc@123"`)

	m.Reset()
	assert.ThatString(m, "abc@123").Must().IsAlphaNumeric("index is 0")
	assert.ThatString(t, m.String()).Equal(`fatal# Assertion failed: expected string to contain only letters and digits, but it does not
  actual: "abc@123"
 message: "index is 0"`)
}

func TestString_IsEmail(t *testing.T) {
	m := new(internal.MockTestingT)

	m.Reset()
	assert.ThatString(m, "test@example.com").IsEmail()
	assert.ThatString(t, m.String()).Equal("")

	m.Reset()
	assert.ThatString(m, "invalid-email").IsEmail()
	assert.ThatString(t, m.String()).Equal(`error# Assertion failed: expected string to be a valid email, but it is not
  actual: "invalid-email"`)

	m.Reset()
	assert.ThatString(m, "invalid-email").Must().IsEmail("index is 0")
	assert.ThatString(t, m.String()).Equal(`fatal# Assertion failed: expected string to be a valid email, but it is not
  actual: "invalid-email"
 message: "index is 0"`)
}

func TestString_IsURL(t *testing.T) {
	m := new(internal.MockTestingT)

	m.Reset()
	assert.ThatString(m, "https://www.example.com").IsURL()
	assert.ThatString(t, m.String()).Equal("")

	m.Reset()
	assert.ThatString(m, "invalid-url").IsURL()
	assert.ThatString(t, m.String()).Equal(`error# Assertion failed: expected string to be a valid URL, but it is not
  actual: "invalid-url"`)

	m.Reset()
	assert.ThatString(m, "invalid-url").Must().IsURL("index is 0")
	assert.ThatString(t, m.String()).Equal(`fatal# Assertion failed: expected string to be a valid URL, but it is not
  actual: "invalid-url"
 message: "index is 0"`)
}

func TestString_IsIPv4(t *testing.T) {
	m := new(internal.MockTestingT)

	m.Reset()
	assert.ThatString(m, "192.168.1.1").IsIPv4()
	assert.ThatString(t, m.String()).Equal("")

	m.Reset()
	assert.ThatString(m, "invalid-ip").IsIPv4()
	assert.ThatString(t, m.String()).Equal(`error# Assertion failed: expected string to be a valid IP, but it is not
  actual: "invalid-ip"`)

	m.Reset()
	assert.ThatString(m, "invalid-ip").Must().IsIPv4("index is 0")
	assert.ThatString(t, m.String()).Equal(`fatal# Assertion failed: expected string to be a valid IP, but it is not
  actual: "invalid-ip"
 message: "index is 0"`)
}

func TestString_IsHex(t *testing.T) {
	m := new(internal.MockTestingT)

	m.Reset()
	assert.ThatString(m, "abcdef123456").IsHex()
	assert.ThatString(t, m.String()).Equal("")

	m.Reset()
	assert.ThatString(m, "abcdefg").IsHex()
	assert.ThatString(t, m.String()).Equal(`error# Assertion failed: expected string to be a valid hexadecimal, but it is not
  actual: "abcdefg"`)

	m.Reset()
	assert.ThatString(m, "abcdefg").Must().IsHex("index is 0")
	assert.ThatString(t, m.String()).Equal(`fatal# Assertion failed: expected string to be a valid hexadecimal, but it is not
  actual: "abcdefg"
 message: "index is 0"`)
}

func TestString_IsBase64(t *testing.T) {
	m := new(internal.MockTestingT)

	m.Reset()
	assert.ThatString(m, "SGVsbG8gd29ybGQ=").IsBase64()
	assert.ThatString(t, m.String()).Equal("")

	m.Reset()
	assert.ThatString(m, "invalid-base64").IsBase64()
	assert.ThatString(t, m.String()).Equal(`error# Assertion failed: expected string to be a valid Base64, but it is not
  actual: "invalid-base64"`)

	m.Reset()
	assert.ThatString(m, "invalid-base64").Must().IsBase64("index is 0")
	assert.ThatString(t, m.String()).Equal(`fatal# Assertion failed: expected string to be a valid Base64, but it is not
  actual: "invalid-base64"
 message: "index is 0"`)
}
