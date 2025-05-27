<script setup>
import BaseSearch from './BaseSearch.vue';
import BaseBadgeList from './BaseBadgeList.vue';
import BaseOption from './BaseOption.vue';

const emit = defineEmits(['auth','logout','checkItem','searchItem']);

const prop = defineProps({
    openFilter:Boolean,
    username:String,
    openBlure:Boolean,
    isDrop:Boolean,
    choiceItem:Array,
})
const selectedItem = ref([])
const selectedSortOption = ref(false)
let isShowed = ref(false);
let searchResults = ref([]);
const handleGet = (val)=>{
    try{
        if (searchResults.value.length > 10)searchResults.value.shift();
        if (val != null && val !== '')searchResults.value.push(val);
        const object = {
            "name":val,
            "category":selectedItem.value,
            "price":selectedSortOption.value
        }
        emit('searchItem',object)
    }catch(err){
        console.error(err)
    }
};
const showFilter = ()=>{
    isShowed.value = !isShowed.value;
};
const handleItem = (item)=>{
    val.value = item;
}
const val = ref("");
const handleOutside = ()=>{
    isShowed.value = false;
}
</script>
<template>
    <BaseHeadBar :username="username" :is-show="isDrop" @logout="$emit('logout')" @is-dropuser="$emit('auth')" @checkItem="$emit('checkItem')"
    @auth="$emit('auth')" :class="openBlure ? 'blur-xs':''"/>
    <section class=" flex justify-center flex-row gap-5 items-start" :class="openBlure ? 'blur-xs':''">
        <div class=" flex flex-col w-2xl min-h-[50px]">
            <BaseSearch @search="handleGet" v-model="val"/>
            <BaseBadgeList :badges="searchResults" @item="handleItem" class=" overflow-x-hidden"/>
        </div>
        <div class=" relative ">
            <BaseOption @open-filter="showFilter" :flag="!isShowed" class="hover:cursor-pointer"/>
            <BaseTypeList v-if="isShowed"  v-model:selectedItem="selectedItem" v-model:selectedSortOption="selectedSortOption"  v-click-outside="handleOutside"
             mode="screening" :product="prop.choiceItem" class=" absolute max-h-30 overflow-y-auto drop-shadow-lg z-10"></BaseTypeList>
        </div>
        <!-- this div is drop to center when i add something at abseBadgeList -->
    </section>
    
</template>