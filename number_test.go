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
	"github.com/go-spring/assert/internal"
)

func TestNumber_Equal(t *testing.T) {
	m := new(internal.MockTestingT)
	assert.ThatNumber(m, 5).Equal(5)
	assert.ThatString(t, m.String()).Equal("")

	m.Reset()
	assert.ThatNumber(m, 5).Equal(10)
	assert.ThatString(t, m.String()).Equal(`error# Assertion failed: expected number to be equal to 5, but got 10`)

	m.Reset()
	assert.ThatNumber(m, 5).Must().Equal(10, "index is 0")
	assert.ThatString(t, m.String()).Equal(`fatal# Assertion failed: expected number to be equal to 5, but got 10
message: "index is 0"`)
}

func TestNumber_NotEqual(t *testing.T) {
	m := new(internal.MockTestingT)
	assert.ThatNumber(m, 5).NotEqual(10)
	assert.ThatString(t, m.String()).Equal("")

	m.Reset()
	assert.ThatNumber(m, 5).NotEqual(5)
	assert.ThatString(t, m.String()).Equal(`error# Assertion failed: expected number not to be equal to 5, but it is`)

	m.Reset()
	assert.ThatNumber(m, 5).Must().NotEqual(5, "index is 0")
	assert.ThatString(t, m.String()).Equal(`fatal# Assertion failed: expected number not to be equal to 5, but it is
message: "index is 0"`)
}

func TestNumber_GreaterThan(t *testing.T) {
	m := new(internal.MockTestingT)
	assert.ThatNumber(m, 10).GreaterThan(5)
	assert.ThatString(t, m.String()).Equal("")

	m.Reset()
	assert.ThatNumber(m, 5).GreaterThan(10)
	assert.ThatString(t, m.String()).Equal(`error# Assertion failed: expected number to be greater than 5, but got 10`)

	m.Reset()
	assert.ThatNumber(m, 5).Must().GreaterThan(10, "index is 0")
	assert.ThatString(t, m.String()).Equal(`fatal# Assertion failed: expected number to be greater than 5, but got 10
message: "index is 0"`)
}

func TestNumber_GreaterOrEqual(t *testing.T) {
	m := new(internal.MockTestingT)
	assert.ThatNumber(m, 10).GreaterOrEqual(5)
	assert.ThatString(t, m.String()).Equal("")

	m.Reset()
	assert.ThatNumber(m, 5).GreaterOrEqual(10)
	assert.ThatString(t, m.String()).Equal(`error# Assertion failed: expected number to be greater than or equal to 5, but got 10`)

	m.Reset()
	assert.ThatNumber(m, 5).Must().GreaterOrEqual(10, "index is 0")
	assert.ThatString(t, m.String()).Equal(`fatal# Assertion failed: expected number to be greater than or equal to 5, but got 10
message: "index is 0"`)
}

func TestNumber_LessThan(t *testing.T) {
	m := new(internal.MockTestingT)
	assert.ThatNumber(m, 5).LessThan(10)
	assert.ThatString(t, m.String()).Equal("")

	m.Reset()
	assert.ThatNumber(m, 10).LessThan(5)
	assert.ThatString(t, m.String()).Equal(`error# Assertion failed: expected number to be less than 10, but got 5`)

	m.Reset()
	assert.ThatNumber(m, 10).Must().LessThan(5, "index is 0")
	assert.ThatString(t, m.String()).Equal(`fatal# Assertion failed: expected number to be less than 10, but got 5
message: "index is 0"`)
}

func TestNumber_LessOrEqual(t *testing.T) {
	m := new(internal.MockTestingT)
	assert.ThatNumber(m, 5).LessOrEqual(10)
	assert.ThatString(t, m.String()).Equal("")

	m.Reset()
	assert.ThatNumber(m, 10).LessOrEqual(5)
	assert.ThatString(t, m.String()).Equal(`error# Assertion failed: expected number to be less than or equal to 10, but got 5`)

	m.Reset()
	assert.ThatNumber(m, 10).Must().LessOrEqual(5, "index is 0")
	assert.ThatString(t, m.String()).Equal(`fatal# Assertion failed: expected number to be less than or equal to 10, but got 5
message: "index is 0"`)
}

