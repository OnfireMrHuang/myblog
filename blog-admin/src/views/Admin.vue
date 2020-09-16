<template>
  <div>
    <nav-header></nav-header>
    <el-row type="flex" justify="space-between" style="min-width:1200px;">
        <el-col :span="3">
            <nav-left class="navleft"></nav-left>
        </el-col>
        <el-col :span="21" class="routerView">
            <router-view></router-view>
        </el-col>
    </el-row>
  </div>
</template>
<script>
import NavHeader from '@/views/NavHeader'
import NavLeft from '@/views/NavLeft'
export default {
  data () {
    return {
      value: ''
    }
  },
  components: {
    NavHeader,
    NavLeft
  }
}
</script>
<style lang="less">
    @import './../assets/css/admin/style.less';
    body {
        min-width:1200px;
    }
</style>
