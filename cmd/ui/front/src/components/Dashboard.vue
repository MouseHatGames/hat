<template lang="pug">
.container
    .tile.is-ancestor(v-for="row in widgets")
        .tile.is-parent(v-for="(widget, path) in row")
            Widget(:widget="widget" :path="path")
</template>

<script lang="ts">
import { computed, defineComponent, provide, reactive, ref } from "vue";
import axios from "axios";
import WidgetComponent from "./Widget.vue"
import { Widget } from "../types/widget";

type WidgetRow = { [path: string]: Widget };

export default defineComponent({
    components: {
        Widget: WidgetComponent
    },
    setup() {
        const widgets = ref<WidgetRow[]>()
        const data = ref<Record<string, any>>()
        
        axios.get("/api/data").then(resp => {
            widgets.value = resp.data.widgets;
            data.value = reactive(resp.data.data);
        })
        
        provide("data", data)

        return { widgets }
    }
})
</script>
