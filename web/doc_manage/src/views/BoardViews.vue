<script setup>
import { useRouter } from 'vue-router';
import { storeToRefs } from 'pinia';

import Board from '@/components/Board.vue';

import { useDocsStore } from '@/stores';
import { useAuthStore } from '@/stores';

const authStore = useAuthStore();
const { getAllDocs } = useDocsStore();
const { items } = storeToRefs(useDocsStore());

const router = useRouter();

function createForm() {
  // redirect to create page
  router.push({ name: 'add', replace: true });
}

getAllDocs()

</script>

<template>
  <section>
    <!-- Component Board -->
    <div class="col-md-10 offset-md-1">
      <button v-show="authStore.user" @click="createForm" type="button" class="btn btn-primary mb-3">Create</button>
      <Board :items="items.data" />
    </div>
  </section>
  <div v-if="items.error" class="col-md-10">
    <p>Ооопля, пусто однако!</p>
  </div>
</template>


