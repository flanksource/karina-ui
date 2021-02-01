<template>
  <span v-if="shape == 'none'" class="parent" />

  <span v-else class="parent">
   <i :class="'icon-'+shape" :style="iconStyles" class="icon" :title="title" ></i>
    <v-icon color="white" class="overlay">
      {{ overlay }}
    </v-icon>
  </span>
</template>

<script>


export default {
  name: "IndicatorIcon",

  data: () => ({
    hover: false,
  }),


  props: {
    id: String,
    color: String,
    shape: String,
    overlay: String,
    //hover: Boolean,
    title: String,
  },

  computed: {
    iconStyles() {
      return {
        color: this.color,
      };
    },
  },
};
</script>

<style scoped>
.parent {
  position: relative;
  top: 0;
  left: 0;
}

.icon {
  position: absolute;
  top: 0;
  left: 0;
}

.overlay {
  position: relative;
  top: -3px;
  left: 2px;
  font-size: 12px;
}

@import '../assets/fonts/karina-ui-icons/icons.css';

</style>
