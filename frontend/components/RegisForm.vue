<script setup>
const emit = defineEmits(['submitRegister','submitLogin','submitForget']);
const {handleSubmit, isSubmit} = useForm({
    validationSchema:useRegisterValidationSchema(),
    validateOnInput:true,
});
const onSubmit = handleSubmit(values=>{
    emit('submitRegister',values);
})
const onForget = ()=>{
    emit('submitForget');
}
const onLogin = ()=>{
    emit('submitLogin');
}

</script>
<template>
    <form @submit.prevent="onSubmit" class="min-h-screen flex justify-center items-center bg-neutral-50 py-12 px-4 sm:px-6 lg:px-8">
        <div class="max-w-md w-full space-y-8 bg-white p-10 rounded-2xl shadow-xl border border-neutral-100 flex flex-col items-center">
            <h1 class="text-center text-3xl font-extrabold text-neutral-900 mb-8">Sign-up</h1>
            <BaseInput name="username" placeholder="Username..." height="h-12" width="w-full"/>
            <BaseInput name="password_1" placeholder="Password..." type="password" height="h-12" width="w-full"/>
            <BaseInput name="password_2" placeholder="Verified password..." type="password" height="h-12" width="w-full"/>
            <div class="flex justify-between items-center w-full mt-6 space-x-4">
                <BaseButton size="small" theme="fourth" @click="onLogin" type="button" class="whitespace-nowrap">Sign-in</BaseButton>
                <BaseButton v-if="!isSubmit" size="large" theme="first" type="submit" class="w-full">Sign-up</BaseButton>
                <BaseButton size="small" theme="fourth" @click="onForget" type="button" class="text-xs text-right">Forget<br>password</BaseButton>
            </div>
        </div>
    </form>
</template>