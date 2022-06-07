# nogisched

乃木坂46の公式サイトのスケジュールをスクレピングして，標準出力へ出力するCLI．

usage:
```bash
./nogisched -d 20220601 -notifyで通知を行う場合は使用するトークンを環境変数 `NOGI
```

options:

| name  | 説明  | デフォルト値 |
|---|---| --- |
| d  | スケージュールを取得したい日付のyyyyMMdd形式(e.g. 20220605))  | 端末の実行時の日付 |
| notify  | LINE Notifyを使って通知するか(通知する場合は後述の環境変数が必要)  | false |


出力形式
```txt
05 Sun
ラジオ 18:00〜18:30 文化放送「乃木坂46の「の」」筒井あやめ
TV 24:00〜24:30 テレビ東京系「乃木坂工事中」
```

LINE Notifyで通知を行う場合は使用するトークンを環境変数 `NOGISCHED_NOTIFY_TOKEN` にLINE Notifyのトークンを設定する．