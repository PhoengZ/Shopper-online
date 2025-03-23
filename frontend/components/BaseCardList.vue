<script setup>
defineProps({
    product:Array
});
</script>
<template>
    <ul class="[&>li] grid grid-cols-3 gap-4 ">
        <li v-for="item in product">
            <BaseCardItem :model="item" class=" bg-white rounded-2xl drop-shadow-lg p-3" />
        </li>
    </ul>
</template>