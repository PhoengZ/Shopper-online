<script setup>
import { getProfile, validateToken } from '~/repositories/auth'
import { useAuthStore } from '~/Stores/auth'

definePageMeta({ layout: false })
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
</script>

<template>
  <div class="min-h-screen w-full bg-gray-100 py-16 px-4 flex justify-center items-start">
    <section class="bg-white w-full max-w-4xl rounded-3xl shadow-xl p-8 flex flex-col gap-8">
        <BaseButton size="small" theme="circular" @click="handleBack"  class="absolute" ><IconBackArrow color="#000000"  class="absolute"/></BaseButton>
        <h1 class="text-center text-4xl font-bold text-gray-800 flex items-center justify-center gap-5">
            <IconUser color="#000000"/>Profile</h1>
        <div class="grid grid-cols-1 md:grid-cols-2 gap-6 items-center">
        <div class="flex justify-center">
          <BaseImage
            url="https://f.ptcdn.info/090/014/000/1388837662-pantiptalk-o.png"
            class="w-32 h-32 rounded-full border-4 border-blue-400 hover:scale-105 transition-transform duration-300"
          />
        </div>

        <div>
          <h2 class="text-xl font-semibold text-gray-700 mb-2 text-center md:text-left">Delivery Address</h2>
          <textarea
            class="w-full h-24 border-2 border-gray-300 rounded-xl px-4 py-2 resize-none text-sm"
            placeholder="Enter your delivery address..."
            :value="address"
            :disabled="!isEditAddress"
          ></textarea>
        </div>
      </div>

      <div class="grid grid-cols-1 md:grid-cols-2 gap-4">
        <div>
          <label class="block text-gray-600 font-medium mb-1">Username</label>
          <input
            type="text"
            class="w-full border-2 border-gray-300 rounded-xl px-4 py-2"
            disabled
            :value="username"
          />
        </div>
        <div>
          <label class="block text-gray-600 font-medium mb-1">Password</label>
          <input
            type="password"
            class="w-full border-2 border-gray-300 rounded-xl px-4 py-2"
            disabled
            value="ChangePasswordHere"
          />
        </div>
      </div>

      <div class="flex flex-col md:flex-row justify-center gap-4 mt-4">
        <BaseButton size="small" theme="second" class="px-6" @click="isEditAddress = !isEditAddress">
          {{ isEditAddress ? 'Save Address' : 'Edit Address' }}
        </BaseButton>
        <BaseButton size="small" theme="third" class="px-6">Forget Password</BaseButton>
      </div>

      <div class="text-center">
        <BaseButton size="large" theme="third" class="px-8" @click="signOut">Sign Out</BaseButton>
      </div>

      <div class="w-full bg-gray-100 p-4 rounded-2xl">
        <BaseOption :flag="openFilter" @open-filter="handleOpenfilter" class="mx-auto my-auto">
          <IconShoppingCart color="#000000"/> Order History
        </BaseOption>

        <div v-if="openFilter" class="mt-4">
          <BaseCardList :product="products" mode="profile" />
        </div>
      </div>
    </section>
  </div>
</template>
