# data-interface-for-salesforce-contract-create
data-interface-for-salesforce-contract-create は、salesforce の契約オブジェクトを新規で作成するために必要なデータの整形、および作成時に salesforce から返ってきた response の mysql への格納を行います。

# 動作環境
data-interface-for-salesforce-contract-create は、aion-coreのプラットフォーム上での動作を前提としています。
使用する際は、事前に下記の通りAIONの動作環境を用意してください。

OS: Linux OS  
CPU: ARM/AMD/Intel  
Kubernetes  
AIONのリソース  

## 事前準備
本マイクロサービスは DB に MySQL を使用します。
MySQL に関する設定は 環境変数に記述されています。

| env_name | description |
| --- | --- |
| MYSQL_HOST | ホスト名 |
| MYSQL_PORT | ポート番号 |
| MYSQL_USER | ユーザー名 |
| MYSQL_PASSWORD | パスワード |
| MYSQL_DBNAME | データベース名 |
| MAX_OPEN_CONNECTION | 最大コネクション数 |
| MAX_IDLE_CONNECTION | アイドル状態の最大コネクション数 |
| KANBANADDR: | kanban のアドレス |
| TZ | タイムゾーン |

## kanban との通信
### kanban から受信するデータ
kanban から受信する metadata に下記の情報を含む必要があります。

| key | value |
| --- | --- |
| method | post |
| object | Contract |
| account_id | 操作対象が紐づく Account ID |
| connection_type | request |

具体例 : 
```example
# metadata (map[string]interface{}) の中身

"method": "post"
"object": "Contract"
"account_id": "xxxx"
"connection_type": "request"
```

### kanban に送信するデータ
kanban に送信する metadata は下記の情報を含みます。

| key | type | description |
| --- | --- | --- |
| method | string | 文字列 "post" を指定 |
| object | string | 文字列 "Contract" を指定|
| connection_key | string | 文字列 "contract_create" を指定|

具体例: 
```example
# metadata (map[string]interface{}) の中身

"method": "post"
"object": "Contract"
"connection_key": "contract_create"
```

## kanban から受信するデータ
kanban からの受信可能データは下記の形式です

| key | value |
| --- | --- |
| key | 文字列 "Contract" |
| content | Contract の詳細情報を含む JSON |
| connection_type | 文字列 "response" |

具体例:
```example
# metadata (map[string]interface{}) の中身

"key": "Contract"
"content": "{xxxxxxxxxxxxxx}"
"connection_type": "response"
```
