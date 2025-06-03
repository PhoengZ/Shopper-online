<script setup>
const prop = defineProps(
    {
        items:Array,
        selectedItem:Array
    }
)
const emit = defineEmits(['add', 'remove','update:selectedItem'])
const selectedItem = ref(prop.selectedItem)
const handleAdd = (item)=>{
    emit('add', item)
}
const handleRemove = (item)=>{
    emit('remove',item)
}
watch(selectedItem, (val) =>{
    emit('update:selectedItem',val)
})
</script>
<template>
    <ul class=" flex flex-col overflow-y-scroll gap-y-3 [&>li]:block p-2">
        <li v-for="item in prop.items" :key="item.id" class=" w-full bg-white rounded-2xl drop-shadow-2xl">
            <BaseCartItem class=" font-bold grid grid-cols-6">
                <div class="flex items-center justify-center">
                    <BaseImage url="test"/>
                </div>
                <div class=" flex items-center justify-center text-center">
                    {{item.name}}
                </div>
                <div class=" flex items-center justify-center text-center">
                    {{item.quantity == null ? 1 :item.quantity}}
                </div>
                <div class=" flex items-center justify-center text-center">
                    <BaseButton size="small" theme="second" @click="handleAdd(item)">
                        <IconDropUp width="20" height="20"/>
                    </BaseButton>
                    <BaseButton size="small" theme="third" @click="handleRemove(item)">
                        <IconDropDown width="20" height="20"/>
                    </BaseButton>
                </div>
                <div class=" flex items-center justify-center text-center">
                    {{item.price * (item.quantity == null ? 1 : item.quantity)}}
                </div>
                <input type="checkbox" v-model="selectedItem" class=" h-5 w-5 mx-auto my-auto" :value="{
                    id:item.id,
                    name:item.name,
                    url:item.url,
                    quantity:item.quantity == null ? 1 : item.quantity,
                    price:item.price * (item.quantity == null ? 1 : item.quantity)
                }">
            </BaseCartItem>
        </li>
    </ul>
</template>