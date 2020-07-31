<template>
  <div>
    <el-image style="top: 0px;"
              :src=logoimg>
    </el-image>
    <el-row>
      <el-col :xs="12"
              :sm="12">
        <el-input placeholder="请输入内容"
                  v-model="searchInfo.softName">
        </el-input>
      </el-col>
      <el-col :xs="12"
              :sm="12">
        <el-button type="primary"
                   @click="onSubmit">搜索</el-button>
      </el-col>
    </el-row>
    <el-table :data="tableData"
              :fit="true"
              :show-header="true">
      <el-table-column prop="softImg"
                       label="图标"
                       width="60">
        <template slot-scope="scope">
          <img alt
               class="logoimg"
               :src=scope.row.softImg />

        </template>
      </el-table-column>
      <el-table-column prop="softName"
                       label="软件名"
                       width="120">
      </el-table-column>
      <el-table-column label="描述"
                       prop="softDescription">
      </el-table-column>
      <el-table-column label="操作"
                       fixed="right"
                       width="200px">
        <template slot-scope="scope">
          <el-button mc-type="column-el-button"
                     @click="downloadSoft(scope.row)"
                     size="mini"
                     type="primary">下载</el-button>

        </template>
      </el-table-column>

    </el-table>
  </div>
</template>
<script>
import {
  findNielsenSoftware,
  getNielsenSoftwareList,
} from '@/api/nielsenSoftware' //  此处请自行替换地址
import { formatTimeToStr } from '@/utils/data'
import infoList from '@/components/mixins/infoList'
const path = process.env.VUE_APP_BASE_API
const imgPath = process.env.VUE_APP_IMGPRE
export default {
  name: 'NielsenSoftware',
  mixins: [infoList],
  data() {
    return {
      listApi: getNielsenSoftwareList,
      filePath: imgPath,
      logoimg: require('../../assets/banner.png'),
    }
  },

  methods: {
    downloadSoft(row) {
      debugger
      window.location.href = row.download
    },
    changedir() {
      return 1
    },
    onSubmit() {
      debugger
      this.page = 1
      this.pageSize = 10
      this.getTableData()
    },
  },
  created() {
    this.getTableData()
  },
}
</script>
<style scoped>
.logoimg {
  width: 30px;
  height: 30px;
  vertical-align: middle;
  background: #00adee;
  border-radius: 50%;
  padding: 3px;
}
</style>