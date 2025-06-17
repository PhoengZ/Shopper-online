<script setup>
import { addStoreItem, editStoreItem, getStoreItem, removeStoreItem } from '~/repositories/product'
import { useAuthStore } from '~/Stores/auth'

const {$toast} = useNuxtApp()
useHead({
    title:"My store"
})

const user = useAuthStore()
const prevClick = ()=>{
    navigateTo('/')
}
const {data, error:err, status} = await getStoreItem(user.userID, user.token)
if (status === 'error'){
  $toast.error('ไม่สามารถโหลดสินค้าได้',{
    description:'รายละเอียด: '+err.value
  })
}
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
  try{
    const {message} = await removeStoreItem(id, user.token)
    curProducts.value = curProducts.value.filter(p => p.id !== id)
    outProducts.value = outProducts.value.filter(p => p.id !== id)
    $toast.success(message,{
      style:{
        background:'green',
        color:'white'
      }
    })
  }catch(err){
    console.error("Error removing product:", err.data.error)
    $toast.error('ไม่สามารถลบสินค้าได้',{
      description:'รายละเอียด: '+err.data.error
    })
  }
  
}
const handleEdit = (id)=>{
  editingSet.value.add(id)
}
const handleAdd = ()=>{
  addProduct.value = !addProduct.value
}
const submitAdd = async(item)=>{
  const data = new FormData()
  for (let i in item){
    data.append(i, item[i])
  }
  data.append('uid',user.userID)
  try{
    const {product} = await addStoreItem(data, user.token)
    if (product.quantity > 0) {
      curProducts.value.push(product)
    } else {
      outProducts.value.push(product)
    }
    $toast.success('เพิ่มรายการสินค้าสำเร็จ',{
      style:{
        background:'green',
        color:'white'
      }
    })
  }catch(err){
    console.error("Error adding product:", err.data.error)
    $toast.error('ไม่สามารถลบสินค้าได้',{
      description:'รายละเอียด: '+err.data.error
    })
  } 
  addProduct.value = false
} 
const submitEdit = async(item)=>{
  const data = new FormData()
  for (let i in item){
    data.append(i, item[i])
  }
  try{
    const {product} = await editStoreItem(data,user.token)
    editingSet.value.delete(item.id)
    var found = false;
    if (product.quantity > 0) {
      for (let i in curProducts.value){
        if (curProducts.value[i].id === product.id) {
          curProducts.value[i] = product
          found = true
          break
        }
      }
      if (!found) {
        curProducts.value.push(product)
        outProducts.value = outProducts.value.filter(p => p.id !== product.id)
      }
    } else {
      for (let i in outProducts.value){
        if (outProducts.value[i].id === product.id) {
          outProducts.value[i] = product
          found = true
          break
        }
      }
      if (!found){
        outProducts.value.push(product)
        curProducts.value = curProducts.value.filter(p => p.id !== product.id)
      }
    }
    $toast.success('แก้ไขรายการสินค้าสำเร็จ',{
      style:{
        background:'green',
        color:'white'
      }
    })
  }catch(error){
    console.error("Error editing product:", error?.data?.error);
    $toast.error('ไม่สามารถแก้ไขสินค้าได้',{
      description:'รายละเอียด: '+err.data.error
    })
  }
}
</script>
<template>
  <div class="min-h-screen w-full bg-gray-100 py-16 px-4 flex justify-center items-start">
      <section class="bg-white w-full max-w-4xl rounded-3xl shadow-xl p-8 flex flex-col gap-8">
          <BaseButton class=" absolute" size="small" theme="circular" @click="prevClick">
          <IconBackArrow color="#000000" class="absolute"/>
          </BaseButton>
          <h1 class=" text-center text-3xl font-bold text-gray-700 flex justify-center items-center gap-5"><IconCheckList color="#000000" /> Currently selling</h1>
          <BaseStoreList :products="curProducts" @remove="handleRemove" @edit="handleEdit" @submit="submitEdit" :editing-i-d="editingSet"/>
          <BaseStoreItem v-if="addProduct" :edit-mode="addProduct" :item="newItem" @submit="submitAdd"/>
          <BaseButton size="small" theme="first" class=" flex items-center justify-center" @click="handleAdd"><IconPlus color="#ffffff"/></BaseButton>
          <h1 class=" text-center text-3xl font-bold text-gray-700 flex justify-center items-center gap-5 border-t-2 pt-4"><IconCheckList color="#000000"/>Out of stock</h1>
          <BaseStoreList :products="outProducts" @remove="handleRemove" @edit="handleEdit" @submit="submitEdit" :editing-i-d="editingSet"/>
      </section>
  </div>
    
</template>