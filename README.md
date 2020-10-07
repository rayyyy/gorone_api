# Gorone API

このプロジェクトの責務
* リクエストを受ける窓口
* 登録したリクエストはredisに登録する
* 基本は非同期のAPIで下記の3構成
  * Request登録API
  * Status取得API
  * Result取得API
* キャッシュ機能
  * すでに処理が完了していたらpodに仕事させない



```
sql-migrate up --env=production
```