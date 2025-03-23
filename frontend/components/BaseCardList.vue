<script setup>
defineProps({
    product:Array
});
defineEmits(['click','buy']);

const handleBuy = (item)=>{
    console.log("Buy item: ",item);
};
const handleCancle = (item)=>{
    console.log("Cancle item: ",item);
    
};

</script>
<template>
    <ul class="[&>li] grid grid-cols-3 gap-4 ">
        <li v-for="item in product">
            <BaseCardItem :model="item" @cancle="handleCancle" @buy="handleBuy"/>
        </li>
    </ul>
</template>