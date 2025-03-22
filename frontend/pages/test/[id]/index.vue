<script setup>
const route = useRoute();

console.log(route.params.id);
</script>

<template>
    หน้านี้คือหน้า id: {{route.params.id}}
</template>