# Getting Started with DynamoDB

https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/GettingStartedDynamoDB.html

## Basic Concepts

[Core Components of Amazon DynamoDB \- Amazon DynamoDB](https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/HowItWorks.CoreComponents.html)

### Tables, Items, and Attributes

- **テーブル (table)**: データの集合。他の DB と同じ。テーブルは 0 個以上の項目 (items) を含む。
- **項目 (item)**: 一意に識別可能な、属性 (attributes) の集合。他の DB の行 (rows) やレコード (records) に当たる。 1 つの項目は 1 つ以上の属性から成る。 DynamoDB では 1 つのテーブルに保管できる項目の数に制限はない。
    - 各項目は一意な識別子（**プライマリキー**、primary key）を持つ。
    - プライマリキーを除いては**スキーマレス**である。
- **属性 (attribute)**: これ以上分解する必要のない、基本的なデータ要素。他の DB の列 (columns) やフィールド (fields) に当たる。
    - 殆どの属性は**スカラー** (scalar) であり、単一の値を持つ。文字列や数値など。
    - 属性は**ネストできる**。 DynamoDB では 32 段まで属性をネストさせられる。
    - リスト(?)の属性もある。

### Primary Key

プライマリキーには次の 2 種類がある。

- **パーティションキー** (partition key)
- **パーティションキーとソートキー** (sort key)、複合プライマリキー

パーティションキーは、内部でハッシュ関数の入力として使われる。そのハッシュ関数の出力によって、その項目が（DynamoDB の内部物理ストレージの）どのパーティションに保管されるかが決まる。

複合プライマリキーの場合は、同じパーティションキーを持つ項目を複数持つことができる。それらの項目はソートキーの値で並べられる。
同じパーティションキーを持つ項目のソートキーは、それらの中で一意でなければならない。

プライマリキー属性に限り、スカラー（文字列 (string)、数値 (number)、バイナリ (binary) のいずれか）でなければならない。

プライマリキーはテーブル作成時に指定する。

### Secondary Indexes

プライマリキー以外のキーでクエリするために、**セカンダリインデックス**を設定できる。
インデックスはテーブルに所属し、インデックスが所属するテーブルは（そのインデックスの）**ベーステーブル** (base table) と呼ばれる。

インデックスには次の 2 種類がある。

- **グローバルセカンダリインデックス** (global secondary index; GSI): テーブルのプライマリキーとは異なり得るパーティションキーとソートキーによるインデックス
- **ローカルセカンダリインデックス** (local secondary index; LSI): テーブルのプライマリキーと同じパーティションキーと、テーブルのプライマリキーとは異なるソートキーによるインデックス

DynamoDB では、 1 つのテーブルにつき 20 個（デフォルトクォータ）の GSI と 5 つの LSI を持てる。

インデックスの作成時に、インデックスの属性をベーステーブルから射影 (project) するかコピー (copy) するか選択できる。

### DynamoDB Streams

テーブルのデータ操作イベントを捕捉する機能が DynamoDB Stream

テーブルに以下のイベントが起きた際に、そのイベントに関する情報がストリームレコード (stream record) として流れてくる。

- アイテムの**追加**: そのアイテムの全属性を含むアイテムのイメージ (image)
- アイテムの**更新**: 変更された (modified) 属性の、変更前後のイメージ
- アイテムの**削除**: 削除される前のアイテムのイメージ

このほか、ストリームレコードには次の情報も含まれている。

- テーブル名
- イベントのタイムスタンプ
- その他のメタデータ

各ストリームレコードは 24 時間後にストリームから自動的に削除される。

## References

- [Getting Started with DynamoDB \- Amazon DynamoDB](https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/GettingStartedDynamoDB.html)
- [AWS DynamoDBの概要をざっくり理解する \- Qiita](https://qiita.com/t-shmp/items/222e2f96d5bda4b42bde)
