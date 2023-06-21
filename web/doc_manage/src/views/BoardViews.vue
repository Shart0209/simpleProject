<script setup>
import { useRouter } from 'vue-router';
import { storeToRefs } from 'pinia';

import Board from '@/components/Board.vue';

import { useDocsStore } from '@/stores';
import { useAuthStore } from '@/stores';

const authStore = useAuthStore();
const { getAllDocs } = useDocsStore();
const { items, optionsSelect } = storeToRefs(useDocsStore());

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
      <div class="row justify-content-between gy-5 mb-3">
        <div class="col-sm-6 ">
          <div class="btn-group" role="group">
            <button v-show="authStore.user.isActive" @click="createForm" type="button"
              class="btn btn-primary me-1">Create</button>
            <button type="button" class="btn btn-primary" @click="refreshForm">Обновить</button>
          </div>
        </div>
        <div class="col-sm-3">
          <select class="form-floating form-select">
            <option disabled selected value="">Способ заключения</option>
            <option v-for="item in optionsSelect.categories" :key=item.id :value="item.id">
              {{ item.name }}
            </option>
          </select>
        </div>
        <div class="col-sm-3">
          <select class="form-floating form-select">
            <option disabled selected value="">Группа</option>
            <option v-for="item in optionsSelect.groups" :key=item.id :value="item.id">{{ item.name }}
            </option>
          </select>
        </div>
      </div>
      <Board :items="items.data" />
    </div>
  </section>
  <div v-if="items.error" class="col-md-10">
    <p>Ооопля, пусто однако!</p>
  </div>
</template>


