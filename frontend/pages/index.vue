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
    name.value = token.value;
}
onMounted(async ()=>{
    if (token.value){
        name.value = token.value;
    }
})

const checkAuth = async ()=>{
    if (token.value){
        console.log("isLoggedIn");
    }else{
        navigateTo('/login');
    }
};
</script>

<template>
    <TheHeader :username="name" @auth="checkAuth"/>
    <section class="bg-white max-w-screen-lg m-auto px-3">
        <!-- Part of showing product  -->
         <BaseCardList class="p-6" :product="pd" />
    </section>
</template>