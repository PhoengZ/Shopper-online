<script setup>

defineProps({
    product:Array,
});
</script>
<template>
    <ul class="[&>li] w-2xs block">
        <li v-for="p in product">
            <BaseTypeItem @clik="console.log('test')">{{p}}</BaseTypeItem>
        </li>
    </ul>
</template>