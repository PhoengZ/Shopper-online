<script setup>
definePageMeta({
    layout:'auth'
})
useHead({
    title:"My store"
})
const product = ref([])

const prevClick = ()=>{
    navigateTo('/')
}
const products = ref([
  {
    name: "Wireless Mouse",
    category: "electronics",
    price: 799,
  },
  {
    name: "Mechanical Keyboard",
    category: "keyboards",
    price: 1999,
  },
  {
    name: "Bluetooth Speaker",
    category: "audio",
    price: 1499,
  },
  {
    name: "USB-C Hub",
    category: "computer accessories",
    price: 1099,
  },
])

</script>
<template>
    <BaseButton class=" absolute" size="small" theme="circular" @click="prevClick">
        <IconBackArrow color="#000000" class="absolute"/>
    </BaseButton>
    <h1 class=" text-center text-3xl font-bold text-gray-700 flex justify-center items-center gap-5"><IconCheckList color="#000000" /> Currently selling</h1>
    <BaseCardList mode="main" :product="products"/>
    <h1 class=" text-center text-3xl font-bold text-gray-700 flex justify-center items-center gap-5 border-t-2 pt-4"><IconCheckList color="#000000" />Out of stock</h1>
    <BaseCardList mode="main" :product="products"/>
</template>