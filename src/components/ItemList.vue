<template>
  <span>
    <span class="cluster-title">
      {{ name }}
     </span>
   <span
      v-for="(property, i) in properties"
      :key="i"
    >
      <ItemCard
        :label="property.value" 
        :icon="property.icon" 
        :alerts="getPropertyAlerts(property.alerts)"
        :alertscolour="property.badgecolour"
        class="clickable-icon"
        @click.native="selectProperties(properties.capacity.memory, properties.allocatable.memory)"
      />

    </span>

    <v-dialog v-model="dialog" max-width="50%">
      <v-card>
        <v-card-title class="headline"> </v-card-title>
        <v-card-text>
          <v-row no-gutters>
          </v-row>
        </v-card-text>
      </v-card>
    </v-dialog>

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

    data() {
      return {
        dialog: false,
        selectedProperty: {},
        propertyAllocatable: {},
        propertyCapacity: {},
      };
    },

    methods: {
      getPropertyAlerts(alerts){

        var count = 0;
        if (alerts) {
          count = alerts.length;
        }
        return count;      
      },
      selectProperty(property) {
        this.selectedProperty = property
        this.dialog = true;
      },
      selectProperties(propertycapacity, propertyallocatable) {
        this.propertyAllocatable = propertyallocatable
        this.propertyCapacity = propertycapacity
        this.dialog = true;
      },
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

  .clickable-icon {
    cursor: pointer;
  }

</style>
