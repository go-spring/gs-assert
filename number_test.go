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
	"math"
	"testing"

	"github.com/go-spring/assert"
)

func TestNumber_Equal(t *testing.T) {
	m := new(MockTestingT)
	assert.ThatNumber(m, 5).Equal(5)
	assert.ThatString(t, m.String()).Equal("")

	m.Reset()
	assert.ThatNumber(m, 5).Equal(10)
	assert.ThatString(t, m.String()).Equal(`Assertion failed: values not equal:
    got: (int) 5
 expect: (int) 10`)

	m.Reset()
	assert.ThatNumber(m, 5).Equal(10, "index is 0")
	assert.ThatString(t, m.String()).Equal(`Assertion failed: values not equal:
    got: (int) 5
 expect: (int) 10
message: index is 0`)
}

func TestNumber_NotEqual(t *testing.T) {
	m := new(MockTestingT)
	assert.ThatNumber(m, 5).NotEqual(10)
	assert.ThatString(t, m.String()).Equal("")

	m.Reset()
	assert.ThatNumber(m, 5).NotEqual(5)
	assert.ThatString(t, m.String()).Equal(`Assertion failed: values are equal:
    got: (int) 5
 expect: not equal to (int) 5`)

	m.Reset()
	assert.ThatNumber(m, 5).NotEqual(5, "index is 0")
	assert.ThatString(t, m.String()).Equal(`Assertion failed: values are equal:
    got: (int) 5
 expect: not equal to (int) 5
message: index is 0`)
}

func TestNumber_GreaterThan(t *testing.T) {
	m := new(MockTestingT)
	assert.ThatNumber(m, 10).GreaterThan(5)
	assert.ThatString(t, m.String()).Equal("")

	m.Reset()
	assert.ThatNumber(m, 5).GreaterThan(10)
	assert.ThatString(t, m.String()).Equal(`Assertion failed: value not greater than expected:
    got: (int) 5
 expect: greater than (int) 10`)

	m.Reset()
	assert.ThatNumber(m, 5).GreaterThan(10, "index is 0")
	assert.ThatString(t, m.String()).Equal(`Assertion failed: value not greater than expected:
    got: (int) 5
 expect: greater than (int) 10
message: index is 0`)
}

func TestNumber_GreaterOrEqual(t *testing.T) {
	m := new(MockTestingT)
	assert.ThatNumber(m, 10).GreaterOrEqual(5)
	assert.ThatString(t, m.String()).Equal("")

	m.Reset()
	assert.ThatNumber(m, 5).GreaterOrEqual(10)
	assert.ThatString(t, m.String()).Equal(`Assertion failed: value not greater than or equal to expected:
    got: (int) 5
 expect: greater than or equal to (int) 10`)

	m.Reset()
	assert.ThatNumber(m, 5).GreaterOrEqual(10, "index is 0")
	assert.ThatString(t, m.String()).Equal(`Assertion failed: value not greater than or equal to expected:
    got: (int) 5
 expect: greater than or equal to (int) 10
message: index is 0`)
}

func TestNumber_LessThan(t *testing.T) {
	m := new(MockTestingT)
	assert.ThatNumber(m, 5).LessThan(10)
	assert.ThatString(t, m.String()).Equal("")

	m.Reset()
	assert.ThatNumber(m, 10).LessThan(5)
	assert.ThatString(t, m.String()).Equal(`Assertion failed: value not less than expected:
    got: (int) 10
 expect: less than (int) 5`)

	m.Reset()
	assert.ThatNumber(m, 10).LessThan(5, "index is 0")
	assert.ThatString(t, m.String()).Equal(`Assertion failed: value not less than expected:
    got: (int) 10
 expect: less than (int) 5
message: index is 0`)
}

func TestNumber_LessOrEqual(t *testing.T) {
	m := new(MockTestingT)
	assert.ThatNumber(m, 5).LessOrEqual(10)
	assert.ThatString(t, m.String()).Equal("")

	m.Reset()
	assert.ThatNumber(m, 10).LessOrEqual(5)
	assert.ThatString(t, m.String()).Equal(`Assertion failed: value not less than or equal to expected:
    got: (int) 10
 expect: less than or equal to (int) 5`)

	m.Reset()
	assert.ThatNumber(m, 10).LessOrEqual(5, "index is 0")
	assert.ThatString(t, m.String()).Equal(`Assertion failed: value not less than or equal to expected:
    got: (int) 10
 expect: less than or equal to (int) 5
message: index is 0`)
}

