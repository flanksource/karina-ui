<template>
  <v-card>
    <v-row>
      <v-col>
        <v-icon :color="indigo">mdi-memory</v-icon>
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
