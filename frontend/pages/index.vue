<script setup>
import { useAuthStore } from '~/Stores/auth';
import { addItem, getCartItem, validateToken } from '~/repositories/auth';
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
    }
}

const checkAuth = ()=>{
    if (!isValidToken.value){
        navigateTo('/login');
    }
};

const checkItem = async ()=>{
    if (!isValidToken.value){
        navigateTo('/login');
    }
    if (showList.value){
        showList.value = !showList.value
        return 
    }
    const {products} = await getCartItem(userID.value);
    Item.value = products
    showList.value = !showList.value;
}
const Buying = async (item)=>{
    if (!isValidToken.value){
        navigateTo('/login');
    }
    for (let key in item) {
        console.log(key, item[key]);
    }
    const {message} = await addItem(item,userID.value)
    if (message != "Success adding item"){
        console.log("Error adding item");
        return 
    }
    const {products} = await getCartItem(userID.value);
    Item.value = products
} 
const Cancle = (item)=>{
    if (!isValidToken.value){
        navigateTo('/login');
    }
    let l = Item.value.length;
    for (let i = 0;i<l;i++){
        if (Item.value[i].id == item.id){
            if (Item.value[i].quantity == 1){
                Item.value.splice(i,1);
                return;
            }
            Item.value[i].quantity -= 1;
            return;
        }
    }
    console.log("Item not found");
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
