<script setup>
import BaseSearch from '../components/BaseSearch.vue';
import BaseBadgeList from '../components/BaseBadgeList.vue';
let searchResults = ref([]);

const handleGet = (val)=>{
    try{
        if (searchResults.value.length > 10)searchResults.value.shift();
        searchResults.value.push(val);
    }catch(err){
        console.log(err);
    }
};
</script>
<template>
    <header class="bg-white w-screen-lg px-3 flex items-center justify-center flex-col pt-5">
        <BaseSearch @search="handleGet"/>
        <BaseBadgeList :badges="searchResults"/>
    </header>
    <main class=" bg-white max-w-screen-lg m-auto px-3">
        <slot>
            
        </slot>
    </main>
    <!-- <footer>This is footer</footer> -->
</template>