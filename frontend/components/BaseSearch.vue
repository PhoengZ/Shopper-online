<script setup>
import { string } from 'yup';
const prop = defineProps({
    modelValue:{
        type:String,
        default:''
    }
})
const emit = defineEmits(['search','update:modelValue']);

const val = ref(prop.modelValue);

watch(() => prop.modelValue, (newVal) => {
  val.value = newVal;
});

watch(val, (newVal) => {
  emit('update:modelValue', newVal);
});
const handleClick = () => {
  emit('search', val.value);
  val.value = '';
};
</script>

<template>
    <div class="relative w-2xl">
        <input 
            type="text"
            class="w-full text-xl text-white bg-gray-600 rounded-2xl px-5 py-2 focus:outline-none focus:ring-2 focus:ring-gray-400 placeholder-gray-500"
            placeholder="Search..."
            v-model="val"
        />
        <button class="absolute right-3 top-1/2 transform -translate-y-1/2 text-white hover:bg-gray-400 rounded-2xl" @click="handleClick">
            🔍
        </button>
    </div>
</template>
