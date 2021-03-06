<template lang="pug">
.tile.card.is-flex.is-flex-direction-column
    header.card-header
        p.card-header-title(:class="{'has-text-danger': failed}")
            span(:title="path") {{widget.title}}
            input.input.ml-3.is-small.is-param(v-for="param in widget.params" type="text" v-model="params[param]" :placeholder="param")

        .card-header-icon(v-if="widget.description")
            .dropdown.is-hoverable.is-right
                .dropdown-trigger
                    span.icon
                        icon(icon="info-circle")
                    .dropdown-menu
                        .dropdown-content
                            .dropdown-item {{widget.description}}
        .card-header-icon(v-if="failed")
            .dropdown.is-hoverable.is-right
                .dropdown-trigger
                    span.icon.has-text-danger
                        icon(icon="times")
                    .dropdown-menu
                        .dropdown-content
                            .dropdown-item
                                | There was an error submitting the changes.
                                | Press F5 to reload.

    .card-content.p-1.pt-2(v-if="widget.type == 'group'")
        .tile.is-ancestor
            .tile.is-parent(v-for="child in widget.children")
                Widget(:widget="child")

    .card-content.is-flex-grow-1.is-flex.is-justify-content-center.is-align-items-center(v-else)
        //- On/Off button
        OnOff(v-if="widget.type == 'onoff'" name="nice" :modelValue="value" @update:modelValue="setValue")

        //- Options
        .options-group(v-else-if="widget.type == 'options'")
            //- Few options - addons layout
            .field.has-addons(v-if="widget.options.length < 5")
                .control(v-for="(opt, i) in widget.options")
                    button.button(:class="{'is-info': i == value}" @click="setValue(i)") {{opt}}

            //- Many options - flex layout
            .many.is-flex.is-flex-wrap-wrap.is-justify-content-space-around(v-else)
                button.button.mb-1(v-for="(opt, i) in widget.options" :class="{'is-info': i == value}" @click="setValue(i)") {{opt}}

        //- Text fields (big and small)
        .field.has-addons(v-else-if="widget.type == 'text'" :class="{'w-100': widget.big}")
            //- Small field
            input.input.mr-2(v-if="!widget.big" type="text" :placeholder="widget.placeholder" v-model="value" @focusout="save()")

            //- Big field
            textarea.input.mr-2.h-100.w-100(v-else v-model="value" :placeholder="widget.placeholder" @focusout="save()")
</template>

<script lang="ts">
import axios from 'axios'
import { computed, defineComponent, inject, PropType, Ref, ref, reactive } from 'vue'
import { Widget } from '../types/widget'
import OnOff from "./OnOff.vue"

export default defineComponent({
    components: {
        OnOff
    },
    props: {
        widget: Object as PropType<Widget>,
        path: String,
    },
    setup(props, { emit }) {
        const data = inject<Ref<Record<string, any>>>("data")
        const failed = ref(false);
        const params = props.widget.params && reactive(Object.fromEntries(props.widget.params.map(o => [o, null])))

        const value = computed({
            get: () => data.value[props.path],
            set: val => data.value[props.path] = val
        })
        
        async function save(oldValue?: any) {
            failed.value = false;

            try {
                if (params && Object.entries(params).some(o => o[1] === null)) {
                    throw "empty parameter";
                }

                await axios.post(`/api/widget/${props.path}/value`, JSON.stringify(value.value), {
                    params: params
                });
            } catch (e) {
                failed.value = true;

                if (oldValue !== undefined)
                    value.value = oldValue;
            }
        }

        async function setValue(val: any) {
            var oldValue = value.value;
            value.value = val

            await save(oldValue);
        }
        
        return { save, setValue, value, failed, params }
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
        color: #a8a8a8;

        & ~ button {
            display: none;
        }

        &::placeholder {
            color: #a8a8a8;
        }
    }
}

.options-group {
    max-width: 300px;
}

.is-param {
    max-width: 10rem;
}
</style>