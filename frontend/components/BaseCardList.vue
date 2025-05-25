<script setup>
defineProps({
    product:Array
});
const emit = defineEmits(['buy']);

const handleBuy = (item)=>{
    const object = {
        "id":item._id,
        "name":item.name,
        "price":item.price,
        "url":item.url,
        "quantity":1
    }
    emit('buy',object);
};


</script>
<template>
    <ul class="[&>li] grid grid-cols-3 gap-4 ">
        <li v-for="item in product">
            <BaseCardItem :model="item" @buy="handleBuy"/>
        </li>
    </ul>
</template>