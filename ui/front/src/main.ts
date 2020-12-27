import { createApp } from 'vue'
import App from './App.vue'
import './style/bulma.scss'
import './style/index.scss'

import icons from "./icons"

createApp(App)
    .use(icons)
    .mount('#app')
