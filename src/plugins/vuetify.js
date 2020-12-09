import Vue from 'vue';
import Vuetify from 'vuetify/lib';

Vue.use(Vuetify);


const vuetify = new Vuetify({
  theme: {
    themes: {
      light: {
        primary: '#ffffff',
/*        shade1: '#181818',
        shade2: '#282828',
        shade3: '#363636',
        shade4: '#404040',
        shade5: '#b3b3b3',
        shade6: '#ffffff',*/
      },
      dark: {
        primary: '#363636',
        shade0: '#121212',
        shade1: '#181818',
        shade2: '#282828',
        shade3: '#363636',
        shade4: '#3a3b3c',
        shade5: '#b3b3b3',
        shade6: '#ffffff',

      },
    },
  },
})

export default vuetify