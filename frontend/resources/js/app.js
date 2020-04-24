require('./bootstrap');
import VueAxios from 'vue-axios';
window.Vue = require('vue');

Vue.use(VueAxios,axios);
Vue.component('edit-component', require('./components/EditComponent.vue').default);

const demo = new Vue({
    el: '#app',
});

