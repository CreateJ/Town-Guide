/******/ (() => { // webpackBootstrap
var __webpack_exports__ = {};
// scenic.js
// 获取应用实例
const app = getApp();
Page({
  data: {
    globalUrl: app.globalUrl,
    navBarHeight: app.globalData.navBarHeight,
    isLogin: true,
    openId: '',
    userInfo: {},
    activeType: 'clock',
    showData: [],
    CollectionScenicInfo: [],
    ClockScenicInfo: []
  },

  // 事件处理函数
  onLoad() {
    const isLogin = wx.getStorageSync('isLogin');
    const openId = wx.getStorageSync('openId');
    this.setData({
      isLogin,
      openId,
      userInfo: app.globalData?.userInfo || {}
    });
    console.log(isLogin, openId, this.data.userInfo);
  },

  getUserProfile(e) {
    // 推荐使用wx.getUserProfile获取用户信息，开发者每次通过该接口获取用户个人信息均需用户确认，开发者妥善保管用户快速填写的头像昵称，避免重复弹窗
    wx.getUserProfile({
      desc: '展示用户信息',
      // 声明获取用户个人信息后的用途，后续会展示在弹窗中，请谨慎填写
      success: res => {
        console.log(res, 111);
        this.setData({
          userInfo: res.userInfo,
          hasUserInfo: true
        });
        this.setData({
          userInfo: { ...this.data.userInfo,
            openId: wx.getStorageSync('openId')
          }
        });
        const {
          openId,
          nickName,
          avatarUrl,
          gender
        } = this.data.userInfo;
        wx.request({
          method: 'POST',
          url: this.data.globalUrl + '/user/register',
          data: {
            gender: gender,
            open_id: openId,
            nick_name: nickName,
            avatar: avatarUrl
          },
          success: res => {
            console.log(res);
            this.setData({
              isLogin: true
            });
          }
        });
      },
      fail: err => {
        console.log(err, 222);
      }
    });
  },

  goBack() {
    wx.navigateBack();
  },

  changeActiveType(e) {
    this.setData({
      activeType: e.currentTarget.dataset.type
    });
    this.setShowData(e.currentTarget.dataset.type);
  },

  setShowData(type) {
    if (type === 'clock') {
      this.setData({
        showData: this.data.ClockScenicInfo
      });
    } else {
      this.setData({
        showData: this.data.CollectionScenicInfo
      });
    }
  },

  onShow() {
    this.initData();
  },

  initData() {
    wx.request({
      url: this.data.globalUrl + '/user/getUserDetail/' + wx.getStorageSync('openId'),
      method: 'GET',
      success: res => {
        console.log(res.data.data);
        this.setData({
          CollectionScenicInfo: res.data.data.CollectionScenicInfo,
          ClockScenicInfo: res.data.data.ClockScenicInfo
        });
        this.setShowData(this.data.activeType);
      }
    });
  },

  clickPointItem(e) {
    const that = this;
    console.log(e.currentTarget.dataset.point);
    wx.navigateTo({
      url: '/pages/visits/scenic/scenic',
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
  }

});
/******/ })()
;