func TestNumber_IsZero(t *testing.T) {
	m := new(internal.MockTestingT)
	assert.ThatNumber(m, 0).IsZero()
	assert.ThatString(t, m.String()).Equal("")

	m.Reset()
	assert.ThatNumber(m, 5).IsZero()
	assert.ThatString(t, m.String()).Equal(`error# Assertion failed: expected number to be zero, but got 5`)

	m.Reset()
	assert.ThatNumber(m, 5).Must().IsZero("index is 0")
	assert.ThatString(t, m.String()).Equal(`fatal# Assertion failed: expected number to be zero, but got 5
message: "index is 0"`)
}

func TestNumber_IsNotZero(t *testing.T) {
	m := new(internal.MockTestingT)
	assert.ThatNumber(m, 5).IsNotZero()
	assert.ThatString(t, m.String()).Equal("")

	m.Reset()
	assert.ThatNumber(m, 0).IsNotZero()
	assert.ThatString(t, m.String()).Equal(`error# Assertion failed: expected number not to be zero, but got 0`)

	m.Reset()
	assert.ThatNumber(m, 0).Must().IsNotZero("index is 0")
	assert.ThatString(t, m.String()).Equal(`fatal# Assertion failed: expected number not to be zero, but got 0
message: "index is 0"`)
}

func TestNumber_IsPositive(t *testing.T) {
	m := new(internal.MockTestingT)
	assert.ThatNumber(m, 5).IsPositive()
	assert.ThatString(t, m.String()).Equal("")

	m.Reset()
	assert.ThatNumber(m, -5).IsPositive()
	assert.ThatString(t, m.String()).Equal(`error# Assertion failed: expected number to be positive, but got -5`)

	m.Reset()
	assert.ThatNumber(m, -5).Must().IsPositive("index is 0")
	assert.ThatString(t, m.String()).Equal(`fatal# Assertion failed: expected number to be positive, but got -5
message: "index is 0"`)
}

func TestNumber_IsNegative(t *testing.T) {
	m := new(internal.MockTestingT)
	assert.ThatNumber(m, -5).IsNegative()
	assert.ThatString(t, m.String()).Equal("")

	m.Reset()
	assert.ThatNumber(m, 5).IsNegative()
	assert.ThatString(t, m.String()).Equal(`error# Assertion failed: expected number to be negative, but got 5`)

	m.Reset()
	assert.ThatNumber(m, 5).Must().IsNegative("index is 0")
	assert.ThatString(t, m.String()).Equal(`fatal# Assertion failed: expected number to be negative, but got 5
message: "index is 0"`)
}

func TestNumber_IsNonNegative(t *testing.T) {
	m := new(internal.MockTestingT)
	assert.ThatNumber(m, 5).IsNonNegative()
	assert.ThatString(t, m.String()).Equal("")

	m.Reset()
	assert.ThatNumber(m, -5).IsNonNegative()
	assert.ThatString(t, m.String()).Equal(`error# Assertion failed: expected number to be non-negative, but got -5`)

	m.Reset()
	assert.ThatNumber(m, -5).Must().IsNonNegative("index is 0")
	assert.ThatString(t, m.String()).Equal(`fatal# Assertion failed: expected number to be non-negative, but got -5
message: "index is 0"`)
}

func TestNumber_IsNonPositive(t *testing.T) {
	m := new(internal.MockTestingT)
	assert.ThatNumber(m, -5).IsNonPositive()
	assert.ThatString(t, m.String()).Equal("")

	m.Reset()
	assert.ThatNumber(m, 5).IsNonPositive()
	assert.ThatString(t, m.String()).Equal(`error# Assertion failed: expected number to be non-positive, but got 5`)

	m.Reset()
	assert.ThatNumber(m, 5).Must().IsNonPositive("index is 0")
	assert.ThatString(t, m.String()).Equal(`fatal# Assertion failed: expected number to be non-positive, but got 5
message: "index is 0"`)
}

