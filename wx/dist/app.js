/******/ (() => { // webpackBootstrap
var __webpack_exports__ = {};
// app.js
App({
  onLaunch() {
    // 展示本地存储能力
    const logs = wx.getStorageSync('logs') || [];
    logs.unshift(Date.now());
    wx.setStorageSync('logs', logs);
    const that = this; // 登录

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
    });
  },

  globalData: {
    userInfo: null
  }
});
/******/ })()
;