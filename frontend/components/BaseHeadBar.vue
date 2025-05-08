<script setup>
const prop = defineProps({
    username:String,
})
const emit = defineEmits(['isDropuser','logout']);
const handleClick = ()=>{
    isShow.value = !isShow.value;
    emit('isDropuser');
}
const handlelogout = () => {
    emit('logout');
}
const isShow = ref(!(prop.username === ''));
const proc = ref([prop.username,"Setting","Coin: 0","Logout",]);
</script>

<template>
    <div class="w-full h-auto bg-gray-500 mb-4 drop-shadow-lg flex justify-end px-5 py-2 hover:cursor-pointer">
        <div class=" relative">
            <div class="flex justify-center items-center gap-2 rounded-e-2xl hover:underline" @click="handleClick">
                <IconUser/>
                <h4 class=" text-white">{{username === '' ? 'Sign-in':username}}</h4>
            </div>
            <BaseTypeList v-if="isShow" @logout="handlelogout" :product="proc" class="absolute left-0 top-full mt-2 max-h-30 w-full overflow-y-auto drop-shadow-lg z-10">
            </BaseTypeList>
        </div>
        
    </div>
    
</template>
