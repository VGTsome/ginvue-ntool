<template>
  <div>
    <div class="search-term">
      <el-form :inline="true"
               :model="searchInfo"
               class="demo-form-inline">
        <el-form-item>
          <el-button @click="onSubmit"
                     type="primary">查询</el-button>
        </el-form-item>
        <el-form-item>
          <el-button @click="openDialog"
                     type="primary">新增Nielsen软件表</el-button>
        </el-form-item>
        <el-form-item>
          <el-popover placement="top"
                      v-model="deleteVisible"
                      width="160">
            <p>确定要删除吗？</p>
            <div style="text-align: right; margin: 0">
              <el-button @click="deleteVisible = false"
                         size="mini"
                         type="text">取消</el-button>
              <el-button @click="onDelete"
                         size="mini"
                         type="primary">确定</el-button>
            </div>
            <el-button icon="el-icon-delete"
                       size="mini"
                       slot="reference"
                       type="danger">批量删除</el-button>
          </el-popover>
        </el-form-item>
      </el-form>
    </div>
    <el-table :data="tableData"
              @selection-change="handleSelectionChange"
              border
              ref="multipleTable"
              stripe
              style="width: 100%"
              tooltip-effect="dark">
      <el-table-column type="selection"
                       width="55"></el-table-column>
      <el-table-column label="日期"
                       width="180">
        <template slot-scope="scope">{{scope.row.CreatedAt|formatDate}}</template>
      </el-table-column>

      <el-table-column label="softName字段"
                       prop="softName"
                       width="120"></el-table-column>

      <el-table-column label="softDescription字段"
                       prop="softDescription"
                       width="120"></el-table-column>

      <el-table-column label="download字段"
                       prop="download"
                       width="120"></el-table-column>

      <el-table-column label="softImg字段"
                       prop="softImg"
                       width="120"></el-table-column>

      <el-table-column label="version字段"
                       prop="version"
                       width="120"></el-table-column>

      <el-table-column label="按钮组">
        <template slot-scope="scope">
          <el-button @click="updateNielsenSoftware(scope.row)"
                     size="small"
                     type="primary">变更</el-button>
          <el-popover placement="top"
                      width="160"
                      v-model="scope.row.visible">
            <p>确定要删除吗？</p>
            <div style="text-align: right; margin: 0">
              <el-button size="mini"
                         type="text"
                         @click="scope.row.visible = false">取消</el-button>
              <el-button type="primary"
                         size="mini"
                         @click="deleteNielsenSoftware(scope.row)">确定</el-button>
            </div>
            <el-button type="danger"
                       icon="el-icon-delete"
                       size="mini"
                       slot="reference">删除</el-button>
          </el-popover>
        </template>
      </el-table-column>
    </el-table>

    <el-pagination :current-page="page"
                   :page-size="pageSize"
                   :page-sizes="[10, 30, 50, 100]"
                   :style="{float:'right',padding:'20px'}"
                   :total="total"
                   @current-change="handleCurrentChange"
                   @size-change="handleSizeChange"
                   layout="total, sizes, prev, pager, next, jumper"></el-pagination>

    <el-dialog :before-close="closeDialog"
               :visible.sync="dialogFormVisible"
               title="弹窗操作">
      <el-form ref="elForm"
               :model="formData"
               :rules="rules"
               size="medium"
               label-width="100px">
        <el-form-item label="软件名称"
                      prop="softName">
          <el-input v-model="formData.softName"
                    placeholder="请输入软件名称"
                    clearable
                    :style="{width: '100%'}">
          </el-input>
        </el-form-item>
        <el-form-item label="软件描述"
                      prop="softDescription">
          <el-input v-model="formData.softDescription"
                    type="textarea"
                    placeholder="请输入软件描述"
                    :autosize="{minRows: 4, maxRows: 4}"
                    :style="{width: '100%'}"></el-input>
        </el-form-item>
        <el-form-item label="上传软件"
                      prop="download"
                      required>
          <el-upload ref="download"
                     :file-list="downloadfileList"
                     :action="downloadAction"
                     :before-upload="downloadBeforeUpload"
                     :headers="headinfo"
                     :on-success="updatefilelist"
                     accept="application/x-zip-compressed">
            <el-button size="small"
                       type="primary"
                       icon="el-icon-upload">点击上传</el-button>
          </el-upload>
        </el-form-item>
        <el-form-item label="软件图片"
                      prop="soft_img">
          <el-upload ref="soft_img"
                     :file-list="soft_imgfileList"
                     :action="soft_imgAction"
                     :headers="headinfo"
                     :on-success="updateimglist"
                     accept="image/*"
                     :before-upload="soft_imgBeforeUpload"
                     name="file">
            <el-button size="small"
                       type="primary"
                       icon="el-icon-upload">点击上传</el-button>
          </el-upload>
        </el-form-item>
        <el-form-item label="软件版本"
                      prop="version">
          <el-input v-model="formData.version"
                    placeholder="请输入软件版本"
                    clearable
                    :style="{width: '100%'}">
          </el-input>
        </el-form-item>
      </el-form>
      <div class="dialog-footer"
           slot="footer">
        <el-button @click="closeDialog">取 消</el-button>
        <el-button @click="enterDialog"
                   type="primary">确 定</el-button>
      </div>
    </el-dialog>
  </div>
