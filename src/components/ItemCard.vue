<template>
  <span>
    <div v-if="content === '0'">
      <svg-icon :icon="icon" :hasFill="false" :color="color" class="smile-icon"></svg-icon>

      {{label}}
    </div>

    <div v-else>
      <v-badge
       color="deep-purple accent-4"
       :content="content"
       inline="true"
      >
        <svg-icon :icon="icon" :hasFill="false" :color="color" class="smile-icon"></svg-icon>

        {{label}}
      </v-badge>  
    </div>
  </span>
</template>

<script>
  import SvgIcon from './SVGIcon.vue';

  export default {

    components: {
      'svg-icon': SvgIcon,
    },
    props: {
      source: String,
      icon: String,
      label: String,
      content: String,
    },
  }
</script>