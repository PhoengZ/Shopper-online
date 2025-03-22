<script setup>
defineProps({
    product:Array,
});
</script>
<template>
    <ul class="[&>li]  w-full block rounded-bl-lg rounded-br-lg">
        <li v-for="p in product">
            <BaseTypeItem @onclik="console.log('test')">{{p}}</BaseTypeItem>
        </li>
    </ul>
</template>