<template>
  <v-app>
    <v-main>
      <router-view />
      <p>
        <strong>{{count}}</strong>
      </p>
    </v-main>
  </v-app>
</template>

<script lang="ts">
import Vue from 'vue';
import store from './store';

export default Vue.extend({
  name: 'App',

  data: () => ({}),

  computed: {
    count() {
      return store.state.count;
    },
  },

  mounted() {
    store.commit('reset');
    // console.log('Starting websocket connection!');
    // const connection = new WebSocket('ws://10.30.0.2:3000/ws');
    // console.log('created ws: ', connection);
    // connection.onmessage = (event) => {
    //   console.log('got event: ', event);
    //   store.commit('increment');
    //   console.log('store value now: ', store.state.count);
    // };
  },
});
</script>

<style>
</style>
