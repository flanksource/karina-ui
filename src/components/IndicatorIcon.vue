<template>
  <div class="smile-icon">
    <svg-icon icon="smile" :hasFill="true"></svg-icon>
  </div>
</template>

<script>
import SvgIcon from './SVGIcon.vue';

export default {
  name: 'my-component',
  components: {
    'svg-icon': SvgIcon,
  },
      data: () => ({
       items: [
        {
          color: '#1F7087',
          icon: 'https://cdn.vuetifyjs.com/images/cards/foster.jpg',
          title: 'Supermodel',
          artist: 'Foster the People',
          indicator_icon: 'smile',
        },
        {
          color: '#952175',
          icon: 'https://cdn.vuetifyjs.com/images/cards/halcyon.png',
          title: 'Halcyon Days',
          artist: 'Ellie Goulding',
          indicator_icon: 'test'
        },
      ],
    }),
}
</script>

<style scoped>
.smile-icon {
  font-size: 24px;
  color: #000; 

}
</style>
