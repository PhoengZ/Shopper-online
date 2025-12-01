<script setup>
import { IconDropUp } from '#components';

let IsClick = ref(false);

defineProps({
    flag:Boolean,
});

const emit = defineEmits(['openFilter']);

const changeClick = ()=>{
    IsClick.value = !IsClick.value;
    emit('openFilter');
};
</script>

<template>
    <button class="flex items-center justify-center p-2 rounded-lg text-neutral-500 hover:text-primary-600 hover:bg-neutral-50 transition-colors duration-200 focus:outline-none focus:ring-2 focus:ring-primary-500/20" @click="changeClick">
        <span class="mr-2 text-sm font-medium">Filter</span>
        <IconDropDown v-if="flag" class="w-5 h-5"/>
        <IconDropUp v-else class="w-5 h-5"/>
    </button>
</template>