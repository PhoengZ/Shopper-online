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
    <form @submit.prevent="onSubmit" class="h-screen flex justify-center items-center bg-gray-200">
        <div class=" h-8/12 w-6/12 bg-white rounded-4xl flex justify-center items-center flex-col">
            <h1 class=" text-black font-bold text-6xl mb-10">Sign-up</h1>
            <BaseInput name="username" placeholder="Username..."/>
            <BaseInput name="password_1" placeholder="Password..." type="password"/>
            <BaseInput name="password_2" placeholder="Verified password..." type="password"/>
            <div class=" flex justify-between w-10/12 mt-5">
                <BaseButton size="small" theme="fourth" @click="onLogin" type="button">Sign-in</BaseButton>
                <BaseButton v-if="!isSubmit" size="large" theme="first" type="submit">Sign-up</BaseButton>
                <BaseButton size="small" theme="fourth" @click="onForget" type="button">Forget
                    <br>password
                </BaseButton>
            </div>
        </div>
    </form>
</template>