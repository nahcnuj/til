# Perl 6 コンパイラのインストール (MoarVM)

Windows Subsystem for Linux (WSL) に Perl 6 コンパイラの [Rakudo](https://rakudo.org/) をソースコードからコンパイルしてインストールする。
今回は Rakudo のバックエンドに **MoarVM** を選択した（[Java Virtual Machine (JVM) を選択する場合はこちら](./install-for-jvm.md)）。

## インストール方法

以下のコマンドを実行する。インストール先は `$HOME/usr`（3 行目の `--prefix=$HOME/usr`）。

```bash
git clone git@github.com:rakudo/rakudo.git
cd rakudo
perl Configure.pl --gen-moar --gen-nqp --backends=moar --prefix=$HOME/usr
make
make install
```

## 動作確認

```bash
$ perl6 -e 'say "Hello, Perl 6 World!"'
Hello, Perl 6 World!
```

## 所感

JVM と比べて爆速ですごい。

```bash
$ time perl6 -e ''

real    0m0.370s
user    0m0.297s
sys     0m0.172s

$ time perl6 -v
This is Rakudo version 2018.12-311-gd656381 built on MoarVM version 2018.12-110-g04982f6
implementing Perl 6.d.

real    0m0.086s
user    0m0.000s
sys     0m0.047s
```