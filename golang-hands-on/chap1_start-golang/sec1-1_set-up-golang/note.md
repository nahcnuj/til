# 1-1. Go のセットアップ

## `go run` で実行

```console
$ go run hello.go
Hello, Go-lang!
```

## `go build` でビルド（バイナリファイルの生成）

```console
$ ls
hello.go  note.md
$ go build hello.go
$ ls
hello  hello.go  note.md
$ ./hello
Hello, Go-lang!
```

`go run` も内部でビルドして実行している。
Go はインタプリタ言語ではない。