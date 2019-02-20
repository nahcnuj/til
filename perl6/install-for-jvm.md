# Perl 6 コンパイラのインストール (JVM)

Windows Subsystem for Linux (WSL) に Perl 6 コンパイラの [Rakudo](https://rakudo.org/) をソースコードからコンパイルしてインストールする。
今回は Rakudo のバックエンドに **Java Virtual Machine (JVM)** を選択した（[MoarVM を選択する場合はこちら](./install-for-moar.md)）。

## インストール方法

インストール先は `$HOME/usr`（3 行目の `--prefix=$HOME/usr`）。

```bash
git clone git@github.com:rakudo/rakudo.git
cd rakudo
perl Configure.pl --gen-nqp --backends=jvm --prefix=$HOME/usr
make
make install
```

## 動作確認

```bash
$ perl6 -e 'say "Hello, Perl 6 World!"'
Hello, Perl 6 World!
```

## 所感

JVM の起動が遅い...。

```bash
$ time perl6 -e ''

real    0m9.235s
user    0m18.531s
sys     0m3.016s

$ time perl6 -v
This is Rakudo version 2018.12-311-gd656381 built on JVM
implementing Perl 6.d.

real    0m3.451s
user    0m8.063s
sys     0m1.328s
```