func TestNumber_IsZero(t *testing.T) {
	m := new(MockTestingT)
	assert.ThatNumber(m, 0).IsZero()
	assert.ThatString(t, m.String()).Equal("")

	m.Reset()
	assert.ThatNumber(m, 5).IsZero()
	assert.ThatString(t, m.String()).Equal(`Assertion failed: value is not zero:
    got: (int) 5
 expect: zero`)

	m.Reset()
	assert.ThatNumber(m, 5).IsZero("index is 0")
	assert.ThatString(t, m.String()).Equal(`Assertion failed: value is not zero:
    got: (int) 5
 expect: zero
message: index is 0`)
}

func TestNumber_IsNotZero(t *testing.T) {
	m := new(MockTestingT)
	assert.ThatNumber(m, 5).IsNotZero()
	assert.ThatString(t, m.String()).Equal("")

	m.Reset()
	assert.ThatNumber(m, 0).IsNotZero()
	assert.ThatString(t, m.String()).Equal(`Assertion failed: value is zero:
    got: (int) 0
 expect: non-zero`)

	m.Reset()
	assert.ThatNumber(m, 0).IsNotZero("index is 0")
	assert.ThatString(t, m.String()).Equal(`Assertion failed: value is zero:
    got: (int) 0
 expect: non-zero
message: index is 0`)
}

func TestNumber_IsPositive(t *testing.T) {
	m := new(MockTestingT)
	assert.ThatNumber(m, 5).IsPositive()
	assert.ThatString(t, m.String()).Equal("")

	m.Reset()
	assert.ThatNumber(m, -5).IsPositive()
	assert.ThatString(t, m.String()).Equal(`Assertion failed: value is not positive:
    got: (int) -5
 expect: positive`)

	m.Reset()
	assert.ThatNumber(m, -5).IsPositive("index is 0")
	assert.ThatString(t, m.String()).Equal(`Assertion failed: value is not positive:
    got: (int) -5
 expect: positive
message: index is 0`)
}

func TestNumber_IsNegative(t *testing.T) {
	m := new(MockTestingT)
	assert.ThatNumber(m, -5).IsNegative()
	assert.ThatString(t, m.String()).Equal("")

	m.Reset()
	assert.ThatNumber(m, 5).IsNegative()
	assert.ThatString(t, m.String()).Equal(`Assertion failed: value is not negative:
    got: (int) 5
 expect: negative`)

	m.Reset()
	assert.ThatNumber(m, 5).IsNegative("index is 0")
	assert.ThatString(t, m.String()).Equal(`Assertion failed: value is not negative:
    got: (int) 5
 expect: negative
message: index is 0`)
}

func TestNumber_IsNonNegative(t *testing.T) {
	m := new(MockTestingT)
	assert.ThatNumber(m, 5).IsNonNegative()
	assert.ThatString(t, m.String()).Equal("")

	m.Reset()
	assert.ThatNumber(m, -5).IsNonNegative()
	assert.ThatString(t, m.String()).Equal(`Assertion failed: value is negative:
    got: (int) -5
 expect: non-negative`)

	m.Reset()
	assert.ThatNumber(m, -5).IsNonNegative("index is 0")
	assert.ThatString(t, m.String()).Equal(`Assertion failed: value is negative:
    got: (int) -5
 expect: non-negative
message: index is 0`)
}

func TestNumber_IsNonPositive(t *testing.T) {
	m := new(MockTestingT)
	assert.ThatNumber(m, -5).IsNonPositive()
	assert.ThatString(t, m.String()).Equal("")

	m.Reset()
	assert.ThatNumber(m, 5).IsNonPositive()
	assert.ThatString(t, m.String()).Equal(`Assertion failed: value is positive:
    got: (int) 5
 expect: non-positive`)

	m.Reset()
	assert.ThatNumber(m, 5).IsNonPositive("index is 0")
	assert.ThatString(t, m.String()).Equal(`Assertion failed: value is positive:
    got: (int) 5
 expect: non-positive
message: index is 0`)
}

