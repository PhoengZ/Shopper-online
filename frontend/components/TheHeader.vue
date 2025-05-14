<script setup>
import BaseSearch from './BaseSearch.vue';
import BaseBadgeList from './BaseBadgeList.vue';
import BaseOption from './BaseOption.vue';

defineEmits(['auth','logout','checkItem','cancle']);

const prop = defineProps({
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
const handleItem = (item)=>{
    val.value = item;
}
const val = ref("");
</script>
<template>
    <BaseHeadBar :username="username" @logout="$emit('logout')" @is-dropuser="$emit('auth')" @checkItem="$emit('checkItem')"
    @auth="$emit('auth')"/>
    <section class=" flex justify-center flex-row gap-5 items-start ">
        <div class=" flex flex-col w-2xl min-h-[50px]">
            <BaseSearch @search="handleGet" v-model="val"/>
            <BaseBadgeList :badges="searchResults" @item="handleItem" class=" overflow-x-hidden"/>
        </div>
        <div class=" relative ">
            <BaseOption @open-filter="showFilter" :flag="!isShowed" class="hover:cursor-pointer"/>
            <BaseTypeList v-if="isShowed" :product="['a','b','c','d','e','f','g',]" class=" absolute max-h-30 overflow-y-auto drop-shadow-lg z-10"></BaseTypeList>
        </div>
        <!-- this div is drop to center when i add something at abseBadgeList -->
    </section>
    
</template>