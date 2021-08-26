<template>
  <div id="detail">
    <el-page-header @back="goBack" title="返回" :content="article.name">
    </el-page-header>
    <div class="markdown" v-html="article.content">
    </div>
  </div>
</template>

<script>
  export default {
    name: "Detail",
    components: {},
    data() {
      return {
        code: this.$route.query.code,
        data: "",
        article: "",
      }
    },
    created() {
      this.getArticleDetail(this.code)
    },
    methods: {
      getArticleDetail(code) {
        this.$axios.get("/blog/v1/article/detail", {
          params: {
            "code": code,
          },
        }).then((res) => {
          this.article = res.entity
        }).catch((err) => {
          this.data = err
        })
      },
      goBack() {
        this.$router.go(-1);
      }
    }
  }
</script>


<style>
  @import '../../static/css/markdown.css';

  #detail .el-page-header {
    margin: 20px 0;
  }
</style>
