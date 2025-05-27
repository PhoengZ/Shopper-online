<script setup>
defineProps({
    product:Array,
    mode:String,
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
    <template v-if="mode === 'main'">
        <ul class="grid grid-cols-3 gap-4 ">
            <li v-for="item in product">
                <BaseCardItem :model="item" @buy="handleBuy" height="h-full" :mode="mode"/>
            </li>
        </ul>
    </template>
    <template v-else>
        <ul class=" overflow-x-auto overflow-y-hidden whitespace-nowrap space-x-4 flex ">
            <li v-for="item in product">
                <BaseCardItem :model="item" @buy="handleBuy" height="h-full" :mode="mode"/>
            </li>
        </ul>
    </template>
</template>