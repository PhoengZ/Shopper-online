<script setup>
const emit = defineEmits(['remove','add'])
const prop = defineProps(
    {
        item:Array,
    }
)
const addItem = (id)=>{
    emit('add',id)
}
const removeItem = (id)=>{
    emit('remove',id)
}
const totalPrice = computed(() =>
  prop.item.reduce((sum, i) => sum + i.price * (i.quantity == null ? 1 : i.quantity),0)
)
</script>
<template>
    <div class=" w-2/3 h-2/3 bg-white rounded-2xl mx-auto> fixed top-1/2 left-1/2 transform -translate-x-1/2 -translate-y-1/2 drop-shadow-2xl flex flex-col">
        <BaseCartList :items="item" @add="addItem" @remove="removeItem"/>
        <div class=" mt-auto text-center text-2xl font-bold pt-4 border-t border-gray-200">
            Total price: {{totalPrice}}
        </div>
    </div>
</template>