# 2-2. 制御構文

## `if`-`else`

```go
if n%2 == 0 {
    fmt.Println("even")
} else {
    fmt.Println("odd")
}
```

条件式にカッコは付けない。
`{}` は必須。
`else if` もある。

### ショートステートメント付き `if`

条件の前に 1 つだけ文を置ける。
その文で宣言された変数は `if`-`else` の中でのみ使える。

```go
x := input("type a number")
if n, err := strconv.Atoi(x); err == nil {
    if n%2 == 0 {
        fmt.Println("even")
    } else {
        fmt.Println("odd")
    }
} else {
    fmt.Println(err)
}
// `n` cannot be used here
```

## `switch`-`case`

```go
switch n {
case 1:
    fmt.Println("1")
case 2:
    fmt.Println("2")
default:
    fmt.Println("other")
}
```

デフォルトでは fallthrough しない。`fallthrough` を置くと次の `case` も実行させられる。

`if` と同様にショートステートメントも使える。

```go
switch {    // switch true { と同じ
case n%2 == 0:
    fmt.Println("even")
case n%2 == 1:
    fmt.Println("odd")
}
```

## `for`

```go
n := 0
for n < 5 {
    fmt.Println(n)
    n++
}
```

条件が真である間実行する。他のプログラミング言語の `while` 相当。

`n++` はインクリメント **文** である。式ではない。デクリメントも同様。

C 言語スタイルの `for` もある。

```go
for i := 1; i <= 5; i++ {
    fmt.Println(i)
}
```

無限ループも `for` で書く。

```go
for {   // for true { と同じ
    ...
}
```

残りの処理をスキップして次のループを始める `continue`、ループを抜ける `break` がある。

## `goto`

ある。
