<template>
  <div>
  	<nav-header :active="active"></nav-header>
	<el-row type="flex" justify="center">
		<el-col :span="10">
			<h1>更新日志</h1>
		</el-col>
	</el-row>
	<el-row type="flex" class="version" justify="center" v-for="item in list" :key="item._id">
		<el-col :span="10">
			<el-card class="box-card">
			  <div slot="header" class="clearfix">
			    <span style="font-weight:bold;">{{item.version}}</span>
			    <span style="float:right">{{item.time}}</span>
			  </div>
			  <div v-html="item.content">
                  {{item.content}}
              </div>
			</el-card>
		</el-col>
	</el-row>
  </div>
</template>

<script>
import NavHeader from '~/components/NavHeader.vue'
// import {baseurl} from '~/plugins/url.js'
export default {
    data() {
        return {
            active:'version'
        }
    },
    async asyncData({app}) {
        // let result = await app.$axios.get(`${baseurl}/api/version/getVersion`);
		// let {error,list} = result.data;
		let list = [
			{
				version:"1.0.0",
				time: "2020-08-24",
				content: " 第一个版本",
			},
			{
				version:"2.0.0",
				time: "2020-08-24",
				content: " 第二个版本",
			}
		]
        return {list}
    },
    components:{
        NavHeader
    },
    head() {
		return {
			title:"更新日志-Jessie‘s Blog",
			meta:[
				{hid:'description',name:'description',content:''},
				{hid:'keywords',name:'keywords',content:''},
				{hid:'author',content:'jessie'}
			]
		}
	}
}
</script>

<style>
    .version {
        margin-bottom:2rem;
    }
</style>
