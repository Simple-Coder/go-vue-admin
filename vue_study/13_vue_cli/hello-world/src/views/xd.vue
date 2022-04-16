<template>
  <div>
    store: {{ $store.state.myNum }} <br />
    vue: {{ xdNum }} <br />
    <button @click="sub()">-</button> <br />
    <button @click="add()">+</button>
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
      this.$store.commit("addNum");
      this.xdNum++;
    },
    sub() {
      this.$store.commit("subNum");
      this.xdNum--;
    },
  },
};
</script>
<style lang='css' scoped>
</style>