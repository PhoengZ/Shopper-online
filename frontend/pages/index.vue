<script setup>
import { useAuthStore } from '~/Stores/auth';
import { validateToken } from '~/repositories/auth';
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
const { data: validateData, error: validateError } = await validateToken(token.value)
const isValidToken = computed(() => validateData.value?.message === 'Valid')
if (isValidToken)name.value = user.Username
onMounted(()=>{
    if (isValidToken){
        name.value = user.Username;
    }
})
const checkLogout = ()=>{
    if (isValidToken.value){
        token.value = null;
        name.value = '';
        user.Username = '';
        navigateTo('/login');
    }
}

const checkAuth = ()=>{
    if (!isValidToken.value){
        navigateTo('/login');
    }
};

const checkItem = ()=>{
    if (!isValidToken.value){
        navigateTo('/login');
    }
    showList.value = !showList.value;
}
const Buying = (item)=>{
    if (!isValidToken.value){
        navigateTo('/login');
    }
    let l = Item.value.length;
    for (let i = 0;i<l;i++){
        if (Item.value[i].id == item.id){
            Item.value[i].quantity += 1;
            return;
        }
    }
    item["quantity"] = 1;
    Item.value.push(item);
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
