# gs-assert

<div>
   <img src="https://img.shields.io/github/license/go-spring/gs-assert" alt="license"/>
   <img src="https://img.shields.io/github/go-mod/go-version/go-spring/gs-assert" alt="go-version"/>
   <img src="https://img.shields.io/github/v/release/go-spring/gs-assert?include_prereleases" alt="release"/>
   <a href="https://codecov.io/gh/go-spring/gs-assert" > 
      <img src="https://codecov.io/gh/go-spring/gs-assert/graph/badge.svg?token=SX7CV1T0O8" alt="test-coverage"/>
   </a>
   <a href="https://deepwiki.com/go-spring/gs-assert"><img src="https://deepwiki.com/badge.svg" alt="Ask DeepWiki"></a>
</div>

[English](README.md) | [中文](README_CN.md)

gs-assert 是一个用于 Go 单元测试的断言库，旨在提高测试代码的可读性和可维护性，
提供功能性和流畅性的断言风格，支持多种数据结构的断言，包括通用值、字符串、数字、切片、
映射和 `error`。

## 安装指南

要使用 `gs-assert` 库，首先确保你已经安装了 Go。然后，你可以通过以下命令安装 `gs-assert`：

```bash
go get github.com/go-spring/gs-assert
```

## 功能特性

- **功能性断言**：提供丰富的断言函数，如 `Nil`、`NotNil`、`Equal`、`NotEqual`等。
- **流畅性断言**：通过链式调用提高代码可读性，如 `assert.That(t).NotNil(obj).Equal(...)`。
- **数据结构支持**：支持对通用值、map、slice、string、number、`error` 等常见数据结构的断言。
- **`require` 模块**：提供更严格的断言方法，当断言失败时会立即终止测试。

### 支持的数据结构和断言方法

#### 通用值断言 (assert.Assertion)

通过 `assert.That(t, value)` 创建，支持以下方法：

- `True() / False()` - 断言布尔值
- `Nil() / NotNil()` - 断言值是否为 nil
- `Equal(expect) / NotEqual(expect)` - 断言值相等/不等
- `Same(expect) / NotSame(expect)` - 断言值相同/不同
- `TypeOf(expect)` - 断言类型兼容性
- `Implements(expect)` - 断言接口实现
- `Has(expect)` - 断言包含元素（调用 Has 方法）
- `Contains(expect)` - 断言包含元素（调用 Contains 方法）

#### 字符串断言 (assert.StringAssertion)

通过 `assert.ThatString(t, value)` 创建，支持以下方法：

- `Length(length)` - 断言字符串长度
- `Blank() / NonBlank()` - 断言字符串是否为空白
- `Equal(expect) / NotEqual(expect)` - 断言字符串相等/不等
- `EqualFold(expect)` - 断言字符串忽略大小写相等
- `JSONEqual(expect)` - 断言 JSON 字符串结构相等
- `Matches(pattern)` - 断言字符串匹配正则表达式
- `HasPrefix(prefix) / HasSuffix(suffix)` - 断言前缀/后缀
- `Contains(substr)` - 断言包含子字符串
- `IsLowerCase() / IsUpperCase()` - 断言大小写
- `IsNumeric() / IsAlpha() / IsAlphaNumeric()` - 断言字符类型
- `IsEmail() / IsURL() / IsIPv4() / IsHex() / IsBase64()` - 断言特定格式

#### 数字断言 (assert.NumberAssertion)

通过 `assert.ThatNumber(t, value)` 创建，支持以下方法：

- `Equal(expect) / NotEqual(expect)` - 断言数值相等/不等
- `GreaterThan(expect) / GreaterOrEqual(expect)` - 断言大于/大于等于
- `LessThan(expect) / LessOrEqual(expect)` - 断言小于/小于等于
- `Zero() / NonZero()` - 断言为零/非零
- `Positive() / NotPositive()` - 断言为正数/非正数
- `Negative() / NotNegative()` - 断言为负数/非负数
- `Between(lower, upper) / NotBetween(lower, upper)` - 断言在/不在范围内
- `InDelta(expect, delta)` - 断言在 delta 范围内
- `IsNaN() / IsInf(sign) / IsFinite()` - 断言特殊数值状态

#### 切片断言 (assert.SliceAssertion)

通过 `assert.ThatSlice(t, value)` 创建，支持以下方法：

- `Length(length)` - 断言切片长度
- `Nil() / NotNil()` - 断言切片为 nil/非 nil
- `Equal(expect) / NotEqual(expect)` - 断言切片相等/不等
- `Contains(element) / NotContains(element)` - 断言包含/不包含元素
- `ContainsSlice(sub) / NotContainsSlice(sub)` - 断言包含/不包含子切片
- `HasPrefix(prefix) / HasSuffix(suffix)` - 断言前缀/后缀
- `AllUnique()` - 断言所有元素唯一
- `AllMatches(fn) / AnyMatches(fn) / NoneMatches(fn)` - 断言元素匹配条件

#### 映射断言 (assert.MapAssertion)

通过 `assert.ThatMap(t, value)` 创建，支持以下方法：

- `Length(length)` - 断言映射长度
- `Nil() / NotNil()` - 断言映射为 nil/非 nil
- `Equal(expect) / NotEqual(expect)` - 断言映射相等/不等
- `ContainsKey(key) / NotContainsKey(key)` - 断言包含/不包含键
- `ContainsValue(value) / NotContainsValue(value)` - 断言包含/不包含值
- `ContainsKeyValue(key, value)` - 断言包含键值对
- `ContainsKeys(keys) / NotContainsKeys(keys)` - 断言包含/不包含多个键
- `ContainsValues(values) / NotContainsValues(values)` - 断言包含/不包含多个值
- `SubsetOf(expect) / SupersetOf(expect)` - 断言为子集/超集
- `HasSameKeys(expect) / HasSameValues(expect)` - 断言有相同键/值

#### 错误断言 (assert.ErrorAssertion)

通过 `assert.ThatError(t, value)` 创建，支持以下方法：

- `Nil() / NotNil()` - 断言错误为 nil/非 nil
- `Is(target) / NotIs(target)` - 断言错误匹配/不匹配目标错误
- `Matches(expr)` - 断言错误信息匹配正则表达式

#### Panic 断言

通过 `assert.Panic(t, fn, expr)` 断言函数会 panic 且 panic 信息匹配表达式。

## 使用示例

```go
import "github.com/go-spring/gs-assert/assert"
import "github.com/go-spring/gs-assert/require"

func TestAssertExample(t *testing.T) {
    // 使用 assert 模块进行断言，如果断言失败会继续执行测试
    assert.That(t, t).NotNil().NotSame(nil)
    assert.That(t, 1+1).Equal(2).NotEqual(3)
}

func TestRequireExample(t *testing.T) {
    // 使用 require 模块进行断言，如果断言失败会立即终止测试
    require.That(t, t).NotNil().NotSame(nil)
    require.That(t, 1+1).Equal(2).NotEqual(3)
}

func TestPanicExample(t *testing.T) {
    // 断言函数会 panic，并且 panic 信息匹配指定的表达式
    require.Panic(t, func() {
        panic("something went wrong")
    }, "something went wrong")
}
```

## 贡献指南

我们欢迎任何形式的贡献！如果你有任何建议或发现 bug，请提交 issue 或 pull request。在提交贡献之前，请确保你的代码符合项目的编码规范，并通过所有测试。

## 许可证

The Go-Spring is released under version 2.0 of the Apache License.