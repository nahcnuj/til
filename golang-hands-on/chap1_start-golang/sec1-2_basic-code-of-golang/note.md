# 1-2. Go の基本コード

## Go のソースコード

- メインプログラム（エントリポイントとなるプログラム）は `main` パッケージ
- `import ( ... )` でパッケージ読み込み
    - 複数パッケージを読み込む場合は改行区切り
- エントリポイントは `main` パッケージの `main` 関数
    - `func main() { ... }`
    - 引数も戻り値もない

## Go の書き方

- case sensitive
- インデントは **タブ** で
    - ※本リポジトリの `note.md` においては、スペース 4 つになっている。 ~~タブに置き換えるのが面倒くさい~~
- 名前は、公開されるものは大文字で、それ以外は小文字で始める
- 演算子の前後にスペースを入れない
    - 優先順位を明確にするためにスペースを入れる

## GOPATH

```console
$ echo $GOPATH
/home/user/go/1.16.3
```

anyenv と goenv 使ってると少し様子が違うっぽい
```console
$ which go
/home/user/.anyenv/envs/goenv/shims/go
$ ls ~/.anyenv/envs/goenv/versions/1.16.3/
AUTHORS  CONTRIBUTING.md  CONTRIBUTORS  LICENSE  PATENTS  README.md  SECURITY.md  VERSION  api  bin  doc  favicon.ico  lib  misc  pkg  robots.txt  src  test
```