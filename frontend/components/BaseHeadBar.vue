<script setup>
const emit = defineEmits(['isDropuser']);
const handleClick = ()=>{
    emit('isDropuser');
    console.log("isDroped");
}

</script>

<template>
    <div class="w-full h-auto bg-gray-500 mb-4 drop-shadow-lg flex justify-end px-5 py-2">
        <div class="flex justify-center items-center gap-2 rounded-e-2xl hover:bg-gray-900" @click="handleClick">
            <IconUser/>
            <h4 class=" text-white">Sign-in</h4>
        </div>
        
    </div>
    
</template>
