import { library } from '@fortawesome/fontawesome-svg-core'
import { faInfoCircle, faTimes } from '@fortawesome/free-solid-svg-icons'
import { FontAwesomeIcon } from '@fortawesome/vue-fontawesome'

library.add(faInfoCircle, faTimes)

import { App } from 'vue';

export default {
    install(app: App) {
        app.component("icon", FontAwesomeIcon);
    }
}