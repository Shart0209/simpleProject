<script setup>
import {onMounted, onUnmounted, ref} from "vue";
import {Modal} from "bootstrap";

const emit = defineEmits(['create'])
defineExpose({ show: _show });

const props = defineProps({
  attr: {type: Object},
  docsArr: {type: Object}
})

let modalEle = ref(null);
let thisModalObj = null;
function _show() {
  thisModalObj.show();
}
function create() {
  props.attr.id = Date.now()
  emit('create',  props.attr)

  thisModalObj.hide()
}

onMounted(() => {
  thisModalObj = new Modal(modalEle.value);
});

</script>

<template>
  <div class="modal fade" ref="modalEle" data-bs-backdrop="static" tabindex="-1" aria-hidden="true">
  <div class="modal-dialog">
      <div class="modal-content">
        <div class="modal-header">
          <h5 id="infoModalLabel">Добавление записи</h5>
        </div>
        <div class="modal-body">
          <form method="POST" class="needs-validation" novalidate enctype="multipart/form-data" @submit.prevent>
            <div class="row g-3">
              <div class="col-12 mb-2">
                <input
                    type="text"
                    class="form-control"
                    v-model="attr.title"
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
                    v-model="attr.number"
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
                    v-model="attr.price"
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
                    v-model="attr.date"
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
                    v-model="attr.start_date"
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
                    v-model="attr.end_date"
                >
                <div class="invalid-feedback">
                  Valid date is required.
                </div>
              </div>
              <div class="col-sm-6">
                <select
                    class="form-floating form-select"
                    required
                    v-model="attr.category"
                >
                  <option disabled selected>Способ заключения</option>
                  <option
                      v-for="item in  docsArr.categories"
                      :key=item.id
                  >{{ item.name }}
                  </option>
                </select>
              </div>
              <div class="col-sm-6">
                <select
                    class="form-floating form-select"
                    required
                    v-model=" attr.distributor">
                  <option disabled selected>Поставщик</option>
                  <option
                      v-for="item in  docsArr.distributors"
                      :key=item.id
                  >{{ item.name }}
                  </option>
                </select>
              </div>
              <div class="col-sm-12">
            <textarea
                class="form-control"
                v-model=" attr.description"
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
          </form>
        </div>
        <div class="modal-footer">
          <button @click="create" class="btn btn-primary">Ok</button>
          <button type="button" class="btn btn-secondary" @click="thisModalObj.hide()">Close</button>
        </div>
      </div>
    </div>
  </div>
</template>

<style lang="scss" scoped>

</style>