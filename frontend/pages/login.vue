<script setup>
import RegisForm from '~/components/RegisForm.vue';
import { useAuthStore } from '~/Stores/auth';

definePageMeta({
    layout:false,
});
useHead({
    title:"Sign-in",
});
let reg = ref(false);
const user = useAuthStore();
async function onLogin(value){
    const check = await user.Login(value.username,value.password);
    if (check){
        await navigateTo('/');
    }
}
function onRegsiter(){
    console.log('yeah');
}
function onForgetpass(){
    console.log('yo');
    
}
</script>
<template>
    <LoginForm @submitLogin="onLogin" @gotoRegister="onRegsiter" @gotoForgetpass="onForgetpass"/>
</template>