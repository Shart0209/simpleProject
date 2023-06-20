<script setup>
import { computed } from 'vue';
import { storeToRefs } from 'pinia';
import { useDocsStore } from '@/stores';

const props = defineProps({
  updItem: {
    required: true
  }
});

const { optionsSelect, error } = storeToRefs(useDocsStore());

const checkedNames = computed(() => props.updItem.status === true ? 'Действителен' : 'Не действителен')

function onChangeFile(e) {
  if (e.target.files) {
    for (let i = 0; i < e.target.files.length; i++) {
      props.updItem.files.push(e.target.files[i])
    }
  }
}

</script>

<template>
  <div class="col-md-6">
    <h4 class="mb-2">Редактирование </h4>
    <form method="POST" class="was-validated" enctype="multipart/form-data" @submit.prevent>
      <div class="row g-3">
        <div class="col-12 mb-2">
          <input type="text" class="form-control" maxlength="100" v-model.trim="updItem.title" placeholder="Наименование"
            required>
          <div class="invalid-feedback">
            Valid title is required.
          </div>
        </div>
        <div class="col-sm-6">
          <input type="text" class="form-control" maxlength="50" v-model.trim="updItem.number" placeholder="Номер"
            required>
          <div class="invalid-feedback">
            Valid number is required.
          </div>
        </div>
        <div class="col-sm-6">
          <input type="number" class="form-control" v-model.trim="updItem.price" placeholder="Цена контракта" required>
          <div class="invalid-feedback">
            Valid price is required.
          </div>
        </div>

        <div class="col-sm-4 mb-2">
          <label>Дата заключения</label>
          <input type="date" class="form-control" v-model="updItem.date" required>
          <div class="invalid-feedback">
            Valid date is required.
          </div>
        </div>
        <div class="col-sm-4 mb-2">
          <label>Дата начала действия</label>
          <input type="date" class="form-control" v-model="updItem.start_date" required>
          <div class="invalid-feedback">
            Valid date is required.
          </div>
        </div>
        <div class="col-sm-4 mb-2">
          <label>Дата окончания действия</label>
          <input type="date" class="form-control" v-model="updItem.end_date" required>
          <div class="invalid-feedback">
            Valid date is required.
          </div>
        </div>
        <div class="col-sm-6">
          <select class="form-floating form-select" v-model="updItem.category" required>
            <option disabled selected value="">Способ заключения</option>
            <option v-for="data in optionsSelect.categories" :key=data.id :value="data">
              {{ data.name }}
            </option>
          </select>
        </div>
        <div class="col-sm-6">
          <select class="form-floating form-select" v-model="updItem.group" required>
            <option disabled selected value="">Группа</option>
            <option v-for="data in optionsSelect.groups" :key=data.id :value="data">{{ data.name }}
            </option>
          </select>
        </div>
        <div class="col-sm-6">
          <select class="form-floating form-select" v-model="updItem.supplier" required>
            <option disabled selected value="">Поставщик</option>
            <option v-for="data in optionsSelect.suppliers" :key=data.id :value="data">{{ data.name }}
            </option>
          </select>
        </div>
        <div class="col-sm-6">
          <div class="form-check">
            <input class="form-check-input" type="checkbox" value="" v-model="updItem.status" id="checkbox">
            <label class="form-check-label" for="checkbox">
              Статус контракта: {{ checkedNames }}
            </label>
          </div>
        </div>
        <div class="col-sm-12">
          <textarea class="form-control" maxlength="150" v-model.trim="updItem.description" row="3"
            placeholder="Примечание"></textarea>
        </div>
        <div class="col-sm-12 mb-2">
          <input type="file" @change="onChangeFile($event)" ref="file" class="form-control" multiple>
        </div>
      </div>
      <div class="d-grid gap-2 d-md-flex justify-content-md-end mt-2">
        <button type="submit" class="btn btn-primary" @click="$emit('ok')">OK</button>
        <button type="button" class="btn btn-primary" @click="$emit('close')">Close</button>
      </div>
    </form>
    <div v-show="error">
      <span>{{ error }}</span>
    </div>
  </div>
</template>