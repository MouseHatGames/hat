import { onMounted, onUnmounted } from "vue";

export function onWindowKeyDown(callback: (ev: KeyboardEvent) => void) {
    onMounted(() => {
        window.addEventListener("keydown", callback)
    })
    onUnmounted(() => {
        window.removeEventListener("keydown", callback)
    })
}