<template>
  <div>
    <el-row type="flex" justify="start">
      <el-col :span="12" index="1">
        <span>版本号：</span>
        <el-input v-model="version" size="medium" placeholder="输入版本号，例如:1.0.0"></el-input>
      </el-col>
    </el-row>
    <el-row type="flex" justify="center">
      <el-col :span="22">
        <!-- <mavon-editor @change="changeContent" class="version_content" v-model="content" fontSize="18px" placeholder="# 发布版本内容" style="min-height:600px;" /> -->
        <el-button type="success" class="version_button" @click="sublimtContent">发布新版本</el-button>
      </el-col>
    </el-row>
  </div>
</template>

<script>
export default {
  data () {
    return {
      version: '1.0.0',
      content: '',
      htmlContent: ''
    }
  },
  methods: {
    changeContent (content, render) {
      this.htmlContent = render
    },
    sublimtContent () {
      let json = {
        version: this.version,
        content: this.htmlContent
      }
      this.$axios.post('/api/version/insert', json).then(res => {
        let {error} = res.data
        if (Object.is(error, 0)) {
          this.success('文章发布成功', '', true)
            [this.version, this.content, this.htmlContent] = ['']
        } else {
          this.error('发布失败', '未知原因', false)
        }
      })
    }
  }
}
</script>
<style lang="less">
  .version_content {
    margin-top:2rem;
  }
  .version_button {
    float:right;
    margin-top:2rem;
  }
</style>
