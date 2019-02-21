# Perl 6 の実行のボトルネックはどこか

`--stagestats` オプションを付けて、JVM と MoarVM でそれぞれ `perl6` を実行し、ステージ毎の実行時間を見てみました。

```bash
# on JVM
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

# on MoarVM
$ perl6-m --stagestats -e ''
Stage start      :   0.000
Stage parse      :   0.283
Stage syntaxcheck:   0.000
Stage ast        :   0.000
Stage optimize   :   0.002
Stage mast       :   0.013
Stage mbc        :   0.001
Stage moar       :   0.000
```

どちらもソースコードのパースに最も時間が掛かっていました。
パースの段階では MoarVM も JVM も処理は変わらないはずなのにここまでの違いになるのはなぜなんでしょうか...。
