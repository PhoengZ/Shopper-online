<script setup>
const pops = defineProps({
    size:String,
    theme:String,
});

const emit = defineEmits(['click']);

const s = differentButtonSize(pops.size);
const v = differentVariant(pops.theme);
const handleClick = () => {
    emit('click');
};
</script>

<template>
    <button :class="[s,v]" @click="handleClick">
        <slot>

        </slot>
    </button>
</template>