func TestNumber_IsBetween(t *testing.T) {
	m := new(MockTestingT)
	assert.ThatNumber(m, 5).IsBetween(1, 10)
	assert.ThatString(t, m.String()).Equal("")

	m.Reset()
	assert.ThatNumber(m, 0).IsBetween(1, 10)
	assert.ThatString(t, m.String()).Equal(`Assertion failed: value not within range:
    got: (int) 0
 expect: between (int) 1 and (int) 10`)

	m.Reset()
	assert.ThatNumber(m, 0).IsBetween(1, 10, "index is 0")
	assert.ThatString(t, m.String()).Equal(`Assertion failed: value not within range:
    got: (int) 0
 expect: between (int) 1 and (int) 10
message: index is 0`)
}

func TestNumber_IsNotBetween(t *testing.T) {
	m := new(MockTestingT)
	assert.ThatNumber(m, 0).IsNotBetween(1, 10)
	assert.ThatString(t, m.String()).Equal("")

	m.Reset()
	assert.ThatNumber(m, 5).IsNotBetween(1, 10)
	assert.ThatString(t, m.String()).Equal(`Assertion failed: value is within range:
    got: (int) 5
 expect: not between (int) 1 and (int) 10`)

	m.Reset()
	assert.ThatNumber(m, 5).IsNotBetween(1, 10, "index is 0")
	assert.ThatString(t, m.String()).Equal(`Assertion failed: value is within range:
    got: (int) 5
 expect: not between (int) 1 and (int) 10
message: index is 0`)
}

func TestNumber_IsInDelta(t *testing.T) {
	m := new(MockTestingT)
	assert.ThatNumber(m, 5.2).IsInDelta(5.0, 0.3)
	assert.ThatString(t, m.String()).Equal("")

	m.Reset()
	assert.ThatNumber(m, 5.6).IsInDelta(5.0, 0.3)
	assert.ThatString(t, m.String()).Equal(`Assertion failed: value not within delta:
    got: (float64) 5.6
 expect: within ±(float64) 0.3 of (float64) 5`)

	m.Reset()
	assert.ThatNumber(m, 5.6).IsInDelta(5.0, 0.3, "index is 0")
	assert.ThatString(t, m.String()).Equal(`Assertion failed: value not within delta:
    got: (float64) 5.6
 expect: within ±(float64) 0.3 of (float64) 5
message: index is 0`)
}

func TestNumber_IsNaN(t *testing.T) {
	m := new(MockTestingT)
	assert.ThatNumber(m, math.NaN()).IsNaN()
	assert.ThatString(t, m.String()).Equal("")

	m.Reset()
	assert.ThatNumber(m, 5.0).IsNaN()
	assert.ThatString(t, m.String()).Equal(`Assertion failed: value is not NaN:
    got: (float64) 5
 expect: NaN`)

	m.Reset()
	assert.ThatNumber(m, 5.0).IsNaN("index is 0")
	assert.ThatString(t, m.String()).Equal(`Assertion failed: value is not NaN:
    got: (float64) 5
 expect: NaN
message: index is 0`)
}

func TestNumber_IsInf(t *testing.T) {
	m := new(MockTestingT)
	assert.ThatNumber(m, math.Inf(1)).IsInf(1)
	assert.ThatString(t, m.String()).Equal("")

	m.Reset()
	assert.ThatNumber(m, math.Inf(-1)).IsInf(-1)
	assert.ThatString(t, m.String()).Equal("")

	m.Reset()
	assert.ThatNumber(m, 5.0).IsInf(1)
	assert.ThatString(t, m.String()).Equal(`Assertion failed: value is not infinite:
    got: (float64) 5
 expect: infinite with sign 1`)

	m.Reset()
	assert.ThatNumber(m, 5.0).IsInf(1, "index is 0")
	assert.ThatString(t, m.String()).Equal(`Assertion failed: value is not infinite:
    got: (float64) 5
 expect: infinite with sign 1
message: index is 0`)
}

func TestNumber_IsFinite(t *testing.T) {
	m := new(MockTestingT)
	assert.ThatNumber(m, 5.0).IsFinite()
	assert.ThatString(t, m.String()).Equal("")

	m.Reset()
	assert.ThatNumber(m, math.Inf(1)).IsFinite()
	assert.ThatString(t, m.String()).Equal(`Assertion failed: value is not finite:
    got: (float64) +Inf
 expect: finite`)

	m.Reset()
	assert.ThatNumber(m, math.Inf(-1)).IsFinite("index is 0")
	assert.ThatString(t, m.String()).Equal(`Assertion failed: value is not finite:
    got: (float64) -Inf
 expect: finite
message: index is 0`)
}
