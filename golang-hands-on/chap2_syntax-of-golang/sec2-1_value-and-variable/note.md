# 2-1. 値と変数

## 型

- 整数
    - `int`: `int32` or `int64`
    - `int8`: 8 bits
    - `int16`: 16 bits
    - `int32`: 32 bits
    - `int64`: 64 bits
- 符号なし整数
    - `uint`: `uint32` or `uint64`
    - `uint8`: 8 bits
    - `uint16`: 16 bits
    - `uint32`: 32 bits
    - `uint64`: 64 bits
- ポインタ用符号なし整数 `uintptr`
- バイト `byte`（`uint8` の別名）
- Unicode のコードポイント `rune`（`int32` の別名）
- 浮動小数点数
    - `float32`: 32 bits
    - `float64`: 64 bits
- 複素数
    - `complex64`: 実部と虚部を `float32`
    - `complex128`: 実部と虚部を `float64`
- 真偽値 `bool`
- 文字列 `string`
    - 1 文字だけの値を表すには `rune` 型

## リテラル

- 整数（`int`）
    - 10 進数 `123`
    - 8 進数 `0o123`
    - 16 進数 `0x123`
    - 2 進数 `0b1111`
- 実数（`float64`）
    - `3.141592`
    - `6.02e23` = 6.02 × 10²³
- 文字（`rune`）
    - `'a'`
- 文字列（`string`）
    - `"hello"`
- 改行を含む文字列（`string`）
    ```go
    `hoge
    fuga`
    ```
- 真偽値（`bool`）
    - `true`
    - `false`

> \`〇〇\`というように前後を **「`'`」** でくくると、その間のテキストは途中で改行しても一つの値として認識されます。

意味がわからない...（p.55）

## 基本的な演算

- 数値
    - `+`,`-`,`*`,`/`,`%`
- 文字列
    - `+` で結合

`bool` 等は算術演算不可。
`rune` は `int32` なので算術演算できるが、結果は整数値（`int32`）になる。

## 変数

```go
var a int
var x, y, z float32

a = 42
x, y, z = 0.1, 1.2, 2.3

var pi = 3.141592
e := 2.71828
```

lexical scope

Go では変数名は短いほうが良い

## 型変換

```go
var x int32 = 100
var y = int64(x)
var z = float32(y)
```

Go に暗黙の型変換はない。

## 文字列⇆数値

`strconv` パッケージの `Atoi` 関数と `Iota` 関数

```go
n, err := strconv.Atoi("123")

s, err := strconv.Iota(123)
```

Go では宣言された変数に未使用のものがあるとコンパイルエラーになる。

## 定数

```go
const n int = 42
const pi = 3.141592
```

暗黙の型変換がないため、以下はコンパイルエラーになる。

```go
const n int = 42
x := n * 1.1    // Error!
```

型を明示せずに定義すると、型が使用時に推論されてエラーにならない。

```go
const n = 42
x := n * 1.1    // OK
```

## factored `import` statement

```go
import (
    "fmt"
    "strconv"
)
```
のような書き方のこと。

分けて書くこともできる。

```go
import "fmt"
import "strconv"
```
