<template lang="pug">
.container
    transition(name="fade")
        div(v-if="widgets")
            .tile.is-ancestor(v-for="row in widgets")
                .tile.is-parent(v-for="(widget, path) in row")
                    Widget(:widget="widget" :path="path")
        template(v-else)
            progress.progress.is-small.is-dark.mt-5
</template>

<script lang="ts">
import { computed, defineComponent, provide, reactive, ref } from "vue";
import axios from "axios";
import WidgetComponent from "./Widget.vue"
import { Widget } from "../types/widget";
import { onWindowKeyDown } from "../utils";

type WidgetRow = { [path: string]: Widget };

export default defineComponent({
    components: {
        Widget: WidgetComponent
    },
    setup() {
        const widgets = ref<Widget[][]>()
        const data = ref<Record<string, any>>()
        
        function refresh() {
            widgets.value = null;
            data.value = null;
            
            axios.get<{
                widgets: Widget[]
                columns: number
                data: any
            }>("/api/data").then(resp => {
                var rows: WidgetRow[] = []
                var currentRow: WidgetRow = {}
                var i = 0;

                for (const widget of resp.data.widgets) {
                    if (i++ == resp.data.columns) {
                        rows.push(currentRow);
                        currentRow = {};
                    }

                    currentRow[widget.path] = widget;
                }
                rows.push(currentRow);

                widgets.value = rows;
                data.value = reactive(resp.data.data);
            })
        }
        refresh();

        onWindowKeyDown(ev => {
            if (ev.key == "F5") {
                ev.preventDefault();

                refresh();
            }
        })

        provide("data", data)

        return { widgets }
    }
})
</script>

<style lang="scss" scoped>
.fade-enter-active,
.fade-leave-active {
    transition: opacity 0.3s ease;
    position: absolute;
    width: 100%;
}

.fade-enter-from,
.fade-leave-to {
    opacity: 0;
}
</style>