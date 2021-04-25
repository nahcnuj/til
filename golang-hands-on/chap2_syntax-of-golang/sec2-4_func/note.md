# 2-4. 関数について

## 関数の基本

```
func 関数名(引数) 戻り値 {
    ...
}
```

外部に**公開する関数名は大文字で始める**

戻り値を複数返すこともできる。

```go
func f(n int) (int, string) {
    return n, strconv.Itoa(n)
}
```

### 名前付き戻り値

戻り値に名前をつけることもできる。
この場合、戻り値が 1 つでも括弧 `()` で囲む必要がある。

```go
func double(n int) (d int) {
    d = 2 * n
    return  // return は必須
}
```

### 可変長引数

型名の前に `...` を付けると、可変長の引数を受け取れる。

```go
func sum(a ...int) (s int) {    // a はスライスになる
    for _, n := range a {
        s += n
    }
    return
}
```

### 無名関数

Go では関数も値であり、変数に入れられる。

```go
func main() {
    f := func(s []string) (string, []string) {
        return s[0], s[1:]
    }

    s := []string{"one", "two", "three"}
    for len(s) > 0 {
        var v string
        v, s = f(s)
        fmt.Println("popped "+v+", remaining: ", s)
    }
}
```

実行結果は次のようになる。
```
poped one, remaining:  [two three]
poped two, remaining:  [three]
poped three, remaining:  []
```

### 高階関数

関数を引数に取る

```go
func main() {
    square := func(n int) int {
        return n*n
    }
    a := []int{1, 2, 3}
    fmt.Println(sum(a, square)) // => 14
}

func sum(a []int, f func(int) int) (s int) {
    for _, n := range a {
        s += f(n)
    }
    return
}
```

関数を返すこともできる

```go
func main() {
    hello := wrap()
    fmt.Println(hello())
}

func wrap() func() string {
    return func() string {
        return "hello"
    }
}
```

### クロージャ

```go
func main() {
    k := 2
    f := func(n int) int {
        return k * n
    }

    fmt.Println(f(3))   // => 6
    
    k = 3
    fmt.Println(f(3))   // => 9
}
```

### 無名関数の即時実行

```go
fmt.Println(func(n int) int {
    return 2 * n
}(3))   // => 6
```