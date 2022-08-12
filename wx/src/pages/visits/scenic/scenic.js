// scenic.js
// 获取应用实例
const app = getApp()

Page({
  data: {
    globalUrl: app.globalUrl,
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
        background: data.data.banner.split('|').map(item => app.globalUrl + '/utils/getPic/' + item),
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
  },
  handleClock (e) {
    if(!wx.getStorageSync('isLogin')) {
      wx.showToast({
        title: '打卡功能需要先登录',
        duration: 2000
      })
      return
    }
    const that = this
    const item = this.data.scenicData
    const targetState = item.user_clock_state === 2 ? 1 : 2
    wx.request({
      url: app.globalUrl + '/user/action/clock',
      method: 'POST',
      data: {
        open_id: wx.getStorageSync('openId'),
        scenic_id: item.id,
        action_state: targetState
      },
      success(res) {
        if (res.data.code === 2) {
          that.setData({
            scenicData: {
              ...that.data.scenicData,
              user_clock_state: 2
            }
          })
          // const eventChannel = that.getOpenerEventChannel()
          // that.eventChannel.emit('refreshPage')
          wx.showToast({
            title: targetState === 2 ? '打卡成功' : '取消打卡成功',
            icon: 'success',
            duration: 2000
          })
        }
      }
    })
  }



})
