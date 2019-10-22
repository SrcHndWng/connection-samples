# tcp

## About This

TCPサーバのサンプル。
[8.1 Socketプログラミング](https://astaxie.gitbooks.io/build-web-application-with-golang/ja/08.1.html)を参考に
`quit`などを追加・改修したもの。

## Usage

### Server

```
$ go build
$ ./main
```

### Client

クライアントにはtelnetを使用。

```
$ telnet 127.0.0.1 8888
```

## example

### Server

```
$ ./main
2019/10/22 20:22:45 received length = 5, message = aaa
2019/10/22 20:22:52 received length = 5, message = bbb
2019/10/22 20:22:57 received length = 6, message = quit
2019/10/22 20:22:57 quit!!
2019/10/22 20:23:05 received length = 5, message = ccc
2019/10/22 20:24:02 read error. err = read tcp 127.0.0.1:8888->127.0.0.1:61516: i/o timeout
```

### Client

```
$ telnet 127.0.0.1 8888
Trying 127.0.0.1...
Connected to 127.0.0.1.
Escape character is '^]'.
aaa                                        # your input
[2019/10/22 20:22:45] your message is aaa  # return result
bbb
[2019/10/22 20:22:52] your message is bbb
quit                                       # input for quit
Connection closed by foreign host.
$ telnet 127.0.0.1 8888                    # re-connect telnet
Trying 127.0.0.1...
Connected to 127.0.0.1.
Escape character is '^]'.
ccc
[2019/10/22 20:23:05] your message is ccc
Connection closed by foreign host.
```
