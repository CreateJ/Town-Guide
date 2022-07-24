// index.js
// 获取应用实例
const app = getApp()

Page({
  data: {
    msg: '你好'
  },
  // 事件处理函数
  onLoad() {
    console.log(this.data.msg)
  }
})
