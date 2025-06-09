<script setup>

const prop = defineProps({
    product:Array,
    mode:String,
    selectedItem:Array,
    selectedSortOption:Boolean
});
const emit = defineEmits(['logout','update:selectedItem','update:selectedSortOption','profile','product','topup']);

const handleClick = (item)=>{
    if (item === 'Logout' && prop.mode === 'account'){
        emit('logout');
        return
    }
    if (item === 'Profile' && prop.mode === 'account'){
        emit('profile')
        return
    }
    if (item === 'Product' && prop.mode === 'account'){
        emit('product')
        return
    }
    if (item == 'Top up' && prop.mode === 'account'){
        emit('topup')
        return
    }

}
const selectedItem = ref(prop.selectedItem)
const selectedSortOption = ref(prop.selectedSortOption)

watch(selectedItem, val => {
  emit('update:selectedItem', val)
})
watch(selectedSortOption, val => {
  emit('update:selectedSortOption', val)
})
</script>
<template>
    <ul class=" rounded-2xl [&>li] w-2xs block">
        <li v-for="p in prop.product" :key="p">
            <template v-if="prop.mode === 'account'">
                <BaseTypeItem @click="() => handleClick(p)">{{p}}</BaseTypeItem>
            </template>
            <template v-else>
                <div class=" grid grid-cols-2">
                    <BaseTypeItem >{{p}}</BaseTypeItem>
                    <div class=" bg-white flex">
                        <template v-if="p === 'Price'">
                            <label class=" flex text-xs mx-auto my-auto">
                                <input type="radio" class=" mx-auto my-auto" :value="false" v-model="selectedSortOption">
                                มากไปน้อย
                            </label>
                            <label class=" flex text-xs mx-auto">
                                <input type="radio" class=" mx-auto my-auto" :value="true" v-model="selectedSortOption">
                                น้อยไปมาก
                            </label>
                        </template>
                        <template v-else>
                            <label class=" flex text-xs my-auto mx-auto">
                                <input type="checkbox" class=" mx-auto my-auto" :value="p" v-model="selectedItem">
                            </label>
                        </template>
                    </div>
                </div>
            </template>
        </li>
    </ul>
</template>