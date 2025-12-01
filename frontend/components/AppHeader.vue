<script setup>
const props = defineProps({
  coin: { type: Number, default: 0 },
  choiceItem: { type: Array, default: () => [] },
  username: { type: String, default: '' },
  isDrop: { type: Boolean, default: false },
  openBlure: { type: Boolean, default: false }
});

const emit = defineEmits(['topup', 'getproduct', 'profile', 'searchItem', 'logout', 'auth', 'checkItem']);

const search = ref('');

const handleSearch = () => {
  emit('searchItem', { name: search.value, category: selectedItem.value, price: selectedSortOption.value });
};

// Filter Logic
const isShowed = ref(false);
const selectedItem = ref([]);
const selectedSortOption = ref(false);

const showFilter = () => {
  isShowed.value = !isShowed.value;
};

const handleOutside = () => {
  isShowed.value = false;
};

</script>

<template>
  <header class="sticky top-0 z-50 bg-white/80 backdrop-blur-md border-b border-neutral-200 shadow-sm">
    <div class="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8">
      <div class="flex justify-between items-center h-16">
        <!-- Logo -->
        <div class="flex-shrink-0 flex items-center cursor-pointer" @click="navigateTo('/')">
          <h1 class="text-2xl font-bold bg-gradient-to-r from-primary-600 to-primary-400 bg-clip-text text-transparent">
            Shopper<span class="text-neutral-900">Online</span>
          </h1>
        </div>

        <!-- Search Bar -->
        <div class="hidden md:flex flex-1 max-w-lg mx-8">
          <div class="relative w-full">
            <div class="absolute inset-y-0 left-0 pl-3 flex items-center pointer-events-none">
              <svg class="h-5 w-5 text-neutral-400" xmlns="http://www.w3.org/2000/svg" viewBox="0 0 20 20" fill="currentColor">
                <path fill-rule="evenodd" d="M8 4a4 4 0 100 8 4 4 0 000-8zM2 8a6 6 0 1110.89 3.476l4.817 4.817a1 1 0 01-1.414 1.414l-4.816-4.816A6 6 0 012 8z" clip-rule="evenodd" />
              </svg>
            </div>
            <input
              v-model="search"
              @keyup.enter="handleSearch"
              type="text"
              class="block w-full pl-10 pr-3 py-2 border border-neutral-300 rounded-lg leading-5 bg-neutral-50 placeholder-neutral-500 focus:outline-none focus:placeholder-neutral-400 focus:border-primary-500 focus:ring-1 focus:ring-primary-500 sm:text-sm transition duration-150 ease-in-out"
              placeholder="Search products..."
            />
          </div>
        </div>
        <div class="relative">
            <BaseOption @open-filter="showFilter" :flag="!isShowed" class="hover:cursor-pointer"/>
            <BaseTypeList v-if="isShowed" v-model:selectedItem="selectedItem" v-model:selectedSortOption="selectedSortOption"  v-click-outside="handleOutside"
             mode="screening" :product="props.choiceItem" class=" absolute top-full right-0 mt-2 max-h-60 overflow-y-auto drop-shadow-xl z-50"></BaseTypeList>
        </div>
        <!-- Right Side Actions -->
        <div class="flex items-center space-x-4">
          <template v-if="username">
            <!-- Coin Display -->
            <div class="hidden sm:flex items-center px-3 py-1 bg-yellow-50 text-yellow-700 rounded-full border border-yellow-200 text-sm font-medium cursor-pointer hover:bg-yellow-100 transition" @click="$emit('topup')">
              <span class="mr-1">💰</span>
              {{ coin.toLocaleString() }}
            </div>

            <!-- Selling Button -->
            <button @click="$emit('getproduct')" class="text-neutral-600 hover:text-primary-600 font-medium text-sm transition">
              Start Selling
            </button>

            <!-- Cart Button -->
            <button @click="$emit('checkItem')" class="relative p-2 text-neutral-600 hover:text-primary-600 transition">
              <svg xmlns="http://www.w3.org/2000/svg" class="h-6 w-6" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M3 3h2l.4 2M7 13h10l4-8H5.4M7 13L5.4 5M7 13l-2.293 2.293c-.63.63-.184 1.707.707 1.707H17m0 0a2 2 0 100 4 2 2 0 000-4zm-8 2a2 2 0 11-4 0 2 2 0 014 0z" />
              </svg>
            </button>

            <!-- Profile Dropdown -->
            <div class="relative ml-3">
              <div>
                <button @click="$emit('auth')" class="flex items-center max-w-xs text-sm rounded-full focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-primary-500" id="user-menu-button" aria-expanded="false" aria-haspopup="true">
                  <span class="sr-only">Open user menu</span>
                  <div class="h-8 w-8 rounded-full bg-primary-100 flex items-center justify-center text-primary-700 font-bold">
                    {{ username.charAt(0).toUpperCase() }}
                  </div>
                </button>
              </div>

              <!-- Dropdown menu -->
              <div v-if="isDrop" class="origin-top-right absolute right-0 mt-2 w-48 rounded-md shadow-lg py-1 bg-white ring-1 ring-black ring-opacity-5 focus:outline-none z-50" role="menu" aria-orientation="vertical" aria-labelledby="user-menu-button" tabindex="-1">
                <div class="px-4 py-2 border-b border-neutral-100">
                  <p class="text-sm text-neutral-500">Signed in as</p>
                  <p class="text-sm font-medium text-neutral-900 truncate">{{ username }}</p>
                </div>
                <a href="#" @click.prevent="$emit('profile')" class="block px-4 py-2 text-sm text-neutral-700 hover:bg-neutral-50" role="menuitem">Your Profile</a>
                <a href="#" @click.prevent="$emit('topup')" class="block px-4 py-2 text-sm text-neutral-700 hover:bg-neutral-50" role="menuitem">Top Up</a>
                <a href="#" @click.prevent="$emit('logout')" class="block px-4 py-2 text-sm text-red-600 hover:bg-red-50" role="menuitem">Sign out</a>
              </div>
            </div>
          </template>

          <template v-else>
            <button @click="navigateTo('/login')" class="text-neutral-600 hover:text-primary-600 font-medium text-sm transition">
              Log in
            </button>
            <button @click="navigateTo('/signUp')" class="ml-3 inline-flex items-center px-4 py-2 border border-transparent text-sm font-medium rounded-md shadow-sm text-white bg-primary-600 hover:bg-primary-700 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-primary-500 transition">
              Sign up
            </button>
          </template>
        </div>
      </div>
    </div>
  </header>
</template>
