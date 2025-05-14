<script setup>
defineProps({
    product:Array
});
const emit = defineEmits(['buy']);

const handleBuy = (item)=>{
    emit('buy',item);
};


</script>
<template>
    <ul class="[&>li] grid grid-cols-3 gap-4 ">
        <li v-for="item in product">
            <BaseCardItem :model="item" @buy="handleBuy"/>
        </li>
    </ul>
</template>