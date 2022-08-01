<template>
  <div class="page-container">
    <div class="search-box" />
    <div class="action-box">
      <el-button type="primary" @click="handleAdd">添加景区</el-button>
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
    <div class="pagination-box">123</div>
  </div>
</template>

<script>
import api from '@/api'

export default {
  name: 'InformationManagement',
  data() {
    return {
      tableData: [],
      tableColumn: [
        { label: '景点名称', key: 'name' },
        { label: '地点描述', key: 'description' },
        { label: '标签', key: 'tag' },
        { label: '简介', key: 'intro' }
      ]
    }
  },
  mounted() {
    this.getTableData()
  },
  methods: {
    getTableData() {
      api.scenic.getScenicList().then(res => {
        console.log(res)
        this.tableData = res.data
      })
    },
    handleAdd() {
      this.$router.push('/information-management/scenic-add')
    },
    handleModify(row) {
      this.$router.push({
        name: 'scenicModify',
        query: row
      })
    },
    handleDelete(row) {
      this.$confirm('确认删除该景点?', {
        title: '删除',
        type: 'info',
        confirmButtonText: '删除',
        cancelButtonText: '取消'
      }).then(_ => {
        api.scenic.deleteScenic({ id: row.id }).then(res => {
          console.log(res)
          this.$message.success('删除成功！')
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
  overflow: auto;

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
