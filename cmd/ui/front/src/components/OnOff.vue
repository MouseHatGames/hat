<template lang="pug">
.onoff
    input(type="checkbox" :id="name" v-model="value")
    label(:for="name").toggleWrapper
        .toggle
</template>

<script lang="ts">
import { computed, defineComponent, readonly, ref } from 'vue'

export default defineComponent({
    props: {
        modelValue: Boolean
    },
    setup(props, { emit }) {
        const name = readonly(ref(String(Math.floor(Math.random() * 10000))));
        const value = computed({
            get: () => props.modelValue,
            set: newValue => emit("update:modelValue", newValue)
        });
        
        return { name, value }
    }
})
</script>

<style lang="scss" scoped>
.onoff {
    display: grid;
    align-items: center;
    justify-content: center;
    height: 100px;
}

input {
    display: none;
}

.toggleWrapper {
    z-index: 3;
    display: flex;
    align-items: center;
    justify-content: center;
    cursor: pointer;
    transition: all 0.2s;
    width: 100px;
    height: 100px;
    border-radius: 50%;
    background-color: #fe4551;
    box-shadow: 0 7px 7px 0 rgba(#fe4551, 0.3);

    &:active {
        width: 95px;
        height: 95px;
        box-shadow: 0 11px 11px 0 rgba(#fe4551, 0.5);

        .toggle {
            height: 17px;
            width: 17px;
        }
    }

    .toggle {
        transition: all 0.2s ease-in-out;
        height: 20px;
        width: 20px;
        background-color: transparent;
        border: 10px solid #fff;
        border-radius: 50%;
        cursor: pointer;

        animation: red 0.1s linear forwards;
    }
}

input:checked + .toggleWrapper {
    background-color: #48e98a;
    box-shadow: 0 11px 11px 0 rgba(#48e98a, 0.3);

    .toggle {
        width: 0;
        background-color: #fff;
        border-color: transparent;
        border-radius: 30px;
        animation: green 0.1s linear forwards !important;
    }
}

@keyframes red {
    0% {
        height: 35px;
        width: 0;
        border-width: 5px;
    }
    100% {
        height: 35px;
        width: 35px;
        border-width: 10px;
        // opacity: 0.7;
    }
}

@keyframes green {
    0% {
        height: 35px;
        width: 35px;
        border-width: 10px;
    }
    100% {
        height: 35px;
        width: 0;
        border-width: 5px;
    }
}
</style>