<script setup>
import { useAuthStore } from '~/Stores/auth';
import { addItem, getCartItem, removeItem, updateProfile, validateToken } from '~/repositories/auth';
import { getProduct, getProductBySearching } from '~/repositories/product';
import { getCategories } from '../repositories/categories';
definePageMeta({
    layout:false,
});
useHead({
    title:"Shopper-Online"
})
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
let showSetting = ref(false);
const choiceItem = ref([]);
const {data:response, er, status} = await getCategories();
if (status.value === 'error'){
    console.error('Failed to fetch categories', er.value);
    choiceItem.value = [];
} else {
    choiceItem.value = response.value.categories;
    if (choiceItem.value[0] != "Price"){
        choiceItem.value.unshift("Price")
    }
} 
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
    showSetting.value = !showSetting.value
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
const Adding = async (item)=>{
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
const SearchItem = async (block)=>{
    try{
        const object = await getProductBySearching(block)
        pd.value = object
    }catch(err){
        console.error(err)
    }
}
const handleProfile = ()=>{
    if (!isValidToken.value){
        navigateTo('/login');
        return 
    }
    navigateTo('/profile');
}
const handleBuyItem = async(item,totalPrice) =>{
    if (!isValidToken.value){
        navigateTo('/login')
        return
    }
    const object = {
        "password":"",
        "address":"",
        "coin":-totalPrice,
        "history":item
    }
    try{
        const {message} = await updateProfile(userID.value,object,token.value)
        const {products} = await getCartItem(userID.value, token.value)
        Item.value = products
    }catch(err){
        console.error(err);
    }
}
const handleOutside = ()=>{
    showList.value = false;
}
const handleProduct = ()=>{
    if (!isValidToken){
        navigateTo('/login')
    }
    navigateTo('/selling')
}   
const handleTopup = ()=>{
    if (!isValidToken.value){
        navigateTo('/login')
        return
    }
    navigateTo('/topup');
}
</script>

<template>
    <TheHeader :choiceItem="choiceItem" :username="name" :isDrop="showSetting" :openBlure="showList" @topup="handleTopup" @getproduct="handleProduct" @profile="handleProfile" @searchItem="SearchItem" @logout="checkLogout" @auth="checkAuth" @checkItem="checkItem"/>
    <section class="bg-white max-w-screen-lg m-auto px-3" :class="showList ? 'blur-xs':''">
         <BaseCardList class="p-6" :product="pd" @buy="Adding" mode="main" />
    </section>
    <CartForm v-if="showList" :item="Item" v-click-outside="handleOutside" @buy="handleBuyItem" @add="Adding" @remove="Cancle"/>
</template>
