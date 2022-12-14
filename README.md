# Web_app
提供通用的 web 项目开发项目搭建范例。

目前项目中完成了：
+ 通用模块划分
+ controller
+ routes
+ middleware
+ logic
+ models
+ dao
+ pkg
+ conf
+ setting
+ logger

已提供功能接口：
+ signup
+ login
+ ~~ping~~
+ 
# `Air`

添加提高开发效率工具，完成源码下载之后，运行如下代码安装并运行 `air` ：
```Bash
./start_air.sh
```

## 已完成接口
```API
POST   /api/v1/signup            --> webapp.io/controllers/userHanlder.UserSignUpHandler (3 handlers)
POST   /api/v1/login             --> webapp.io/controllers/userHanlder.UserLoginHandler (3 handlers)
GET    /api/v1/community         --> webapp.io/controllers/community.GetCommunityHandler (4 handlers)
GET    /api/v1/community/:id     --> webapp.io/controllers/community.GetCommunityDetailHandler (4 handlers)
POST   /api/v1/post              --> webapp.io/controllers/post.CreatePostHandler (4 handlers)
GET    /api/v1/post/:id          --> webapp.io/controllers/post.GetPostDetailHandler (4 handlers)
GET    /api/v1/posts             --> webapp.io/controllers/post.GetPostListHandler (4 handlers)
GET    /api/v1/postssorted       --> webapp.io/controllers/post.GetPostListSortedHandler (4 handlers)
POST   /api/v1/vote              --> webapp.io/controllers/vote.PostVoteHandler (4 handlers)
```

## 添加压测工具 go-wrk
下载 wrk 仓库代码到本地，然后编译代码使用编译出来的可执行文件，使用下方命令完成压测。
```Bash
https://github.com/adjust/go-wrk.git

go mod init
go mod tidy
go build
```
```Bash
go-wrk -t=8 -c=100 -n=10000 "http://127.0.0.1:8001/api/v1/posts?size=1&page=1"
```
