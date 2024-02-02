# docker images

## MySQL

```shell
docker run -d -p 3306:3306 --name=mysql -v mysql_data_volume:/var/lib/mysql -e MYSQL_ROOT_PASSWORD=123456 mysql:8.3.0
```

## Redis

```shell
docker run -d -p 6379:6379 --name=redis redis:7.2.4
```