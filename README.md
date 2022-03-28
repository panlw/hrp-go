# 最简单的 HTTP 反向代理服务

## dotenv 文件方式

```
# .env
HRP_INNER_HOST=localhost:1234
HRP_SERVE_PORT=4321
```

## 无配置文件方式

```shell
HRP_INNER_HOST=localhost:1234 HRP_SERVE_PORT=4321 HRP
```

## 使用 caddy 给出同样的能力

```shell
caddy reverse-proxy -insecure --from :4321 --to 127.0.0.1:1234
```
