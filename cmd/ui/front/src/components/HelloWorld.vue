<template lang="pug">
.container
    .tile.is-ancestor(v-for="row in widgets")
        .tile.is-parent(v-for="widget in row")
            Widget(:widget="widget")
</template>

<script lang="ts">
import { computed, defineComponent, reactive, ref } from "vue";
import WidgetComponent from "./Widget.vue"
import { Widget } from "../types/widget";

export default defineComponent({
    components: {
        Widget: WidgetComponent
    },
    setup() {
        const widgets = reactive<Widget[][]>([
            [
                {
                    title: "What's your name?",
                    type: "text",
                    value: "John Doe",
                    placeholder: "Nice"
                },
                {
                    title: "What's your story?",
                    type: "text",
                    value: "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Quisque vel venenatis nisl. Maecenas feugiat enim in erat mattis, non commodo ante porta.",
                    placeholder: "Nice",
                    big: true
                },
                {
                    title: "Lots of options!",
                    type: "options",
                    description: "Nice",
                    options: [ "One", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten" ],
                    value: 5
                },
            ],
            [
                {
                    title: "I'm a parent",
                    type: "group",
                    children: [
                        {
                            title: "Just a few options",
                            type: "options",
                            description: "Nice",
                            options: [ "One", "Two", "Three" ],
                            value: 0
                        },
                        {
                            title: "I have a couple children",
                            type: "group",
                            children: [
                                {
                                    title: "On or off?",
                                    type: "onoff",
                                    value: true
                                },
                                {
                                    title: "How about now?",
                                    type: "onoff",
                                    value: false
                                },
                            ]
                        },
                    ]
                }
            ]
        ])

        return { widgets }
    }
})
</script>
