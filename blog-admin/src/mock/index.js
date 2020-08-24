
// 使用 Mock
var Mock = require('mockjs')

// 用户名密码验证
Mock.mock("Login/Token",'post',function(req){
    const obj= {
        success: false,
        token: "",
    };
    let json = JSON.parse(req.body)
    if (json.name == "admin" && json.pass == "123456") {
        obj.success = true
        obj.token = 'testkfhksgkshgshghgbdakfbdkhfahfdkjbgkagkahdkfhakfhahgabdfkafksahdofhakfdakjghfkashfkabfkafbkasdbfkdasfbakfbkasdbfakfbfsgsgfsdfnslfnslgnslgn'
    }
    return obj
})

export default Mock

