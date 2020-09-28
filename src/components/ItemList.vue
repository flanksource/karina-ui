<template>
  <span>
    <span class="cluster-title">
      {{ name }}
     </span>
    <span
      v-for="(property, i) in properties"
      :key="i"
      class="no-events"
    >
      <ItemCard
        :label="property.label" 
        :icon="property.type" 
        :alerts="getPropertyAlerts(property.alerts)"
        :alertscolour="property.badgecolour"
      />

    </span>
  </span>
</template>

<script>
import ItemCard from "./ItemCard.vue";

  export default {
    name: 'ItemList',
    
    components:{
      ItemCard
    },

    props: {
      name: String,
      properties: Object,
    },

    methods: {
      getPropertyAlerts(alerts){

        var count = 0;
        if (alerts) {
          count = alerts.length;
        }
        return count;      
      }
    }
  }
</script>

<style scoped>
  .no-events{
    pointer-events: none
  }
  
.cluster-title {
  text-decoration: underline;
  font-weight: bold;
}
</style>
