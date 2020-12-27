import { library } from '@fortawesome/fontawesome-svg-core'
import { faInfoCircle } from '@fortawesome/free-solid-svg-icons'
import { FontAwesomeIcon } from '@fortawesome/vue-fontawesome'

library.add(faInfoCircle)

import { App } from 'vue';

export default {
    install(app: App) {
        app.component("icon", FontAwesomeIcon);
    }
}