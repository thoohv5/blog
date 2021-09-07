<template>
  <div id="info">
    <!--    <Header></Header>-->
    info
    <div @click="getApi">获取数据</div>
    <div>{{data}}</div>
  </div>
</template>

<script>

  export default {
    name: "Info",
    components: {},
    data() {
      return {
        data: "初始化"
      }
    },
    methods: {
      getApi() {
        this.$axios.get("/blog/v1/article/list?keyword=Utdolore&start=0&limit=1").then((res) => {
          this.data = res
        }).catch((err) => {
          this.data = err
        })
      }
    }
  }
</script>


<style>
  #info{
    height: 1000px;
  }
</style>
