<script setup>
defineProps({
    products:Array,
})
defineProps(['remove','edit'])
</script>
<template>
    <ul class=" flex flex-col">
        <li v-for="item in products">
            <BaseStoreItem>
                
            </BaseStoreItem>
        </li>
    </ul>
</template>