</template>

<script>
import {
  createNielsenSoftware,
  deleteNielsenSoftware,
  deleteNielsenSoftwareByIds,
  updateNielsenSoftware,
  findNielsenSoftware,
  getNielsenSoftwareList
} from '@/api/nielsenSoftware' //  此处请自行替换地址
import { formatTimeToStr } from '@/utils/data'
import infoList from '@/components/mixins/infoList'
const path = process.env.VUE_APP_BASE_API
import { store } from '@/store/index'
const token = store.getters['user/token']
const user = store.getters['user/userInfo']
export default {
  name: 'NielsenSoftware',
  mixins: [infoList],
  data() {
    return {
      headinfo: {
        'x-token': token,
        'x-user-id': user.ID
      },
      listApi: getNielsenSoftwareList,
      dialogFormVisible: false,
      visible: false,
      type: '',
      deleteVisible: false,
      multipleSelection: [],
      formData: {
        softName: null,
        softDescription: null,
        download: null,
        softImg: null,
        version: null
      },
      rules: {
        softName: [
          {
            required: true,
            message: '请输入软件名称',
            trigger: 'blur'
          }
        ],
        softDescription: [],
        version: [
          {
            required: true,
            message: '请输入软件版本',
            trigger: 'blur'
          }
        ]
      },
      downloadAction: path + '/fileUploadAndDownload/upload?keepname=1',
      downloadfileList: [],
      soft_imgAction: path + '/fileUploadAndDownload/upload',
      soft_imgfileList: []
    }
  },
  filters: {
    formatDate: function(time) {
      if (time != null && time != '') {
        var date = new Date(time)
        return formatTimeToStr(date, 'yyyy-MM-dd hh:mm:ss')
      } else {
        return ''
      }
    },
    formatBoolean: function(bool) {
      if (bool != null) {
        return bool ? '是' : '否'
      } else {
        return ''
      }
    }
  },
  methods: {
    updatefilelist(response) {
      this.formData.download = response.data.file.url
    },
    updateimglist(response) {
      this.formData.softImg = response.data.file.url
    },
    downloadBeforeUpload(file) {
      let isRightSize = file.size / 1024 / 1024 < 300
      if (!isRightSize) {
        this.$message.error('文件大小超过 300MB')
      }
      debugger
      let isAccept = new RegExp('application/.*zip.*').test(file.type)
      if (!isAccept) {
        this.$message.error('应该选择zip类型的文件')
      }
      return isRightSize && isAccept
    },
    soft_imgBeforeUpload(file) {
      let isRightSize = file.size / 1024 / 1024 < 5
      if (!isRightSize) {
        this.$message.error('文件大小超过 5MB')
      }
      return isRightSize
    },
    //条件搜索前端看此方法
    onSubmit() {
      this.page = 1
      this.pageSize = 10
      this.getTableData()
    },
    handleSelectionChange(val) {
      this.multipleSelection = val
    },
    async onDelete() {
      const ids = []
      this.multipleSelection &&
        this.multipleSelection.map(item => {
          ids.push(item.ID)
        })
      const res = await deleteNielsenSoftwareByIds({ ids })
      if (res.code == 0) {
        this.$message({
          type: 'success',
          message: '删除成功'
        })
        this.deleteVisible = false
        this.getTableData()
      }
    },
    async updateNielsenSoftware(row) {
      const res = await findNielsenSoftware({ ID: row.ID })
      this.type = 'update'
      if (res.code == 0) {
        this.formData = res.data.rens
        this.dialogFormVisible = true
      }
    },
    closeDialog() {
      this.dialogFormVisible = false
      this.formData = {
        softName: null,
        softDescription: null,
        download: null,
        softImg: null,
        version: null
      }
      this.$refs.download.clearFiles()
      this.$refs.soft_img.clearFiles()
    },
    async deleteNielsenSoftware(row) {
      this.visible = false
      const res = await deleteNielsenSoftware({ ID: row.ID })
      if (res.code == 0) {
        this.$message({
          type: 'success',
          message: '删除成功'
        })
        this.getTableData()
      }
    },
    async enterDialog() {
      let res
      switch (this.type) {
        case 'create':
          res = await createNielsenSoftware(this.formData)
          break
        case 'update':
          res = await updateNielsenSoftware(this.formData)
          break
        default:
          res = await createNielsenSoftware(this.formData)
          break
      }
      if (res.code == 0) {
        this.$message({
          type: 'success',
          message: '创建/更改成功'
        })
        this.closeDialog()
        this.getTableData()
      }
    },
    openDialog() {
      this.type = 'create'
      this.dialogFormVisible = true
    }
  },
  created() {
    this.getTableData()
  }
}
</script>

<style>
</style>