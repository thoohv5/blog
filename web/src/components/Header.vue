<template>
  <div id="header">
    <el-row>
      <el-col :span="1" :offset="5">
        <el-image
          style="width: 50px; height: 50px; border-radius: 50%;margin-top: 5px;"
          :src="url"
          fit="fill"></el-image>
      </el-col>
      <el-col :span="5" :offset="8">
        <el-menu :default-active="activeIndex" router class="el-menu-demo" mode="horizontal" @select="handleSelect">
          <el-menu-item index="/">首页</el-menu-item>
          <el-menu-item index="/list">列表</el-menu-item>
          <el-menu-item index="/info">信息</el-menu-item>
        </el-menu>
      </el-col>

      <el-col :span="4">
        <div style="height: 60px; line-height: 60px;text-align: left;">
          <Box></Box>
        </div>
      </el-col>
    </el-row>
  </div>
</template>

<script>
  import Box from "./Box";

  export default {
    name: "Header",
    components: {
      Box,
    },
    data() {
      return {
        url: require('../assets/logo.jpg'),
        activeIndex: this.$route.path,
        nav: [{
          name: '主页'
        }, {
          name: '列表',
          selected: true
        }, {
          name: '简介'
        }],
      }
    },
    methods: {
      handleSelect(key, keyPath) {
        this.activeIndex = key
        console.log(key, keyPath);
      }
    },
    watch: {
      $route(to, from) {
        console.log(to, from)
        this.activeIndex = to.path
      }
    },
  }
</script>


<style >
  #header {

  }

  #header .el-menu {
    background: transparent;
  }

  #header .el-menu-item {
    font-size: 18px;
  }
</style>
