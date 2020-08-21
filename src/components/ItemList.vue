<template>
  <span>
    <v-text><b><u>Panel {{ id }}</u></b></v-text>
    <span
      v-for="(item, i) in items"
      :key="i"
      class="no-events"
    >
      <ItemCard  :label="item.label" :icon="item.icon" :content="item.badge" :colour="item.badgecolour"/>
    </span>
  </span>
</template>

<script>
  import ItemCard from './ItemCard.vue'

  export default {
    name: 'ItemList',
    components:{
      ItemCard
    },
    props: {
      id: Number,
      items: Object,
    },  
  }
</script>

<style scoped>
  .no-events{
    pointer-events: none
  }
</style>