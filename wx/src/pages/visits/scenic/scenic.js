// scenic.js
// 获取应用实例
const app = getApp()

Page({
  data: {
    navBarHeight: app.globalData.navBarHeight,
    indicatorDots: true,
    vertical: false,
    autoplay: false,
    interval: 2000,
    duration: 500,
    background: [],
    scenicData: null,
    displayType: 'video'
  },
  // 事件处理函数
  onLoad(option) {
    const eventChannel = this.getOpenerEventChannel()
    // eventChannel.emit('refreshPage', {data: 1})
    eventChannel.on('getPointData',  (data) => {
      console.log(data)
      this.setData({
        background: data.data.banner.split('|').map(item => 'https://guide.time-traveler.cn/utils/getPic/' + item),
        scenicData: data.data
      })
    })
  },

  onShareAppMessage() {
    return {
      title: 'swiper',
      path: 'page/component/pages/swiper/swiper'
    }
  },
  goBack() {
    console.log(11)
    wx.navigateBack()
  },

  changeIndicatorDots() {
    this.setData({
      indicatorDots: !this.data.indicatorDots
    })
  },

  changeAutoplay() {
    this.setData({
      autoplay: !this.data.autoplay
    })
  },

  intervalChange(e) {
    this.setData({
      interval: e.detail.value
    })
  },

  durationChange(e) {
    this.setData({
      duration: e.detail.value
    })
  }
})
