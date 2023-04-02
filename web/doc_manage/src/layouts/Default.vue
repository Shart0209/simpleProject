<script setup>
import Nav from '@/components/Nav.vue'
import Footer from '@/components/Footer.vue'
</script>

<template>
  <header>
    <Nav />
  </header>

  <main>
    <section class="pt-4 pb-0">
      <div class="container">
        <div class="row">
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