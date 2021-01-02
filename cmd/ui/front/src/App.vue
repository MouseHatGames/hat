<template lang="pug">
Suspense
    template(#default)
        HelloWorld(msg="Hello Vue 3.0 + Vite")
    template(#fallback)
        h1 Loading
</template>

<script>
import HelloWorld from './components/HelloWorld.vue'

export default {
    name: 'App',
    components: {
        HelloWorld
    }
}
</script>
