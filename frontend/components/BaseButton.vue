<script setup>
const pops = defineProps({
    size:String,
    theme:String,
});

const s = differentButtonSize(pops.size);
const v = differentVariant(pops.theme);
const handleClick = () => {
  console.log('Button clicked!');
};
</script>

<template>
    <button :class="[s,v]" @click="handleClick">
        <slot>

        </slot>
    </button>
</template>