<script setup>

import { formatDate } from '@/utils/common'

let counter_id = 1

const props = defineProps({
    items: {
        required: true
    }
})
</script>

<template>
    <div class="col-sm-6">
        <slot name="menu"></slot>
    </div>
    <table v-if="items" class="table table-hover table-striped table-sm">
        <thead>
            <tr>
                <th scope="col">#</th>
                <th scope="col">Номер</th>
                <th scope="col">Наименование</th>
                <th scope="col">Дата заключения</th>
                <th scope="col">Категория</th>
                <th scope="col">Цена</th>
                <th scope="col">Статус</th>
            </tr>
        </thead>
        <tbody>
            <tr v-for="item in items" :key="item.id">
                <th scope="row">{{ counter_id++ }}</th>
                <td>
                    <router-link :to="{ name: 'board_detail_id', params: { id: item.id } }">
                        {{ item.number }}
                    </router-link>
                </td>
                <td>{{ item.title }}</td>
                <td>{{ formatDate(item.date) }}</td>
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
    <div v-else>Загрузка...</div>
</template>

<style lang="scss" scoped></style>
