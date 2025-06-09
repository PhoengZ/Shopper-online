<script setup>
definePageMeta({
    layout:'auth',
})
useHead({
    title:'Top up'
})
const {handleSubmit, isSubmiting} = useForm({
    validationSchema:usePaymentValidationSchema(),
    validateOnInput:true,
})
const money = ref()
const prevClick = ()=>{
    navigateTo("/")
}
const confirmClick = handleSubmit(async value=>{

})
</script>
<template>
    <BaseButton class=" absolute" size="small" theme="circular" @click="prevClick">
        <IconBackArrow color="#000000" class="absolute"/>
    </BaseButton>
    <h3 class="text-center text-3xl rounded-t-2xl p-3 text-gray-700 font-bold flex justify-center items-center gap-5"><img src="../assets/Image/thai_qr_payment.png" alt="" class="h-10 w-10">Payment</h3>
    <div class=" grid grid-cols-2 px-10">
        <div class="flex items-center text-left font-bold text-lg">Also Pay</div>
        <div class="font-bold text-lg flex justify-end"><img src="../assets/Image/PromptPay.png" alt="" class=" h-10 w-20 my-auto"></div>
        <hr class=" col-span-2 border-t-2 border-gray-600 my-5">
        <div class="flex items-center text-left font-bold text-lg">Amount to be paid</div>
        <div class="font-bold text-lg flex justify-end">
            <BaseInput name="amount" :-update="true" width="w-full" placeholder="Amount to be paid" v-model:modelvalue="money" type="number"/>
        </div>
    </div>
    <BaseButton size="large" theme="second" @click="">{{isSubmiting ? 'Confirming' : 'Confirm'}}</BaseButton>
    
</template>