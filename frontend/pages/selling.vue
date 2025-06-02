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
const newItem = ref({
  name: '',
  description: '',
  price: '',
  quantity: '',
  category: []
})
const editingSet = ref(new Set())

const addProduct = ref(false)
for (let i in products){
  if (products[i].quantity > 0){
    curProducts.value.push(products[i])
  }else{
    outProducts.value.push(products[i])
  }
}
const handleRemove = async(id) =>{
  console.log("testremove",id);
  //เรียก api remove สินค้าแล้วลบออกจาก curProducts หรือ outProducts
}
const handleEdit = async(id)=>{
  console.log("testEdit",id);
  editingSet.value.add(id)
}
const handleAdd = ()=>{
  addProduct.value = !addProduct.value
}
const submitAdd = async(item)=>{
  // เรียกapi add สินค้า
  console.log(item);
} 
const submitEdit = async(item)=>{
  // เรียกapi edit สินค้า
  editingSet.value.delete(item.id)
}
</script>
<template>
    <BaseButton class=" absolute" size="small" theme="circular" @click="prevClick">
        <IconBackArrow color="#000000" class="absolute"/>
    </BaseButton>
    <h1 class=" text-center text-3xl font-bold text-gray-700 flex justify-center items-center gap-5"><IconCheckList color="#000000" /> Currently selling</h1>
    <BaseStoreList :products="curProducts" @remove="handleRemove" @edit="handleEdit" @submit="submitEdit" :editing-i-d="editingSet"/>
    <BaseStoreItem v-if="addProduct" :edit-mode="addProduct" :item="newItem" @submit="submitAdd"/>
    <BaseButton size="small" theme="first" class=" flex items-center justify-center" @click="handleAdd"><IconPlus color="#ffffff"/></BaseButton>
    <h1 class=" text-center text-3xl font-bold text-gray-700 flex justify-center items-center gap-5 border-t-2 pt-4"><IconCheckList color="#000000"/>Out of stock</h1>
    <BaseStoreList :products="outProducts" @remove="handleRemove" @edit="handleEdit" @submit="submitEdit" :editing-i-d="editingSet"/>
</template>