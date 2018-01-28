# 変更履歴

## v0.0.5

### esa regexpコマンドを実装
- esaのpost（記事）を検索keywordで検索し、正規表現で抽出した文字列をまとめてクリップボードにコピーしてくれる
    - 用途例として、直近20件の日報の `悪かったこと` に書いた内容をまとめて振り返りたい等
- ※注意※ページネーション対応してないので20件までしかとれません... :pray:
- `<username> <検索keyword> <正規表現> <抽出した文字列のprefix (Optional)> `:
    - 直近20件の自分の日報の悪かったことを抽出する（例: `esa regexp "kyokomi" "日報" "# 悪かったこと([\s\S]*)# 所感" "## "` ）
    - 悪かったことの次のセクションが所感なので↑の例になっている


## v0.0.4

### todayコマンドに以下の引数を対応
- `today <username> <number of hours ago>`:
    - 指定した時間前の日付を基準に取得（例: `esa today kyokomi 48` ）
- `today <username> <yyyy-MM-dd>`:
    - 指定した日付を基準に取得（例: `esa today kyokomi 2017-02-01` ）

### デフォルトを24時間前ではなく、0時基準で今日の日付に変更

### 現在時刻が 2017/02/08 09:32:00 だった場合 => 2017/02/08 00:00:00 を基準日とする

## v0.0.3

### esa todayコマンドを実装

## v0.0.2

### bugfixなど

## v0.0.1

### esa setupコマンドを実装
Get Personal Token which have a scope for read and put it on:

`https://<your team name>.esa.io/user/applications`

- `esa setup <accessToken> <teamName>`:

### esa searchコマンドを実装
- `esa search <query>`:
    - 指定した条件で記事を検索します（例: `esa search 日報 -ポエム`）
    - 検索条件はこちらを参照ください https://docs.esa.io/posts/104
