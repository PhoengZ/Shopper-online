<script setup>
import { useAuthStore } from '~/Stores/auth';

definePageMeta({
    layout:false,
});
useHead({
    title:"Sign-in",
});
const user = useAuthStore();
async function onLogin({ values, setFieldError }){
    const check = await user.Login(values.username,values.password);
    console.log(check);
    if (check){
        await navigateTo('/');
    }else{
        setFieldError('username', 'Invalid username or password');
        setFieldError('password', 'Invalid username or password');
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