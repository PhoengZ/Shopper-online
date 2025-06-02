<script setup>

const props = defineProps({
  item: Object,
  editMode:{
    type:Boolean,
    default:false
  }
})
const emit = defineEmits(['remove', 'edit','submit'])
const product = ref(props.item)
const handleRemove = ()=>{
    emit('remove')
}
const handleEdit = ()=>{
    emit('edit')
}
const submitAdd = ()=>{
  emit('submit', product.value)
}
const categoryAdd = (value)=>{
  product.value.category = value
}

</script>

<template>
  <template v-if="!editMode">
    <div class="relative border p-5 bg-white rounded-xl shadow-lg hover:shadow-2xl transition-shadow flex gap-5 w-full">
      <BaseImage
        class="w-32 h-32 object-cover rounded-lg"
        url="https://www.istudio.store/cdn/shop/files/iPhone_16_Pro_White_Titanium_TH_1_1ac3f824-7153-43b1-8211-a2f67cdd5e68.jpg?v=1725929670&width=823"
      />
      <div class="flex flex-col justify-between w-full">
        <div>
          <h3 class="font-bold text-xl text-gray-800">{{ item.name }}</h3>
          <p class="text-sm text-gray-600 mt-1">{{ item.description }}</p>
          <div class="text-sm text-gray-500 mt-2">
            Category:
            <BaseBadgeList :badges="item.category" />
          </div>
        </div>

        <div class="flex items-center justify-between mt-4">
          <p class="font-bold text-lg text-green-600">{{ item.price }}฿</p>
          <p class="font-bold text-lg text-gray-700" :class="item.quantity > 0 ? 'text-gray-700':'text-red-600' ">{{item.quantity}} piece</p>
          <div class="flex gap-2">
            <BaseButton size="small" theme="third" @click="handleRemove">Remove</BaseButton>
          </div>
        </div>
      </div>

      <BaseButton
        size="small"
        theme="circular"
        class="absolute right-3 top-3 bg-white hover:bg-gray-100 shadow" @click="handleEdit">
        <IconEdit color="#000000" />
      </BaseButton>
    </div>
  </template>
  <template v-else>
    <div class=" border-2 p-5 rounded-2xl flex flex-row items-start gap-3 bg-white">
      <BaseInput placeholder="Add your image" type="file" width="w-2/6" class="h-full"/>
      <div class=" flex flex-col w-full">
        <BaseInput  placeholder="Product name" width="w-full" v-model:modelvalue="product.name"/>
        <BaseInput  placeholder="Description" width="w-full" v-model:modelvalue="product.description"/>
        <BaseInput  placeholder="Price" type="number" width="w-full" v-model:modelvalue="product.price"/>
        <BaseInput  placeholder="Quantity" type="number" width="w-full" v-model:modelvalue="product.quantity"/>
        <CategoryForm :categories="product.category" @submit="categoryAdd"/>
      </div>
      <BaseButton size="small" theme="second" class="mt-auto mx-auto" @click="submitAdd">Submit</BaseButton>
    </div>
  </template>
</template>
