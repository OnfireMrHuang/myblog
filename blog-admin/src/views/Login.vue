<template>
  <div class="login_images">
    <el-row type="flex" justify="center">
      <el-col :span="8" style="margin-top:10%;">
        <el-card class="box-card" shadow="always">
            <div slot="header" class="clearfix">
                <span>欢迎登陆!</span>
            </div>
            <el-form ref="loginForm" :model="user" :rules="rules" status-icon label-width="100px">
                <el-form-item label="账号" prop="name">
                    <el-input v-model="user.name"></el-input>
                </el-form-item>
                <el-form-item label="密码" prop="pass">
                    <el-input v-model="user.pass" type="password"></el-input>
                </el-form-item>
                <el-form-item>
                    <el-button type="primary" style="margin-left:25%" @click="login">登陆</el-button>
                </el-form-item>
            </el-form>
        </el-card>
      </el-col>
    </el-row>
  </div>
</template>

<script>

export default {
    data() {
        return {
            user: {},
            rules: {
                name: [{ required: true,message: "用户名不能为空", trigger: "blur"}],
                pass: [{ required: true,message: "密码不能为空", trigger: "blur"}]
            }
        }
    },
    created() {
        if (window.localStorage.Token&&window.localStorage.Token.length>=128) {
            this.isLogin = true;
        }
    },
    methods: {
        login() {
            let that = this;
            that.$store.commit("saveToken","") // 清理token
            this.$refs.loginForm.validate(valid=>{
                if (valid) {
                    this.$api.post(
                        "Login/Token",
                        { name: that.user.name, pass: that.user.pass },
                        r => {
                            if(r.success) { // 如果响应成功
                                var token = r.token;
                                that.$store.commit("saveToken",token) // 将获取的token保存下来
                                this.$notify({
                                    type: "success",
                                    message: "欢迎你," + this.user.name + "!",
                                    duration: 3000
                                });
                                this.$router.replace("/")
                            } else { // 否则就是失败
                                this.$message({
                                    type: "error",
                                    message: "用户名或密码错误",
                                    showClose: true
                                });
                            }
                        }       
                    )
                } else {
                    return false;
                }
            });
        },
        logout() {
            this.isLogin = false;
            this.$store.commit("saveToken",""); // 清理token
        }
    }
}
</script>

<style lang="less">
  .login_images {
    width:100%;
    height:100vh;
    background:url('../../static/images/login.jpg') no-repeat;
    background-size:cover;
    background-position: center center;
  }
</style>