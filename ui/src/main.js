import g from 'guark'
import Vue from 'vue'
import App from './App.vue'
import store from './store'
import router from './router'
import vuetify from './plugins/vuetify';

Vue.config.productionTip = false

new Vue({
    store,
    render: h => h(App),
    created: () => g.hook("created"),
    router,
    vuetify,
    mounted: () => g.hook("mounted")
}).$mount('#app')