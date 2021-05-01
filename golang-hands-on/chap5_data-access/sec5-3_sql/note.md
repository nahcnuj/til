# 5-3. データベースアクセス

```console
$ docker-compose run --rm db data.sqlite3 <./sql/init.sql
$ docker-compose run --rm db data.sqlite3 "select * from mydata"
1|jun|nahcnuj.work@gmail.com|25
2|taro|taro@example.com|39
3|hanako|hanako@example.com|28
4|sachiko|sachiko@example.com|17
5|jiro|jiro@example.com|6
```

```console
$ go mod init sqlite-practice
$ go mod tidy
```

```console
$ go run .
id: 1
<"1:jun" nahcnuj.work@gmail.com,25>
id: 2
<"2:taro" taro@example.com,39>
id: 3
<"3:hanako" hanako@example.com,28>
id: 4
<"4:sachiko" sachiko@example.com,17>
id: 5
<"5:jiro" jiro@example.com,6>
id: 6
id: a
strconv.Atoi: parsing "a": invalid syntax
$
```

`id := strconv.Atoi(s)` せず、`db.Query()` に直接 `s` を渡した場合

```console
$ go run .
id: 1
<"1:jun" nahcnuj.work@gmail.com,25>
id: 2
<"2:taro" taro@example.com,39>
id: 3
<"3:hanako" hanako@example.com,28>
id: 4
<"4:sachiko" sachiko@example.com,17>
id: 5
<"5:jiro" jiro@example.com,6>
id: 6
id: a
id: 1 OR 1 == 1
id: 
$
```