<script setup>
import { useAuthStore } from '~/Stores/auth';

definePageMeta({
    layout:false,
});
useHead({
    title:"Sign-in",
});
const user = useAuthStore();
async function onLogin(value){
    const check = await user.Login(value.username,value.password);
    if (check){
        await navigateTo('/');
    }
}
function onRegsiter(){
    navigateTo("/signUp");
}
function onForgetpass(){
    console.log('yo');
    
}
</script>
<template>
    <LoginForm @submitLogin="onLogin" @gotoRegister="onRegsiter" @gotoForgetpass="onForgetpass"/>
</template>