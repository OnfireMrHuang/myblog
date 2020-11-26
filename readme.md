# 我的个人博客系统

## 接口定义

### 获取文章信息
**method:**  GET

**url:**   /api/v1/articles?id=文章ID 

>
reponse
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
        "commment": [],
    }
}
```
### 获取文章列表

[GET] /api/v1/articles

```json
{
    "error": 0,
    "count": 6,
    "list": [
        {
            "comment": [], // 评论列表
            "id": "文章ID",
            "title": "文章标题",
            "time": "时间",
            "des": "文章描述",
        }
    ]
}
```

### 发布评论

POST /api/comment

req:
```json
{
    "username":"用户名",
    "email": "邮箱",
    "content": "评论内容",
    "time": "时间戳",
    "ip": "IP地址"
}
```

rsp:
```json
{
  "status": "0000",
  "success": 1,
  "result": {
    "n": 1,
    "nModified": 1,
    "ok": 1
  }
}
```

#### 获取文章评论

POST /api/articleComments

req:
```json
{
    "id": "文章ID",
}
```

rsp:
```json
{
    "error": 0,
    "count": 13,
    "result": {
        "id": "文章ID号",
        "title": "文章标题",
        "comment" : [
            {
                "username": "xixi",
                "email": "2292553208@qq.com",
                "content": "测试一下",
                "time": 1604745002246,
                "ip": "::ffff:172.25.0.1"
            }
        ]
    }
}
```