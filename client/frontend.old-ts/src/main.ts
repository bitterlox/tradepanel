import { createApp } from 'vue';
import App from './App.vue';
import router from './router';
import * as Wails from '@wailsapp/runtime';

// @ts-ignore
import BalmUI from 'balm-ui';

Wails.Init(() => {
    createApp(App).use(router, BalmUI).mount('#app');
});
