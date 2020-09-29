
// 使用 Mock
var Mock = require('mockjs')

/**
 * POST /login/token
 * req : {
 *      "name" : 用户名,
 *      "pass" : 密码
 * }
 * 
 * rsp : {
 *      "success" : 是否成功
 *      "token" : token值
 *      "msg" : 备注消息
 * }
 */
Mock.mock("login/token",'post',function(req){
    const obj= {
        success: false,
        token: "",
    };
    let json = JSON.parse(req.body)
    if (json.name == "admin" && json.pass == "123456") {
        obj.success = true
        obj.token = 'e0c9035898dd52fc65c41454cec9c4d2611bfb37'
    }
    return obj
})

/**
 * GET /api/system
 * 
 * rsp : {
 *      "exception" : int //服务器是否异常
 *      "exception_info" : string // 异常说明
 *      "platform" : string 操作系统
 *      "hostname" : string 主机名
 *      "freemem" : string 剩余可用内存数
 *      "totalmem" : string 总内存
 *      "percentage" : int 百分比
 *      "cpu" : array [
 *          {
 *              "model": string
 *              "speed": string
 *              "times": string
 *          }
 *      ]
 * }
 */
Mock.mock("/api/system",'get',function(){
    const obj= {
        exception: 1,
        exception_info: "网络出错",
        platform: "linux",
        hostname: "127.0.0.1",
        freemem: "2G",
        totalmem: "8G",
        percentage: 40,
        cpu: [
            {
                model: "Intel(R) Xeon(R) Platinum 8163 CPU @ 2.50GHz ",
                speed: "2.50GHz",
                times: "test"
            }
        ]
    };
    return obj
})

export default Mock

