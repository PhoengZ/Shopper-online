<script setup>
import User from '~/components/Icon/User.vue'
import { updateProfile } from '~/repositories/auth'
import { requestTransaction } from '~/repositories/topup'
import { useAuthStore } from '~/Stores/auth'

const {$toast} = useNuxtApp()

useHead({
    title:'Top up'
})
const user = useAuthStore()
const money = ref(1)
const {handleSubmit, isSubmiting} = useForm({
    validationSchema:usePaymentValidationSchema(),
    validateOnInput:true,
    initialValues:{
        amount:money.value || 1
    }
})
const prevClick = ()=>{
    navigateTo("/")
}
const confirmClick = handleSubmit(async (value)=>{
    const object = {
        "password":"",
        "address":"",
        "coin":value.amount*2,
        "history":[]
    }
    try{
        const {message_1} = await requestTransaction(user.userID, value.amount * 2, user.token)
        const {message_2} = await updateProfile(user.userID, object, user.token)
        $toast.success('ทำรายการสำเร็จโปรดรอการตรวจสอบจาก Admin',{
            style:{
                background:'green',
                color:'white'
            }
        })
    }catch(err){
        console.error(err.data)
        $toast.error('ทำรายการไม่สำเร็จโปรดลองใหม่อีกครั้ง',{
            description:'รายละเอียด: '+err.data.error
        })
    }
})  

const totalCoin = ref(money.value * 2)
watch(money, newValue =>{
    totalCoin.value = newValue * 2;
})
</script>
<template>
    <div class="min-h-screen w-full bg-neutral-50 py-12 px-4 sm:px-6 lg:px-8">
        <section class="bg-white w-full max-w-4xl mx-auto rounded-2xl shadow-sm border border-neutral-100 p-8 flex flex-col gap-8 relative">
            <div class="absolute top-8 left-8">
                <BaseButton size="small" theme="circular" @click="prevClick"><IconBackArrow color="#4b5563"/></BaseButton>
            </div>
            
            <h3 class="text-center text-3xl font-extrabold text-neutral-900 flex justify-center items-center gap-3 mt-2">
                <img src="../assets/Image/thai_qr_payment.png" alt="" class="h-10 w-10 object-contain">Payment
            </h3>
            
            <div class="max-w-2xl mx-auto w-full bg-neutral-50 rounded-xl border border-neutral-200 p-8 space-y-6">
                <div class="flex items-center justify-between">
                    <div class="font-bold text-lg text-neutral-700">Also Pay</div>
                    <div class="flex justify-end"><img src="../assets/Image/PromptPay.png" alt="" class="h-8 object-contain"></div>
                </div>
                
                <hr class="border-neutral-200">
                
                <div class="grid grid-cols-1 md:grid-cols-2 gap-6 items-center">
                    <div class="font-bold text-lg text-neutral-700">Amount to be paid</div>
                    <div class="w-full">
                        <BaseInput name="amount" :-update="true" width="w-full" placeholder="Amount to be paid" v-model:modelvalue="money" type="number"/>
                    </div>
                </div>
                
                <div class="flex items-center justify-between pt-4">
                    <div class="font-bold text-lg text-neutral-700">Amount coin</div>
                    <div class="font-bold text-2xl text-primary-600">{{totalCoin}} 🪙</div>
                </div>
            </div>
            
            <div class="flex justify-center mt-4">
                <BaseButton size="large" theme="second" class="px-12 py-3 text-lg shadow-lg hover:shadow-xl transform hover:-translate-y-1 transition-all" @click="confirmClick">
                    {{isSubmiting ? 'Confirming...' : 'Confirm Payment'}}
                </BaseButton> 
            </div>
        </section>
    </div>
</template>