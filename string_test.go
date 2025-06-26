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
	"testing"

	"github.com/go-spring/assert"
	"github.com/go-spring/assert/internal"
)

func TestString_Length(t *testing.T) {
	m := new(internal.MockTestingT)
	assert.ThatString(m, "0").Length(1)
	assert.ThatString(t, m.String()).Equal("")

	m.Reset()
	assert.ThatString(m, "0").Length(0)
	assert.ThatString(t, m.String()).Equal(`error# Assertion failed: expected string to have length 0, but it has length 1
    got: "0"`)

	m.Reset()
	assert.ThatString(m, "0").Must().Length(0, "index is 0")
	assert.ThatString(t, m.String()).Equal(`fatal# Assertion failed: expected string to have length 0, but it has length 1
    got: "0"
message: "index is 0"`)
}

func TestString_Equal(t *testing.T) {
	m := new(internal.MockTestingT)
	assert.ThatString(m, "0").Equal("0")
	assert.ThatString(t, m.String()).Equal("")

	m.Reset()
	assert.ThatString(m, "0").Equal("1")
	assert.ThatString(t, m.String()).Equal(`error# Assertion failed: expected strings to be equal, but they are not
    got: "0"
 expect: "1"`)

	m.Reset()
	assert.ThatString(m, "0").Must().Equal("1", "index is 0")
	assert.ThatString(t, m.String()).Equal(`fatal# Assertion failed: expected strings to be equal, but they are not
    got: "0"
 expect: "1"
message: "index is 0"`)
}

func TestString_NotEqual(t *testing.T) {
	m := new(internal.MockTestingT)
	assert.ThatString(m, "0").NotEqual("1")
	assert.ThatString(t, m.String()).Equal("")

	m.Reset()
	assert.ThatString(m, "0").NotEqual("0")
	assert.ThatString(t, m.String()).Equal(`error# Assertion failed: expected strings to be different, but they are equal
    got: "0"
 expect: "0"`)

	m.Reset()
	assert.ThatString(m, "0").Must().NotEqual("0", "index is 0")
	assert.ThatString(t, m.String()).Equal(`fatal# Assertion failed: expected strings to be different, but they are equal
    got: "0"
 expect: "0"
message: "index is 0"`)
}

func TestString_JSONEqual(t *testing.T) {
	m := new(internal.MockTestingT)
	assert.ThatString(m, `{"a":0,"b":1}`).JSONEqual(`{"b":1,"a":0}`)
	assert.ThatString(t, m.String()).Equal("")

	m.Reset()
	assert.ThatString(m, `this is an error`).JSONEqual(`[{"b":1},{"a":0}]`)
	assert.ThatString(t, m.String()).Equal(`error# Assertion failed: expected strings to be JSON-equal, but failed to unmarshal got value
    got: "this is an error"
  error: "invalid character 'h' in literal true (expecting 'r')"`)

	m.Reset()
	assert.ThatString(m, `{"a":0,"b":1}`).Must().JSONEqual(`this is an error`)
	assert.ThatString(t, m.String()).Equal(`fatal# Assertion failed: expected strings to be JSON-equal, but failed to unmarshal expect value
 expect: "this is an error"
  error: "invalid character 'h' in literal true (expecting 'r')"`)

	m.Reset()
	assert.ThatString(m, `{"a":0,"b":1}`).JSONEqual(`[{"b":1},{"a":0}]`)
	assert.ThatString(t, m.String()).Equal(`error# Assertion failed: expected strings to be JSON-equal, but they are not
    got: "{\"a\":0,\"b\":1}"
 expect: "[{\"b\":1},{\"a\":0}]"`)

	m.Reset()
	assert.ThatString(m, `{"a":0}`).Must().JSONEqual(`{"a":1}`, "index is 0")
	assert.ThatString(t, m.String()).Equal(`fatal# Assertion failed: expected strings to be JSON-equal, but they are not
    got: "{\"a\":0}"
 expect: "{\"a\":1}"
message: "index is 0"`)
}

func TestString_Matches(t *testing.T) {
	m := new(internal.MockTestingT)
	assert.ThatString(m, "this is an error").Matches("this is an error")
	assert.ThatString(t, m.String()).Equal("")

	m.Reset()
	assert.ThatString(m, "this is an error").Matches("an error (")
	assert.ThatString(t, m.String()).Equal(`error# Assertion failed: expected string to match the pattern, but it does not
    got: "this is an error"
pattern: "an error ("
  error: "error parsing regexp: missing closing ): ` + "`an error (`\"")

	m.Reset()
	assert.ThatString(m, "there's no error").Matches("an error")
	assert.ThatString(t, m.String()).Equal(`error# Assertion failed: expected string to match the pattern, but it does not
    got: "there's no error"
pattern: "an error"`)

	m.Reset()
	assert.ThatString(m, "there's no error").Must().Matches("an error", "index is 0")
	assert.ThatString(t, m.String()).Equal(`fatal# Assertion failed: expected string to match the pattern, but it does not
    got: "there's no error"
pattern: "an error"
message: "index is 0"`)
}

