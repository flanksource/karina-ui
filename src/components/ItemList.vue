<template>
    <v-list class="item-list">
      <v-list-item-group v-model="item" color="primary">
        <v-list-item
          v-for="(item, i) in items"
          :key="i"
        >
          <ItemCard :label="item.label" :icon="item.icon" :content="item.badge"/>
        </v-list-item>
      </v-list-item-group>
    </v-list>
</template>

<script>
  import ItemCard from './ItemCard.vue'

  export default {
    components:{
      ItemCard
    },
    props: {
      items: Object,
    },  
  }
</script>

<style>
  .item-list  {
    border-right: 2px solid black !important;
    width: 25%;
  }
</style>