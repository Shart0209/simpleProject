<script setup>
import Footer from '@/components/Footer.vue'
</script>

<template>

  <main>
    <section class="pt-4 pb-0">
      <div class="container">
        <div class="row justify-content-center">
            <router-view />
        </div>
      </div>
    </section>
  </main>

  <footer>
    <Footer />
  </footer>
</template>

<style lang="scss" scoped></style>