func TestString_EqualFold(t *testing.T) {
	m := new(internal.MockTestingT)
	assert.ThatString(m, "hello, world!").EqualFold("Hello, World!")
	assert.ThatString(t, m.String()).Equal("")

	m.Reset()
	assert.ThatString(m, "hello, world!").EqualFold("xxx")
	assert.ThatString(t, m.String()).Equal(`error# Assertion failed: expected strings to be equal (case-insensitive), but they are not
    got: "hello, world!"
 expect: "xxx"`)

	m.Reset()
	assert.ThatString(m, "hello, world!").Must().EqualFold("xxx", "index is 0")
	assert.ThatString(t, m.String()).Equal(`fatal# Assertion failed: expected strings to be equal (case-insensitive), but they are not
    got: "hello, world!"
 expect: "xxx"
message: "index is 0"`)
}

func TestString_HasPrefix(t *testing.T) {
	m := new(internal.MockTestingT)
	assert.ThatString(m, "hello, world!").HasPrefix("hello")
	assert.ThatString(t, m.String()).Equal("")

	m.Reset()
	assert.ThatString(m, "hello, world!").HasPrefix("xxx")
	assert.ThatString(t, m.String()).Equal(`error# Assertion failed: expected string to start with the specified prefix, but it does not
    got: "hello, world!"
 prefix: "xxx"`)

	m.Reset()
	assert.ThatString(m, "hello, world!").Must().HasPrefix("xxx", "index is 0")
	assert.ThatString(t, m.String()).Equal(`fatal# Assertion failed: expected string to start with the specified prefix, but it does not
    got: "hello, world!"
 prefix: "xxx"
message: "index is 0"`)
}

func TestString_HasSuffix(t *testing.T) {
	m := new(internal.MockTestingT)
	assert.ThatString(m, "hello, world!").HasSuffix("world!")
	assert.ThatString(t, m.String()).Equal("")

	m.Reset()
	assert.ThatString(m, "hello, world!").HasSuffix("xxx")
	assert.ThatString(t, m.String()).Equal(`error# Assertion failed: expected string to end with the specified suffix, but it does not
    got: "hello, world!"
 suffix: "xxx"`)

	m.Reset()
	assert.ThatString(m, "hello, world!").Must().HasSuffix("xxx", "index is 0")
	assert.ThatString(t, m.String()).Equal(`fatal# Assertion failed: expected string to end with the specified suffix, but it does not
    got: "hello, world!"
 suffix: "xxx"
message: "index is 0"`)
}

func TestString_Contains(t *testing.T) {
	m := new(internal.MockTestingT)
	assert.ThatString(m, "hello, world!").Contains("hello")
	assert.ThatString(t, m.String()).Equal("")

	m.Reset()
	assert.ThatString(m, "hello, world!").Contains("xxx")
	assert.ThatString(t, m.String()).Equal(`error# Assertion failed: expected string to contain the specified substring, but it does not
    got: "hello, world!"
 substr: "xxx"`)

	m.Reset()
	assert.ThatString(m, "hello, world!").Must().Contains("xxx", "index is 0")
	assert.ThatString(t, m.String()).Equal(`fatal# Assertion failed: expected string to contain the specified substring, but it does not
    got: "hello, world!"
 substr: "xxx"
message: "index is 0"`)
}

func TestString_Empty(t *testing.T) {
	m := new(internal.MockTestingT)
	assert.ThatString(m, "").Empty()
	assert.ThatString(t, m.String()).Equal("")

	m.Reset()
	assert.ThatString(m, "hello").Empty()
	assert.ThatString(t, m.String()).Equal(`error# Assertion failed: expected string to be empty, but it is not
    got: "hello"`)

	m.Reset()
	assert.ThatString(m, "hello").Must().Empty("index is 0")
	assert.ThatString(t, m.String()).Equal(`fatal# Assertion failed: expected string to be empty, but it is not
    got: "hello"
message: "index is 0"`)
}

