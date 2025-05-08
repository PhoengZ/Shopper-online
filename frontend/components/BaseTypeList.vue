<script setup>

defineProps({
    product:Array,
});
const emit = defineEmits(['logout']);

const handleClick = (item)=>{
    if (item === 'Logout'){
        emit('logout');
    }
}

</script>
<template>
    <ul class="[&>li] w-2xs block">
        <li v-for="p in product">
            <BaseTypeItem @clik="handleClick(p)">{{p}}</BaseTypeItem>
        </li>
    </ul>
</template>