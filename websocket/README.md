# udp

## About This

WebSocketのサンプル。
[8.2 WebSocket](https://astaxie.gitbooks.io/build-web-application-with-golang/ja/08.2.html)を参考に改修したもの。

## Usage

### Server

```
$ go build
$ ./main
```

### Client

index.htmlをブラウザで開く。
開発者ツールなどでコンソールを表示する。

## example

### Server

```
$ ./main
2019/10/28 21:13:34 Server received : WebSocket Sample!!
2019/10/28 21:13:44 Server received : WebSocket Sample!!
^C # Ctrl + Cなどでサーバを停止する
```

### Client

開発者ツールのコンソールの表示例。

```
onload
index.html:15 connected to ws://127.0.0.1:8888
index.html:23 [2019/10/28 21:13:34] Server received : WebSocket Sample!!
index.html:23 [2019/10/28 21:13:44] Server received : WebSocket Sample!!
index.html:19 connection closed (1006)
```
