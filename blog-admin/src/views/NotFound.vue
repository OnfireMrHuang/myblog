<template>
  <div class="error_box">
    <el-row type="flex" justify="center" align="middle" class="error_log">
      <!-- 左侧 -->
        <div class="error_logo"></div>
    </el-row>
    <el-row>
        <div class="error_text">未经管理员授权禁止访问授权地址！</div>
        <div class="error_text error_back"><router-link :to="{path:'/'}">返回首页</router-link></div>
    </el-row>
  </div>
</template>
<style lang="less" scoped>
  .error_box {
    margin-top: 100px;
    .error_logo {
      width:800px;
      height:155px;
      background:url('../../static/images/stop.png') no-repeat;
      background-size:cover;
      background-position:center center;
    }
    .error_text {
      text-align:center;
      font-weight:bold;
      font-size:30px;
      padding-top:2rem;
      text-align: center;/*让文字水平居中*/
    }
    .error_back{
      font-size:20px;
      a {
        color:#495060;
        text-decoration: underline
      }
    }
  }
</style>
