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

	m.Reset()
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

	m.Reset()
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

	m.Reset()
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

	m.Reset()
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

	m.Reset()
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

	m.Reset()
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

func TestNumber_Zero(t *testing.T) {
	m := new(internal.MockTestingT)

	m.Reset()
	assert.ThatNumber(m, 0).Zero()
	assert.ThatString(t, m.String()).Equal("")

	m.Reset()
	assert.ThatNumber(m, 5).Zero()
	assert.ThatString(t, m.String()).Equal(`error# Assertion failed: expected number to be zero, but got 5`)

	m.Reset()
	assert.ThatNumber(m, 5).Must().Zero("index is 0")
	assert.ThatString(t, m.String()).Equal(`fatal# Assertion failed: expected number to be zero, but got 5
 message: "index is 0"`)
}

func TestNumber_NotZero(t *testing.T) {
	m := new(internal.MockTestingT)

	m.Reset()
	assert.ThatNumber(m, 5).NotZero()
	assert.ThatString(t, m.String()).Equal("")

	m.Reset()
	assert.ThatNumber(m, 0).NotZero()
	assert.ThatString(t, m.String()).Equal(`error# Assertion failed: expected number not to be zero, but got 0`)

	m.Reset()
	assert.ThatNumber(m, 0).Must().NotZero("index is 0")
	assert.ThatString(t, m.String()).Equal(`fatal# Assertion failed: expected number not to be zero, but got 0
 message: "index is 0"`)
}

func TestNumber_Positive(t *testing.T) {
	m := new(internal.MockTestingT)

	m.Reset()
	assert.ThatNumber(m, 5).Positive()
	assert.ThatString(t, m.String()).Equal("")

	m.Reset()
	assert.ThatNumber(m, -5).Positive()
	assert.ThatString(t, m.String()).Equal(`error# Assertion failed: expected number to be positive, but got -5`)

	m.Reset()
	assert.ThatNumber(m, -5).Must().Positive("index is 0")
	assert.ThatString(t, m.String()).Equal(`fatal# Assertion failed: expected number to be positive, but got -5
 message: "index is 0"`)
}

func TestNumber_NotPositive(t *testing.T) {
	m := new(internal.MockTestingT)

	m.Reset()
	assert.ThatNumber(m, -5).NotPositive()
	assert.ThatString(t, m.String()).Equal("")

	m.Reset()
	assert.ThatNumber(m, 5).NotPositive()
	assert.ThatString(t, m.String()).Equal(`error# Assertion failed: expected number to be non-positive, but got 5`)

	m.Reset()
	assert.ThatNumber(m, 5).Must().NotPositive("index is 0")
	assert.ThatString(t, m.String()).Equal(`fatal# Assertion failed: expected number to be non-positive, but got 5
 message: "index is 0"`)
}

func TestNumber_Negative(t *testing.T) {
	m := new(internal.MockTestingT)

	m.Reset()
	assert.ThatNumber(m, -5).Negative()
	assert.ThatString(t, m.String()).Equal("")

	m.Reset()
	assert.ThatNumber(m, 5).Negative()
	assert.ThatString(t, m.String()).Equal(`error# Assertion failed: expected number to be negative, but got 5`)

	m.Reset()
	assert.ThatNumber(m, 5).Must().Negative("index is 0")
	assert.ThatString(t, m.String()).Equal(`fatal# Assertion failed: expected number to be negative, but got 5
 message: "index is 0"`)
}

func TestNumber_NotNegative(t *testing.T) {
	m := new(internal.MockTestingT)

	m.Reset()
	assert.ThatNumber(m, 5).NotNegative()
	assert.ThatString(t, m.String()).Equal("")

	m.Reset()
	assert.ThatNumber(m, -5).NotNegative()
	assert.ThatString(t, m.String()).Equal(`error# Assertion failed: expected number to be non-negative, but got -5`)

	m.Reset()
	assert.ThatNumber(m, -5).Must().NotNegative("index is 0")
	assert.ThatString(t, m.String()).Equal(`fatal# Assertion failed: expected number to be non-negative, but got -5
 message: "index is 0"`)
}

