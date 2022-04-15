<template>
  <div>
    我是用户页面
    <p>Url传参:{{ this.$route.params.id }}</p>
    <br />
    <p>声明Url传参:{{ this.$route.params.id }}</p>
    <br />
    <p>声明query传参:{{ this.$route.query.id }}</p>
  </div>
</template>

<script>
export default {
  data() {
    return {};
  },

  components: {},

  computed: {},

  methods: {},
};
</script>
<style lang='css' scoped>
</style>