模块二作业：
编写一个 HTTP 服务器（大家视个人不同情况决定完成到哪个环节，但尽量把1都做完）
- 接收客户端 request，并将 request 中带的 header 写入 response header
- 读取当前系统的环境变量中的 VERSION 配置，并写入 response header
- Server 端记录访问日志包括客户端 IP，HTTP 返回码，输出到 server 端的标准输出
- 当访问 localhost/healthz 时，应返回 200
  提交链接🔗：https://jinshuju.net/f/PlZ3xg
  截止时间：10月7日晚23:59前

模块三作业：
构建本地镜像。
- 编写 Dockerfile 将练习 2.2 编写的 httpserver 容器化（请思考有哪些最佳实践可以引入到 Dockerfile 中来）。
- 将镜像推送至 Docker 官方镜像仓库。
- 通过 Docker 命令本地启动 httpserver。
- 通过 nsenter 进入容器查看 IP 配置。
- 作业需编写并提交 Dockerfile 及源代码。
  提交链接：https://jinshuju.net/f/rxeJhn
  截止日期：10月17日晚23:59之前
  提示💡
  1、自行选择做作业的地址，只要提交的链接能让助教老师打开即可
  2、自己所在的助教答疑群是几组，提交作业就选几组

## 第二次作业
> 参考
> [go_mod](https://ld246.com/article/1565244079032)
> [docker官网](https://docs.docker.com/language/golang/build-images/)
```shell
#启动
$ cd ch7 && ls
Dockerfile go.mod     go.sum     main.go
$docker build .  -t liaoyingjie/gohttp_mu:v1
$ docker images
REPOSITORY              TAG        IMAGE ID       CREATED             SIZE
liaoyingjie/gohttp_mu   v1         7a03a46b6371   33 minutes ago      25.5MB
#push到dockerhub后删除本地
$docker run -it -p 1080:1080 liaoyingjie/gohttp_mu:v1
Unable to find image 'liaoyingjie/gohtlstp_mu:v1' locally
v1: Pulling from liaoyingjie/gohttp_mu
cfb92865f5ba: Already exists 
8dd350b5e0d5: Already exists 
010c251fdc94: Already exists
Digest: sha256:d0a4fda68fc4601ce4dfe0571c2230df04f1c82c3bd5df53996c2169f3b92741
Status: Downloaded newer image for liaoyingjie/gohttp_mu:v1
INFO[0021] ip: 172.17.0.1:60600, http code: 200         
INFO[0022] ip: 172.17.0.1:60600, http code: 200   
```

