<script setup>
defineProps({
    model:Object
});
defineEmits(['buy'])

</script>

<template>
    <div class=" relative w-2xs h-120 flex items-center justify-start flex-col bg-white rounded-2xl drop-shadow-lg p-3">
        <BaseImage url="https://www.istudio.store/cdn/shop/files/iPhone_16_Pro_White_Titanium_TH_1_1ac3f824-7153-43b1-8211-a2f67cdd5e68.jpg?v=1725929670&width=823" 
        class=" w-1xs h-2/4 "/>
        <div class=" bg-gray-400 rounded-2xl w-full h-full flex justify-start items-center flex-col p-3">
            <h4 class=" font-bold">{{model.name}}</h4>
            <p class=" w-full">{{model.description}}</p>
        </div>
        <div class=" mt-2 flex w-full justify-center gap-2">
            <p class=" text-2xl font-bold">{{model.price}}</p>
            <BaseButton class=" font-bold ml-auto" size="small" theme="second" @click="$emit('buy',model)">+</BaseButton>
        </div>
    </div>
</template>