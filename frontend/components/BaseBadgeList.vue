<script setup>
defineProps({
    badges:{
        type:Array,
    },
});
defineEmits(['item'])
</script>
<template>
    <ul class="[&>li]:inline-block">
        <li class=" mr-1 md:mr-2" v-for="item in badges" key="item">
            <BaseBadgeItem class=" font-bold" @click="()=>{
                $emit('item',item);
                // used for emit to filter product from the old serching
            }">{{item}}</BaseBadgeItem>
        </li>
    </ul>
</template>
