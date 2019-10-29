# redis-chat

## About This

Redisを使ったのサンプル。
[GoとRedisにおける簡単なチャットアプリケーション](https://medium.com/eureka-engineering/go-redis-application-28c8c793a652)を参考に改修したもの。

## Usage

### build

```
$ go build
```

### example

ターミナルを2つ開く。それぞれをA、Bとし、以下の例はA、Bの順に交互に入力した場合となる。

#### Terminal A

```
$ ./main Mr.A
Mr.A>Mr.B has joined
hello, Mr.B
Mr.A>Mr.A:hello, Mr.B
Mr.B:hello Mr.A
how are you?
Mr.A>Mr.A:how are you?
Mr.B:I'm fine. and you?
fine.
Mr.A>Mr.A:fine.
Mr.B:bye!!
/exit
```

#### Terminal B


```
$ ./main Mr.B
Mr.B>Mr.A:hello, Mr.B
hello Mr.A
Mr.B>Mr.B:hello Mr.A
Mr.A:how are you?
I'm fine. and you?
Mr.B>Mr.B:I'm fine. and you?
Mr.A:fine.
/who
Mr.B>Online members are Mr.A,Mr.B.
bye!!
Mr.B>Mr.B:bye!!
Mr.A has left
/exit
```
