<script setup>
import { useAuthStore } from '~/Stores/auth';
import { addItem, getCartItem, removeItem, validateToken } from '~/repositories/auth';
import { getProduct } from '~/repositories/product';
definePageMeta({
    layout:false,
});
let showList = ref(false);

const { data: products, error } = await getProduct();
const pd = ref(products.value || []);

if (error.value) {
  console.error('Failed to fetch products', error.value);
}
let Item = ref([]);
const token = useCookie('token');
const user = useAuthStore();
const name = ref('');
const userID = ref('');
const {error:err, data: validateData} = await validateToken(token.value)
const isValidToken = computed(() => validateData.value?.message === 'Valid')
if (isValidToken){
    name.value = user.Username
    userID.value = user.userID;
}
onMounted(()=>{
    if (isValidToken){
        name.value = user.Username;
        userID.value = user.userID;
    }
})
const checkLogout = ()=>{
    if (isValidToken.value){
        token.value = null;
        name.value = '';
        user.Username = '';
        userID.value = '';
        navigateTo('/login');
        return 
    }
}

const checkAuth = ()=>{
    if (!isValidToken.value){
        navigateTo('/login');
        return 
    }
};

const checkItem = async ()=>{
    if (!isValidToken.value){
        navigateTo('/login');
        return 
    }
    if (showList.value){
        showList.value = !showList.value
        return 
    }
    try{
        const {products} = await getCartItem(userID.value,token.value);
        Item.value = products
        showList.value = !showList.value;
    }catch(err){
        if (err.response?.status == 401){
            navigateTo('/login')
        }else{
            console.error("Unexpected error:", err);
        }
    }
}
const Buying = async (item)=>{
    if (!isValidToken.value){
        navigateTo('/login');
        return 
    }
    try{
        const {message} = await addItem(item,userID.value,token.value)
        const {products} = await getCartItem(userID.value,token.value)
        Item.value = products
    }catch(err){
        if (err.response?.status == 401){
            navigateTo('/login')
        }else{
            console.error("Unexpected error:", err)
        }
    }
} 
const Cancle = async (item)=>{
    if (!isValidToken.value){
        navigateTo('/login');
        return
    }
    try{
        const {message} = await removeItem(userID.value,item.id,token.value)
        const {products} = await getCartItem(userID.value,token.value)
        Item.value = products
    }catch(err){
        if (err.response?.status == 401){
            navigateTo('/login')
        }else{
            console.error("Unexpected error:", err)
        }
    }
}
const handleOutside = ()=>{
    showList.value = false;
}
</script>

<template>
    <TheHeader :username="name" :openBlure="showList" @logout="checkLogout" @auth="checkAuth" @checkItem="checkItem"/>
    <section class="bg-white max-w-screen-lg m-auto px-3" :class="showList ? 'blur-xs':''">
        <!-- Part of showing product  -->
         <BaseCardList class="p-6" :product="pd" @buy="Buying" />
    </section>
    <CartForm v-if="showList" :item="Item" v-click-outside="handleOutside" @add="Buying" @remove="Cancle"/>
</template>
