// index.js
// 获取应用实例
const app = getApp()

Page({
  data: {
    msg: '你好',
    defaultData: {
      title: '我的主页'
    }
  },
  // 事件处理函数
  onLoad() {

    // wx.request({
    //   url: 'https://guide.time-traveler.cn/user/1',
    //   method: 'GET',
    //   success(res) {
    //     console.log(res)
    //   }
    // })
  },
  onReady() {
  }
})
