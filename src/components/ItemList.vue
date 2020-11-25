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
        :alerts="getAlerts(property.alerts)"
        :alertscolour="getAlertsColour(property.alerts)"
        class="clickable-icon"
        @click.native="selectProperty(property.alerts)"
      />

    </span>

    <v-dialog v-model="dialog" max-width="50%">
      <v-card>
        <v-card-title class="headline"> </v-card-title>
        <v-card-text>
         
           <v-row>

            <span
              v-for="(selectedAlert, i) in selectedAlerts"
              :key="i"
            > 

            <b>Alert {{ i + 1 }} </b> level {{ selectedAlert.level }}
             since:  {{ selectedAlert.since }} <br/>
             Message:  {{ selectedAlert.message }}


          </span>

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
        selectedAlerts: {},
        propertyAllocatable: {},
        propertyCapacity: {},
      };
    },

    methods: {

      getAlertsColour(alerts) {
        if (alerts) {
          var highest = 'one'

          for (var i = 0; i < alerts.length; i++) {
            var level = alerts[i].level
            highest = this.setHighest(highest, level)
          }
          return this.pickColour(highest)
        }       
      },

      getAlerts(alerts) {
        var count = 0
        if (alerts) {
          count = alerts.length
        }
        return count;      
      },

      setHighest(highest, level) {
        if (highest == 'three' || level == 'three') {
          highest = 'three'
        
        } else if (highest == 'two' || level == 'two') {
          highest = 'two'

        } else if (highest == 'one' || level == 'one') {
          highest = 'one'

        }
        return highest
      },

      pickColour(highest) {
        var colour
        if (highest == 'three') {
          colour = '#990000'

        } else if (highest == 'two') {
          colour = '#f68c1f'

        } else {
          colour = '#eae229'

        }
        return colour
      },

      selectProperty(alerts) {
        if (alerts) {
          this.selectedAlerts = alerts
          this.dialog = true
        }
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