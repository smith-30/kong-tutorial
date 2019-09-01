## kong とは

OpenResty をベースに作られたAPI Gateway
マイクロサービスに必要であろうビルトインプラグインが豊富  
足りないものは開発者がプラグインを作ることもできる  

(Open Resty は nginx の Lua 拡張)

## 基本アクセス

- 8000
  - リバプロ用、アップストリームに流す
- 8443
  - https用
- 8001
  - admin api、kongの設定に使う
- 8444
  - 8001 のhttpsの口

## バックエンドサービスの登録

[mockbin](https://mockbin.com/) を登録する  
自由にコールさせてくれるサービスっぽい  

