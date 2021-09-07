<template>
  <div id="app">
    <el-container>
      <el-header>
        <Header></Header>
      </el-header>
      <el-container>
        <el-main class="content">
          <Backtop></Backtop>
          <router-view></router-view>
          <el-footer>
            <Footer></Footer>
          </el-footer>
        </el-main>
      </el-container>
    </el-container>
  </div>
</template>

<script>
  import Header from "./components/Header";
  import Footer from "./components/Footer";
  import Backtop from "./components/backtop";

  export default {
    name: 'App',
    components: {
      Backtop,
      Footer,
      Header,
    },
    data() {
      return {}
    },
  }
</script>

<style>
  @import '../static/css/global.css';

  .el-header {
    background-color: #f5f5f5;
    border-bottom: 1px solid #eaeaea;
  }

  .el-container {
  }

  .el-main {
    position: absolute;
    overflow-y: scroll;
    width: 100%;
    height: 100vh;

    background-color: #fdfdfd;
    color: #333;
  }

  .el-footer {
    position: relative;
    bottom: 0;
    width: 100%;
  }

</style>
