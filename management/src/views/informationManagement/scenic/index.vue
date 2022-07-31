<template>
  <div class="page-container">
    <div class="search-box"></div>
    <div class="action-box">
      <el-button type="primary" @click="handleAdd">添加景区</el-button>123
    </div>
    <div class="table-box">
      <el-table :data="tableData" class="table">
        <el-table-column
          v-for="(item, index) in tableColumn"
          :prop="item.key"
          :key="index"
          :label="item.label"
        >
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
