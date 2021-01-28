<template>
  <div>
      <nav-header :active="active"></nav-header>
      <el-row type="flex" justify="center">
        <el-col :span="14" class="detail_title">
            <div>{{title}}</div>
            <div class="time">发布时间：{{time}}&nbsp;&nbsp;&nbsp;&nbsp;</div>
        </el-col>

      </el-row>
      <el-row type="flex" justify="center">
        <el-col :span="14" class="detail_content">
            <el-card>
                <div v-show="!content">暂无文章数据...</div>
                <div v-html="content" class="md markdown-body"></div>
            </el-card>
        </el-col>
      </el-row>
  </div>
</template>

<script>
import NavHeader from '~/components/NavHeader.vue';
import {baseurl} from '~/plugins/url.js';
export default {
	data() {
		return {
            active:'Backarticle'
        }
	},
	async asyncData({app,params}) {
        let {data} = await app.$axios.get(`${baseurl}/api/v1/articles/${params.id}`);
        // let {code,msg} = data;
		if(data.code != 200) {
			console.log("错误信息: ",data.msg)
			return {}
        }
        let title = data.data.title
        let des = data.data.desc
        let content = data.data.content
        let time = data.data.modified_by
		return {title,des,content,time}
	},
    head() {
		return {
			title:this.title,
            meta:[
				{hid:'description',name:'description',content:`${this.des}`},
				{hid:'author',content:'brian'}
			]
		}
	},
    components:{
        NavHeader
    }
}
</script>
<style lang="less">
    @import './../../assets/css/Index/Detail.less';
</style>
