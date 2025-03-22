<script setup>
let IsClick = ref(false);

const emit = defineEmits(['openFilter']);

const changeClick = ()=>{
    IsClick.value = !IsClick.value;
    emit('openFilter');
};
</script>

<template>
    <button class=" text-center flex flex-col items-center">
        ทั้งหมด
        <IconDropDown class=" rounded-2xl hover:bg-gray-300"  @click="changeClick"/>
    </button>
</template>