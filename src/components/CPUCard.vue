<template>
  <v-card>
    <v-row>
      <v-col>
        <v-icon :color="indigo"> mdi-cpu-64-bit</v-icon>
        <v-text>
          CPU Usage
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
        :value="cpuInstances"
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
    cpuInstances: [],
  }),
  created() {
    this.takeInstance(false);
  },
  methods: {
    cpuInstance() {
      return Math.ceil(Math.random() * 100);
    },
    async takeInstance(inhale = true) {
      this.checking = true;
      inhale && (await exhale(1000));
      this.cpuInstances = Array.from({ length: 20 }, this.cpuInstance);
      this.checking = false;
    },
  },
};
</script>
