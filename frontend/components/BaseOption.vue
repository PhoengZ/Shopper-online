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
    <button class=" text-center flex flex-col items-center">
        ทั้งหมด
        <IconDropDown class=" rounded-2xl hover:bg-gray-300"  @click="changeClick" v-if="flag"/>
        <IconDropUp class="rounded-2xl hover:bg-gray-300"  @click="changeClick" v-else/>
    </button>
</template>