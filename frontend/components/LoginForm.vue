<script setup>
const emit = defineEmits(['submitLogin','gotoRegister','gotoForgetpass']);
const {handleSubmit, isSubmitting , setFieldError} = useForm({
    validationSchema: useLoginValidationSchema(),
    validateOnInput:false,
});
const onSubmit = handleSubmit(values =>{
    emit('submitLogin',{values,setFieldError})
});

let regisForm = ref(false);
let forgetPassword = ref(false);

const handleRegister = ()=>{
    emit('gotoRegister');
};
const handleForgetPassword = ()=>{
    emit('gotoForgetpass');
}

</script>
<template>
    <form  @submit.prevent="onSubmit" class="min-h-screen flex justify-center items-center bg-neutral-50 py-12 px-4 sm:px-6 lg:px-8">
        <div class="max-w-md w-full space-y-8 bg-white p-10 rounded-2xl shadow-xl border border-neutral-100 flex flex-col items-center">
            <h1 v-if="!regisForm && !forgetPassword" class="text-center text-3xl font-extrabold text-neutral-900 mb-8">Sign-in</h1>
            <BaseInput name="username" placeholder="Username" width="w-full" height="h-12"/>
            <BaseInput name="password" placeholder="Password" type="password" width="w-full" height="h-12"/>
            <div class="flex justify-between items-center w-full mt-6 space-x-4">
                <BaseButton v-if="!regisForm" size="small" theme="fourth" @click="handleRegister" type="button" class="whitespace-nowrap">Sign-Up</BaseButton>  
                <BaseButton v-if="!isSubmitting" size="large" theme="first" type="submit" class="w-full">Sign-in</BaseButton>
                <BaseButton size="small" theme="fourth" @click="handleForgetPassword" type="button" class="text-xs text-right">Forget<br>password</BaseButton> 
            </div>
            
        </div>
    </form>
</template>
