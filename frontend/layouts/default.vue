<script setup>
import TheHeader from '../components/TheHeader.vue';
</script>
<template>
    <TheHeader class="bg-white w-screen-lg px-3 flex items-center justify-center flex-col pt-5">
        <BaseSearch @search="handleGet"/>
        <BaseBadgeList :badges="searchResults"/>
    </TheHeader>
    <main class=" bg-white max-w-screen-lg m-auto px-3">
        <slot>
            
        </slot>
    </main>
</template>