<template>
<div>
<el-row>
  <el-col span="19">
    <label for="title" class="article">撰写新文章</label>
    <el-input v-model="title" size="medium" placeholder="在此输入文章标题" name="title" class="article_title"></el-input>
    <label for="title" class="article">文章简介</label>
    <el-input v-model="des" size="medium" placeholder="在此输入文章简介" name="title" class="article_title"></el-input>
    <mavon-editor ref="mavonEditor" @change="changeContent" class="article_content" v-model="content" fontSize="18px" @imgDel="imgDel" @imgAdd="imgAdd" placeholder="开始编写文章内容(上传图片为七牛云，删除功能暂无)..." style="min-height:600px;" />
    <el-button type="success" class="article_button" @click="submitArticle">发布文章</el-button>
  </el-col>
  <el-col span="4" offset="1" class="content_right">
    <el-card class="box-card">
      <label for="date" class="article" slot="header">发布日期</label>
        <el-date-picker 
        @on-change="dateContent" 
        :v-model="date" 
        type="date" 
        name="date" 
        size="large" 
        class="data_picker" 
        placeholder="选择日期" 
        style="width:100%;">
        </el-date-picker>
    </el-card>
    <br><br>
    <el-card class="box-card">
        <p slot="header">分类目录</p>
         <el-radio-group v-model="radio">
            <el-radio :label="Back">
                <span class="list_menu">后端开发</span>
            </el-radio>
        </el-radio-group>
    </el-card>
  </el-col>
</el-row>
</div>
</template>

<script>
export default {
    data () {
    return {
        title: '',
        content: '',
        htmlContent: '',
        date: FormatDate(new Date()),
        radio: 'Front',
        des: '',
        originalContent: '',
        uploadToken: ''
    }
  },
}
</script>
<style lang="less">
  @import './../assets/css/admin/article.less';
</style>
