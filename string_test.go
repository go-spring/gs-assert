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
)

func TestString_Length(t *testing.T) {
	m := new(MockTestingT)
	assert.ThatString(m, "0").Length(1)
	assert.ThatString(t, m.String()).Equal("")

	m.Reset()
	assert.ThatString(m, "0").Length(0)
	assert.ThatString(t, m.String()).Equal(`Assertion failed: length mismatch:
    got: length 1 (string) "0"
 expect: length 0`)
}

func TestString_Equal(t *testing.T) {
	m := new(MockTestingT)
	assert.ThatString(m, "0").Equal("0")
	assert.ThatString(t, m.String()).Equal("")

	m.Reset()
	assert.ThatString(m, "0").Equal("1")
	assert.ThatString(t, m.String()).Equal(`Assertion failed: strings not equal:
    got: (string) "0"
 expect: (string) "1"`)

	m.Reset()
	assert.ThatString(m, "0").Equal("1", "index is 0")
	assert.ThatString(t, m.String()).Equal(`Assertion failed: strings not equal:
    got: (string) "0"
 expect: (string) "1"
message: index is 0`)
}

func TestString_NotEqual(t *testing.T) {
	m := new(MockTestingT)
	assert.ThatString(m, "0").NotEqual("1")
	assert.ThatString(t, m.String()).Equal("")

	m.Reset()
	assert.ThatString(m, "0").NotEqual("0")
	assert.ThatString(t, m.String()).Equal(`Assertion failed: strings are equal:
    got: (string) "0"
 expect: not equal to "0"`)

	m.Reset()
	assert.ThatString(m, "0").NotEqual("0", "index is 0")
	assert.ThatString(t, m.String()).Equal(`Assertion failed: strings are equal:
    got: (string) "0"
 expect: not equal to "0"
message: index is 0`)
}

func TestString_JSONEqual(t *testing.T) {
	m := new(MockTestingT)
	assert.ThatString(m, `{"a":0,"b":1}`).JSONEqual(`{"b":1,"a":0}`)
	assert.ThatString(t, m.String()).Equal("")

	m.Reset()
	assert.ThatString(m, `this is an error`).JSONEqual(`[{"b":1},{"a":0}]`)
	assert.ThatString(t, m.String()).Equal(`Assertion failed: invalid JSON in got value:
    got: (string) "this is an error"
 expect: (string) "[{\"b\":1},{\"a\":0}]"
  error: invalid character 'h' in literal true (expecting 'r')`)

	m.Reset()
	assert.ThatString(m, `{"a":0,"b":1}`).JSONEqual(`this is an error`)
	assert.ThatString(t, m.String()).Equal(`Assertion failed: invalid JSON in expect value:
    got: (string) "{\"a\":0,\"b\":1}"
 expect: (string) "this is an error"
  error: invalid character 'h' in literal true (expecting 'r')`)

	m.Reset()
	assert.ThatString(m, `{"a":0,"b":1}`).JSONEqual(`[{"b":1},{"a":0}]`)
	assert.ThatString(t, m.String()).Equal(`Assertion failed: JSON structures are not equal:
    got: (string) "{\"a\":0,\"b\":1}"
 expect: (string) "[{\"b\":1},{\"a\":0}]"`)

	m.Reset()
	assert.ThatString(m, `{"a":0}`).JSONEqual(`{"a":1}`, "index is 0")
	assert.ThatString(t, m.String()).Equal(`Assertion failed: JSON structures are not equal:
    got: (string) "{\"a\":0}"
 expect: (string) "{\"a\":1}"
message: index is 0`)
}

func TestString_Matches(t *testing.T) {
	m := new(MockTestingT)
	assert.ThatString(m, "this is an error").Matches("this is an error")
	assert.ThatString(t, m.String()).Equal("")

	m.Reset()
	assert.ThatString(m, "this is an error").Matches("an error (")
	assert.ThatString(t, m.String()).Equal(`Assertion failed: string does not match the pattern:
    got: (string) "this is an error"
 expect: to match regex "an error ("
  error: error parsing regexp: missing closing ): ` + "`an error (`")

	m.Reset()
	assert.ThatString(m, "there's no error").Matches("an error")
	assert.ThatString(t, m.String()).Equal(`Assertion failed: string does not match the pattern:
    got: (string) "there's no error"
 expect: to match regex "an error"`)

	m.Reset()
	assert.ThatString(m, "there's no error").Matches("an error", "index is 0")
	assert.ThatString(t, m.String()).Equal(`Assertion failed: string does not match the pattern:
    got: (string) "there's no error"
 expect: to match regex "an error"
message: index is 0`)
}

