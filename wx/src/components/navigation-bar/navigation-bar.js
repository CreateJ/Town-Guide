const app = getApp()
Component({
  properties: {
    // defaultData（父页面传递的数据-就是引用组件的页面）
    defaultData: {
      type: Object,
      value: {
        title: "我是默认标题"
      },
      observer: function (newVal, oldVal) {
      }
    }
  },
  data: {
    navBarHeight: app.globalData.navBarHeight,
    menuRight: app.globalData.menuRight,
    menuTop: app.globalData.menuTop,
    menuHeight: app.globalData.menuHeight,
    active: null,
    button1Bg: '../../assets/image/point.png',
    button2Bg: '../../assets/image/list.png'
  },
  attached: function () {
  },
  methods: {
    clickCategory() {
      if (this.data.active === 0) {
        this.data.active = null
        this.setData({
          button1Bg: '../../assets/image/point.png'
        })
      } else {
        this.data.active = 0
        this.setData({
          button1Bg: '../../assets/image/point-active.png',
          button2Bg: '../../assets/image/list.png',
        })
      }
    },
    clickList() {
      if (this.data.active === 1) {
        this.data.active = null
        this.setData({
          button2Bg: '../../assets/image/list.png',
          button1Bg: '../../assets/image/point.png'
        })
      } else {
        this.data.active = 1
        this.setData({
          button2Bg: '../../assets/image/list-active.png',
          button1Bg: '../../assets/image/point.png',
        })
      }
    }
  },
  lifetimes: {},
})
