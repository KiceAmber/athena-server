# 目前已经完成的功能

- 文章模块：
    - [ ] 添加文章
    - [ ] 查询文章列表
    - [ ] 删除文章
- 标签模块：
    - [x] 查询标签列表
    - [x] 添加标签
    - [x] 删除标签
    - [x] 更新标签
- 分类模块：
    - [ ] 查询分类列表
    - [ ] 添加分类
    - [ ] 删除分类
    - [ ] 更新分类

# Third party libraries

## GoFrame

官网地址：https://goframe.org/

## Markdown parser

```shell
go get -u "github.com/russross/blackfriday"
```

# Docker images

## MySQL

```shell
docker run -d -p 3306:3306 --name=mysql -v mysql_data_volume:/var/lib/mysql -e MYSQL_ROOT_PASSWORD=123456 mysql:8.3.0
```

## Redis

```shell
docker run -d -p 6379:6379 --name=redis redis:7.2.4
```
