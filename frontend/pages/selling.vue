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
const ItemAdd = ref({
  name: '',
  description: '',
  price: '',
  quantity: '',
  category: ''
})
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
}
const handleEdit = async(id)=>{
  console.log("testEdit",id);
  
}
const handleAdd = ()=>{
  addProduct.value = !addProduct.value
}
const summitAdd = async()=>{
  for (let i in ItemAdd.value){
    console.log("Itemadd:",ItemAdd.value[i]);
  }
}

</script>
<template>
    <BaseButton class=" absolute" size="small" theme="circular" @click="prevClick">
        <IconBackArrow color="#000000" class="absolute"/>
    </BaseButton>
    <h1 class=" text-center text-3xl font-bold text-gray-700 flex justify-center items-center gap-5"><IconCheckList color="#000000" /> Currently selling</h1>
    <BaseStoreList :products="curProducts" @remove="handleRemove" @edit="handleEdit"/>
    <BaseStoreItem v-if="addProduct" :edit-mode="addProduct">
      <div class=" border-2 p-5 rounded-2xl flex flex-row items-start gap-3 bg-white">
        <BaseInput placeholder="Add your image" type="file" width="w-2/6" class="h-full"/>
        <div class=" flex flex-col w-full">
          <BaseInput  placeholder="Product name" width="w-full" v-model:modelvalue="ItemAdd.name"/>
          <BaseInput  placeholder="Description" width="w-full" v-model:modelvalue="ItemAdd.description"/>
          <BaseInput  placeholder="Price" type="number" width="w-full" v-model:modelvalue="ItemAdd.price"/>
          <BaseInput  placeholder="Quantity" type="number" width="w-full" v-model:modelvalue="ItemAdd.quantity"/>
          <BaseInput  placeholder="Category (comma separated)" width="w-full" v-model:modelvalue="ItemAdd.category"/>
        </div>
        <BaseButton size="small" theme="second" class="mt-auto mx-auto" @click="summitAdd">Summit</BaseButton>
      </div>
    </BaseStoreItem>
    <BaseButton size="small" theme="first" class=" flex items-center justify-center" @click="handleAdd"><IconPlus color="#ffffff"/></BaseButton>
    <h1 class=" text-center text-3xl font-bold text-gray-700 flex justify-center items-center gap-5 border-t-2 pt-4"><IconCheckList color="#000000"/>Out of stock</h1>
    <BaseStoreList :products="outProducts" @remove="handleRemove" @edit="handleEdit"/>
</template>