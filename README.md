# 实现一个简单的web server

1. 初始化配置文件
2. 初始化日志库
3. 连接mysql数据库
4. 连接redis数据库
5. 注册路由
6. 启动server

---

## 学习参考

- [基于gin实现一个简单的web server](https://juejin.cn/post/7033793209360711716)
- [构建你的docker镜像](https://docs.docker.com/language/golang/build-images/)

## TODO

- [x] 初始化配置文件
- [ ] 基于docker配置开发环境
- [ ] 初始化日志库
- [ ] 连接mysql、redis
- [ ] 注册路由
- [ ] 集成swagger
- [ ] 表单校验
- [ ] crud

## use in docker

### build

```bash
docker build --tag gin-study .
```

### run

```bash
docker run -p 3000:3000 gin-study
```

### test

```bash
curl http://localhost:3000/ping
{"message":"pong"}%
```
