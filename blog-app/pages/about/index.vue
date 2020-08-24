<template>
<div>
	<nav-header :active="active"></nav-header>
	<el-row type="flex" justify="center" class="about_content">
		<el-col :span="15" class="about">
			<el-card class="box-card">
				<h1>About Blog</h1>
				<hr class="line"/>
    			<p>简单介绍下博客使用技术：</p>
    			<p class="about_title">博客前端：</p>
    			<p class="about_text"><b>采用服务器端渲染(ssr)</b>,基于 Vue.js 的通用应用框架:<b>Nuxt.js + Element-UI + @nuxtjs/axios + Less</b></p>
    			<p class="about_title">博客后端:</p>
    			<p class="about_text">UI层：使用 <b>Vue + Vue-Router + Vue-cli + Vue + Axios + Less + Iview</b></p>
    			<p class="about_text">服务层：使用 <b>Go + Gin + MonogoDB + Docker</b></p>
    			<h1>About Me</h1>
				<hr class="line"/>
				<p class="about_title">个人简介</p>
				<ul>
					<li>姓名：黄文武</li>
					<li>英文名：Jessie</li>
					<li>爱好：编程、户外</li>
					<li>座右铭：乐观是保持不断前行必需的营养</li>
				</ul>
				<p class="about_title">技能</p>
				<ul>
					<li>前端：HTML、CSS、Javacript、Vue、微信小程序开发</li>
					<li>UI框架：Element-UI</li>
					<li>后端：GO、C/C++、PHP、Mysql、Redis、Kafaka、MongoDB数据库、Docker、K8s 等等</li>
				</ul>
				<p class="about_title">个人评价</p>
				<p class="about_text">能从管理、业务主营、运行效率多个角度考虑问题，专注编写代码。</p>
				<p class="about_text">未来1-3年规划：熟练掌握前端开发，保持对产品的快速实现;</p>
                <p class="about_text">深入Go在物联网、大数据、网络安全的研究开发；学习熟悉金融和实际操作手段。</p>
				<p class="about_title">联系方式</p>
				<p class="about_text">邮箱：engineer_mrhuang@163.com</p>
			</el-card>
		</el-col>
	</el-row>
</div>
</template>

<script>
import NavHeader from '@/components/NavHeader.vue'
export default {
	data() {
		return {
			active:'about'
		}
	},
	components:{
		NavHeader
	},
	head() {
		return {
			title:"个人简介-Jessie's Blog",
			meta:[
				{hid:'description',name:'description',content:''},
				{hid:'keywords',name:'keywords',content:''},
				{hid:'author',content:'Jessie'}
			]
		}
	}
}
</script>

<style lang="less" scoped>
	@import './../../assets/css/Index/About.less';
</style>
