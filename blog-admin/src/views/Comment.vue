<template>
  <div>
    <el-row type="flex" justify="space-between">
      <el-col :span="11">
        <el-card>
          <div slot="header" style="color:#41b883;">评论开放规则</div>
          <div>
            <span style="font-size=10px;">状态：</span>
            <el-switch width="40" 
            v-model="status" 
            active-text="开启"
            inactive-text="关闭"
            @on-change="switchChange">
            </el-switch>
            <p style="padding-top:1rem;color:#c0c4cc;">后续增加功能</p>
          </div>
        </el-card>
      </el-col>
      <el-col :span="11">
        <el-card>
          <div slot="header" style="color:#41b883;">作者保留字段设置</div>
          <div>
            <div style="padding-bottom:1rem;color:#c0c4cc;text-align: center;">添加的字段会在评论列表中显示 "作者"</div>
            <div style="display:flex;justify-content: space-between;">
              <el-input type="text" style="width:80%;" placeholder="添加保留字段，最好为英文" v-model="author"></el-input> 
              <el-button type="success" @click="addRules">添加保留字</el-button>
            </div>
            <div style="padding-top:0.5rem;">
              <el-tag v-for="(item, index) in authorList" :key="index">{{item}}</el-tag>
            </div>
          </div>
        </el-card>
      </el-col>
    </el-row>
    <el-divider><span style="color:#c0c4cc;margin:2.5rem 0;">评论列表</span></el-divider>
    <el-row type="flex" justify="space-between">
      <el-col :span="24" v-for="(item, index) in commentTable" :key="index">
        <el-card style="margin-bottom:1rem;">
          <p slot="title">
            <span style="color:#41b883;">{{item.title}}</span>
            <span style="float:right;color:#41b883;">{{item.id}}</span>
          </p>
          <div>
            <el-table border :data="item.comment">
                <el-table-column
                  align='center'
                  prop="username"
                  label="用户名">
                </el-table-column>
                <el-table-column
                  align='center'
                  prop="email"
                  label="邮箱">
                </el-table-column>
                <el-table-column
                  align='center'
                  prop="ip"
                  label="IP地址">
                </el-table-column>
                <el-table-column
                  align='center'
                  prop="time"
                  label="时间">
                </el-table-column>
                <el-table-column
                  align='center'
                  prop="content"
                  label="评论内容">
                </el-table-column>
                <el-table-column
                  align='center'
                  prop="action"
                  label="操作">
                </el-table-column>
            </el-table>
          </div>
        </el-card>
      </el-col>
    </el-row>
    <dialog-del :status.sync="modalStatus" @on-del="modalDel"></dialog-del>
  </div>
</template>
<script>
import DialogDel from '@/components/DialogDel.vue'
export default {
  data () {
    return {
      authorList: [
          "huang",
          "zhang"
      ],
      status: true,
      author: '',
      modalStatus: false,
      // 临时存放时间戳
      commentTime: 0,
      // 临时存放id
      commentId: null,
      commentTable: ["fafafsf"],
      authorConfig: []
    }
  },
  created () {
    this.authorLists()
    this.commentList(1, 4)
    this.commentLists()
  },
  components: {
    DialogDel
  },
  methods: {
    async authorLists () {
      // 调用控制字段列表
      try {
        let {data} = await this.$axios.post('/api/comment/config/list')
        if (Object.is(data.error, 0)) {
          this.authorList = data.data.author
          this.status = data.data.status
        }
      } catch (error) {
        // handle error
      }
    },
    async changeCommentConfig (author, status) {
      // 增加控制字段
      try {
        let {data} = await this.$axios.post('/api/comment/config', {author, status})
        if (Object.is(data.nModified, 1)) {
          this.success('添加成功', '添加保留字段成功', false)
          this.author = ''
          this.authorLists()
        } else if (Object.is(data.nModified, 0) || Object.is(data.ok, 1)) {
          this.warning('添加重复', '重复添加保留字段', false)
        } else {
          this.error('添加失败', '添加保留字段失败', false)
        }
      } catch (error) {
        // handle error
      }
    },
    addRules () {
      this.changeCommentConfig(this.author, this.status)
    },
    switchChange (val) {
      console.log(val)
    },
    async commentList (page, pageSize) {
      // 评论列表
      try {
        let {data} = await this.$axios.post('/api/commentsList', {page, pageSize})
        this.commentTable = data.result
      } catch (error) {
        // handle error
      }
    },
    show (row) {
      console.log(row)
      this.$Modal.info({
        title: '评论信息',
        content: `用户名：${row.username}<br>邮箱：${row.email}<br>评论内容：${row.content}`
      })
    },
    removeComment (id, time) {
      this.modalStatus = true
      this.commentTime = time
      this.commentId = id
    },
    async modalDel () {
      try {
        let {data} = await this.$axios.post('/api/comment/delComment', {id: this.commentId, time: this.commentTime})
        if (Object.is(data.error, 0)) {
          // 成功删除
          this.success('删除成功', `成功删除数量：${data.delCount}`, false)
          // 重新调用列表
          this.commentList(1, 4)
        } else {
          this.error(`错误代码：${data.error}`, `${data.data}`, false)
        }
      } catch (error) {
        this.error(`错误信息`, `${error}`, false)
      }
    },
    async commentLists () {
      try {
        let {data: {data}} = await this.$axios.post(`/api/comment/config/list`)
        this.authorConfig = data.author
      } catch (error) {
        // handle error
      }
    }
  }
}
</script>
