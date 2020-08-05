<template>
    <v-list class="item-list">
      <v-list-item-group v-model="item">
        <v-list-item
          v-for="(item, i) in items"
          :key="i"
          class="no-events"
        >
          <ItemCard :label="item.label" :icon="item.icon" :content="item.badge"/>
        </v-list-item>
      </v-list-item-group>
    </v-list>
</template>

<script>
  import ItemCard from './ItemCard.vue'

  export default {
    name: 'ItemList',
    components:{
      ItemCard
    },
    props: {
      items: Object,
    },  
  }
</script>

<style scoped>
  .item-list  {
    border-right: 2px solid black !important;
    width: 25%;
  }
  .no-events{
    pointer-events: none
  }
</style>