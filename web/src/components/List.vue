<template>
  <div id="list">
    <div class="list" v-infinite-scroll="load"
         infinite-scroll-disabled="disabled"
         infinite-scroll-distance="20"
    >
      <ul class="posts-list">
        <li class="posts-item" v-for="(item,index) in articles" :key="index">
          <router-link :to="{ path: '/detail', query: { code: item.code }}">
            {{item.name}}
            <span>作者：{{item.author}} 阅读：{{item.read}}</span>
          </router-link>
        </li>
      </ul>
      <p v-if="loading">加载中...</p>
      <p v-if="noMore">没有更多了</p>
    </div>

  </div>
</template>

<script>

  export default {
    name: "List",
    components: {},
    data() {
      return {
        start: 0,
        articles: [],
        loading: false,
        hasMore: false,
      }
    },
    computed: {
      noMore() {
        return !this.hasMore
      },
      disabled() {
        return this.loading || this.noMore
      },
    },
    created() {
      this.articleList()
    },
    methods: {
      articleList() {
        let params = {
          "start": this.start,
          "limit": 15,
        }
        this.$axios.get("/blog/v1/article/list", {
          params,
        }).then((res) => {
          this.hasMore = res.hasMore
          if (this.hasMore) {
            this.articles = this.articles.concat(res.list)
            this.start = res.start
          }
          this.loading = false
        }).catch((err) => {
          console.log(err)
        })
      },
      load() {
        this.loading = true
        setTimeout(() => {
          this.articleList()
        }, 1000);
      },
    }
  }
</script>


<style scoped>
  .list p {
    text-align: center;
  }
  .posts-list {
    flex-grow: 1;
    margin: 0;
    padding: 0;
    list-style: none;
  }

  .posts-item {
    border-bottom: 1px #7d828a dashed;
  }

  .posts-item a {
    display: flex;
    justify-content: space-between;
    align-items: baseline;
    padding: 12px 0;
  }
</style>
