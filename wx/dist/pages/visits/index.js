/******/ (() => { // webpackBootstrap
var __webpack_exports__ = {};
// scenic.js
// 获取应用实例
const app = getApp();
Page({
  data: {
    globalUrl: app.globalUrl,
    msg: '你好',
    defaultData: {
      title: '地图'
    },
    activeCategory: null,
    categoryList: [],
    allIcon: '/assets/image/categories/all.png',
    allActiveIcon: '/assets/image/categories/all-active.png',
    likeIcon: '/assets/image/like.png',
    unLikeIcon: '/assets/image/unlike.png',
    scenicList: []
  },

  // 事件处理函数
  onLoad() {
    console.log(this.data.msg);
    wx.request({
      url: app.globalUrl + '/category/getAll',
      method: 'GET',
      success: res => {
        let list = res.data.data.map(item => {
          return {
            name: item.name,
            icon: app.globalUrl + '/utils/getPic/' + item.icon,
            iconActive: app.globalUrl + '/utils/getPic/' + item.icon_active,
            id: item.id
          };
        });
        console.log(list);
        this.setData({
          categoryList: list
        });
        this.getScenic(null);
      }
    });
  },

  switchTo(e) {
    wx.switchTab({
      url: e.detail.url
    });
  },

  onShow() {
    this.getScenic(this.data.activeCategory);
  },

  clickCategoryItem(e) {
    if (this.data.activeCategory === e.currentTarget.dataset.type) {
      this.setData({
        activeCategory: null
      });
      this.getScenic(null);
    } else {
      this.setData({
        activeCategory: e.currentTarget.dataset.type
      });
      this.getScenic(e.currentTarget.dataset.type);
    }
  },

  getScenic(id) {
    if (id === null) {
      wx.request({
        url: app.globalUrl + '/scenic/getAll',
        method: 'GET',
        data: {
          open_id: wx.getStorageSync('openId')
        },
        success: res => {
          this.setData({
            scenicList: res.data.data
          });
        }
      });
    } else {
      wx.request({
        url: app.globalUrl + '/scenic/getByCategory/' + id,
        data: {
          open_id: wx.getStorageSync('openId')
        },
        method: 'GET',
        success: res => {
          console.log(res.data.data);
          this.setData({
            scenicList: res.data.data
          });
        }
      });
    }
  },

  clickPointItem(e) {
    const that = this;
    console.log(e.currentTarget.dataset.point);
    wx.navigateTo({
      url: './scenic/scenic',
      events: {
        refreshPage: data => {
          console.log(data);
          that.getScenic(that.data.activeCategory);
        }
      },
      success: function (res) {
        // 通过 eventChannel 向被打开页面传送数据
        res.eventChannel.emit('getPointData', {
          data: e.currentTarget.dataset.point
        });
      }
    });
  },

  changeCollection(e) {
    if (!wx.getStorageSync('isLogin')) {
      wx.showToast({
        title: '收藏功能需要先登录',
        duration: 2000
      });
    }

    const that = this;
    console.log(e.currentTarget.dataset.item);
    const {
      item
    } = e.currentTarget.dataset;
    const targetState = item.user_collection_state === 2 ? 1 : 2;
    wx.request({
      url: app.globalUrl + '/user/action/collection',
      method: 'POST',
      data: {
        open_id: wx.getStorageSync('openId'),
        scenic_id: item.id,
        action_state: targetState
      },

      success(res) {
        if (res.data.code === 2) {
          const targetIndex = that.data.scenicList.findIndex(scenic => scenic.id === item.id);
          const temp = JSON.parse(JSON.stringify(that.data.scenicList));
          temp[targetIndex].user_collection_state = targetState;
          that.setData({
            scenicList: temp
          });
          wx.showToast({
            title: targetState === 2 ? '收藏成功' : '取消收藏成功',
            icon: 'success',
            duration: 2000
          });
        }
      }

    });
  }

});
/******/ })()
;