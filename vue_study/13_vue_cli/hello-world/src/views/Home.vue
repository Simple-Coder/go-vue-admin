<template>
  <div class="home">
    <!-- <img alt="Vue logo" src="../assets/logo.png" />
    <HelloWorld msg="Welcome to Your Vue.js App" /> -->

    <div id="nav">
      <router-link to="/">Home</router-link> |
      <router-link to="/home/about">About</router-link> |
      <router-link to="/home/xd">xiedong</router-link> |

      <ul>
        <li @click="ToLink(key + 1)" v-for="(item, key) in 5" :key="key">
          我是用户{{ key }}
          <!-- <router-link :to="'/home/user/' + (key + 1)">
            我是用户{{ key }}
          </router-link> -->

          <!-- <router-link
            :to="{
              name: 'user',
              query: {
                id: '声明式传参params' + (key + 1),
              },
            }"
          >
            我是用户{{ key }}
          </router-link> -->

          <!-- <router-link
            :to="{
              path: '/home/user',
              query: {
                id: '声明式传参params' + (key + 1),
              },
            }"
          >
            我是用户{{ key }}
          </router-link> -->

          <!-- <router-link :to="ToLink(key + 1)"> 我是用户{{ key }} </router-link> -->
        </li>
      </ul>
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
  methods: {
    ToLink(id) {
      this.$router.push({
        path: "/home/user",
        query: {
          id: id,
        },
      });
    },
  },
};
</script>
