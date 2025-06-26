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
	"errors"
	"testing"

	"github.com/go-spring/assert"
	"github.com/go-spring/assert/internal"
)

func TestError_Nil(t *testing.T) {
	m := new(internal.MockTestingT)

	assert.ThatError(m, nil).Nil()
	assert.ThatString(t, m.String()).Equal("")

	m.Reset()
	assert.ThatError(m, errors.New("this is an error")).Nil()
	assert.ThatString(t, m.String()).Equal(`error# Assertion failed: expected error to be nil, but it is not
    got: this is an error`)

	m.Reset()
	assert.ThatError(m, errors.New("this is an error")).Must().Nil("index is 0")
	assert.ThatString(t, m.String()).Equal(`fatal# Assertion failed: expected error to be nil, but it is not
    got: this is an error
message: "index is 0"`)
}

func TestError_NotNil(t *testing.T) {
	m := new(internal.MockTestingT)

	assert.ThatError(m, errors.New("this is an error")).NotNil()
	assert.ThatString(t, m.String()).Equal("")

	m.Reset()
	assert.ThatError(m, nil).NotNil()
	assert.ThatString(t, m.String()).Equal(`error# Assertion failed: expected error to be non-nil, but it is nil`)

	m.Reset()
	assert.ThatError(m, nil).Must().NotNil("index is 0")
	assert.ThatString(t, m.String()).Equal(`fatal# Assertion failed: expected error to be non-nil, but it is nil
message: "index is 0"`)
}

func TestError_Is(t *testing.T) {
	m := new(internal.MockTestingT)

	err := errors.New("this is an error")
	assert.ThatError(m, err).Is(err)
	assert.ThatString(t, m.String()).Equal("")

	m.Reset()
	assert.ThatError(m, err).Is(errors.New("another error"))
	assert.ThatString(t, m.String()).Equal(`error# Assertion failed: expected error to be equal to target, but they are different 
    got: this is an error
 expect: another error`)

	m.Reset()
	assert.ThatError(m, err).Must().Is(errors.New("another error"), "index is 0")
	assert.ThatString(t, m.String()).Equal(`fatal# Assertion failed: expected error to be equal to target, but they are different 
    got: this is an error
 expect: another error
message: "index is 0"`)
}

func TestError_NotIs(t *testing.T) {
	m := new(internal.MockTestingT)

	err := errors.New("this is an error")
	assert.ThatError(m, err).NotIs(errors.New("another error"))
	assert.ThatString(t, m.String()).Equal("")

	m.Reset()
	assert.ThatError(m, err).NotIs(err)
	assert.ThatString(t, m.String()).Equal(`error# Assertion failed: expected error not to be equal to target, but they are equal 
    got: this is an error
 expect: this is an error`)

	m.Reset()
	assert.ThatError(m, err).Must().NotIs(err, "index is 0")
	assert.ThatString(t, m.String()).Equal(`fatal# Assertion failed: expected error not to be equal to target, but they are equal 
    got: this is an error
 expect: this is an error
message: "index is 0"`)
}

func TestError_ContainsMessage(t *testing.T) {
	m := new(internal.MockTestingT)

	err := errors.New("this is an error")
	assert.ThatError(m, err).ContainsMessage("an error")
	assert.ThatString(t, m.String()).Equal("")

	m.Reset()
	assert.ThatError(m, err).ContainsMessage("not in message")
	assert.ThatString(t, m.String()).Equal(`error# Assertion failed: expected error message to contain "not in message", but it does not
    got: "this is an error"`)

	m.Reset()
	assert.ThatError(m, err).Must().ContainsMessage("not in message", "index is 0")
	assert.ThatString(t, m.String()).Equal(`fatal# Assertion failed: expected error message to contain "not in message", but it does not
    got: "this is an error"
message: "index is 0"`)
}

func TestError_Matches(t *testing.T) {
	m := new(internal.MockTestingT)
	assert.ThatError(m, errors.New("this is an error")).Matches("an error")
	assert.ThatString(t, m.String()).Equal("")

	m.Reset()
	assert.ThatError(m, errors.New("there's no error")).Matches(`an error \`)
	assert.ThatString(t, m.String()).Equal("error# Assertion failed: invalid pattern")

	m.Reset()
	assert.ThatError(m, nil).Matches("an error")
	assert.ThatString(t, m.String()).Equal("error# Assertion failed: expected non-nil error, but got nil")

	m.Reset()
	assert.ThatError(m, nil).Matches("an error", "index is 0")
	assert.ThatString(t, m.String()).Equal(`error# Assertion failed: expected non-nil error, but got nil
message: "index is 0"`)

	m.Reset()
	assert.ThatError(m, errors.New("there's no error")).Must().Matches("an error")
	assert.ThatString(t, m.String()).Equal(`fatal# Assertion failed: got "there's no error" which does not match "an error"`)

	m.Reset()
	assert.ThatError(m, errors.New("there's no error")).Must().Matches("an error", "index is 0")
	assert.ThatString(t, m.String()).Equal(`fatal# Assertion failed: got "there's no error" which does not match "an error"
message: "index is 0"`)
}
