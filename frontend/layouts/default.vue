<template>
  <div>
    <NuxtPage />
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
