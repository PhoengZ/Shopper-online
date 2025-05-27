<script setup>
definePageMeta({ layout: false })
useHead({ title: "Profile" })

const isEditAddress = ref(false)
const openFilter = ref(false)
const handleOpenfilter = () => {
  openFilter.value = !openFilter.value
}

const products = [
  {
    name: "iPhone 16 Pro",
    description: "The latest iPhone with A18 Pro chip, 120Hz ProMotion display, and titanium body.",
    price: 1299,
    quantity: 2
  },
  {
    name: "MacBook Air M3",
    description: "Lightweight and powerful with M3 chip, perfect for everyday work and travel.",
    price: 1099,
    quantity: 1
  },
  {
    name: "Apple Watch Series 10",
    description: "Redesigned with bigger screen, longer battery life, and new health sensors.",
    price: 499,
    quantity: 3
  }
]
</script>

<template>
  <div class="min-h-screen w-full bg-gray-100 py-16 px-4 flex justify-center items-start">
    <section class="bg-white w-full max-w-4xl rounded-3xl shadow-xl p-8 flex flex-col gap-8">
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
            value="ChangeYourUsernameHere"
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
        <BaseButton size="large" theme="third" class="px-8">Sign Out</BaseButton>
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
