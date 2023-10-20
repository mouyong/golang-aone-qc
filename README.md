## 部署说明

```
cp config.example.yaml
cp Makefile-example Makefile
```

修改2个文件中的数据库连接配置

## 迁移与回滚
```
make migrate
make rollback
```

## 创建迁移

```
make create NAME=xxx
```
