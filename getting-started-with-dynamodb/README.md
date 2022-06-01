# Getting Started with DynamoDB

https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/GettingStartedDynamoDB.html

## Basic Concepts

[Core Components of Amazon DynamoDB \- Amazon DynamoDB](https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/HowItWorks.CoreComponents.html)

### Tables, Items, and Attributes

- **テーブル (table)**: データの集合。他の DB と同じ。テーブルは 0 個以上の項目 (items) を含む。
- **項目 (item)**: 一意に識別可能な、属性 (attributes) の集合。他の DB の行 (rows) やレコード (records) に当たる。 1 つの項目は 1 つ以上の属性から成る。 DynamoDB では 1 つのテーブルに保管できる項目の数に制限はない。
    - 各項目は一意な識別子（**プライマリキー**、primary key）を持つ。プライマリキーは **1 つまたは 2 つの属性**で構成される。2 つの場合はそれらの組み合わせで項目を識別する。
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
