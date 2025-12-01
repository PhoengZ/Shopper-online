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
        <ul class="grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-3 xl:grid-cols-4 gap-6">
            <li v-for="item in product" :key="item._id" class="flex justify-center">
                <BaseCardItem :model="item" @buy="handleBuy" height="h-full" :mode="mode" class="w-full"/>
            </li>
        </ul>
    </template>
    <template v-else-if="mode === 'profile'">
        <ul class="flex overflow-x-auto space-x-4 pb-4 scrollbar-hide">
            <li v-for="item in product" :key="item._id" class="flex-shrink-0 w-64">
                <BaseCardItem :model="item" @buy="handleBuy" height="h-full" :mode="mode" class="w-full"/>
            </li>
        </ul>
    </template>
</template>