func TestString_EqualFold(t *testing.T) {
	m := new(MockTestingT)
	assert.ThatString(m, "hello, world!").EqualFold("Hello, World!")
	assert.ThatString(t, m.String()).Equal("")

	m.Reset()
	assert.ThatString(m, "hello, world!").EqualFold("xxx")
	assert.ThatString(t, m.String()).Equal(`Assertion failed: strings are not equal under case-folding:
    got: (string) "hello, world!"
 expect: (string) "xxx"`)

	m.Reset()
	assert.ThatString(m, "hello, world!").EqualFold("xxx", "index is 0")
	assert.ThatString(t, m.String()).Equal(`Assertion failed: strings are not equal under case-folding:
    got: (string) "hello, world!"
 expect: (string) "xxx"
message: index is 0`)
}

func TestString_HasPrefix(t *testing.T) {
	m := new(MockTestingT)
	assert.ThatString(m, "hello, world!").HasPrefix("hello")
	assert.ThatString(t, m.String()).Equal("")

	m.Reset()
	assert.ThatString(m, "hello, world!").HasPrefix("xxx")
	assert.ThatString(t, m.String()).Equal(`Assertion failed: string does not start with the specified prefix:
    got: (string) "hello, world!"
 expect: to have prefix "xxx"`)

	m.Reset()
	assert.ThatString(m, "hello, world!").HasPrefix("xxx", "index is 0")
	assert.ThatString(t, m.String()).Equal(`Assertion failed: string does not start with the specified prefix:
    got: (string) "hello, world!"
 expect: to have prefix "xxx"
message: index is 0`)
}

func TestString_HasSuffix(t *testing.T) {
	m := new(MockTestingT)
	assert.ThatString(m, "hello, world!").HasSuffix("world!")
	assert.ThatString(t, m.String()).Equal("")

	m.Reset()
	assert.ThatString(m, "hello, world!").HasSuffix("xxx")
	assert.ThatString(t, m.String()).Equal(`Assertion failed: string does not end with the specified suffix:
    got: (string) "hello, world!"
 expect: to have suffix "xxx"`)

	m.Reset()
	assert.ThatString(m, "hello, world!").HasSuffix("xxx", "index is 0")
	assert.ThatString(t, m.String()).Equal(`Assertion failed: string does not end with the specified suffix:
    got: (string) "hello, world!"
 expect: to have suffix "xxx"
message: index is 0`)
}

func TestString_Contains(t *testing.T) {
	m := new(MockTestingT)
	assert.ThatString(m, "hello, world!").Contains("hello")
	assert.ThatString(t, m.String()).Equal("")

	m.Reset()
	assert.ThatString(m, "hello, world!").Contains("xxx")
	assert.ThatString(t, m.String()).Equal(`Assertion failed: string does not contain the specified substring:
    got: (string) "hello, world!"
 expect: to contain substring "xxx"`)

	m.Reset()
	assert.ThatString(m, "hello, world!").Contains("xxx", "index is 0")
	assert.ThatString(t, m.String()).Equal(`Assertion failed: string does not contain the specified substring:
    got: (string) "hello, world!"
 expect: to contain substring "xxx"
message: index is 0`)
}

func TestString_IsEmpty(t *testing.T) {
	m := new(MockTestingT)
	assert.ThatString(m, "").IsEmpty()
	assert.ThatString(t, m.String()).Equal("")

	m.Reset()
	assert.ThatString(m, "hello").IsEmpty()
	assert.ThatString(t, m.String()).Equal(`Assertion failed: string is not empty:
    got: (string) "hello"
 expect: empty string`)

	m.Reset()
	assert.ThatString(m, "hello").IsEmpty("index is 0")
	assert.ThatString(t, m.String()).Equal(`Assertion failed: string is not empty:
    got: (string) "hello"
 expect: empty string
message: index is 0`)
}

func TestString_IsNotEmpty(t *testing.T) {
	m := new(MockTestingT)
	assert.ThatString(m, "hello").IsNotEmpty()
	assert.ThatString(t, m.String()).Equal("")

	m.Reset()
	assert.ThatString(m, "").IsNotEmpty()
	assert.ThatString(t, m.String()).Equal(`Assertion failed: string is empty:
    got: (string) ""
 expect: non-empty string`)

	m.Reset()
	assert.ThatString(m, "").IsNotEmpty("index is 0")
	assert.ThatString(t, m.String()).Equal(`Assertion failed: string is empty:
    got: (string) ""
 expect: non-empty string
message: index is 0`)
}

