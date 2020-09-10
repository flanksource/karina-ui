<template>

  <span>
    <div v-if="content === 0">
      <v-badge
        color=" "
        :content="content"
        inline
      >
        <svg-icon
          :icon="icon"
          :hasFill="false"
          :color="colour"
        />

        <span class="pl-1 text-size">
          {{ label }}
        </span>
      </v-badge>  
    </div>

    <div v-else >
      <v-badge
        :color="colour"
        :content="content"
        inline
      >
        <svg-icon
          :icon="icon"
          :hasFill="false"
          :color="colour"
        />

        <span class="pl-1 text-size">
          {{ label }}
        </span>
      </v-badge>  
    </div>
  </span>

</template>

<script>

  import SvgIcon from './SVGIcon.vue';

  export default {
    name: 'ItemCard',

    components: {
      'svg-icon': SvgIcon,
    },

    props: {
      source: String,
      icon: String,
      label: String,
      content: Number,
      colour: String,
    },
  }

</script>

<style scoped>

  .text-size{
    font-size: 12px;
  }

</style>