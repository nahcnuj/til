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

JVM だとめっちゃ遅いですね...。
ちなみに、JVM の起動が遅いのではなく、ソースコードのパースに時間が掛かっているようでした。

```bash
$ perl6-j --stagestats -e ''
Stage start      :   0.002
Stage parse      :   6.021
Stage syntaxcheck:   0.001
Stage ast        :   0.001
Stage optimize   :   0.071
Stage jast       :   0.212
Stage classfile  :   0.035
Stage jar        :   0.000
Stage jvm        :   0.013
```
