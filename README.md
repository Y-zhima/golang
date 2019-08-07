# 服务协议

## 定义

服务协议的`proto`文件在 **src** 目录下定义

1. 每个服务一个目录，服务名和目录名一致
2. 公用部分的协议统一定义在`common`目录，如错误码
3. `proto` 之间可以相互 `import`，避免重复定义
4. `proto` 可以定义路由，路由规则默认: `/v[0-9]/{service}/{action}`, `action` 默认首字母小写驼峰写法, 例如 

```
// 文件服务
service File {
    rpc UploadPlaybook (stream UploadPlaybookRequest) returns (UploadPlaybookResponse) {
        option (google.api.http) = {
            post: "/v1/file/uploadPlaybook"
            body: "*"
        };
    }
}
```

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

#### swagger ui
参考: http://aloyzh.com/2018/06/go-integrate-swagger

1. 合并swagger json
```
swagger mixin swagger/playbook/playbook.swagger.json
```

2. swagger serve
```
swagger serve swagger.json -Fswagger
```

## 提交