<template>
  <v-card>
    <v-row>
      <v-col>
        <v-icon :color="indigo">mdi-memory</v-icon>
        <i class="icon icon-memory"></i>
        <p>HERE</p>
        <v-text class="text-size">
          Memory Usage
          <span v-text="50"></span>
          <span>%</span>
        </v-text>
      </v-col>
    </v-row>

    <v-sheet color="transparent">
      <v-sparkline
        :key="String(avg)"
        :smooth="16"
        :gradient="['#f72047', '#ffd200', '#1feaea']"
        :line-width="3"
        :value="memoryInstances"
        auto-draw
        stroke-linecap="round"
      ></v-sparkline>
    </v-sheet>
  </v-card>
</template>

<script>
const exhale = (ms) => new Promise((resolve) => setTimeout(resolve, ms));

export default {
  data: () => ({
    checking: false,
    memoryInstances: [],
  }),
  created() {
    this.takeInstance(false);
  },
  methods: {
    memoryInstance() {
      return Math.ceil(Math.random() * 100);
    },
    async takeInstance(inhale = true) {
      this.checking = true;
      inhale && (await exhale(1000));
      this.memoryInstances = Array.from({ length: 20 }, this.memoryInstance);
      this.checking = false;
    },
  },
};
</script>

<style>
@font-face {
font-family: "karina-ui-icons";
src: local("icons.eot"), url("../assets/fonts/karina-ui-icons/icons.eot") format("embedded-opentype"),
local("icons.eot"), url("../assets/fonts/karina-ui-icons/icons.woff2") format("woff2"),
local("icons.eot"), url("../assets/fonts/karina-ui-icons/icons.woff") format("woff");
}

i[class^="icon-"]:before, i[class*=" icon-"]:before {
font-family: karina-ui-icons !important;
font-style: normal;
font-weight: normal !important;
font-variant: normal;
text-transform: none;
line-height: 1;
-webkit-font-smoothing: antialiased;
-moz-osx-font-smoothing: grayscale;
}

.icon-triangle-marker:before {
content: "\f102";
}
.icon-memory:before {
content: "\f101";
}
</style>