func TestString_IsBlank(t *testing.T) {
	m := new(MockTestingT)
	assert.ThatString(m, "   ").IsBlank()
	assert.ThatString(t, m.String()).Equal("")

	m.Reset()
	assert.ThatString(m, "hello").IsBlank()
	assert.ThatString(t, m.String()).Equal(`Assertion failed: string contains non-whitespace characters:
    got: (string) "hello"
 expect: blank string`)

	m.Reset()
	assert.ThatString(m, "hello").IsBlank("index is 0")
	assert.ThatString(t, m.String()).Equal(`Assertion failed: string contains non-whitespace characters:
    got: (string) "hello"
 expect: blank string
message: index is 0`)
}

func TestString_IsNotBlank(t *testing.T) {
	m := new(MockTestingT)
	assert.ThatString(m, "hello").IsNotBlank()
	assert.ThatString(t, m.String()).Equal("")

	m.Reset()
	assert.ThatString(m, "   ").IsNotBlank()
	assert.ThatString(t, m.String()).Equal(`Assertion failed: string is blank:
    got: (string) "   "
 expect: non-blank string`)

	m.Reset()
	assert.ThatString(m, "   ").IsNotBlank("index is 0")
	assert.ThatString(t, m.String()).Equal(`Assertion failed: string is blank:
    got: (string) "   "
 expect: non-blank string
message: index is 0`)
}

func TestString_IsLowerCase(t *testing.T) {
	m := new(MockTestingT)
	assert.ThatString(m, "hello").IsLowerCase()
	assert.ThatString(t, m.String()).Equal("")

	m.Reset()
	assert.ThatString(m, "Hello").IsLowerCase()
	assert.ThatString(t, m.String()).Equal(`Assertion failed: string contains uppercase characters:
    got: (string) "Hello"
 expect: lowercase string`)

	m.Reset()
	assert.ThatString(m, "Hello").IsLowerCase("index is 0")
	assert.ThatString(t, m.String()).Equal(`Assertion failed: string contains uppercase characters:
    got: (string) "Hello"
 expect: lowercase string
message: index is 0`)
}

func TestString_IsUpperCase(t *testing.T) {
	m := new(MockTestingT)
	assert.ThatString(m, "HELLO").IsUpperCase()
	assert.ThatString(t, m.String()).Equal("")

	m.Reset()
	assert.ThatString(m, "Hello").IsUpperCase()
	assert.ThatString(t, m.String()).Equal(`Assertion failed: string contains lowercase characters:
    got: (string) "Hello"
 expect: uppercase string`)

	m.Reset()
	assert.ThatString(m, "Hello").IsUpperCase("index is 0")
	assert.ThatString(t, m.String()).Equal(`Assertion failed: string contains lowercase characters:
    got: (string) "Hello"
 expect: uppercase string
message: index is 0`)
}

func TestString_IsNumeric(t *testing.T) {
	m := new(MockTestingT)
	assert.ThatString(m, "12345").IsNumeric()
	assert.ThatString(t, m.String()).Equal("")

	m.Reset()
	assert.ThatString(m, "123a45").IsNumeric()
	assert.ThatString(t, m.String()).Equal(`Assertion failed: string contains non-numeric characters:
    got: (string) "123a45"
 expect: numeric string`)

	m.Reset()
	assert.ThatString(m, "123a45").IsNumeric("index is 0")
	assert.ThatString(t, m.String()).Equal(`Assertion failed: string contains non-numeric characters:
    got: (string) "123a45"
 expect: numeric string
message: index is 0`)
}

func TestString_IsAlpha(t *testing.T) {
	m := new(MockTestingT)
	assert.ThatString(m, "abcdef").IsAlpha()
	assert.ThatString(t, m.String()).Equal("")

	m.Reset()
	assert.ThatString(m, "abc123").IsAlpha()
	assert.ThatString(t, m.String()).Equal(`Assertion failed: string contains non-alphabetic characters:
    got: (string) "abc123"
 expect: alphabetic string`)

	m.Reset()
	assert.ThatString(m, "abc123").IsAlpha("index is 0")
	assert.ThatString(t, m.String()).Equal(`Assertion failed: string contains non-alphabetic characters:
    got: (string) "abc123"
 expect: alphabetic string
message: index is 0`)
}

