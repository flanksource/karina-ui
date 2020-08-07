<template>
  <v-card class="mx-auto" color="white lighten-4" max-width="600">
    <v-card-title>
      <v-row align="start">
        <v-icon :color="indigo" class="mr-3" size="24"> mdi-cpu-64-bit</v-icon>
     
        <div class="caption grey--text text-uppercase">CPU Usage</div>

        <div class="no">
          <span class="display-2 font-weight-black" v-text="50"></span>
          <span>%</span>
        </div>
      </v-row>    
    </v-card-title>

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
  const exhale = ms => new Promise(resolve => setTimeout(resolve, ms))

  export default {
    data: () => ({
      checking: false,
      cpuInstances: [],
    }),
    created () {
      this.takeInstance(false)
    },
    methods: {
      cpuInstance () {
        return Math.ceil(Math.random() * (100))
      },
      async takeInstance (inhale = true) {
        this.checking = true
        inhale && await exhale(1000)
        this.cpuInstances = Array.from({ length: 20 }, this.cpuInstance)
        this.checking = false
      },
    },
  }
</script>