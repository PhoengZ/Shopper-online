<script setup>
defineProps({
  model: Object,
  height: String,
  mode: String,
});
defineEmits(['buy']);
</script>

<template>
  <div
    class="relative w-64 flex flex-col items-center bg-white rounded-3xl shadow-lg overflow-hidden transition-transform hover:scale-105 duration-300"
    :class="height"
  >
    <BaseImage
      url="https://www.istudio.store/cdn/shop/files/iPhone_16_Pro_White_Titanium_TH_1_1ac3f824-7153-43b1-8211-a2f67cdd5e68.jpg?v=1725929670&width=823"
      class="w-full h-48 object-cover"
    />

    <div class="bg-gray-100 w-full p-4 flex flex-col items-start space-y-2">
      <h4 class="text-lg font-semibold text-gray-800">{{ model.name }}</h4>
      <p
        v-if="mode === 'main'"
        class="text-sm text-gray-600 line-clamp-3"
      >
        {{ model.description }}
      </p>
    </div>

    <div class="w-full flex items-center justify-between px-4 py-3 border-t border-gray-200">
      <p class="text-xl font-bold text-gray-800">{{ model.price }}$</p>

      <template v-if="mode === 'main'">
        <BaseButton
          class="font-bold"
          size="small"
          theme="second"
          @click="$emit('buy', model)"
        >
          Buy
        </BaseButton>
      </template>

      <template v-else>
        <p class="text-sm font-semibold text-gray-600">
          Quantity: <span class="text-gray-800">{{ model.quantity }}</span>
        </p>
      </template>
    </div>
  </div>
</template>
