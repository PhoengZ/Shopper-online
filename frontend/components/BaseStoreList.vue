<script setup>
defineProps({
    products:Array,
})
defineEmits(['remove','edit'])
</script>
<template>
    <ul class=" flex flex-col">
        <li v-for="item in products" :key="item.name">
            <BaseStoreItem :item="item" @remove="$emit('remove',item.id)" @edit="$emit('edit',item.id)"/>
        </li>
    </ul>
</template>