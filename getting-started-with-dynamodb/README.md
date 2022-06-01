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

- 項目の**追加**: その項目の全属性を含む項目のイメージ (image)
- 項目の**更新**: 変更された (modified) 属性の、変更前後のイメージ
- 項目の**削除**: 削除される前の項目のイメージ

このほか、ストリームレコードには次の情報も含まれている。

- テーブル名
- イベントのタイムスタンプ
- その他のメタデータ

各ストリームレコードは 24 時間後にストリームから自動的に削除される。

## Prerequisites

[Prerequisites \- Getting Started Tutorial \- Amazon DynamoDB](https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/GettingStarted.SettingUp.DynamoWebService.html)

DynamoDB を使う事前準備

### ローカル環境：DynamoDB Local

[Deploying DynamoDB Locally on Your Computer \- Amazon DynamoDB](https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/DynamoDBLocal.DownloadingAndRunning.html#docker)

Jar ファイル、Docker イメージ、Maven のいずれかで利用できる。

ここでは Docker イメージを使ってみる
[amazon/dynamodb\-local \- Docker Image \| Docker Hub](https://hub.docker.com/r/amazon/dynamodb-local)

#### Docker image

```console
$ docker -v
Docker version 20.10.16, build aa7e414
$ docker compose version
Docker Compose version v2.6.0
$ aws --version
aws-cli/2.7.5 Python/3.9.11 Linux/5.10.102.1-microsoft-standard-WSL2 exe/x86_64.ubuntu.22 prompt/off
$
```

docker-compose.yml を作成して `docker compose up` で起動

```console
$ mkdir data
$ docker compose up
```

適当なダミー認証情報を設定し、エンドポイントの URL に `http://localhost:8000/` を指定して、適当な API を叩いてみる：

```console
$ aws configure
AWS Access Key ID [None]: DUMMY
AWS Secret Access Key [None]: DUMMY
Default region name [ap-northeast-1]: 
Default output format [None]: 
$ aws dynamodb list-tables --endpoint-url=http://localhost:8000/
{
    "TableNames": []
}
$
```

DB ファイルが SQLite だった

```console
$ file data/shared-local-instance.db 
data/shared-local-instance.db: SQLite 3.x database, last written using SQLite version 3008007, page size 1024, file counter 1, database pages 13, cookie 0x8, schema 4, UTF-8, version-valid-for 1
$
```

### ウェブサービス：DynamoDB

[Setting Up DynamoDB \(Web Service\) \- Amazon DynamoDB](https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/SettingUp.DynamoWebService.html)

(TODO)

以降は DynamoDB Local で行う。

## Step 1: Create a Table

(TODO)

## References

- [Getting Started with DynamoDB \- Amazon DynamoDB](https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/GettingStartedDynamoDB.html)
- [AWS DynamoDBの概要をざっくり理解する \- Qiita](https://qiita.com/t-shmp/items/222e2f96d5bda4b42bde)