func TestNumber_Between(t *testing.T) {
	m := new(internal.MockTestingT)

	m.Reset()
	assert.ThatNumber(m, 5).Between(1, 10)
	assert.ThatString(t, m.String()).Equal("")

	m.Reset()
	assert.ThatNumber(m, 0).Between(1, 10)
	assert.ThatString(t, m.String()).Equal(`error# Assertion failed: expected number to be between 1 and 10, but got 0`)

	m.Reset()
	assert.ThatNumber(m, 0).Must().Between(1, 10, "index is 0")
	assert.ThatString(t, m.String()).Equal(`fatal# Assertion failed: expected number to be between 1 and 10, but got 0
 message: "index is 0"`)
}

func TestNumber_NotBetween(t *testing.T) {
	m := new(internal.MockTestingT)

	m.Reset()
	assert.ThatNumber(m, 0).NotBetween(1, 10)
	assert.ThatString(t, m.String()).Equal("")

	m.Reset()
	assert.ThatNumber(m, 5).NotBetween(1, 10)
	assert.ThatString(t, m.String()).Equal(`error# Assertion failed: expected number not to be between 1 and 10, but got 5`)

	m.Reset()
	assert.ThatNumber(m, 5).Must().NotBetween(1, 10, "index is 0")
	assert.ThatString(t, m.String()).Equal(`fatal# Assertion failed: expected number not to be between 1 and 10, but got 5
 message: "index is 0"`)
}

func TestNumber_InDelta(t *testing.T) {
	m := new(internal.MockTestingT)

	m.Reset()
	assert.ThatNumber(m, 5.2).InDelta(5.0, 0.3)
	assert.ThatString(t, m.String()).Equal("")

	assert.ThatNumber(m, 5.2).InDelta(5.5, 0.3)
	assert.ThatString(t, m.String()).Equal("")

	m.Reset()
	assert.ThatNumber(m, 5.6).InDelta(5.0, 0.3)
	assert.ThatString(t, m.String()).Equal(`error# Assertion failed: expected number to be within ±0.3 of 5, but got 5.6`)

	m.Reset()
	assert.ThatNumber(m, 5.6).Must().InDelta(5.0, 0.3, "index is 0")
	assert.ThatString(t, m.String()).Equal(`fatal# Assertion failed: expected number to be within ±0.3 of 5, but got 5.6
 message: "index is 0"`)
}

func TestNumber_IsNaN(t *testing.T) {
	m := new(internal.MockTestingT)

	m.Reset()
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

	m.Reset()
	assert.ThatNumber(m, math.Inf(1)).IsInf(1)
	assert.ThatString(t, m.String()).Equal("")

	m.Reset()
	assert.ThatNumber(m, math.Inf(-1)).IsInf(-1)
	assert.ThatString(t, m.String()).Equal("")

	m.Reset()
	assert.ThatNumber(m, 5.0).IsInf(1)
	assert.ThatString(t, m.String()).Equal(`error# Assertion failed: expected number to be +Inf, but got 5`)

	m.Reset()
	assert.ThatNumber(m, 5.0).Must().IsInf(-1, "index is 0")
	assert.ThatString(t, m.String()).Equal(`fatal# Assertion failed: expected number to be -Inf, but got 5
 message: "index is 0"`)
}

func TestNumber_IsFinite(t *testing.T) {
	m := new(internal.MockTestingT)

	m.Reset()
	assert.ThatNumber(m, 5.0).IsFinite()
	assert.ThatString(t, m.String()).Equal("")

	m.Reset()
	assert.ThatNumber(m, float32(5.0)).IsFinite()
	assert.ThatString(t, m.String()).Equal("")

	m.Reset()
	assert.ThatNumber(m, int64(5)).IsFinite()
	assert.ThatString(t, m.String()).Equal("")

	m.Reset()
	assert.ThatNumber(m, math.Inf(1)).IsFinite()
	assert.ThatString(t, m.String()).Equal(`error# Assertion failed: expected number to be finite, but got +Inf`)

	m.Reset()
	assert.ThatNumber(m, math.Inf(-1)).Must().IsFinite("index is 0")
	assert.ThatString(t, m.String()).Equal(`fatal# Assertion failed: expected number to be finite, but got -Inf
 message: "index is 0"`)
}
