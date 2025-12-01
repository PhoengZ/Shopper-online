<script setup>
import { getProfile, updateProfile, validateToken } from '~/repositories/auth'
import { useAuthStore } from '~/Stores/auth'

const {$toast} = useNuxtApp()
useHead({ title: "Profile" })
const user = useAuthStore()
const username = ref('')
const userID = ref('')
const token = useCookie('token')
const products = ref([])
const address = ref('')
const isEditAddress = ref(false)
const openFilter = ref(false)
const handleOpenfilter = () => {
  openFilter.value = !openFilter.value
}
const {error: err , data:validateData} = await validateToken(token.value)
if (err.value){
    navigateTo('/login')
}else{
    if (validateData.value?.message === 'Valid'){
        username.value = user.Username
        userID.value = user.userID
        const {data: userData, error: userError} = await getProfile(userID.value,token.value)
        if (userError.value){
            console.error('Failed to fetch user profile', userError.value)
        }else{
            products.value = userData.value?.history
            address.value = userData.value?.address
        }
    }else{
        navigateTo('/login')
    }
}
const handleBack = () =>{
    navigateTo('/')
}
const signOut = () =>{
    token.value = null
    user.Logout()
    username.value = ''
    userID.value = ''
    navigateTo('/login')
}
const EditAddress = async ()=>{
  if (isEditAddress.value){
    const object = {
        "password":"",
        "address":address.value,
        "coin":0,
        "history":[]
    }
    try{
      const {message}  = await updateProfile(userID.value, object, token.value)
      $toast.success(message,{
        style:{
          background:'green',
          color:'white'
        }
      })
      isEditAddress.value = !isEditAddress.value
    }catch(err){
      console.error('Failed to update address', err)
      $toast.error('ไม่สามารถแก้ไขที่อยู่ได้',{
        description:'รายละเอียด: '+err.data.error
      })
    }
  }else{
    isEditAddress.value = !isEditAddress.value
  }
}
</script>

<template>
  <div class="min-h-screen w-full bg-neutral-50 py-12 px-4 sm:px-6 lg:px-8">
    <section class="bg-white w-full max-w-4xl mx-auto rounded-2xl shadow-sm border border-neutral-100 p-8 flex flex-col gap-8 relative">
        <div class="absolute top-8 left-8">
             <BaseButton size="small" theme="circular" @click="handleBack"><IconBackArrow color="#4b5563"/></BaseButton>
        </div>
        
        <h1 class="text-center text-3xl font-extrabold text-neutral-900 flex items-center justify-center gap-3 mt-2">
            <IconUser color="#4f46e5" class="w-8 h-8"/>Profile
        </h1>

        <div class="grid grid-cols-1 md:grid-cols-2 gap-8 items-center mt-4">
            <div class="flex justify-center">
              <BaseImage
                url="https://f.ptcdn.info/090/014/000/1388837662-pantiptalk-o.png"
                class="w-40 h-40 rounded-full border-4 border-primary-100 shadow-md hover:scale-105 transition-transform duration-300 object-cover"
              />
            </div>

            <div class="space-y-2">
              <h2 class="text-lg font-semibold text-neutral-700 text-center md:text-left">Delivery Address</h2>
              <textarea
                class="w-full h-32 border border-neutral-300 rounded-lg px-4 py-3 resize-none text-sm focus:outline-none focus:ring-2 focus:ring-primary-500/20 focus:border-primary-500 transition-colors"
                placeholder="Enter your delivery address..."
                v-model="address"
                :disabled="!isEditAddress"
                :class="{'bg-neutral-50 text-neutral-500': !isEditAddress, 'bg-white text-neutral-900': isEditAddress}"
              ></textarea>
            </div>
      </div>

      <div class="grid grid-cols-1 md:grid-cols-2 gap-6">
        <div class="space-y-1">
          <label class="block text-sm font-medium text-neutral-700">Username</label>
          <input
            type="text"
            class="w-full border border-neutral-300 rounded-lg px-4 py-2.5 bg-neutral-50 text-neutral-500 focus:outline-none"
            disabled
            :value="username"
          />
        </div>
        <div class="space-y-1">
          <label class="block text-sm font-medium text-neutral-700">Password</label>
          <input
            type="password"
            class="w-full border border-neutral-300 rounded-lg px-4 py-2.5 bg-neutral-50 text-neutral-500 focus:outline-none"
            disabled
            value="ChangePasswordHere"
          />
        </div>
      </div>

      <div class="flex flex-col sm:flex-row justify-center gap-4 mt-4">
        <BaseButton size="small" theme="second" class="px-8" @click="EditAddress">
          {{ isEditAddress ? 'Save Address' : 'Edit Address' }}
        </BaseButton>
        <BaseButton size="small" theme="third" class="px-8">Forget Password</BaseButton>
      </div>

      <div class="text-center pt-4 border-t border-neutral-100">
        <BaseButton size="large" theme="third" class="px-10" @click="signOut">Sign Out</BaseButton>
      </div>

      <div class="w-full bg-neutral-50 p-6 rounded-xl border border-neutral-200 mt-4">
        <BaseOption :flag="openFilter" @open-filter="handleOpenfilter" class="mx-auto my-auto text-neutral-700 hover:text-primary-600 transition-colors">
          <IconShoppingCart color="currentColor" class="mr-2"/> Order History
        </BaseOption>

        <div v-if="openFilter" class="mt-6">
          <BaseCardList :product="products" mode="profile" />
        </div>
      </div>
    </section>
  </div>
</template>
