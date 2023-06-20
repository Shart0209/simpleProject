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

async function createForm() {
  // redirect to create page
  router.push({ name: 'add', replace: true });
}

async function refreshForm() {
  getAllDocs();
}

getAllDocs();

</script>

<template>
  <section>
    <!-- Component Board -->
    <div class="col-md-10 offset-md-1">
      <div class="d-grid gap-2 d-md-flex mb-3">
        <button v-show="authStore.user.isActive" @click="createForm" type="button" class="btn btn-primary">Create</button>
        <button type="button" class="btn btn-primary" @click="refreshForm">Обновить</button>
      </div>
      <Board :items="items.data" />
    </div>
  </section>
  <div v-if="items.error" class="col-md-10">
    <p>Ооопля, пусто однако!</p>
  </div>
</template>


