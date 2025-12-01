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
    <ul class="bg-white rounded-xl shadow-xl border border-neutral-100 py-2 min-w-[200px]">
        <li v-for="p in prop.product" :key="p" class="px-1">
            <template v-if="prop.mode === 'account'">
                <BaseTypeItem @click="() => handleClick(p)" class="w-full text-left px-4 py-2 rounded-lg text-sm text-neutral-700 hover:bg-neutral-50 hover:text-primary-600 transition-colors">{{p}}</BaseTypeItem>
            </template>
            <template v-else>
                <div class="p-3 hover:bg-neutral-50 rounded-lg transition-colors">
                    <div class="font-medium text-sm text-neutral-900 mb-2">{{p}}</div>
                    <div class="flex flex-col gap-2">
                        <template v-if="p === 'Price'">
                            <label class="flex items-center cursor-pointer group">
                                <input type="radio" class="form-radio text-primary-600 focus:ring-primary-500 border-neutral-300" :value="false" v-model="selectedSortOption">
                                <span class="ml-2 text-sm text-neutral-600 group-hover:text-neutral-900">High to Low</span>
                            </label>
                            <label class="flex items-center cursor-pointer group">
                                <input type="radio" class="form-radio text-primary-600 focus:ring-primary-500 border-neutral-300" :value="true" v-model="selectedSortOption">
                                <span class="ml-2 text-sm text-neutral-600 group-hover:text-neutral-900">Low to High</span>
                            </label>
                        </template>
                        <template v-else>
                            <label class="flex items-center cursor-pointer group">
                                <input type="checkbox" class="form-checkbox rounded text-primary-600 focus:ring-primary-500 border-neutral-300" :value="p" v-model="selectedItem">
                                <span class="ml-2 text-sm text-neutral-600 group-hover:text-neutral-900">Select</span>
                            </label>
                        </template>
                    </div>
                </div>
            </template>
        </li>
    </ul>
</template>