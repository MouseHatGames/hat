<template lang="pug">
.tile.is-child.card.is-flex.is-flex-direction-column
    header.card-header
        p.card-header-title {{widget.title}}
        .card-header-icon(v-if="widget.description")
            .dropdown.is-hoverable.is-right
                .dropdown-trigger
                    span.icon
                        icon(icon="info-circle")
                    .dropdown-menu
                        .dropdown-content
                            .dropdown-item {{widget.description}}
    
    .card-content.p-1.pt-2(v-if="widget.type == 'group'")
        .tile.is-ancestor
            .tile.is-parent(v-for="child in widget.children")
                Widget(:widget="child")

    .card-content.is-flex-grow-1.is-flex.is-justify-content-center.is-align-items-center(v-else)
        //- On/Off button
        OnOff(v-if="widget.type == 'onoff'" name="nice" v-model="widget.value" @update:modelValue="save")

        //- Options
        .options-group(v-else-if="widget.type == 'options'")
            //- Few options - addons layout
            .field.has-addons(v-if="widget.options.length < 5")
                .control(v-for="(opt, i) in widget.options")
                    button.button(:class="{'is-info': i == widget.value}" @click="setValue(i)") {{opt}}

            //- Many options - flex layout
            .many.is-flex.is-flex-wrap-wrap.is-justify-content-space-around(v-else)
                button.button.mb-1(v-for="(opt, i) in widget.options" :class="{'is-info': i == widget.value}" @click="setValue(i)") {{opt}}

        //- Text fields (big and small)
        .field.has-addons(v-else-if="widget.type == 'text'" :class="{'w-100': widget.big}")
            //- Small field
            input.input.mr-2(v-if="!widget.big" type="text" :placeholder="widget.placeholder" v-model="widget.value" @focusout="save")

            //- Big field
            textarea.input.mr-2.h-100.w-100(v-else v-model="widget.value" :placeholder="widget.placeholder" @focusout="save")
</template>

<script lang="ts">
import { defineComponent } from 'vue'
import OnOff from "./OnOff.vue"

export default defineComponent({
    components: {
        OnOff
    },
    props: {
        widget: Object
    },
    setup(props) {
        async function save() {
            console.log("save")
        }

        async function setValue(val: any) {
            props.widget.value = val;
            await save();
        }
        
        return { save, setValue }
    }
})
</script>

<style lang="scss" scoped>
.icon {
    font-size: .9rem;
    transition: opacity .1s;

    &:not(:hover) {
        opacity: .4;
    }
}

textarea, input[type=text] {
    transition: all .2s;
    resize: none;

    &:not(:focus) {
        background: transparent;
        color: #f2f2f2;

        & ~ button {
            display: none;
        }
    }
}

.options-group {
    max-width: 300px;
}
</style>