<script setup>
const emit = defineEmits(['click']);
const func  = ()=>{
    emit('click');
}
</script>
<template>
    <button class=" bg-white hover:bg-gray-300 w-full " @click="func">
        <slot>
        </slot>
    </button>
</template>