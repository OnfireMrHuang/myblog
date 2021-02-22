# 我的个人博客系统

## 文章article模块

### 获取指定文章内容

[GET]  /api/v1/articles/{id}


响应:
```json
{
    "code": 0,
    "msg": "success",
    "data": {
        "id": "文章ID",
        "title": "标题",
        "des": "简要描述",
        "content":"文章内容",
        "time": "文章创建时间",
        "tags": [], // 文章关联的标签
        "commment": [], // 文章关联的评论
    }
}
```
### 获取文章列表

[GET] /api/v1/articles?page={page}&limit={limit}
 
 响应:
```json
{
    "code": 0,
    "msg" : "success",
    "total": 6,
    "list": [
        {
            "id": "文章ID",
            "title": "文章标题",
            "time": "时间",
            "des": "文章描述",
        }
    ]
}
```

### 发表文章

[POST] /api/v1/article

请求
```json
{
    "title": "文章标题",
    "des": "文章简要描述",
    "content": "文章的内容",
    "tags": [] // 设置关联的标签
}
```

响应:
```json
{
    "code" : 0,
    "msg": "success",
    "data": {
        "id": 1,
    }
}
```

### 编辑文章
[PUT] /api/v1/article/{id}
请求:
```json
{
    "title": "文章标题修改 (optional)",
    "des": "文章简要描述 (optional)",
    "content": "修改后文章内容 (optional)",
    "tags": [] // 设置关联的tag
}

响应:
```json
{
    "code" : 0,
    "msg" : "success",
    "data" : {}
}
```

### 删除文章
[DELETE] /api/v1/article/{id}

响应:
```json
{
    "code" : 0,
    "msg" : "success",
    "data" : {}
}
```


## 评论模块

### 发布评论

[POST] /api/v1/comment

请求:
```json
{
    "article_id": "文章ID",
    "username":"用户名",
    "email": "邮箱",
    "content": "评论内容",
    "time": "时间戳",
    "ip": "IP地址"
}
```

响应:
```json
{
  "code": 0,
  "msg": "success",
  "data": {
    "id": 1, // 评论ID号
    "article_id": 2 // 关联的文章ID
  }
}
```

### 删除评论

[DELETE] /api/v1/comment/{id}

响应:
```json
{
    "code": 0,
    "msg" : "success",
    "data": {},
}
```

## 标签模块

### 获取所有标签

[GET] /api/v1/tags

响应:
```json
{
    "code": 0,
    "message": "success",
    "count" : 6,
    "list" : [
        {
            "id" : 0, // 标签的ID号
            "name" : "标签名称",
            "state" : 0 // 标签状态   
        }
    ]
}
```


### 添加标签

[POST] /api/v1/tags

请求:
```json
{
    "name" : "标签名字",
    "state" : 0 // 标签状态
}
```

响应:
```json
{
    "code": 0,
    "msg": "success",
    "data": {}   
}
```

### 修改标签

[PUT] /api/v1/tags/{id}

请求:
```json
{
    "name": "修改后的标签名称",
    "state": 1   
}
```

响应:
```json
{
    "code": 0,
    "msg": "success",
    "data": {}
}
```

### 删除标签

[DELETE] /api/v1/tags/{id}

响应:
```json
{
    "code": 0,
    "msg": "success",
    "data": {}
}
```