func TestNumber_IsBetween(t *testing.T) {
	m := new(internal.MockTestingT)
	assert.ThatNumber(m, 5).IsBetween(1, 10)
	assert.ThatString(t, m.String()).Equal("")

	m.Reset()
	assert.ThatNumber(m, 0).IsBetween(1, 10)
	assert.ThatString(t, m.String()).Equal(`error# Assertion failed: expected number to be between 1 and 10, but got 0`)

	m.Reset()
	assert.ThatNumber(m, 0).Must().IsBetween(1, 10, "index is 0")
	assert.ThatString(t, m.String()).Equal(`fatal# Assertion failed: expected number to be between 1 and 10, but got 0
message: "index is 0"`)
}

func TestNumber_IsNotBetween(t *testing.T) {
	m := new(internal.MockTestingT)
	assert.ThatNumber(m, 0).IsNotBetween(1, 10)
	assert.ThatString(t, m.String()).Equal("")

	m.Reset()
	assert.ThatNumber(m, 5).IsNotBetween(1, 10)
	assert.ThatString(t, m.String()).Equal(`error# Assertion failed: expected number not to be between 1 and 10, but got 5`)

	m.Reset()
	assert.ThatNumber(m, 5).Must().IsNotBetween(1, 10, "index is 0")
	assert.ThatString(t, m.String()).Equal(`fatal# Assertion failed: expected number not to be between 1 and 10, but got 5
message: "index is 0"`)
}

func TestNumber_IsInDelta(t *testing.T) {
	m := new(internal.MockTestingT)
	assert.ThatNumber(m, 5.2).IsInDelta(5.0, 0.3)
	assert.ThatString(t, m.String()).Equal("")

	m.Reset()
	assert.ThatNumber(m, 5.6).IsInDelta(5.0, 0.3)
	assert.ThatString(t, m.String()).Equal(`error# Assertion failed: expected number to be within ±0.3 of 5, but got 5.6`)

	m.Reset()
	assert.ThatNumber(m, 5.6).Must().IsInDelta(5.0, 0.3, "index is 0")
	assert.ThatString(t, m.String()).Equal(`fatal# Assertion failed: expected number to be within ±0.3 of 5, but got 5.6
message: "index is 0"`)
}

func TestNumber_IsNaN(t *testing.T) {
	m := new(internal.MockTestingT)
	assert.ThatNumber(m, math.NaN()).IsNaN()
	assert.ThatString(t, m.String()).Equal("")

	m.Reset()
	assert.ThatNumber(m, 5.0).IsNaN()
	assert.ThatString(t, m.String()).Equal(`error# Assertion failed: expected number to be NaN, but got 5`)

	m.Reset()
	assert.ThatNumber(m, 5.0).Must().IsNaN("index is 0")
	assert.ThatString(t, m.String()).Equal(`fatal# Assertion failed: expected number to be NaN, but got 5
message: "index is 0"`)
}

func TestNumber_IsInf(t *testing.T) {
	m := new(internal.MockTestingT)
	assert.ThatNumber(m, math.Inf(1)).IsInf(1)
	assert.ThatString(t, m.String()).Equal("")

	m.Reset()
	assert.ThatNumber(m, math.Inf(-1)).IsInf(-1)
	assert.ThatString(t, m.String()).Equal("")

	m.Reset()
	assert.ThatNumber(m, 5.0).IsInf(1)
	assert.ThatString(t, m.String()).Equal(`error# Assertion failed: expected number to be +Inf, but got 5`)

	m.Reset()
	assert.ThatNumber(m, 5.0).Must().IsInf(1, "index is 0")
	assert.ThatString(t, m.String()).Equal(`fatal# Assertion failed: expected number to be +Inf, but got 5
message: "index is 0"`)
}

func TestNumber_IsFinite(t *testing.T) {
	m := new(internal.MockTestingT)
	assert.ThatNumber(m, 5.0).IsFinite()
	assert.ThatString(t, m.String()).Equal("")

	m.Reset()
	assert.ThatNumber(m, math.Inf(1)).IsFinite()
	assert.ThatString(t, m.String()).Equal(`error# Assertion failed: expected number to be finite, but got +Inf`)

	m.Reset()
	assert.ThatNumber(m, math.Inf(-1)).Must().IsFinite("index is 0")
	assert.ThatString(t, m.String()).Equal(`fatal# Assertion failed: expected number to be finite, but got -Inf
message: "index is 0"`)
}
