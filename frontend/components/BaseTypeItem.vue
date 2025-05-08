<script setup>
const emit = defineEmits(['clik']);
const func  = ()=>{
    emit('clik');
}
</script>
<template>
    <button class=" bg-white hover:bg-gray-300 w-full " @click="func">
        <slot>
        </slot>
    </button>
</template>