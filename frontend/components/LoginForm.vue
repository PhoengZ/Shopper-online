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
    <form  @submit.prevent="onSubmit" class=" h-screen flex justify-center items-center">
        <div class=" h-8/12 w-6/12 flex justify-center items-center flex-col gap-5 bg-gray-300 rounded-2xl">
            <h1 v-if="!regisForm && !forgetPassword" class=" font-bold text-6xl mb-10">Sign-in</h1>
            <BaseInput name="username" placeholder="Username"/>
            <BaseInput name="password" placeholder="Password" type="password"/>
            <div class=" flex justify-between w-10/12">
                <BaseButton v-if="!regisForm" size="small" theme="fourth" @click="handleRegister" type="button">Sign-Up</BaseButton>  
                <BaseButton v-if="!isSubmitting" size="large" theme="first" type="submit">Sign-in</BaseButton>
                <BaseButton size="small" theme="fourth" @click="handleForgetPassword" type="button">Forget
                    <br>password
                </BaseButton> 
            </div>
            
        </div>
    </form>
</template>