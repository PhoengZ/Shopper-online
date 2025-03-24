<script setup>
let user = ref("");
let password = ref("");
let c = ref(false);
const lg = ()=>{
    console.log("Is login!");
    console.log(user.value);
    console.log(password.value);
};


</script>
<template>
    <form  @submit.prevent="lg" class=" h-screen flex justify-center items-center">
        <div class=" h-8/12 w-6/12 flex justify-center items-center flex-col gap-2 bg-gray-300 rounded-2xl">
            <BaseInput type="text" placeholder="Username"/>
            <BaseInput type="password" placeholder="Password"/>
            <div>
                <BaseButton v-if="c" size="large" theme="first" >Sign-in</BaseButton>
                <BaseButton v-else size="large" theme="first">Sign-Up</BaseButton>  
            </div>
            
        </div>
    </form>
</template>