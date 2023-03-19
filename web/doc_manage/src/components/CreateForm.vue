<script setup>

const emit = defineEmits(['create'])

const props = defineProps({
  docs: {
    type: Object,
    required: true
  }
})

function create() {
  props.docs.attr.id = Date.now()
  emit('create', props.docs.attr)
  props.docs.attr = {
    title: '',
    number: '',
    date: '',
    start_date: '',
    end_date: '',
    category: 'Способ заключения',
    distributor: 'Поставщик',
    description: '',
    price: 0,
    status: true
  }
}
</script>

<template>
  <div class="container">
    <div class="col-md-7 col-lg-8">
      <form method="POST" class="needs-validation" novalidate enctype="multipart/form-data" @submit.prevent>
        <div class="row g-3">
          <div class="col-12 mb-2">
            <input
                type="text"
                class="form-control"
                v-model="docs.attr.title"
                placeholder="Наименование"
            >
            <div class="invalid-feedback">
              Valid title is required.
            </div>
          </div>
          <div class="col-sm-6">
            <input
                type="text"
                class="form-control"
                v-model="docs.attr.number"
                placeholder="Номер"
            >
            <div class="invalid-feedback">
              Valid number is required.
            </div>
          </div>
          <div class="col-sm-6">
            <input
                type="number"
                class="form-control"
                v-model="docs.attr.price"
                placeholder="Цена контракта"
            >
            <div class="invalid-feedback">
              Valid price is required.
            </div>
          </div>
          <div class="col-sm-4 mb-2">
            <label for="date" class="form-label">Дата заключения</label>
            <input
                type="date"
                class="form-control"
                v-model="docs.attr.date"
            >
            <div class="invalid-feedback">
              Valid date is required.
            </div>
          </div>
          <div class="col-sm-4 mb-2">
            <label for="start_date" class="form-label">Дата начала действия</label>
            <input
                type="date"
                class="form-control"
                v-model="docs.attr.start_date"
            >
            <div class="invalid-feedback">
              Valid date is required.
            </div>
          </div>
          <div class="col-sm-4 mb-2">
            <label for="end_date" class="form-label">Дата окончания действия</label>
            <input
                type="date"
                class="form-control"
                v-model="docs.attr.end_date"
            >
            <div class="invalid-feedback">
              Valid date is required.
            </div>
          </div>
          <div class="col-sm-6">
            <select
                class="form-floating form-select"
                required
                v-model="docs.attr.category"
            >
              <option disabled selected>Способ заключения</option>
              <option
                  v-for="item in docs.docsArr.categories"
                  :key=item.id
              >{{ item.name }}
              </option>
            </select>
          </div>
          <div class="col-sm-6">
            <select
                class="form-floating form-select"
                required
                v-model="docs.attr.distributor">
              <option disabled selected>Поставщик</option>
              <option
                  v-for="item in docs.docsArr.distributors"
                  :key=item.id
              >{{ item.name }}
              </option>
            </select>
          </div>
          <div class="col-sm-12">
            <textarea
                class="form-control"
                v-model="docs.attr.description"
                row="3"
                placeholder="Примечание"
            ></textarea>
          </div>
          <div class="col-sm-12 mb-2">
            <input
                type="file"
                class="form-control"
                multiple
            >
          </div>
        </div>
        <button @click="create" class="btn btn-primary">Ok</button>
        <button @click="create" class="btn btn-secondary">Cancel</button>
      </form>
    </div>
  </div>
</template>

<style lang="scss" scoped>
.container {
  max-width: 960px;
}
</style>