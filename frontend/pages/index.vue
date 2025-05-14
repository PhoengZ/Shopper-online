<script setup>
import { useAuthStore } from '~/Stores/auth';
definePageMeta({
    layout:false,
});
let showList = ref(false);
let pd = ref([{name:'IPhone13',des:'The iPhone is a sleek, high-performance smartphone with advanced cameras, a powerful chip, and an intuitive iOS experience.'},
{name:'IPhone14',des:'The iPhone is a sleek, high-performance smartphone with advanced cameras, a powerful chip, and an intuitive iOS experience.'},
{name:'IPhone15',des:'The iPhone is a sleek, high-performance smartphone with advanced cameras, a powerful chip, and an intuitive iOS experience.'}]);
useHead({
    title:"Shopper-online",
});
let Item = ref([]);
const token = useCookie('token');
const user = useAuthStore();
const name = ref('');
if (token.value){
    name.value = user.Username;
}
onMounted(async ()=>{
    if (token.value){
        name.value = user.Username;
    }
})
function CheckCookie (){
    if (!token.value){
        return false;
    }
    return true;
}
const checkLogout = async ()=>{
    if (CheckCookie()){
        token.value = null;
        name.value = '';
        user.Username = '';
        navigateTo('/login');
    }
}

const checkAuth = async ()=>{
    if (!CheckCookie()){
        navigateTo('/login');
    }
};

const checkItem = async ()=>{
    if (!CheckCookie()){
        navigateTo('/login');
    }
    showList.value = !showList.value;
}
const Buying = async (item)=>{
    if (!CheckCookie()){
        navigateTo('/login');
    }
    console.log("Buy item: ",item);
    Item.value.push(item);
} 
const Cancle = async (item)=>{
    if (!CheckCookie()){
        navigateTo('/login');
    }
    console.log("Cancle item: ",item);
}
</script>

<template>
    <TheHeader :username="name" :openBlure="showList" @logout="checkLogout" @auth="checkAuth" @checkItem="checkItem"  @cancle="Cancle"/>
    <section class="bg-white max-w-screen-lg m-auto px-3" :class="showList ? 'blur-xs':''">
        <!-- Part of showing product  -->
         <BaseCardList class="p-6" :product="pd" @buy="Buying" />
    </section>
    <CartForm v-if="showList"/>
</template>