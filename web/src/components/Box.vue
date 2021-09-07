<template>
  <div id="box">
    <el-button v-if="!loginState" size="small" type="text" @click="qrCodeUrl">登陆</el-button>
    <div v-if="loginState" class="block">
      <el-avatar size="small" :src="avatar"></el-avatar>
      <span>{{name}}</span>
      <el-button @click="logout" size="small" type="text">登出</el-button>
    </div>
    <el-dialog
      title="微信扫码登陆"
      :visible.sync="qrCodeState"
      width="25%"
      center
      @close="close"
    >
      <div class="block">
        <el-image :src="url">
          <div slot="placeholder" class="image-slot">
            加载中<span class="dot">...</span>
          </div>
        </el-image>
        <el-result v-if="prompt" icon="success" title="登陆成功" subTitle="将自动关闭并刷新界面"></el-result>
      </div>
      <span slot="footer" class="dialog-footer">
        <el-button type="text" :disabled="disabled" @click="qrCodeUrl">{{buttonName}}</el-button>
      </span>

    </el-dialog>
  </div>
</template>

<script>
  let interval;
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
        prompt: false,
        loginState: false,
        avatar: "",
        name: "",
      }
    },
    mounted() {
      let user = JSON.parse(localStorage.getItem("user"))
      if (user.userCode > 0) {
        this.loginState = true
        this.avatar = user.portrait
        this.name = user.nickName
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
          interval = setInterval(() => {
            this.buttonName = '（' + this.countDown + '秒）后请刷新';
            --this.countDown;
            if (this.countDown < 0) {
              this.buttonName = "重新刷新";
              this.countDown = 30;
              this.disabled = false
              clearInterval(interval);
            }
            this.checkQRCode()
          }, 1000)
        }).catch((err) => {
          this.data = err
        })
      },
      close() {
        clearInterval(interval);
      },
      checkQRCode() {
        let params = {
          "key": "login",
        }
        this.$axios.get("/wechat/v1/base/check-qr-code", {
          params,
        }).then((res) => {
          console.log(res)
          if (!res.result) {
            this.prompt = true
            clearInterval(interval);
            this.$axios.post("/common/v1/user/login", {
              thirdLoginParam: {
                "thirdPartType": 3,
                "thirdPartCode": res.extra,
              },
            }).then((res) => {
              this.loginState = true
              localStorage.setItem("token", res.token)
              this.$axios.get("/common/v1/user/info", {
                params: {"userCode": res.token},
              }).then((res) => {
                this.avatar = res.portrait
                this.name = res.nickName
                this.qrCodeState = false
                localStorage.setItem("user", JSON.stringify(res))
                this.$message('登陆成功');
              }).catch((err) => {
                console.log(err)
              })
            }).catch((err) => {
              console.log(err)
            })
          }
        }).catch((err) => {
          console.log(err)
        })
      },
      logout() {
        localStorage.removeItem("user")
        localStorage.removeItem("token")
        this.loginState = false
        this.$message('登出成功');
      },
    },
  }
</script>


<style>
  #box {
    display: inline-block;
    vertical-align: middle;
  }

  #box .el-result {
    z-index: 999;
    position: absolute;
    top: 30%;
    left: 22%;
  }

  #box .el-result__title p, #box .el-result__subtitle p {
    color: red !important;
  }

  #box .el-avatar {
    display: inline-block;
    vertical-align: middle;
  }
</style>
