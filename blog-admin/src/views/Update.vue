<template>
<div>
  <el-row>
  <!-- 左侧 -->
  <el-col span="19">
    <i class="el-icon-close"></i>
    <label for="title" class="article">撰写新文章</label>
    <el-input v-model="collections.title" 
    size="medium" 
    placeholder="在此输入文章标题" 
    name="title" 
    class="article_title"></el-input>
    <label for="title" class="article">文章简介</label>
    <el-input v-model="collections.des" 
    size="medium" 
    placeholder="在此输入文章标题" 
    name="title"
     class="article_title"></el-input>
    <el-row>
        <el-col span="10" style="margin-bottom:20px;">
            <el-card class="box-card">
              <div v-if="img.path">
                <el-popconfirm @on-ok="confirmDelete" confirm title="您确定要删除么？" style="width:100%;">
                <i class="icon iconfont icon-guanbi"></i>
                </el-popconfirm>
                <div style="text-align:center">
                  <img :src="img.path" alt="" style="background-size:cover;max-width:100%;height:128px;">
                  <h3>{{img.filename}}</h3>
                </div>
              </div>
              <div v-else>
                <h3 style="text-align: center;">上传Banner图</h3>
                <el-upload multiple type="drag" 
                :body-style="{ padding: '20px' }"
                :on-success="uploadSuccess" 
                :on-error="uploadError" 
                :data="{id: $route.params.id, radio: collections.radio}" 
                action="http://localhost:3000/api/upload" 
                :show-upload-list="false" 
                :format="['jpg','jpeg','png']">
                    <div style="float:center;text-align=center;font-size=80;"><i class="el-icon-upload" style="color: #3399ff"></i></div>
                    <div style="float:center;text-align=center;">点击上传或拖拽</div>
                </el-upload>
              </div>
            </el-card>
          </el-col>
        </el-row>
    <mavon-editor ref="mavonEditor" @change="changeContent" @imgDel="imgDel" @imgAdd="imgAdd" @save="save" class="article_content" v-model="collections.content" fontSize="18px" placeholder="开始编写文章内容..." style="min-height:600px;" />
    <el-button type="warning" class="article_button" @click="submitArticle">修改文章</el-button>
  </el-col>
  <!-- 右侧 -->
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
      collections: {
        title: '',
        content: '',
        htmlContent: '',
        date: FormatDate(new Date()),
        radio: '',
        contentValue: '',
        des: '',
        original: '',
        id: ''
      },
      img: {
        path: '',
        filename: ''
      },
      uploadToken: ''
    }
  },
  created () {
    this.init()
    this.getUploadToken()
  },
  methods: {
    init () {
      let id = this.$route.params.id
      this.$axios.get('/api/article/update', {params: {id}}).then(res => {
        let {data: [{_id, title, des, original, time, list}]} = res
        console.log(res.data)
        Object.assign(this.collections, {id: _id, title, des, content: original, date: time, radio: list})
        this.defaultRequest()
      })
    },
    uploadSuccess (file) {
      this.success(`上传成功`, `上传banner图成功`, false)
      Object.assign(this.img, file)
    },
    uploadError (error, file) {
      this.error(`出现错误`, `错误信息：${error},${file}`, false)
    },
    changeContent (value, render) {
      this.collections.htmlContent = render
      this.collections.original = value
    },
    async confirmDelete () {
      /* 删除 */
      try {
        let {data: {status, result: {nModified}}} = await this.$axios.post('/api/deleteFile', {id: this.$route.params.id, radio: this.collections.radio})
        if (Object.is(status, 200)) {
          this.defaultRequest()
          this.success(`删除成功`, `删除${nModified}个文件`, false)
        }
      } catch (error) {
        this.error(`发生错误`, `${error}`, false)
      }
    },
    async defaultRequest () {
      /* 获取显示图片 */
      let {data: {result}} = await this.$axios.post('/api/findOneArticle', {id: this.$route.params.id, radio: this.collections.radio})
      if (Object.is(result.banner, undefined)) {
        Object.assign(this.img, {path: '', filename: ''})
      } else {
        Object.assign(this.img, {path: result.banner, filename: result.imgFileName})
      }
    },
    save (value, render) {
      this.collections.htmlContent = render
      this.collections.original = value
      this.submitArticle()
    },
    submitArticle () {
      if (Object.is(this.title, '')) {
        this.error('文章标题留空无法保存', '请仔细检查文章标题', false)
      } else {
        this.$axios.post(`/api/article/insert${this.collections.radio}`, this.collections).then(res => {
          let {error} = res.data
          console.log(res.data)
          if (Object.is(error, 0)) {
            this.success('修改成功', '修改成功', false)
          } else {
            this.error('修改失败', '未知原因', false)
          }
        })
      }
    },
    dateContent (val) {
      this.date = FormatDate(val)
    },
    async getUploadToken () {
      try {
        let result = await this.$axios.post('/api/article/getToken')
        this.uploadToken = result.data
      } catch (error) {
        this.error(error, error, false)
      }
    },
    imgAdd (pos, file) {
      var formdata = new FormData()
      formdata.append('token', this.uploadToken)
      formdata.append('file', file)
      this.$axios({
        url: '/api/article/upload',
        method: 'post',
        data: formdata,
        headers: {
          'Content-Type': 'multipart/form-data',
          'Accept': '*/*'
        }
      }).then(res => {
        this.$refs.mavonEditor.$img2Url(pos, res.data.img)
      })
    },
    imgDel (pos, file) {
      /* 删除预留 */
      console.log(pos)
      console.log(file)
    }
  }
}
/* 封装格式化日期 */
function FormatDate (strTime) {
  var date = new Date(strTime)
  return `${date.getFullYear()}年${date.getMonth() + 1}月${date.getDate()}日`
}
</script>
<style>
    @import './../assets/css/admin/article.less';
</style>
