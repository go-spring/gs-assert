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

package example

import (
	"testing"

	springAssert "github.com/go-spring/assert"
	"github.com/smartystreets/goconvey/convey"
	testifyAssert "github.com/stretchr/testify/assert"
)

func TestSpringAssert(t *testing.T) {
	springAssert.ThatMap(t, map[string]int{"a": 1}).Must().Equal(map[string]int{"a": 2}, "abc")
}

func TestTestifyAssert(t *testing.T) {
	testifyAssert.Equal(t, 1, 3, "abc")
}

func TestGoConvey(t *testing.T) {
	convey.Convey("Given some integer with a starting value", t, func() {
		x := 1
		convey.Convey("When the integer is incremented", func() {
			x++
			convey.Convey("The value should be greater by one", func() {
				convey.So(x, convey.ShouldEqual, 3)
			})
		})
	})
}
