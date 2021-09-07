<template>
  <div id="home">
    <Carousel></Carousel>
  </div>
</template>

<script>
  import Box from "./Box";
  import Carousel from "./Carousel";
  export default {
    name: "Home",
    components: {
      Carousel,
      Box
    },
    data() {
      return {
        data: "初始化",
      }
    },
    methods: {}
  }
</script>


<style>
</style>
