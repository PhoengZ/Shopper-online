<script setup>

const prop = defineProps({
    username:String,
    isShow:Boolean,
})
const emit = defineEmits(['isDropuser','logout','checkItem','profile','getproduct']);
const handleClick = ()=>{
    emit('isDropuser');
}
const handlelogout = () => {
    emit('logout');
}
const handleProfile = ()=>{
    emit('profile');
}
const handleItem = ()=>{
    emit('checkItem');
}
const handleProduct = ()=>{
    emit('getproduct')
}
let isShow = computed(()=>prop.isShow)
const proc = ref(["Profile","Product","Coin: 0","Logout",]);
</script>

<template>
    <div class="w-full h-auto bg-gray-500 mb-4 drop-shadow-lg flex justify-end px-5 py-2 hover:cursor-pointer">
        <div class=" relative flex flex-row gap-5 justify-center items-center">
            <BaseButton @click="handleItem" size="small" theme="circular">
                <IconShoppingCart color="#000000"  class="absolute"/>
            </BaseButton>
            <BaseButton class="flex justify-center items-center" size="small" theme="first" @click="handleClick">
                <IconUser/>
                <h4 class=" text-white">{{username === '' ? 'Sign-in':username}}</h4>
            </BaseButton>
            <BaseTypeList v-if="isShow" mode="account" @product="handleProduct" @profile="handleProfile" @logout="handlelogout" :product="proc" class="absolute left-0 top-full mt-2 max-h-30 w-full overflow-y-auto drop-shadow-lg z-10">
            </BaseTypeList>
        </div>
        
    </div>
    
</template>
