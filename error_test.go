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
)

func TestError_Matches(t *testing.T) {
	m := new(MockTestingT)
	assert.ThatError(m, errors.New("this is an error")).Matches("an error")
	assert.ThatString(t, m.String()).Equal("")

	m.Reset()
	assert.ThatError(m, errors.New("there's no error")).Matches(`an error \`)
	assert.ThatString(t, m.String()).Equal("Assertion failed: invalid pattern")

	m.Reset()
	assert.ThatError(m, nil).Matches("an error")
	assert.ThatString(t, m.String()).Equal("Assertion failed: expect not nil error")

	m.Reset()
	assert.ThatError(m, nil).Matches("an error", "index is 0")
	assert.ThatString(t, m.String()).Equal(`Assertion failed: expect not nil error
message: index is 0`)

	m.Reset()
	assert.ThatError(m, errors.New("there's no error")).Matches("an error")
	assert.ThatString(t, m.String()).Equal(`Assertion failed: got "there's no error" which does not match "an error"`)

	m.Reset()
	assert.ThatError(m, errors.New("there's no error")).Matches("an error", "index is 0")
	assert.ThatString(t, m.String()).Equal(`Assertion failed: got "there's no error" which does not match "an error"
message: index is 0`)
}
