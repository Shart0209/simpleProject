<script setup>
defineProps({
  data: {type: Array, required: true},
  error: null
})
</script>

<template>
  <div v-if="error">Oops! Error encountered: {{ error.message }}</div>
  <table v-else-if="data" class="table table-hover table-striped table-sm">
    <thead>
    <tr>
      <th scope="col">#</th>
      <th scope="col">Наименование</th>
      <th scope="col">Номер ГК</th>
      <th scope="col">Дата</th>
      <th scope="col">Категория</th>
      <th scope="col">Цена</th>
      <th scope="col">Статус</th>
    </tr>
    </thead>
    <tbody>
    <tr v-for="item in data" :key="item.id">
      <th scope="row">{{ item.id }}</th>
      <td>{{ item.title }}</td>
      <td>{{ item.number }}</td>
      <td>{{ item.date }}</td>
      <td>{{ item.category }}</td>
      <td>{{ item.price }}</td>
      <td v-if="item.status">
        <i class="bi bi-check-circle-fill text-success"></i>
      </td>
      <td v-else>
        <i class="bi bi-x-circle-fill text-danger"></i>
      </td>
    </tr>
    </tbody>
  </table>
  <div v-else>Loading...</div>
</template>

<style lang="scss" scoped>

</style>
