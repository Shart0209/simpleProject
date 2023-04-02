<script setup>

const props = defineProps({
  attrs: Object,
  selectedFiles: Array,
  optionsSelect: Object

})

function onChangeCategory(e) {
  props.attrs.category = e.target.value;
}

function onChangeDescription(e) {
  props.attrs.distributor = e.target.value;
}

function onChangeFile(e) {

  if (e.target.files) {
    for (let i = 0; i < e.target.files.length; i++) {
      props.selectedFiles.push(e.target.files[i])
    }
    // TODO добаляется какое то поле value. Решено удалить.
    delete props.selectedFiles['value']
  }
}

</script>

<template>
  <form method="POST" class="needs-validation" novalidate enctype="multipart/form-data" @submit.prevent>
    <div class="row g-3">
      <div class="col-12 mb-2">
        <input type="text" class="form-control" v-model="attrs.title" placeholder="Наименование">
        <div class="invalid-feedback">
          Valid title is required.
        </div>
      </div>
      <div class="col-sm-6">
        <input type="text" class="form-control" v-model="attrs.number" placeholder="Номер">
        <div class="invalid-feedback">
          Valid number is required.
        </div>
      </div>
      <div class="col-sm-6">
        <input type="number" class="form-control" v-model="attrs.price" placeholder="Цена контракта">
        <div class="invalid-feedback">
          Valid price is required.
        </div>
      </div>

      <div class="col-sm-4 mb-2">
        <label for="date" class="form-label">Дата заключения</label>
        <input type="date" class="form-control" v-model="attrs.date">
        <div class="invalid-feedback">
          Valid date is required.
        </div>
      </div>
      <div class="col-sm-4 mb-2">
        <label for="start_date" class="form-label">Дата начала действия</label>
        <input type="date" class="form-control" v-model="attrs.start_date">
        <div class="invalid-feedback">
          Valid date is required.
        </div>
      </div>
      <div class="col-sm-4 mb-2">
        <label for="end_date" class="form-label">Дата окончания действия</label>
        <input type="date" class="form-control" v-model="attrs.end_date">
        <div class="invalid-feedback">
          Valid date is required.
        </div>
      </div>
      <div class="col-sm-6">
        <select class="form-floating form-select" @change="onChangeCategory($event)">
          <option disabled selected>Способ заключения</option>
          <option v-for="item in optionsSelect.categories" :key=item.id :value="item.id">{{ item.name }}
          </option>
        </select>
      </div>
      <div class="col-sm-6">
        <select class="form-floating form-select" @change="onChangeDescription($event)">
          <option disabled selected>Поставщик</option>
          <option v-for="item in optionsSelect.distributors" :key=item.id :value="item.id">{{ item.name }}
          </option>
        </select>
      </div>
      <div class="col-sm-12">
        <textarea class="form-control" v-model="attrs.description" row="3" placeholder="Примечание"></textarea>
      </div>
      <div class="col-sm-12 mb-2">
        <input type="file" @change="onChangeFile($event)" ref="file" class="form-control" multiple>
      </div>
    </div>
  </form>
</template>

<style lang="scss" scoped></style>