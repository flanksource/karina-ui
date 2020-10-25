<template>
  <span>
    <span class="cluster-title">
      {{ id }}
    </span>


    <span v-if="alerts">

      <ItemCard
        label="alerts"
        icon="threat"
        :alerts="countPromAlerts(alerts)"
        :alertscolour="getPromColour(alerts)"
        class="clickable-icon"
        @click.native="viewPromAlerts(alerts)"
      />
    </span>

        <v-dialog v-model="dialog" max-width="50%">
          <v-card>
            <v-card-title class="headline"></v-card-title>

            <v-card-text>

              <v-row>
                 <span v-for="(alert, i) in alerts" :key="i">

                  <b>{{i + 1}} {{alert.labels.alertname}}</b>
                   {{alert.labels.severity}}<br/>
                  {{alert.annotations.message}}
                  since:{{alert.since}}<br/>
                 
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
  name: "ItemList",

  components: {
    ItemCard,
  },

  props: {
    id: String,
    items: Array,
    alerts: Array,
  },

  data() {
    return {
      dialog: false,
    };
  },

  methods: {
    countPromAlerts(alerts) {
      var count = 0
      if (alerts) {
        count = alerts.length
      }
      return count;      
    },
      
    viewPromAlerts(alerts) {
      this.dialog = true;
      this.alerts = alerts;
    },

    getPromColour(alerts) {
        if (alerts) {
        var highest = 'none'
        for (var i = 0; i < alerts.length; i++) {
          var level = alerts[i].labels.severity
        
          highest = this.getPromHighest(highest, level)
          
        }
        return this.pickPromColour(highest)
      }       
    },

    getPromHighest(highest, level) {
      if (highest == 'critical' || level == 'critical') {
        highest = 'critical'
      } else if (highest == 'warning' || level == 'warning') {
        highest = 'warning'
        
      } else {
        highest = 'info'
       
      }
      return highest
    },

     pickPromColour(highest) {
        var colour
        if (highest == 'critical') {
          colour = '#990000'
        } else if (highest == 'warning') {
          colour = '#f68c1f'
        } else {
          colour = '#eae229'
        }
        console.log(colour)
        return colour
      },
  }
};
</script>

<style scoped>
.no-events {
  pointer-events: none;
}

.cluster-title {
  text-decoration: underline;
  font-weight: bold;
}

  .clickable-icon {
    cursor: pointer;
  }
</style>