func TestString_NotEmpty(t *testing.T) {
	m := new(internal.MockTestingT)
	assert.ThatString(m, "hello").NotEmpty()
	assert.ThatString(t, m.String()).Equal("")

	m.Reset()
	assert.ThatString(m, "").NotEmpty()
	assert.ThatString(t, m.String()).Equal(`error# Assertion failed: expected string to be non-empty, but it is empty
    got: ""`)

	m.Reset()
	assert.ThatString(m, "").Must().NotEmpty("index is 0")
	assert.ThatString(t, m.String()).Equal(`fatal# Assertion failed: expected string to be non-empty, but it is empty
    got: ""
message: "index is 0"`)
}

func TestString_Blank(t *testing.T) {
	m := new(internal.MockTestingT)
	assert.ThatString(m, "   ").Blank()
	assert.ThatString(t, m.String()).Equal("")

	m.Reset()
	assert.ThatString(m, "hello").Blank()
	assert.ThatString(t, m.String()).Equal(`error# Assertion failed: expected string to contain only whitespace, but it does not
    got: "hello"`)

	m.Reset()
	assert.ThatString(m, "hello").Must().Blank("index is 0")
	assert.ThatString(t, m.String()).Equal(`fatal# Assertion failed: expected string to contain only whitespace, but it does not
    got: "hello"
message: "index is 0"`)
}

func TestString_NotBlank(t *testing.T) {
	m := new(internal.MockTestingT)
	assert.ThatString(m, "hello").NotBlank()
	assert.ThatString(t, m.String()).Equal("")

	m.Reset()
	assert.ThatString(m, "   ").NotBlank()
	assert.ThatString(t, m.String()).Equal(`error# Assertion failed: expected string to be non-blank, but it is blank
    got: "   "`)

	m.Reset()
	assert.ThatString(m, " \n  ").Must().NotBlank("index is 0")
	assert.ThatString(t, m.String()).Equal(`fatal# Assertion failed: expected string to be non-blank, but it is blank
    got: " \n  "
message: "index is 0"`)
}

func TestString_IsLowerCase(t *testing.T) {
	m := new(internal.MockTestingT)
	assert.ThatString(m, "hello").IsLowerCase()
	assert.ThatString(t, m.String()).Equal("")

	m.Reset()
	assert.ThatString(m, "Hello").IsLowerCase()
	assert.ThatString(t, m.String()).Equal(`error# Assertion failed: expected string to be all lowercase, but it is not
    got: "Hello"`)

	m.Reset()
	assert.ThatString(m, "Hello").Must().IsLowerCase("index is 0")
	assert.ThatString(t, m.String()).Equal(`fatal# Assertion failed: expected string to be all lowercase, but it is not
    got: "Hello"
message: "index is 0"`)
}

func TestString_IsUpperCase(t *testing.T) {
	m := new(internal.MockTestingT)
	assert.ThatString(m, "HELLO").IsUpperCase()
	assert.ThatString(t, m.String()).Equal("")

	m.Reset()
	assert.ThatString(m, "Hello").IsUpperCase()
	assert.ThatString(t, m.String()).Equal(`error# Assertion failed: expected string to be all uppercase, but it is not
    got: "Hello"`)

	m.Reset()
	assert.ThatString(m, "Hello").Must().IsUpperCase("index is 0")
	assert.ThatString(t, m.String()).Equal(`fatal# Assertion failed: expected string to be all uppercase, but it is not
    got: "Hello"
message: "index is 0"`)
}

func TestString_IsNumeric(t *testing.T) {
	m := new(internal.MockTestingT)
	assert.ThatString(m, "12345").IsNumeric()
	assert.ThatString(t, m.String()).Equal("")

	m.Reset()
	assert.ThatString(m, "123a45").IsNumeric()
	assert.ThatString(t, m.String()).Equal(`error# Assertion failed: expected string to contain only digits, but it does not
    got: "123a45"`)

	m.Reset()
	assert.ThatString(m, "123a45").Must().IsNumeric("index is 0")
	assert.ThatString(t, m.String()).Equal(`fatal# Assertion failed: expected string to contain only digits, but it does not
    got: "123a45"
message: "index is 0"`)
}

func TestString_IsAlpha(t *testing.T) {
	m := new(internal.MockTestingT)
	assert.ThatString(m, "abcdef").IsAlpha()
	assert.ThatString(t, m.String()).Equal("")

	m.Reset()
	assert.ThatString(m, "abc123").IsAlpha()
	assert.ThatString(t, m.String()).Equal(`error# Assertion failed: expected string to contain only letters, but it does not
    got: "abc123"`)

	m.Reset()
	assert.ThatString(m, "abc123").Must().IsAlpha("index is 0")
	assert.ThatString(t, m.String()).Equal(`fatal# Assertion failed: expected string to contain only letters, but it does not
    got: "abc123"
message: "index is 0"`)
}

