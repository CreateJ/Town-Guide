/******/ (() => { // webpackBootstrap
var __webpack_exports__ = {};
// app.js
App({
  onLaunch() {
    const that = this; // 展示本地存储能力

    const logs = wx.getStorageSync('logs') || [];
    logs.unshift(Date.now());
    wx.setStorageSync('logs', logs); // 登录

    wx.login({
      success: res => {
        console.log(res, 1111); // 发送 res.code 到后台换取 openId, sessionKey, unionId

        wx.request({
          url: `https://guide.time-traveler.cn/user/getUserInfo?code=${res.code}`,
          method: 'GET',
          success: userInfo => {
            console.log(userInfo.data.data);
            const newUserInfo = userInfo.data.data;
            wx.setStorage({
              key: 'openId',
              data: newUserInfo.open_id
            });

            if (newUserInfo.avatar.length) {
              console.log('保存数据');
              wx.setStorage({
                key: 'isLogin',
                data: true
              });
            } else {
              wx.setStorage({
                key: 'isLogin',
                data: false
              });
            }
          }
        });
      },
      fail: err => {
        console.log(err);
      }
    }); // wx.request({
    //   url: 'https://guide.time-traveler.cn/category/getAll',
    //   method:'GET',
    //   success(res) {
    //     console.log(res.data)
    //     that.globalData.categoriesList = res.data.data
    //   }
    // })

    const systemInfo = wx.getSystemInfoSync(); // 胶囊按钮位置信息

    const menuButtonInfo = wx.getMenuButtonBoundingClientRect(); // 导航栏高度 = 状态栏高度 + 44

    that.globalData.navBarHeight = systemInfo.statusBarHeight + 44;
    that.globalData.menuRight = systemInfo.screenWidth - menuButtonInfo.right;
    that.globalData.menuTop = menuButtonInfo.top;
    that.globalData.menuHeight = menuButtonInfo.height;
  },

  globalData: {
    userInfo: null,
    navBarHeight: 0,
    // 导航栏高度
    menuRight: 0,
    // 胶囊距右方间距（方保持左、右间距一致）
    menuTop: 0,
    // 胶囊距底部间距（保持底部间距一致）
    menuHeight: 0,
    // 胶囊高度（自定义内容可与胶囊高度保证一致）,
    categoriesList: []
  }
});
/******/ })()
;