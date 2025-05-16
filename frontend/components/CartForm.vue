<script setup>
const emit = defineEmits(['close','remove'])
const prop = defineProps(
    {
        item:Array,
    }
)
</script>
<template>
    <div class=" w-2/3 h-2/3 bg-white rounded-2xl mx-auto> fixed top-1/2 left-1/2 transform -translate-x-1/2 -translate-y-1/2 drop-shadow-2xl">
        <BaseCartList :items="item"/>
    </div>
</template>