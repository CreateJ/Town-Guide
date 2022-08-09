// scenic.js
// 获取应用实例
const app = getApp()

Page({
  data: {
    msg: '你好',
    navBarHeight: app.globalData.navBarHeight,
  },
  // 事件处理函数
  onLoad() {
    console.log(this.data.msg)
  }
})
