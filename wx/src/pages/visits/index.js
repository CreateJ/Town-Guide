// scenic.js
// 获取应用实例
const app = getApp()

Page({
  data: {
    msg: '你好',
    defaultData: {
      title: '地图'
    },
    activeCategory: null,
    categoryList: [],
    allIcon: '../../assets/image/categories/all.png',
    allActiveIcon: '../../assets/image/categories/all-active.png',
    likeIcon: '../../assets/image/like.png',
    unLikeIcon: '../../assets/image/unlike.png',
    scenicList: []
  },
  // 事件处理函数
  onLoad() {
    console.log(this.data.msg)
    wx.request({
      url: 'https://guide.time-traveler.cn/category/getAll',
      method: 'GET',
      success: (res) => {
        let list = res.data.data.map(item => {
          return {
            name: item.name,
            icon: 'https://guide.time-traveler.cn/utils/getPic/' + item.icon,
            iconActive: 'https://guide.time-traveler.cn/utils/getPic/' + item.icon_active,
            id: item.id
          }
        })
        console.log(list)
        this.setData({
          categoryList: list
        })
        this.getScenic(null)
      }
    })
  },
  switchTo(e) {
    wx.switchTab({
      url: e.detail.url
    })
  },
  clickCategoryItem(e) {
    if (this.data.activeCategory === e.currentTarget.dataset.type) {
      this.setData({
        activeCategory: null
      })
      this.getScenic(null)
    } else {
      this.setData({
        activeCategory: e.currentTarget.dataset.type
      })
      this.getScenic(e.currentTarget.dataset.type)
    }
  },
  getScenic(id) {
    if (id === null) {
      wx.request({
        url: 'https://guide.time-traveler.cn/scenic/getAll',
        method: 'GET',
        success: (res) => {
          console.log(res.data.data)
          this.setData({
            scenicList: res.data.data
          })
        }
      })
    } else {
      wx.request({
        url: 'https://guide.time-traveler.cn/scenic/getByCategory/' + id,
        method: 'GET',
        success: (res) => {
          console.log(res.data.data)
          this.setData({
            scenicList: res.data.data
          })
        }
      })
    }
  },
  clickPointItem(e) {
    console.log(e.currentTarget.dataset.point)
    wx.navigateTo({
      url: './scenic/scenic',
      events: {
        refreshPage: (data) => {
          console.log(data)
        }
      },
      success: function (res) {
        // 通过 eventChannel 向被打开页面传送数据
        res.eventChannel.emit('getPointData', {
          data: e.currentTarget.dataset.point
        })
      }
    })
  }
})
