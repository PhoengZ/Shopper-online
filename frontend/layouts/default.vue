<template>
  <div class="min-h-screen flex flex-col font-sans text-neutral-900 bg-neutral-50">
    <!-- The Header is now handled within pages or we can make it global if we pass props correctly. 
         However, the original code passed many props to TheHeader. 
         For a cleaner redesign, we should ideally use a store or composable for state, 
         but to keep logic untouched as requested, we might need to keep the header in the page 
         OR adapt the layout to handle these props if they are global.
         
         Looking at index.vue, it passes specific props to TheHeader. 
         If I move Header to layout, I need to ensure those props are available.
         
         The original default.vue just had <NuxtPage />.
         The Header was in index.vue.
         
         To make it "professional" and "consistent", a global header is better.
         But since logic shouldn't be touched, and props like 'coin', 'username' come from page logic,
         it's safer to keep the Header in the Page for now, OR use slots/teleport.
         
         However, the plan said "Add a persistent Header and Footer".
         Let's add the Footer here as it's likely static or simple.
         For the Header, since it relies on page-level data (coin, user), 
         we might need to leave it in the page OR refactor the data fetching to be global (e.g. in app.vue or layout).
         
         Given "do not touch logic", I will add the Footer here and wrap NuxtPage.
         The Header in index.vue will be replaced by AppHeader.
    -->
    
    <main class="flex-grow">
      <NuxtPage />
    </main>
    
    <AppFooter />

    <ClientOnly>
      <Toaster
        richColors
        :toastOptions="{
          duration: 4000,
          style: {
            position: 'fixed',
            left: '50%',
            top: '20px',
            transform: 'translateX(-50%)',
            zIndex: 9999,
            background: '#dc2626',
            color: 'white',
            borderRadius: '12px',
            padding: '16px 20px',
            boxShadow: '0 10px 25px rgba(0,0,0,0.2)'
          }
        }"
      />
    </ClientOnly>
  </div>
</template>

<style>
[data-sonner-toast] {
  animation: fadeIn 0.4s ease, fadeOut 0.4s ease 3.6s forwards;
}

@keyframes fadeIn {
  from {
    opacity: 0;
    transform: translate(-50%, -10px); /* รวม x, y */
  }
  to {
    opacity: 1;
    transform: translate(-50%, 0);
  }
}

@keyframes fadeOut {
  from {
    opacity: 1;
    transform: translate(-50%, 0);
  }
  to {
    opacity: 0;
    transform: translate(-50%, -10px);
  }
}

</style>
