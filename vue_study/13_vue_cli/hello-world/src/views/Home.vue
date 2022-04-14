<template>
  <div class="home">
    <!-- <img alt="Vue logo" src="../assets/logo.png" />
    <HelloWorld msg="Welcome to Your Vue.js App" /> -->

    <div id="nav">
      <router-link to="/">Home</router-link> |
      <router-link to="/home/about">About</router-link> |
      <router-link to="/home/xd">xiedong</router-link>
    </div>
    <p>下面是home的router view -------</p>
    <router-view />
  </div>
</template>

<script>
// @ is an alias to /src
import HelloWorld from "@/components/HelloWorld.vue";

export default {
  name: "Home",
  components: {
    HelloWorld,
  },
};
</script>
