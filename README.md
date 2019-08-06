# 服务协议

## 定义

服务协议的`proto`文件在**src**目录下定义

```
cd src
```

1. 每个服务一个目录，目录名对应gitlab项目名
2. `proto`之间可以相互`import`，避免重复定义
3. `proto`的`package`定义避免使用`-`等特殊符号，统一使用英文小写（可以跟目录不一致）
4. 公用部分的协议统一定义在`common`目录，如错误码

## 编译

```
make
```

**Windows环境下可以直接运行docker命令**

```
docker run -it -v $(PWD):/opt/protos 192.168.0.103:8080/axe/protos:v0.0.1 sh gen.sh
```

### 自动生成目录

1. `goout`: 根据`proto`生成的go代码
2. `swagger`: 根据`proto`生成的swagger json文件

### 合并所有的swagger

```
swagger mixin swagger/playbook/playbook.swagger.json
```

### Serve the API

```
swagger serve swagger.json -Fswagger
```


[参考] <http://aloyzh.com/2018/06/go-integrate-swagger/>
## 提交