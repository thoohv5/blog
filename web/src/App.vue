<template>
  <div id="app">
    <el-container>
      <el-header>
        <Header></Header>
      </el-header>
      <el-container style="width: 80%;">
        <el-main class="content">
          <el-backtop target=".content" :bottom="100">
            <div
              style="{
        height: 100%;
        width: 100%;
        background-color: #f2f5f6;
        box-shadow: 0 0 6px rgba(0,0,0, .12);
        text-align: center;
        line-height: 40px;
        color: #1989fa;
      }"
            >
              UP
            </div>
          </el-backtop>
          <router-view></router-view>
          <footer class="footer">
            <ul class="footer-links">
              <section class="copyright"><a href="https://thooh.com">速虎前沿</a> © 2014 - 2021 | <a
                href="http://www.beian.miit.gov.cn" target="_blank">蜀ICP备2020036949号-1</a></section>
            </ul>
          </footer>
        </el-main>
      </el-container>
    </el-container>

  </div>
</template>

<script>
  import Header from "./components/Header";

  export default {
    name: 'App',
    components: {
      Header,
    },
    data() {
      return {}
    },
  }
</script>

<style scoped>
  @import '../static/css/global.css';

  .el-header {
    position: relative;
    width: 100%;
    height: 60px;

    background-color: #f5f5f5;
    border-bottom: 1px solid #eaeaea;

  }


  .el-main {
    position: absolute;
    left: 300px;
    right: 300px;
    top: 60px;
    bottom: 0;
    overflow-y: scroll;

    background-color: #fdfdfd;
    color: #333;
  }

  .footer {
    padding: 25px 0;
    text-align: center;
  }

</style>
