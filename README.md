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

### 生成go mock文件
```
make gen-mock
```

### 自动生成目录

- `goout`: 根据`proto`生成的go代码
- `javaout`: 根据`proto`生成的java代码
- `swagger`: 根据`proto`生成的swagger json文件
- `gooutmock`: 根据`goout`生成的mock文件

#### swagger ui
```
make swagger-ui
```

## 提交