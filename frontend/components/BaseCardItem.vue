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
    class="relative w-full flex flex-col bg-white rounded-xl shadow-sm border border-neutral-200 overflow-hidden transition-all duration-300 hover:shadow-md hover:-translate-y-1 group"
    :class="height"
  >
    <div class="relative overflow-hidden aspect-[4/3]">
      <BaseImage
        :url="model.url"
        class="w-full h-full object-cover transition-transform duration-500 group-hover:scale-110"
      />
    </div>

    <div class="flex flex-col flex-grow p-4">
      <h4 class="text-lg font-semibold text-neutral-900 mb-1 line-clamp-1">{{ model.name }}</h4>
      
      <div v-if="mode === 'main'" class="flex-grow">
        <p class="text-sm text-neutral-500 line-clamp-2 mb-3">
          {{ model.description }}
        </p>
      </div>

      <div class="mt-auto pt-4 flex items-center justify-between border-t border-neutral-100">
        <p class="text-xl font-bold text-primary-600">{{ model.price.toLocaleString() }} $</p>

        <template v-if="mode === 'main'">
          <BaseButton
            class="!rounded-lg !px-4 !py-2 !text-sm !font-medium transition-colors"
            size="small"
            theme="second"
            @click="$emit('buy', model)"
          >
            Add to Cart
          </BaseButton>
        </template>

        <template v-else>
          <p class="text-sm font-medium text-neutral-600 bg-neutral-100 px-2 py-1 rounded">
            Qty: <span class="text-neutral-900">{{ model.quantity }}</span>
          </p>
        </template>
      </div>
    </div>
  </div>
</template>
