一応、自分用メモでもある。


## 開発手順


### 初回のみ準備

1. buildして.envrcに `export alfred_workflow_data="./"` とかを追加する（configファイルが存在するpathを設定する）
1. config.jsonを配置する

### 機能追加時

1. `cmd/alfred-esa-workflow` 下に新しいコマンドを追加する（他のコマンドを参考にXxxxServiceのstructを作る）
    - 基本的にXxxxServiceのstructの中にメンバ関数とかを組み込むようにする（他コマンドで依存しないようにする）
1. go buildして適当にCLI上で実行してテストする
1. `package.json` のversionをインクリメントする（semverな感じで雰囲気でやる）
1. `wercker.yml` の方のversionもインクリメントする（これはgithubのタグ用）
    - TODO: package.jsonの内容を読み込むようにしたさある...
1. CHANGELOG.mdに追加した機能について説明を書く
    - これがalfred-workflowをupdateするときの説明に出てくる
1. `make build` してdev版のworkflowを作る
1. CHANGELOGの内容などを確認して、新しいflowが必要な追加して確認する
1. `make release` してrelease版のworkflowを作る
1. 先程dev版で作ったflowをコピーして貼り付ける
1. Open in Finderで開いてinfo.plistをコピーしてresources下にペーストする
1. これら全てをgit commitしてpushし、masterにマージしたらci経由でgithubのreleaseにuploadされる


## config.json例

```json
{"accessToken":"<esa上で発行したtoken>","teamName":"<チーム名>"}
```

