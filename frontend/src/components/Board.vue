<script setup>
import { computed, ref } from 'vue';
const props = defineProps({
  items: Array,
  filterKeySP: Number,
  filterKeyGR: Number,
  FilterKeyST: Boolean,
});

const filteredData = computed(() => {
  let { items, filterKeySP, filterKeyGR, FilterKeyST } = props
  items = filterKeySP != 0 ? items.filter((row) => row.category['id'] === filterKeySP) : items
  items = filterKeyGR != 0 ? items.filter((row) => row.group['id'] === filterKeyGR) : items
  items = FilterKeyST != undefined ? items.filter((row) => row.status === FilterKeyST) : items
  return items;
});

</script>

<template>
  <div class="table-responsive table-fixed fs-6">
    <table class="table table-striped table-success align-middle">
      <thead>
        <tr>
          <th scope="col">Номер</th>
          <th scope="col">Наименование</th>
          <th scope="col">Дата<br> заключения</th>
          <th scope="col">Способ закупки</th>
          <th scope="col">Услуга</th>
          <th scope="col">Поставшик</th>
          <th scope="col">Цена</th>
          <th scope="col">Статус</th>
        </tr>
      </thead>
      <tbody>
        <tr v-for="item in filteredData" :key="item.id">
          <td class="w-auto text-break">
            <router-link :to="`/board/${item.id}`">
             {{ item.number }}
            </router-link>
          </td>
          <td class="w-auto text-break">{{ item.title }}</td>
          <td>{{ item.date }}</td>
          <td class="w-auto text-break">{{ item.category['name'] }}</td>
          <td class="w-auto">{{ item.group['name'] }}</td>
          <td class="w-auto text-break">{{ item.supplier['name'] }}</td>
          <td class="w-auto">{{ item.price }}</td>
          <td class="w-auto text-break text-center" v-if="item.status">
            <i class="bi bi-check-circle-fill text-success"></i>
          </td>
          <td class="w-auto text-center" v-else>
            <i class="bi bi-x-circle-fill text-danger"></i>
          </td>
        </tr>
      </tbody>
    </table>
  </div>
</template>

<style lang="scss" scoped></style>
