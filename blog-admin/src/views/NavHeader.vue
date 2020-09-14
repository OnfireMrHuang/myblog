<template>
    <div>
      <el-row type="flex" justify="center">
      <el-col :span="16">
        <div class="logo">
            Jessie's Blogs
        </div>
      </el-col>
      <el-col :span="8">
        <div class="nav-right">
            <el-menu
            :default-active="activeIndex"
            class="el-menu-demo"
            mode="horizontal"
            @select="handleSelect">
                <el-menu-item index="1">
                    <i class="el-icon-edit"></i>
                    待定内容1
                </el-menu-item>
                <el-menu-item index="2">
                    <i class="el-icon-share"></i>
                    待定内容2
                </el-menu-item>
                <el-menu-item index="3">
                    <i class="el-icon-setting"></i>
                    待定内容3
                </el-menu-item>
            </el-menu>
        </div>
      </el-col>
    </el-row>
  </div>
</template>

<script>
  export default {
    data() {
      return {
        activeIndex: '1',
        activeIndex2: '1'
      };
    },
    methods: {
      handleSelect(key, keyPath) {
        console.log(key, keyPath);
      }
    }
  }
</script>

<style lang="less">
    .logo {
        font-size: 25px;
        font-weight: bold;
        float: left;
        color: #41b883;
    }
    .nav-right {
        float: right;
        line-height: 58px;
    }
</style>