<script setup>
import { useAuthStore } from '~/Stores/auth';
definePageMeta({
    layout:false,
});
let showList = ref(false);
let pd = ref([
  {
    name: "IPhone13",
    des: "The iPhone is a sleek, high-performance smartphone with advanced cameras, a powerful chip, and an intuitive iOS experience.",
    price: 1000,
    id: 1,
  },
  {
    name: "IPhone14",
    des: "The iPhone is a sleek, high-performance smartphone with advanced cameras, a powerful chip, and an intuitive iOS experience.",
    price: 2000,
    id: 2,
  },
  {
    name: "IPhone15",
    des: "The iPhone is a sleek, high-performance smartphone with advanced cameras, a powerful chip, and an intuitive iOS experience.",
    price: 3000,
    id: 3,
  }
]);

let Item = ref([
  {
    name: "IPhone13",
    des: "The iPhone is a sleek, high-performance smartphone with advanced cameras, a powerful chip, and an intuitive iOS experience.",
    price: 1000,
    id: 1,
    quantity: 1
  },
  {
    name: "IPhone14",
    des: "The iPhone is a sleek, high-performance smartphone with advanced cameras, a powerful chip, and an intuitive iOS experience.",
    price: 2000,
    id: 2,
    quantity: 2
  },
  {
    name: "IPhone15",
    des: "The iPhone is a sleek, high-performance smartphone with advanced cameras, a powerful chip, and an intuitive iOS experience.",
    price: 3000,
    id: 3,
    quantity: 5
  }
]);

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
    console.log("adding Item");
    let l = Item.value.length;
    for (let i = 0;i<l;i++){
        if (Item.value[i].id == item.id){
            Item.value[i].quantity += 1;
            return;
        }
    }
    Item.value.push(item);
} 
const Cancle = async (item)=>{
    if (!CheckCookie()){
        navigateTo('/login');
    }
    let l = Item.value.length;
    console.log("removing Item");
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
    <TheHeader :username="name" :openBlure="showList" @logout="checkLogout" @auth="checkAuth" @checkItem="checkItem"  @cancle="Cancle"/>
    <section class="bg-white max-w-screen-lg m-auto px-3" :class="showList ? 'blur-xs':''">
        <!-- Part of showing product  -->
         <BaseCardList class="p-6" :product="pd" @buy="Buying" />
    </section>
    <CartForm v-if="showList" :item="Item" v-click-outside="handleOutside" @add="Buying" @remove="Cancle"/>
</template>
