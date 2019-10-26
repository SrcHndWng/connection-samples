# udp

## About This

UDPサーバのサンプル。
[8.1 Socketプログラミング](https://astaxie.gitbooks.io/build-web-application-with-golang/ja/08.1.html)を参考に改修したもの。

## Usage

### Server

```
$ go build
$ ./main
```

### Client

クライアントにはnetcatを使用。

```
$ netcat -u 127.0.0.1 8888
```

## example

### Server

```
$ ./main
2019/10/26 20:45:26 received length = 4, message = aaa
2019/10/26 20:45:30 received length = 4, message = bbb
2019/10/26 20:45:33 received length = 4, message = ccc
```

### Client

```
$ netcat -u 127.0.0.1 8888
aaa
[2019/10/26 20:45:26] your message is aaa
bbb
[2019/10/26 20:45:30] your message is bbb
ccc
[2019/10/26 20:45:33] your message is ccc
```
