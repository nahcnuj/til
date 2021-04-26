# 2-3. 複雑な値

## 配列

Go の配列は固定長。
0-indexed。

```go
// int 型の 3 要素の配列 a
var a [3]int

// 宣言と同時に初期化
var b [3]int = [3]int{1, 2, 3}

// 使い方
fmt.Println(b[1])   // => 2
b[0] = 42
```

> ```go
> var 変数 [ 要素数 ] 型 { 値1, 値2, …… }
> ```

じゃないじゃん（p.85）

```go
var 変数 = [ 要素数 ] 型 { 値1, 値2, …… }
```
なら通る。

配列 `a` に対して `len(a)` で要素数を得られる。

### `range`

`range` を使うことで、インデックスと値を得られる。
```go
a := [...]int{4, 5, 6}
for i, x := range a {
    fmt.Println(i, x)
}
```
とすると、次のような出力になる。
```
0 4
1 5
2 6
```

インデックスが要らない場合はアンダースコア `_` にして捨てる。
```go
for _, x := range a {
    fmt.Println(x)
}
```

> Go では、アンダースコアで始まる変数は、他の変数とは少し違います。これは、 **「使わなくてもいい変数」** なのです。

エラーになったが？（p.90）

```console
$ go run for-range-2.go
# command-line-arguments
./for-range-2.go:7:6: _i declared but not used
```

## スライス

動的な配列

配列の宣言で、要素数を省略するとスライスになる。
```go
slice := []int {1,2,3}
```

配列 `a` に対して `a[i:j]` で `a[i]` から `a[j-1]` までの要素のスライスができる。
```go
a := [5]int {1,2,3,4,5}
b := a[1:3] // => [2, 3]
```

`b` がスライス。

`a[i:]` で `a[i]` から末尾までの、`a[:j]` で先頭から `a[j-1]` までのスライスになる。
`a[:0]` は空のスライスになる。

スライスは元の配列の参照

```go
a := [5]int {1,2,3,4,5}
b := a[1:3]     // => [2, 3]

a[2] = 5
fmt.Println(a)  // => [1 2 5 4 5]
fmt.Println(b)  // => [2 5]

b[0] = 0
fmt.Println(a)  // => [1 0 5 4 5]
fmt.Println(b)  // => [0 5]
```

### 長さと容量

- スライス `b` の **長さ** `len(b)`
    - スライスが持つ要素数
- スライス `b` の **容量** `cap(b)`
    - スライスが参照する配列の長さ
    - 配列の途中からスライスを作ると、スライスの先頭から配列の末尾までの長さ

```go
a := [5]int {1,2,3,4,5}
b := a[1:3]

fmt.Println(len(b)) // => 2
fmt.Println(cap(b)) // => 4
```

### スライスに対する操作

スライスの末尾に値を追加する `append` 関数がある。
むしろ、それ以外は用意されていない。
代わりに、`[i:j]` を駆使して実現する。

```go
a := [5]int {1,2,3,4,5}
b := a[1:3]
fmt.Println(b)  // => [2 3]

b = append(b, 6, 7)
fmt.Println(b)  // => [2 3 6 7]
fmt.Println(a)  // => [1 2 3 6 7]
```

`append` によって、配列から作ったスライスが元の配列の要素数を越えたりすると、配列の参照ではなくなるようだ。

```go
a := [5]int{1, 2, 3, 4, 5}
b := a[1:3]

fmt.Println(b) // => [2 3]
fmt.Println(len(b)) // => 2
fmt.Println(cap(b)) // => 4

b = append(b, 6, 7, 8)
fmt.Println(b) // => [2 3 6 7 8]
fmt.Println(len(b)) // => 5
fmt.Println(cap(b)) // => 8
fmt.Println(a) // => [1 2 3 4 5]
```

#### push, pop, shift, unshift, insert, remove を作る

`[]int` 型のスライスに対して各種操作を行う関数を、以下のような仕様で作ってみる。
戻り値はいずれも操作後のスライス。

- `push(s, v)`: スライスの末尾に 1 つ値を追加する
- `pop(s)`: スライスの末尾の値を削除する
- `uhshift(s, v)`: スライスの先頭に 1 つ値を追加する
- `shift(s)`: スライスの先頭の値を 1 つ削除する
- `insert(s, v, i)`: スライスの指定した位置に 1 つ値を追加する
- `remove(s, i)`: スライスの指定した位置の値を削除する

[`./slice.go`](./slice.go)

##### `push`

スライス `s` に値 `v` を `append` すれば良い。

##### `pop`

スライス `s` の先頭から末尾の一つ前までのスライスを取れば良い。

##### `unshift`

追加する値 `v` だけを持つスライスを作り、与えられたスライス `s` のすべての要素を `append` すれば良い。
スライス `s` に対して `s...` とすると、`s[0], s[1], ..., s[len(s)]` というように展開される。

```go
append([]int{v}, s...)
```
は
```
append([]int{v}, s[0], s[1], ..., s[len(s)])
```
と同様の意味になる。

##### `shift`

スライス `s` の（0-indexed で）1 番目から末尾までのスライスを取れば良い。

##### `insert`

追加する値 `v` だけを持つスライスに、`s` の `i` 番目から末尾までの値を `append` したものを、
スライス `s` の先頭から（0-indexed で）`i-1` 番目までのスライスに `append` すれば良い。

下記のようには**書けない**ようだ。
`too many arguments to append` と言われてしまう。
```
append(s[:i], v, s[i:]...)    // error!
```

##### `remove`

`s[:i]` と `s[i+1:]` を `append` で結合すれば良い。

## マップ `map`

`string` 型のキーに対して `int` 型の値を持つマップは次のように宣言する。
```go
var m map[string] int
```

```go
m := map[string] int {
    "key": 42,
    "a":   334,
}
m["sum"] = m["key"] + m["a"]
fmt.Println(m)  // => map[a:334 key:42 sum:376]

delete(m, "a")  // append と異なり、引数のマップ自体が変更される
fmt.Println(m)  // => map[key:42 sum:376]

delete(m, "b")  // エラーにならない

for k, v := range m {
    fmt.Println(k + ": ", v)  // => key:  42
}                             //    sum:  376
```

マップのキー順序は不定。