# assert

<div>
   <img src="https://img.shields.io/github/license/go-spring/assert" alt="license"/>
   <img src="https://img.shields.io/github/go-mod/go-version/go-spring/assert" alt="go-version"/>
   <img src="https://img.shields.io/github/v/release/go-spring/assert?include_prereleases" alt="release"/>
   <a href="https://codecov.io/gh/go-spring/assert" > 
      <img src="https://codecov.io/gh/go-spring/assert/graph/badge.svg?token=SX7CV1T0O8" alt="test-coverage"/>
   </a>
   <a href="https://deepwiki.com/go-spring/assert"><img src="https://deepwiki.com/badge.svg" alt="Ask DeepWiki"></a>
</div>

[English](README.md) | [中文](README_CN.md)

Go-Spring::assert is an assertion library for Go unit tests,
designed to enhance test code readability and maintainability.
It provides functional and fluent assertion styles and
supports assertions for multiple data structures,
including generic values, strings, numbers, slices, maps, and `error`.

## Installation Guide

To use the assert library, ensure you have Go installed.
Then, install assert using the following command:

```bash
go get github.com/go-spring/assert
```

## Features

- **Functional Assertions**：
  Offers a rich set of assertion functions such as `Nil`, `NotNil`, `Equal`, `NotEqual`, etc.
- **Fluent Assertions**：
  Improves code readability through method chaining, e.g., `assert.That(t).NotNil(obj).Equal(...)`.
- **Data Structure Support**：
  Supports assertions for common data structures like generic values, maps, slices, strings, numbers, and `error`.
- **`require` Module**：
  Provides stricter assertion methods that immediately terminate the test when an assertion fails.

### Supported Data Structures and Assertion Methods

#### Generic Value Assertions (assert.Assertion)

Created via `assert.That(t, value)`, supports the following methods:

- `True() / False()` - Assert boolean values.
- `Nil() / NotNil()` - Assert whether a value is nil or not.
- `Equal(expect) / NotEqual(expect)` - Assert equality or inequality.
- `Same(expect) / NotSame(expect)` - Assert identity or non-identity.
- `TypeOf(expect)` - Assert type compatibility.
- `Implements(expect)` - Assert interface implementation.
- `Has(expect)` - Assert element inclusion（via Has method）
- `Contains(expect)` - Assert element inclusion（via Contains method）

#### String Assertions (assert.StringAssertion)

Created via `assert.ThatString(t, value)`, supports the following methods:

- `Length(length)` - Assert string length.
- `Blank() / NonBlank()` - Assert whether the string is blank or not.
- `Equal(expect) / NotEqual(expect)` - Assert string equality or inequality.
- `EqualFold(expect)` - Assert case-insensitive equality.
- `JSONEqual(expect)` - Assert JSON structure equality.
- `Matches(pattern)` - Assert regex match.
- `HasPrefix(prefix) / HasSuffix(suffix)` - Assert prefix/suffix.
- `Contains(substr)` - Assert substring inclusion.
- `IsLowerCase() / IsUpperCase()` - Assert case status.
- `IsNumeric() / IsAlpha() / IsAlphaNumeric()` - Assert character type.
- `IsEmail() / IsURL() / IsIPv4() / IsHex() / IsBase64()` - Assert specific formats.

#### Number Assertions (assert.NumberAssertion)

Created via `assert.ThatNumber(t, value)`, supports the following methods:

- `Equal(expect) / NotEqual(expect)` - Assert numeric equality or inequality.
- `GreaterThan(expect) / GreaterOrEqual(expect)` - Assert greater than or equal.
- `LessThan(expect) / LessOrEqual(expect)` - Assert less than or equal.
- `Zero() / NonZero()` - Assert zero or non-zero.
- `Positive() / NotPositive()` - Assert positive or non-positive.
- `Negative() / NotNegative()` - Assert negative or non-negative.
- `Between(lower, upper) / NotBetween(lower, upper)` - Assert within or outside a range.
- `InDelta(expect, delta)` - Assert within a delta range.
- `IsNaN() / IsInf(sign) / IsFinite()` - Assert special numeric states.

#### Slice Assertions (assert.SliceAssertion)

Created via `assert.ThatSlice(t, value)`, supports the following methods:

- `Length(length)` - Assert slice length.
- `Nil() / NotNil()` - Assert whether the slice is nil or not.
- `Equal(expect) / NotEqual(expect)` - Assert slice equality or inequality.
- `Contains(element) / NotContains(element)` - Assert element inclusion/exclusion.
- `ContainsSlice(sub) / NotContainsSlice(sub)` - Assert sub-slice inclusion/exclusion.
- `HasPrefix(prefix) / HasSuffix(suffix)` - Assert prefix/suffix.
- `AllUnique()` - Assert all elements are unique.
- `AllMatches(fn) / AnyMatches(fn) / NoneMatches(fn)` - Assert element conditions.

#### Map Assertions (assert.MapAssertion)

Created via `assert.ThatMap(t, value)`, supports the following methods:

- `Length(length)` - Assert map length.
- `Nil() / NotNil()` - Assert whether the map is nil or not.
- `Equal(expect) / NotEqual(expect)` - Assert map equality or inequality.
- `ContainsKey(key) / NotContainsKey(key)` - Assert key inclusion/exclusion.
- `ContainsValue(value) / NotContainsValue(value)` - Assert value inclusion/exclusion.
- `ContainsKeyValue(key, value)` - Assert key-value pair inclusion.
- `ContainsKeys(keys) / NotContainsKeys(keys)` - Assert multiple key inclusions/exclusions.
- `ContainsValues(values) / NotContainsValues(values)` - Assert multiple value inclusions/exclusions.
- `SubsetOf(expect) / SupersetOf(expect)` - Assert subset/superset.
- `HasSameKeys(expect) / HasSameValues(expect)` - Assert same keys/values.

#### Error Assertions (assert.ErrorAssertion)

Created via `assert.ThatError(t, value)`, supports the following methods:

- `Nil() / NotNil()` - Assert error is nil or not.
- `Is(target) / NotIs(target)` - Assert error matches or does not match a target.
- `Matches(expr)` - Assert error message matches regex.

#### Panic Assertions

Use `assert.Panic(t, fn, expr)` to assert a function panics and
the panic message matches an expression.

## Usage Examples

```go
func TestAssertExample(t *testing.T) {
    // Use assert module for assertions, test continues on failure
    assert.That(t, t).NotNil().NotSame(nil)
    assert.That(t, 1+1).Equal(2).NotEqual(3)
}

func TestRequireExample(t *testing.T) {
    // Use require module for assertions, test stops on failure
    require.That(t, t).NotNil().NotSame(nil)
    require.That(t, 1+1).Equal(2).NotEqual(3)
}

func TestPanicExample(t *testing.T) {
    // Assert function panics and message matches the expression
    require.Panic(t, func () {
        panic("something went wrong")
    }, "something went wrong")
}
```

## Contribution Guide

We welcome any form of contribution!
If you have suggestions or find bugs, please submit an issue or pull request.
Before submitting contributions, ensure your code adheres to
the project's coding standards and passes all tests.

## License

The Go-Spring is released under version 2.0 of the Apache License.