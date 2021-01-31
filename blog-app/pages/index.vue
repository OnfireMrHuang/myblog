<template>
<div>
	<nav-header :active="active"></nav-header>
	<!-- 左侧文章内容 -->
	<el-row type="flex" justify="center" class="content-blog">
		<el-col :span="10">
			<nuxt-link  v-for="item in list" :key="item.id" :to="{name:'article-id',params:{id:item.id}}" class="box-href">
				<el-card class="box-card" shadow="hover">
					<h2 class="box-title">{{item.title}}</h2>
					<div class="box-icon">
						<span><i class="el-icon-date"></i>&nbsp;{{item.time}}</span>
					</div>
					<div class="box-content">{{item.des}}</div>
				</el-card>
			</nuxt-link>
			<el-pagination class="pagination" @current-change="pagination" background layout="prev, pager, next" :page-size="5" :total="count" v-show="count >= 5"></el-pagination>
		</el-col>
		<!-- 右侧关于我 -->
        <el-col :span="5" :offset="1">
			<el-card class="about">
				<div class="about-title">about Me</div>
				<div class="about-content">
					<p>网名：Jessie</p>

					<p>职业：Go后端工程师</p>

					<p>邮箱：engineer_mrhuang@163.com</p>
				</div>
			</el-card>
			<!-- 近期文章开始 -->
			<el-card class="article">
				<div class="article-title">近期文章</div>
				<hr>
				<nuxt-link v-for="item in lately" :key="item.id" :to="{name:'article-id',params:{id:item.id}}" class="article-link">
					<i class="el-icon-edit"></i>&nbsp;&nbsp;{{item.title}}
				</nuxt-link>
			</el-card>
			<!-- 近期文章结束 -->

			<!-- 友情链接开始 -->
			<el-card class="link">
				<div class="link-title">友情链接</div>
				<hr>
				<div class="link-content">
					<a href="/" target="_blank" class="link-url">虚位以待</a>
					<a href="/" target="_blank" class="link-url">虚位以待</a>
				</div>
			</el-card>
			<!-- 友情链接结束 -->
		</el-col>
	</el-row>
	<nav-footer />
</div>
</template>

<script>
import NavHeader from '~/components/NavHeader.vue'
import NavFooter from '~/components/Footer.vue'
import {baseurl} from '~/plugins/url.js'
export default {
	data() {
		return {
			active:'index',
			list: [],
			count: 5,
			lately: 1,
		}
	},
	async asyncData({app}) {
		// 服务器端渲染数据
		let json = {page:1,limit:5}
		let {data} =await app.$axios.get(`${baseurl}/api/v1/articles`,{params:json});
		let {code,msg,total,list} = data;
		if(code != 200) {
			console.log("错误信息: ",msg)
			return {}
		}
		return {list: list,count:total}
	},
	methods: {
		pagination(page) {
			let json = {page,limit:5}
			this.$axios.get(`${baseurl}/api/v1/articles`,{params:json},{headers: { 'Access-Control-Allow-Origin': '*' }}).then(res=>{
				let {code,msg,total,list} = res.data;
				if(code != 200) {
					console.log("查询分页报错: ",msg)
					return 
				}
				console.log(total)
				console.log(list)
				this.count = total
				this.list =list;
			});
		}
	},
	components: {
		NavHeader,
		NavFooter
	},
	head() {
		return {
			title:'Jessie的个人博客',
			meta:[
				{hid:'description',name:'description',content:''},
				{hid:'keywords',name:'keywords',content:''},
				{hid:'author',content:'jessie'}
			]
		}
	}
}
</script>

<style lang="less">
    @import './../assets/css/Index/Content.less';
</style>