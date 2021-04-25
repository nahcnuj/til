# 3-2. 構造体

## 構造体の定義

```go
var me struct {
    Name string
    Age  int
}
me.Name = "jun"
me.Age = 25
fmt.Println(me) // => {jun 25}
```

## `type` で型として定義

```go
type person struct {
    Name string
    Age  int
}
```

構造体の値

```go
person{"Taro", 23}
```

```go
person{
    Name: "Hanako",
    Age:  17,
}
```

`}` を別行にする場合は、最後のフィールドの値の後ろに `,` が必要。
省くと `syntax error: unexpected newline, expecting comma or }` というエラーが出る。

構造体を関数に値渡しする場合、コピーコストがかかる。

## `new`

指定した型の値を生成し、そのポインタを返す関数。

```go
pp := new(person)
```

## `make`

配列、スライス、マップ、チャンネル（3-4 節の並行処理で登場）の値を作成し、初期化する関数。
スライスを生成する場合、長さと容量を指定することができる。

```go
s := make([]int, 3, 5)
fmt.Println(s)               // => [0 0 0]
fmt.Println(len(s), cap(s))  // => 3 5
```

## メソッド

関数名の前に型を指定すると、その型に関数を組み込める。この指定する型をレシーバ、型に組み込まれた関数をメソッドという。

```go
func (p person) printName() {
    fmt.Println(p.Name)
}

func (p *person) incrementAge() {
    p.Age++
}

func main() {
    someone = person{"Taro", 17}
    someone.printName()
    someone.incrementAge()
    fmt.Println("Happy birthday, "+someone.Name+" is now ", someone.Age)
}
```