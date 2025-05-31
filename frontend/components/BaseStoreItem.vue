<script setup>
defineProps({
    item:Object,
})
</script>
<template>
    <div class=" border-2 p-5 bg-gray-200 drop-shadow-2xl flex">
        <slot>
        </slot>
    </div>
</template>