func TestString_IsAlphaNumeric(t *testing.T) {
	m := new(internal.MockTestingT)
	assert.ThatString(m, "abc123").IsAlphaNumeric()
	assert.ThatString(t, m.String()).Equal("")

	m.Reset()
	assert.ThatString(m, "abc@123").IsAlphaNumeric()
	assert.ThatString(t, m.String()).Equal(`error# Assertion failed: expected string to contain only letters and digits, but it does not
    got: "abc@123"`)

	m.Reset()
	assert.ThatString(m, "abc@123").Must().IsAlphaNumeric("index is 0")
	assert.ThatString(t, m.String()).Equal(`fatal# Assertion failed: expected string to contain only letters and digits, but it does not
    got: "abc@123"
message: "index is 0"`)
}

func TestString_IsEmail(t *testing.T) {
	m := new(internal.MockTestingT)
	assert.ThatString(m, "test@example.com").IsEmail()
	assert.ThatString(t, m.String()).Equal("")

	m.Reset()
	assert.ThatString(m, "invalid-email").IsEmail()
	assert.ThatString(t, m.String()).Equal(`error# Assertion failed: expected string to be a valid email, but it is not
    got: "invalid-email"`)

	m.Reset()
	assert.ThatString(m, "invalid-email").Must().IsEmail("index is 0")
	assert.ThatString(t, m.String()).Equal(`fatal# Assertion failed: expected string to be a valid email, but it is not
    got: "invalid-email"
message: "index is 0"`)
}

func TestString_IsURL(t *testing.T) {
	m := new(internal.MockTestingT)
	assert.ThatString(m, "https://www.example.com").IsURL()
	assert.ThatString(t, m.String()).Equal("")

	m.Reset()
	assert.ThatString(m, "invalid-url").IsURL()
	assert.ThatString(t, m.String()).Equal(`error# Assertion failed: expected string to be a valid URL, but it is not
    got: "invalid-url"`)

	m.Reset()
	assert.ThatString(m, "invalid-url").Must().IsURL("index is 0")
	assert.ThatString(t, m.String()).Equal(`fatal# Assertion failed: expected string to be a valid URL, but it is not
    got: "invalid-url"
message: "index is 0"`)
}

func TestString_IsIP(t *testing.T) {
	m := new(internal.MockTestingT)
	assert.ThatString(m, "192.168.1.1").IsIP()
	assert.ThatString(t, m.String()).Equal("")

	m.Reset()
	assert.ThatString(m, "invalid-ip").IsIP()
	assert.ThatString(t, m.String()).Equal(`error# Assertion failed: expected string to be a valid IP, but it is not
    got: "invalid-ip"`)

	m.Reset()
	assert.ThatString(m, "invalid-ip").Must().IsIP("index is 0")
	assert.ThatString(t, m.String()).Equal(`fatal# Assertion failed: expected string to be a valid IP, but it is not
    got: "invalid-ip"
message: "index is 0"`)
}

func TestString_IsHex(t *testing.T) {
	m := new(internal.MockTestingT)
	assert.ThatString(m, "abcdef123456").IsHex()
	assert.ThatString(t, m.String()).Equal("")

	m.Reset()
	assert.ThatString(m, "abcdefg").IsHex()
	assert.ThatString(t, m.String()).Equal(`error# Assertion failed: expected string to be a valid hexadecimal, but it is not
    got: "abcdefg"`)

	m.Reset()
	assert.ThatString(m, "abcdefg").Must().IsHex("index is 0")
	assert.ThatString(t, m.String()).Equal(`fatal# Assertion failed: expected string to be a valid hexadecimal, but it is not
    got: "abcdefg"
message: "index is 0"`)
}

func TestString_IsBase64(t *testing.T) {
	m := new(internal.MockTestingT)
	assert.ThatString(m, "SGVsbG8gd29ybGQ=").IsBase64()
	assert.ThatString(t, m.String()).Equal("")

	m.Reset()
	assert.ThatString(m, "invalid-base64").IsBase64()
	assert.ThatString(t, m.String()).Equal(`error# Assertion failed: expected string to be a valid Base64, but it is not
    got: "invalid-base64"`)

	m.Reset()
	assert.ThatString(m, "invalid-base64").Must().IsBase64("index is 0")
	assert.ThatString(t, m.String()).Equal(`fatal# Assertion failed: expected string to be a valid Base64, but it is not
    got: "invalid-base64"
message: "index is 0"`)
}
