<script setup>
import { getStoreItem } from '~/repositories/product'
import { useAuthStore } from '~/Stores/auth'

definePageMeta({
    layout:'auth'
})
useHead({
    title:"My store"
})

const user = useAuthStore()
const prevClick = ()=>{
    navigateTo('/')
}
const {data, error:err, status} = await getStoreItem(user.userID, user.token)
const products = data.value?.products
var curProducts = ref([])
var outProducts = ref([])
for (let i in products){
  if (products[i].quantity > 0){
    curProducts.value.push(products[i])
  }else{
    outProducts.value.push(products[i])
  }
}
// const products = ref([
//   {
//     name: "Wireless Mouse",
//     category: ["electronics"],
//     price: 799,
//     description: "A sleek wireless mouse with ergonomic design and long battery life.",
//     quantity: 25,
//   },
//   {
//     name: "Mechanical Keyboard",
//     category: ["keyboards"],
//     price: 1999,
//     description: "A durable mechanical keyboard with customizable RGB lighting.",
//     quantity: 10,
//   },
//   {
//     name: "Bluetooth Speaker",
//     category: ["audio"],
//     price: 1499,
//     description: "Portable Bluetooth speaker with high-quality sound and waterproof design.",
//     quantity: 15,
//   },
//   {
//     name: "USB-C Hub",
//     category: ["computer accessories"],
//     price: 1099,
//     description: "Multiport USB-C hub with HDMI, USB 3.0, and SD card reader support.",
//     quantity: 30,
//   },
// ])

const handleRemove = async(id) =>{
  console.log("testremove",id);
}
const handleEdit = async(id)=>{
  console.log("testEdit",id);
  
}
const handleAdd = async()=>{
  for (let i = 0;i<products.value.length;i++){
    console.log(products.value[i].name);
  }
}
</script>
<template>
    <BaseButton class=" absolute" size="small" theme="circular" @click="prevClick">
        <IconBackArrow color="#000000" class="absolute"/>
    </BaseButton>
    <BaseButton class=" absolute top-4 right-4" size="small" theme="circular" @click="handleAdd"><IconPlus color="#000000"/>
    </BaseButton>
    <h1 class=" text-center text-3xl font-bold text-gray-700 flex justify-center items-center gap-5"><IconCheckList color="#000000" /> Currently selling</h1>
    <BaseStoreList :products="curProducts" @remove="handleRemove" @edit="handleEdit"/>
    <h1 class=" text-center text-3xl font-bold text-gray-700 flex justify-center items-center gap-5 border-t-2 pt-4"><IconCheckList color="#000000" />Out of stock</h1>
    <BaseStoreList :products="outProducts" @remove="handleRemove" @edit="handleEdit"/>
</template>