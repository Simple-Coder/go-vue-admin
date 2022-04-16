<template>
  <div>
    store: {{ $store.state.xd.myNum }} <br />
    vue: {{ xdNum }} <br />
    <button @click="sub()">-</button> <br />
    <button @click="add()">+</button> <br />
    getters: {{ $store.getters["xd/getNum"] }} <br />
  </div>
</template>

<script>
export default {
  data() {
    return {
      message: "谢冬",
      xdNum: 0,
    };
  },

  components: {},

  computed: {},

  methods: {
    add() {
      this.$store.commit("xd/addNum", {
        num: 5,
      });
      this.xdNum++;
      this.$store.dispatch("xd/asyncAdd", {
        num: 5,
      });
    },
    sub() {
      this.$store.commit("xd/subNum");
      this.xdNum--;
    },
  },
};
</script>
<style lang='css' scoped>
</style>