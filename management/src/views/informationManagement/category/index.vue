<template>
  <div class="page-container">
    <!--    <div class="search-box"></div>-->
    <div class="action-box">
      <el-button type="primary" @click="handleAdd">添加分类</el-button>
    </div>
    <div class="table-box">
      <el-table :data="tableData" class="table">
        <el-table-column
          v-for="(item, index) in tableColumn"
          :key="index"
          :prop="item.key"
          :label="item.label"
        />
        <el-table-column
          label="操作"
        >
          <template v-slot:default="{row}">
            <el-button type="primary" @click="handleModify(row)">修改</el-button>
            <el-button type="default" @click="handleDelete(row)">删除</el-button>
          </template>
        </el-table-column>
      </el-table>
    </div>
    <!--    <div class="pagination-box"></div>-->
    <el-dialog :title="categoryForm.id ? '修改' : '新增'" :visible.sync="dialogVisible" @closed="dialogClose">
      <el-form :model="categoryForm" label-width="120">
        <el-form-item prop="name" label="分类名称">
          <el-input v-model="categoryForm.name"/>
        </el-form-item>
        <el-form-item>
          <el-button type="primary" @click="onSubmit">保存</el-button>
          <el-button @click="dialogVisible = false">取消</el-button>
        </el-form-item>
      </el-form>
    </el-dialog>
  </div>
</template>

<script>
import api from '@/api'

export default {
  name: 'CategoryManagement',
  data () {
    return {
      tableData: [],
      tableColumn: [
        { label: '分类名称', key: 'name' }
      ],
      dialogVisible: false,
      categoryForm: {
        id: null,
        name: ''
      },
      currentRow: null
    }
  },
  mounted () {
    this.getTableData()
  },
  methods: {
    getTableData () {
      api.category.getCategoryList().then(res => {
        this.tableData = res.data
      })
    },
    handleAdd () {
      console.log('add')
      this.dialogVisible = true
    },
    handleModify (row) {
      this.dialogVisible = true
      this.currentRow = row
      this.categoryForm = this.currentRow
    },
    dialogClose () {
      this.categoryForm = this.$options.data().categoryForm
    },
    onSubmit () {
      console.log(this.categoryForm.id)
      api.category.addCategory(this.categoryForm).then(res => {
        console.log(res)
      }).finally(_ => {
        this.getTableData()
        this.dialogVisible = false
      })
    },
    handleDelete (row) {
      console.log(this.currentRow)
      this.$confirm('确认删除该分类?', {
        title: '删除',
        type: 'info',
        confirmButtonText: '删除',
        cancelButtonText: '取消'
      }).then(_ => {
        api.category.deleteCategory({ id: row.id }).then(res => {
          console.log(res)
        }).finally(_ => {
          this.getTableData()
        })
      }).catch(_ => {
        console.log('close')
      })
    }
  }
}
</script>

<style scoped lang="scss">
.page-container {
  height: calc(100vh - 84px);
  box-sizing: border-box;
  display: flex;
  flex-flow: column nowrap;
  padding: 0 15px;

  .search-box {
    background: lightblue;
  }

  .action-box {
    padding: 10px 0;
    border-bottom: 1px solid #cccccc
  }

  .table-box {
    flex: 1;
    background: lightgoldenrodyellow;

    .table {
      height: 100%;
    }
  }

  .pagination-box {
    background: lightcyan;
  }
}
</style>
