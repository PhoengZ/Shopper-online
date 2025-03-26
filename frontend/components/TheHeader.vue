<script setup>
import BaseSearch from './BaseSearch.vue';
import BaseBadgeList from './BaseBadgeList.vue';
import BaseOption from './BaseOption.vue';

defineEmits(['auth']);

defineProps({
    openFilter:Boolean,
    username:String,
})
let isShowed = ref(false);
let searchResults = ref([]);
const handleGet = (val)=>{
try{
    if (searchResults.value.length > 10)searchResults.value.shift();
    if (val != null && val !== '')searchResults.value.push(val);
}catch(err){
    console.log(err);
}
};
const showFilter = ()=>{
    isShowed.value = !isShowed.value;
};
</script>
<template>
    <BaseHeadBar :username="username" @is-dropuser="$emit('auth')"/>
    <section class=" flex justify-center flex-row gap-5 items-start">
        <div class=" flex justify-center flex-col w-2xl ">
            <BaseSearch @search="handleGet"/>
            <BaseBadgeList :badges="searchResults" class=" overflow-x-hidden"/>
        </div>
        <div class=" relative">
            <BaseOption @open-filter="showFilter" :flag="!isShowed"/>
            <BaseTypeList v-if="isShowed" :product="['a','b','c','d','e','f','g',]" class=" absolute max-h-30 overflow-y-auto drop-shadow-lg z-10"></BaseTypeList>
        </div>
        <!-- this div is drop to center when i add something at abseBadgeList -->
    </section>
    
</template>