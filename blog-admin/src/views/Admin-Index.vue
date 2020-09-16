<template>
  <div>
      <el-row type="flex" justify="space-around">
            <el-col span="9">
                <el-card>
                    <div slot="header" style="color: #41b883;">服务器状态</div>
                    <div class="server_status">
                      <span>运行状态：</span>
                      <el-tag type="dot" :color="Object.is(constants, '0') ? 'success' : 'error'">{{Object.is(constants, '0') ? '服务器运行中' : '服务器出现错误'}}</el-tag>
                    </div>
                    <div class="server_status">
                        <span>服务器发行版本：</span>
                        <el-tag size="small" checkable color="blue">{{release}}</el-tag>
                    </div>
                </el-card>
            </el-col>
            <el-col span="11">
                <el-card>
                    <div slot="header" style="color: #41b883;">服务器信息</div>
                    <a href="#" slot="extra" @click.prevent="changeInit">
                        <i type="ios-loop-strong"></i>
                        刷新
                    </a>
                    <el-row type="flex" justify="space-between">
                        <el-col span="12">
                            <div class="server_status">
                                <span>服务器主机名：</span>
                                <el-tag size="small" color="#495060">{{hostname}}</el-tag>
                            </div>
                            <div class="server_status">
                                <span>操作系统：</span>
                                <el-tag size="small" color="#ff9900">{{type}}</el-tag>
                            </div>
                            <div class="server_status">
                                <span>服务器总内存数：</span>
                                <el-tag size="small" color="#19be6b">{{totalmem}}</el-tag>
                            </div>
                            <div class="server_status">
                                <span>服务器可用内存数：</span>
                                <el-tag size="small" color="#19be6b">{{freemem}}</el-tag>
                            </div>
                        </el-col>
                        <el-col span="11">
                            <div>
                                <el-progress type="circle" 
                                :percentage="percentage"
                                :stroke-linecap="square"
                                :stroke-width="8"
                                :width="200">
                                </el-progress>
                            </div>
                        </el-col>
                    </el-row>
                </el-card>
            </el-col>
      </el-row>
      <el-row type="flex" justify="space-around" style="margin-top:4rem;">
          <el-col :span="22">
            <el-card>
              <div slot="header" style="color: #41b883;">CPU信息</div>
              <div class="cpu_status">
                  <span>逻辑CPU内核信息：</span>
              </div>
              <el-table border stripe :data="cpu">
                   <el-table-column
                    prop="model"
                    label="CPU内核模型"
                    width="180">
                </el-table-column>
                <el-table-column
                    prop="speed"
                    label="CPU频率(GHz)"
                    width="180">
                </el-table-column>
                <el-table-column
                    prop="times"
                    label="CPU执行模式[毫秒]( user:用户 | nice:良好 | sys:系统 | idle:空闲 | irq:中断 )">
                </el-table-column>
              </el-table>
            </el-card>
          </el-col>
      </el-row>
  </div>
</template>

<script>
export default {
  data () {
    return {
      constants: null,
      release: null,
      platform: null,
      hostname: null,
      type: null,
      freemem: null,
      totalmem: null,
      percentage: "25",
      cpu: [],
    }
  },
  created () {
    this.init()
    console.log(process.env.NODE_ENV)
  },
  methods: {
    init () {
      this.$axios.post('/api/system').then(res => {
        let { constants, release, platform, hostname, type, freemem, totalmem, percentage, cpu } = res.data;
        [this.constants, this.release, this.platform, this.hostname, this.type, this.freemem, this.totalmem, this.percentage, this.cpu] = [constants, release, platform, hostname, type, freemem, totalmem, percentage, cpu]
      })
    },
    changeInit () {
      this.init()
    }
  }
}
</script>

<style lang="less">
 @import './../assets/css/admin/index.less';
</style>
