<script setup>
import { useAuthStore } from '~/Stores/auth';

definePageMeta({
    layout:false,
});
useHead({
    title:"Sign-up",
});
const user = useAuthStore();
async function onRegister(values) {
    const check = await user.Register(values.username,values.password_1);
    if (check){
        navigateTo('/login')
    }
}
const onForget = ()=>{
    console.log("Onforgetpassword");
}

const onLogin = ()=>{
    navigateTo('/login');
}

</script>
<template>
    <RegisForm @submitRegister="onRegister" @submitLogin="onLogin" @submitForget="onForget"/>
</template>