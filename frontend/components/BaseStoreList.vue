<script setup>
defineProps({
    products:Array,
    editingID:{
        type:Set,
        default:new Set()
    }
})
defineEmits(['remove','edit','submit'])
</script>
<template>
    <ul class=" flex flex-col">
        <li v-for="item in products" :key="item.id">
            <BaseStoreItem :item="item" @remove="$emit('remove',item.id)" @submit="val => $emit('submit',val)" @edit="$emit('edit',item.id)" :edit-mode="editingID.has(item.id)"/>
        </li>
    </ul>
</template>