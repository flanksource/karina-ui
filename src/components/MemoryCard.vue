<template>
  <v-card class="mx-auto" color="white lighten-4" max-width="600">
    <v-card-title>
      <v-row align="start">
        <v-icon :color="indigo" class="mr-3" size="24">mdi-memory</v-icon>
     
        <div class="caption grey--text text-uppercase">Memory Usage</div>

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
        :value="memoryInstances"
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
      memoryInstances: [],
    }),
    created () {
      this.takeInstance(false)
    },
    methods: {
      memoryInstance () {
        return Math.ceil(Math.random() * (100))
      },
      async takeInstance (inhale = true) {
        this.checking = true
        inhale && await exhale(1000)
        this.memoryInstances = Array.from({ length: 20 }, this.memoryInstance)
        this.checking = false
      },
    },
  }
</script>