# koebuta

ランダムに女性声優の画像を取得するbotです。

# deploy

AWS Lambda環境でのdeployを想定しています。
WEBサーバーを用意すれば別環境でもOKです。

Lambdaの場合は、これをそのままuploadしてください。

## 環境変数

| key | content | example |
| --- | --- | --- |
| KB_MODE | 動作モードを指定します | "outgoing" |
| KB_USER | slackへのpostに使うユーザーです | koebuta |
| KB_URL | slackへのpost用URLです | https://hooks.slack.com/services/xxx/yyy/zzzz |
| KB_ICON | slackへのpostにつかうICON指定です。 | :pig_nose: |
| KB_CHANNEL | slackへのpostにchannel指定です。 | #example |
| KB_SLACK_TOKEN | slackのoutgoin webhockの認証トークンです。 | hoge |

## incoming bot

slackのincomig webhockを有効にして各種環境変数設定するだけでOK。
`KB_MODE` は `incoming` に設定。
単にrequestを投げるだけなのでサーバーいらず。

## outgoing bot

- slackのoutgoing webhockを有効にする。
- tokenを環境変数に設定する。
- `KB_MODE` は `outgoing` に設定。
- API GWを設定してLambdaに繋ぎこむ
- Slackのoutgoingwebhockのリクエストは `x-www-form-urlencoded` なので、API GWの前処理でjsonに変換してあげる

## simple

単純な画像URL取得コマンドとして使う場合はsimpleモードを活用ください。

`KB_MODE=simple`