func TestString_IsAlphaNumeric(t *testing.T) {
	m := new(MockTestingT)
	assert.ThatString(m, "abc123").IsAlphaNumeric()
	assert.ThatString(t, m.String()).Equal("")

	m.Reset()
	assert.ThatString(m, "abc@123").IsAlphaNumeric()
	assert.ThatString(t, m.String()).Equal(`Assertion failed: string contains non-alphanumeric characters:
    got: (string) "abc@123"
 expect: alphanumeric string`)

	m.Reset()
	assert.ThatString(m, "abc@123").IsAlphaNumeric("index is 0")
	assert.ThatString(t, m.String()).Equal(`Assertion failed: string contains non-alphanumeric characters:
    got: (string) "abc@123"
 expect: alphanumeric string
message: index is 0`)
}

func TestString_IsEmail(t *testing.T) {
	m := new(MockTestingT)
	assert.ThatString(m, "test@example.com").IsEmail()
	assert.ThatString(t, m.String()).Equal("")

	m.Reset()
	assert.ThatString(m, "invalid-email").IsEmail()
	assert.ThatString(t, m.String()).Equal(`Assertion failed: string is not a valid email:
    got: (string) "invalid-email"
 expect: valid email address`)

	m.Reset()
	assert.ThatString(m, "invalid-email").IsEmail("index is 0")
	assert.ThatString(t, m.String()).Equal(`Assertion failed: string is not a valid email:
    got: (string) "invalid-email"
 expect: valid email address
message: index is 0`)
}

func TestString_IsURL(t *testing.T) {
	m := new(MockTestingT)
	assert.ThatString(m, "https://www.example.com").IsURL()
	assert.ThatString(t, m.String()).Equal("")

	m.Reset()
	assert.ThatString(m, "invalid-url").IsURL()
	assert.ThatString(t, m.String()).Equal(`Assertion failed: string is not a valid URL:
    got: (string) "invalid-url"
 expect: valid URL`)

	m.Reset()
	assert.ThatString(m, "invalid-url").IsURL("index is 0")
	assert.ThatString(t, m.String()).Equal(`Assertion failed: string is not a valid URL:
    got: (string) "invalid-url"
 expect: valid URL
message: index is 0`)
}

func TestString_IsIP(t *testing.T) {
	m := new(MockTestingT)
	assert.ThatString(m, "192.168.1.1").IsIP()
	assert.ThatString(t, m.String()).Equal("")

	m.Reset()
	assert.ThatString(m, "invalid-ip").IsIP()
	assert.ThatString(t, m.String()).Equal(`Assertion failed: string is not a valid IP:
    got: (string) "invalid-ip"
 expect: valid IP address`)

	m.Reset()
	assert.ThatString(m, "invalid-ip").IsIP("index is 0")
	assert.ThatString(t, m.String()).Equal(`Assertion failed: string is not a valid IP:
    got: (string) "invalid-ip"
 expect: valid IP address
message: index is 0`)
}

func TestString_IsHex(t *testing.T) {
	m := new(MockTestingT)
	assert.ThatString(m, "abcdef123456").IsHex()
	assert.ThatString(t, m.String()).Equal("")

	m.Reset()
	assert.ThatString(m, "abcdefg").IsHex()
	assert.ThatString(t, m.String()).Equal(`Assertion failed: string is not a valid hexadecimal:
    got: (string) "abcdefg"
 expect: valid hexadecimal number`)

	m.Reset()
	assert.ThatString(m, "abcdefg").IsHex("index is 0")
	assert.ThatString(t, m.String()).Equal(`Assertion failed: string is not a valid hexadecimal:
    got: (string) "abcdefg"
 expect: valid hexadecimal number
message: index is 0`)
}

func TestString_IsBase64(t *testing.T) {
	m := new(MockTestingT)
	assert.ThatString(m, "SGVsbG8gd29ybGQ=").IsBase64()
	assert.ThatString(t, m.String()).Equal("")

	m.Reset()
	assert.ThatString(m, "invalid-base64").IsBase64()
	assert.ThatString(t, m.String()).Equal(`Assertion failed: string is not a valid Base64:
    got: (string) "invalid-base64"
 expect: valid Base64 encoded string`)

	m.Reset()
	assert.ThatString(m, "invalid-base64").IsBase64("index is 0")
	assert.ThatString(t, m.String()).Equal(`Assertion failed: string is not a valid Base64:
    got: (string) "invalid-base64"
 expect: valid Base64 encoded string
message: index is 0`)
}
