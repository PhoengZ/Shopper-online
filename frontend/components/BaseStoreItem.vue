<script setup>
defineProps({
  item: Object,
})
const emit = defineEmits(['remove', 'edit'])
const handleRemove = ()=>{
    emit('remove')
}
const handleEdit = ()=>{
    emit('edit')
}
</script>

<template>
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
