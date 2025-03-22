<script setup>
import BaseSearch from './BaseSearch.vue';
import BaseBadgeList from './BaseBadgeList.vue';
import BaseOption from './BaseOption.vue';

defineProps({
    openFilter:Boolean,
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
    <section class=" flex justify-center flex-row gap-5">
        <div class=" flex justify-center">
            <BaseSearch @search="handleGet"/>
            <BaseBadgeList :badges="searchResults"/>
        </div>
        <div class=" relative">
            <BaseOption @open-filter="showFilter"/>
            <BaseTypeList v-if="isShowed" :product="['a','b','c']" class=" absolute"></BaseTypeList>
        </div>
    </section>
    
</template>