<template>
  <div id="box">
    <el-button type="text" @click="qrCodeUrl">登陆</el-button>
    <el-dialog
      title="微信扫码登陆"
      :visible.sync="qrCodeState"
      width="30%"
      center>
      <div class="block">
        <el-image :src="url">
          <div slot="placeholder" class="image-slot">
            加载中<span class="dot">...</span>
          </div>
        </el-image>
      </div>
      <span slot="footer" class="dialog-footer">
        <el-button type="text" :disabled="disabled" @click="qrCodeUrl">{{buttonName}}</el-button>
      </span>
    </el-dialog>
  </div>
</template>

<script>
  export default {
    name: "Box",
    components: {},
    data() {
      return {
        qrCodeState: false,
        disabled: true,
        countDown: 30,
        buttonName: "",
        url: "",
      }
    },
    methods: {
      qrCodeUrl() {
        this.$axios.get("/wechat/v1/base/qr-code", {
          params: {
            "expireSeconds": "30",
            "actionName": "QR_STR_SCENE",
            "sceneStr": "login"
          },
        }).then((res) => {
          this.url = res.qrCode
          this.qrCodeState = true
          let me = this;
          let interval = window.setInterval(function () {
            me.buttonName = '（' + me.countDown + '秒）后请刷新';
            --me.countDown;
            if (me.countDown < 0) {
              me.buttonName = "重新刷新";
              me.countDown = 30;
              me.disabled = false
              window.clearInterval(interval);
            }
          }, 1000)
        }).catch((err) => {
          this.data = err
        })
      },
    },
  }
</script>


<style>
  #box {

  }
</style>
