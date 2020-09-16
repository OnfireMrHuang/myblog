<template>
  <div>
    <el-menu
      default-active="2"
      class="el-menu-vertical-demo"
      @select="getIndex"
      @open="handleOpen"
      @close="handleClose">
      <el-submenu index="1">
        <template slot="title">
          <i class="el-icon-s-tools"></i>
          <span>控制面板</span>
        </template>
        <el-menu-item index="admin"><i class="el-icon-info"></i>系统概览</el-menu-item>
        <el-menu-item index="article"><i class="el-icon-reading"></i>发布文章</el-menu-item>
        <el-menu-item index="articleList"><i class="el-icon-files"></i>文章列表</el-menu-item>
      </el-submenu>
      <el-submenu index="2">
         <template slot="title">
          <i class="el-icon-coin"></i>
          <span>版本控制</span>
        </template>
        <el-menu-item index="version"><i class="el-icon-edit-outline"></i>发布版本</el-menu-item>
      </el-submenu>
      <el-submenu index="3">
        <template slot="title">
          <i class="el-icon-chat-dot-square"></i>
          <span>评论管理</span>
        </template>
        <el-menu-item index="comment"><i class="el-icon-chat-line-square"></i>发布版本</el-menu-item>
      </el-submenu>
    </el-menu>
</div>
</template>

<script>
  export default {
    methods: {
      handleOpen(key, keyPath) {
        console.log(key, keyPath);
      },
      handleClose(key, keyPath) {
        console.log(key, keyPath);
      },
      getIndex(key, keyPath) {
        this.$router.push({name:key})
        // console.log("key:",key,"key path:", keyPath);
      }
    }
  }
</script>