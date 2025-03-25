<script setup>
import { useAuthStore } from '~/Stores/auth';
definePageMeta({
    layout:false,
});

let pd = ref([{name:'IPhone13',des:'The iPhone is a sleek, high-performance smartphone with advanced cameras, a powerful chip, and an intuitive iOS experience.'},
{name:'IPhone14',des:'The iPhone is a sleek, high-performance smartphone with advanced cameras, a powerful chip, and an intuitive iOS experience.'},
{name:'IPhone15',des:'The iPhone is a sleek, high-performance smartphone with advanced cameras, a powerful chip, and an intuitive iOS experience.'}]);
useHead({
    title:"Shopper-online",
});

const token = useCookie('token');
const user = useAuthStore();
const name = ref('');
if (token.value){
    name.value = user.Username;
}


</script>

<template>
    <!-- นี้คือหน้าหลัก
    <BaseBadgeList :badges="['a', 'b']" />
    <BaseButton size="small" theme="third">
        <IconCart></IconCart>
        Hello
    </BaseButton>
    <BaseImage url=""/> -->
    <TheHeader :username="name"/>
    <section class="bg-white max-w-screen-lg m-auto px-3">
        <!-- Part of showing product  -->
         <BaseCardList class="p-6" :product="pd" />
    </section>
</template>