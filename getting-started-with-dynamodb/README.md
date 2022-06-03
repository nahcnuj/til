# Getting Started with DynamoDB

https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/GettingStartedDynamoDB.html

## Basic Concepts

[Core Components of Amazon DynamoDB \- Amazon DynamoDB](https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/HowItWorks.CoreComponents.html)

### Tables, Items, and Attributes

- **テーブル (table)**
- **項目 (item)**
  - 各項目は**プライマリキー**によって一意に識別可能
  - プライマリキーを除いて**スキーマレス**
- **属性 (attribute)**
  - データ型
    - 参照：[Supported data types and naming rules in Amazon DynamoDB \- Amazon DynamoDB](https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/HowItWorks.NamingRulesDataTypes.html#HowItWorks.DataTypes)
      - スカラー：数値、文字列、バイナリ、真偽値、null
      - ドキュメント：リスト、マップ
      - 集合 (set)：数値/文字列/バイナリ の集合（順序なし、重複なし）

### Primary Key

- **パーティションキー** (partition key)
- **パーティションキーとソートキー** (sort key)、複合プライマリキー

パーティションキーによって、その項目が（DynamoDB の内部物理ストレージの）どのパーティションに保管されるかが決められる。

複合プライマリキーの場合は、同じパーティションキーを持つ項目を複数持たせられ、ソートキーの値で並べられる。

プライマリキー属性は文字列、数値、バイナリのいずれかでなければならない。

### Secondary Indexes

プライマリキー以外のキーでクエリするための索引。

- **グローバルセカンダリインデックス** (global secondary index; GSI): テーブルのプライマリキーとは異なり得るパーティションキーとソートキーによるインデックス
- **ローカルセカンダリインデックス** (local secondary index; LSI): テーブルのプライマリキーと同じパーティションキーと、テーブルのプライマリキーとは異なるソートキーによるインデックス

1 つのテーブルにつき 20 個（デフォルトクォータ）の GSI と 5 つの LSI を持てる。

インデックスの属性は、テーブルから射影 (project) するかコピー (copy) するかを選択可能。

### DynamoDB Streams

テーブルのデータ操作イベントを捕捉する機能

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

```sh
aws dynamodb create-table \
  --table-name Music \
  --attribute-definitions \
    AttributeName=Artist,AttributeType=S \
    AttributeName=SongTitle,AttributeType=S \
  --key-schema \
    AttributeName=Artist,KeyType=HASH \
    AttributeName=SongTitle,KeyType=RANGE \
  --provisioned-throughput \
    ReadCapacityUnits=1,WriteCapacityUnits=1 \
  --table-class STANDARD \
  --endpoint="http://localhost:8000"
```

結果：

```json
{
    "TableDescription": {
        "AttributeDefinitions": [
            {
                "AttributeName": "Artist",
                "AttributeType": "S"
            },
            {
                "AttributeName": "SongTitle",
                "AttributeType": "S"
            }
        ],
        "TableName": "Music",
        "KeySchema": [
            {
                "AttributeName": "Artist",
                "KeyType": "HASH"
            },
            {
                "AttributeName": "SongTitle",
                "KeyType": "RANGE"
            }
        ],
        "TableStatus": "ACTIVE",
        "CreationDateTime": "2022-06-03T02:41:23.983000+09:00",
        "ProvisionedThroughput": {
            "LastIncreaseDateTime": "1970-01-01T09:00:00+09:00",
            "LastDecreaseDateTime": "1970-01-01T09:00:00+09:00",
            "NumberOfDecreasesToday": 0,
            "ReadCapacityUnits": 1,
            "WriteCapacityUnits": 1
        },
        "TableSizeBytes": 0,
        "ItemCount": 0,
        "TableArn": "arn:aws:dynamodb:ddblocal:000000000000:table/Music"
    }
}
```

ローカルで実行したので既に TableStatus が ACTIVE になっているが、
おそらく AWS でやると CREATING になるのだろう。 (TODO)

```console
$ aws dynamodb list-tables --endpoint-url=http://localhost:8000/
{
    "TableNames": [
        "Music"
    ]
}
$ aws dynamodb describe-table --table-name Music --endpoint-url=http://localhost:8000/
# 省略：先の create-table と同じ結果
```

### テーブルの作り方

[Basic Operations on DynamoDB Tables \- Amazon DynamoDB](https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/WorkingWithTables.Basics.html#WorkingWithTables.Basics.CreateTable) の "Creating a Table"

必須設定：

- テーブル名: `--table-name`
  - 命名規則
    - 参照：[DynamoDB の命名規則](https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/HowItWorks.NamingRulesDataTypes.html#HowItWorks.NamingRules)
      - `/^[a-zA-Z0-9_.-]{3,255}$/`
      - case-sensitive
      - UTF-8
  - AWS アカウントとリージョンで一意でなければならない
    - 同じテーブル名でも、リージョンが違えば全く別のテーブルになる
- プライマリキー: `--attribute-definitions`, `--key-schema`
  - 属性名: `AttributeName`
    - 命名規則
        - 参照：[DynamoDB の命名規則](https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/HowItWorks.NamingRulesDataTypes.html#HowItWorks.NamingRules)
        - `/^[a-zA-Z0-9_.-]+$/` かつ 64 KB を超えない
        - 次の属性は 255 文字を超えない
          - セカンダリインデックス（パーティションキー・ソートキー）
          - ユーザー指定の射影属性 (projected attributes)
        - case-sensitive
        - UTF-8
  - データ型: `AttributeType` （`--attribute-definitions` で指定）
    - 参照：[AWS::DynamoDB::Table AttributeDefinition \- AWS CloudFormation](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dynamodb-attributedef.html)
      - 文字列型: `S`
      - 数値型: `N`
      - バイナリ型: `B`
  - 属性の役割: `KeyType` （`--key-schema` で指定）
    - パーティションキー: `HASH`
    - ソートキー: `RANGE`
- スループット: `--provisioned-throughput` （プロビジョンドキャパシティモードの場合）
  - 読み取り容量ユニット: `ReadCapacityUnits`
  - 書き込み容量ユニット: `WriteCapacityUnits`

読み取り/書き込み容量ユニットについては料金ページを参照：
[プロビジョニング済みキャパシティーの料金 \- Amazon DynamoDB \| AWS](https://aws.amazon.com/jp/dynamodb/pricing/provisioned/)

DynamoDB Local では、スループットの設定は必要だが実際には使われないので注意。
参照：[DynamoDB Local の使用に関する注意事項 \- Amazon DynamoDB](https://docs.aws.amazon.com/ja_jp/amazondynamodb/latest/developerguide/DynamoDBLocal.UsageNotes.html#DynamoDBLocal.Differences)

また、 `--table-class` はテーブルクラスで料金に影響する。
テーブルクラスには次の 2 種類がある：

- Standard: `STANDARD`
  - 標準の料金体系
- Standard-Infrequent Access (Standard-IA): `STANDARD_INFREQUENT_ACCESS`
  - アクセスが頻繁でないデータ向けに、ストレージコストが安い（読み書きコストが高い）料金体系

詳細は料金ページを参照：
- [プロビジョニング済みキャパシティーの料金 \- Amazon DynamoDB \| AWS](https://aws.amazon.com/jp/dynamodb/pricing/provisioned/)
- [オンデマンドキャパシティーの料金 \- Amazon DynamoDB \| AWS](https://aws.amazon.com/jp/dynamodb/pricing/on-demand/)

## Step 2: Write Data to a Table Using the Console or AWS CLI

Step 1 で作成した Music テーブルに項目を挿入してみる。

```console
$ aws dynamodb put-item --endpoint-url http://localhost:8000/ \
  --table-name Music \
  --item \
    '{"Artist": {"S": "Nao Toyama"}, "SongTitle": {"S": "True Destiny"}, "Year": {"N": "2017"}}'
$ aws dynamodb put-item --endpoint-url http://localhost:8000/ \
  --table-name Music \
  --item \
    '{"Artist": {"S": "Nao Toyama"}, "SongTitle": {"S": "Chain the world"}, "Year": {"N": "2017"}, "B-side": {"BOOL": true}}'
$ aws dynamodb put-item --endpoint-url http://localhost:8000/ \
  --table-name Music \
  --item \
    '{"Artist": {"S": "Nao Toyama"}, "SongTitle": {"S": "Imakoko"}, "Year": {"N": "2018"}, "B-side": {"BOOL": false}, "InAlbums": {"SS": ["Rainbow"]}}'
$ aws dynamodb put-item --endpoint-url http://localhost:8000/ \
  --table-name Music \
  --item \
    '{"Artist": {"S": "ぽかぽかイオン"}, "SongTitle": {"S": "やじるし→"}, "Year": {"N": "2022"}, "B-side": {"BOOL": false}, "Singers": {"L": [{"S": "Nao Toyama"}, {"S": "Kiyono Yasuno"}]}}'
```

### データの挿入

詳細は：[PutItem \- Amazon DynamoDB](https://docs.aws.amazon.com/amazondynamodb/latest/APIReference/API_PutItem.html)

挿入するデータは以下のような形で指定する。
詳細は：[AttributeValue \- Amazon DynamoDB](https://docs.aws.amazon.com/amazondynamodb/latest/APIReference/API_AttributeValue.html)

```json
{
  "属性名": {
    "型": 値
  },
  ...
}
```

特に、次の型の値の書き方が特徴的。
- 数値型（`N`）：文字列 `"3.14"`
- バイナリ型（`B`）：データを Base64 でエンコードした文字列
- NULL 型（`NULL`）： `true`

なお、put-item は、指定されたプライマリキーを持つ項目が既に存在した場合、デフォルトでは新たに指定した項目で**既存の項目を置き換える**。
これを防ぐ（そのような場合にエラーにする）ためには、 `--condition-expression` オプションで `attribute_not_exists` 関数を使って条件付き挿入にする必要がある。

```console
$ aws dynamodb put-item --endpoint-url http://localhost:8000/ \
  --table-name Music \
  --item '{"Artist": {"S": "Nao Toyama"}, "SongTitle": {"S": "Imakoko"}}' \
  --condition-expression 'attribute_not_exists(Artist) AND attribute_not_exists(SongTitle)'

An error occurred (ConditionalCheckFailedException) when calling the PutItem operation: The conditional request failed
$ echo $?
254
```

参考：[Condition Expressions \- Amazon DynamoDB](https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/Expressions.ConditionExpressions.html#Expressions.ConditionExpressions.PreventingOverwrites)

## Step 3: Read Data from a Table

Step 2 で挿入したデータを読み出してみる。
ここではプライマリキーを指定して単一の項目を取得する。

```console
$ aws dynamodb get-item --endpoint-url http://localhost:8000/ \
  --table-name Music \
  --key \
    '{"Artist": {"S": "Nao Toyama"}, "SongTitle": {"S": "True Destiny"}}' \
  | jq -c
{"Item":{"Artist":{"S":"Nao Toyama"},"Year":{"N":"2017"},"SongTitle":{"S":"True Destiny"}}}
$ aws dynamodb get-item --endpoint-url http://localhost:8000/ \
  --table-name Music \
  --key \
    '{"Artist": {"S": "Nao Toyama"}, "SongTitle": {"S": "Imakoko"}}' \
  | jq -c
{"Item":{"Artist":{"S":"Nao Toyama"},"Year":{"N":"2018"},"B-side":{"BOOL":false},"SongTitle":{"S":"Imakoko"},"InAlbums":{"SS":["Rainbow"]}}}
$ aws dynamodb get-item --endpoint-url http://localhost:8000/ \
  --table-name Music \
  --key \
    '{"Artist": {"S": "ぽかぽかイオン"}, "SongTitle": {"S": "やじるし→"}}' \
  | jq -c
{"Item":{"Artist":{"S":"ぽかぽかイオン"},"Year":{"N":"2022"},"B-side":{"BOOL":false},"SongTitle":{"S":"やじるし→"},"Singers":{"L":[{"S":"Nao Toyama"},{"S":"Kiyono Yasuno"}]}}}
```

※各出力を [jq](https://stedolan.github.io/jq/) を使って縮めていることに注意。

## Step 4: Update Data in a Table

Step 2 で挿入したデータを更新してみる。

```console
$ aws dynamodb update-item --endpoint-url "http://localhost:8000" \
  --table-name Music \
  --key '{"Artist": {"S": "Nao Toyama"}, "SongTitle": {"S": "True Destiny"}}' \
  --update-expression 'SET #SideB = :newval' \
  --expression-attribute-values '{":newval": {"BOOL": false}}' \
  --expression-attribute-names '{"#SideB": "B-side"}' \
  --return-values ALL_NEW \
  | jq -c
{"Attributes":{"Artist":{"S":"Nao Toyama"},"Year":{"N":"2017"},"B-side":{"BOOL":false},"SongTitle":{"S":"True Destiny"}}}
```

※出力を [jq](https://stedolan.github.io/jq/) を使って縮めていることに注意。

`--update-expression` に項目を更新するための式を指定する。

ただし、設定する値は `:newval` のように `:` で始まる変数名として、`--expression-attribute-values` でマッピングする。
直接値を指定しようとすると次のように文法エラーとなる。

```console
$ aws dynamodb update-item --endpoint-url "http://localhost:8000" \
  --table-name Music \
  --key '{"Artist": {"S": "Nao Toyama"}, "SongTitle": {"S": "True Destiny"}}' \
  --update-expression 'SET NewAttr = {"NULL": true}' \
  --return-values ALL_NEW

An error occurred (ValidationException) when calling the UpdateItem operation: Invalid UpdateExpression: Syntax error; token: "{", near: "= {""
```

また、属性名が `.` `-` を含んだり、予約語であったりする場合は直接書けないため、 `#` で始まる変数名として `--expression-attribute-names` でマッピングする。
`-` を含む属性名を直接書いた場合は、次のようなエラーが発生する。

```console
$ aws dynamodb update-item --endpoint-url "http://localhost:8000" \
  --table-name Music \
  --key '{"Artist": {"S": "Nao Toyama"}, "SongTitle": {"S": "True Destiny"}}' \
  --update-expression 'SET Invalid-Name = :newval' \
  --expression-attribute-values '{":newval": {"NULL": true}}' \
  --return-values UPDATED_NEW

An error occurred (ValidationException) when calling the UpdateItem operation: Invalid UpdateExpression: Syntax error; token: "-", near: "Invalid-Name"
```

属性名に `.` を含む場合は、直接書くとネストされた属性のパスになってしまい、次のようなエラーになる。

```console
$ aws dynamodb update-item --endpoint-url "http://localhost:8000" \
  --table-name Music \
  --key '{"Artist": {"S": "Nao Toyama"}, "SongTitle": {"S": "True Destiny"}}' \
  --update-expression 'SET Invalid.AttrName = :newval' \
  --expression-attribute-values '{":newval": {"NULL": true}}' \
  --return-values UPDATED_NEW

An error occurred (ValidationException) when calling the UpdateItem operation: The document path provided in the update expression is invalid for update
```

詳細：[Update Expressions \- Amazon DynamoDB](https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/Expressions.UpdateExpressions.html)

## References

- [Getting Started with DynamoDB \- Amazon DynamoDB](https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/GettingStartedDynamoDB.html)
- [AWS DynamoDBの概要をざっくり理解する \- Qiita](https://qiita.com/t-shmp/items/222e2f96d5bda4b42bde)
