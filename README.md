## 部署说明

```bash
cp config.example.yaml config.yaml
cp Makefile-example Makefile
```

注意：
- 修改2个文件中的数据库连接配置
- 默认提供的迁移命令 migrate.darwin-amd64 是 mac 的版本，需要自行去下载系统对应的版本。下载好后，更新 Makefile 的迁移指令所使用的命令行


## 迁移与回滚

```baash
make migrate
make rollback
```

## 创建迁移

```bash
# NAME 参数传递的是创建的表名，生成的文件在 db/migrations 目录下，使用 sql 写法。
# 为了方便多服务器、多环境同步表结构，请使用原生 sql 处理表结构变更。
make create NAME=xxx
```

## 本地开发

- air 命令可以热更新，前提，代码没有报错。报错了没有提示，看不出来因为报错导致请求失败
- 启动命令是 make，需要热更新写代码，使用 air 命令。使用 make 命令，需要改动代码后，手动停止命令并启动 http 服务。

相关命令参考 Makefile

```bash
git clone git@codeup.aliyun.com:cblink/flow/china-life/data-distribute/data-distribute-api.git data-distribute-api

make
```

## 目录结构

```
.
├── Dockerfile                                                              # 容器化 Dockerfile
├── Makefile                                                                # 辅助命令，请查看有什么可以执行的命令，需要配置数据库才可以执行迁移
├── Makefile-example                                                        # Makefile example 文件
├── README.md                                                               # README 自叙文件
├── cmd                                                                     # 命令行、入口目录
│   ├── main.go                                                             # 入口文件
│   └── server                                                              # server 命令所在目录
│       └── cmd.go                                                          # server 命令，初始化配置、启动引导。请手动控制第三方服务的引导启动顺序，api 接口定义也在这里
├── config.example.yaml                                                     # 配置文件参考
├── config.yaml                                                             # 配置文件
├── db                                                                      # 数据库相关的目录
│   └── migrations                                                          # 数据库表结构相关的目录
│       ├── 20231019124741_create_qc_tasks_table.down.sql
│       ├── 20231019124741_create_qc_tasks_table.up.sql
│       ├── 20231021031811_create_qc_task_samples_table.down.sql
│       └── 20231021031811_create_qc_task_samples_table.up.sql
├── go.mod                                                                  # go module 文件
├── go.sum                                                                  # go sum 文件
├── internal                                                                # 内部目录
│   ├── handlers                                                            # 类似于 MVC 的控制器
│   │   ├── common.go                                                       # 公共结构体、函数
│   │   ├── data.go                                                         # 控制器作用，handle 请求处理，data 数据请求相关接口
│   │   ├── qc_task.go                                                      # 控制器作用，handle 请求处理
│   │   └── qc_task_sample.go                                               # 控制器作用，handle 请求处理
│   ├── initialization                                                      # 项目初始化相关文件
│   │   ├── config.go                                                       # 初始化配置，读取 yaml 配置文件
│   │   └── db.go                                                           # 初始化 db 连接
│   └── models                                                              # models 模型，gorm 模型
│       ├── qc_task.go                                                      # 模型 qc_task
│       └── qc_task_sample.go                                               # 模型 qc_task_sample
├── migrate.darwin-amd64                                                    # mac 的数据库迁移工具
├── pkg                                                                     # 第三方服务扩展封装
│   └── rabbitmq                                                            # rabbitmq 队列
│       └── client.go                                                       # rabbitmq 队列的 client，可以发送队列、消费消息
└── tmp                                                                     # 临时目录，通过 air 编译、热更新。
    └── main                                                                # 实时编译的入口文件
```