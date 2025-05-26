<script setup>

const prop = defineProps({
    username:String,
    isShow:Boolean,
})
const emit = defineEmits(['isDropuser','logout','checkItem']);
const handleClick = ()=>{
    emit('isDropuser');
}
const handlelogout = () => {
    emit('logout');
}
const handleItem = ()=>{
    emit('checkItem');
}
const handleOutside = ()=>{
    isShow.value = false;
}
let isShow = computed(()=>prop.isShow)
const proc = ref([prop.username,"Setting","Coin: 0","Logout",]);
</script>

<template>
    <div class="w-full h-auto bg-gray-500 mb-4 drop-shadow-lg flex justify-end px-5 py-2 hover:cursor-pointer">
        <div class=" relative flex flex-row gap-5 justify-center items-center">
            <BaseButton @click="handleItem" size="small" theme="circular">
                <IconShoppingCart />
            </BaseButton>
            <BaseButton class="flex justify-center items-center" size="small" theme="first" @click="handleClick">
                <IconUser/>
                <h4 class=" text-white">{{username === '' ? 'Sign-in':username}}</h4>
            </BaseButton>
            <BaseTypeList v-if="isShow" @logout="handlelogout" mode="account" :product="proc" class="absolute left-0 top-full mt-2 max-h-30 w-full overflow-y-auto drop-shadow-lg z-10"
            v-click-outside="handleOutside">
            </BaseTypeList>
        </div>
        
    </div>
    
</template>
