<script setup>
const prop = defineProps({
    totalPage:Number,
    nowPage:Number
})
defineEmits(['changePage'])
</script>
<template>
    <ul class="w-full flex justify-center items-center gap-3 my-5">
        <li class="w-fit h-fit" v-for="i in prop.totalPage" :key="i">
            <BaseButton size="small" :theme="prop.nowPage == i ? 'first':'second'" @click="$emit('changePage',i)">{{i}}</BaseButton>
        </li>
        
    </ul>
</template>