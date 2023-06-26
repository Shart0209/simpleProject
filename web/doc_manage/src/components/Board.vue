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
    <div class="col-sm-6">
        <slot name="menu"></slot>
    </div>
    <table class="table table-sm table-striped table-success align-middle">
        <thead>
            <tr>
                <th scope="col">#</th>
                <th scope="col">Номер</th>
                <th scope="col">Наименование</th>
                <th scope="col">Дата заключения</th>
                <th scope="col">Способ закупки</th>
                <th scope="col">Услуга</th>
                <th scope="col">Поставшик</th>
                <th scope="col">Цена</th>
                <th scope="col">Статус</th>
            </tr>
        </thead>
        <tbody>
            <tr v-for="item in filteredData" :key="item.id">
                <th scope="row">#</th>
                <td>
                    <router-link :to="`/board/${item.id}`">
                        {{ item.number }}
                    </router-link>
                </td>
                <td>{{ item.title }}</td>
                <td>{{ item.date }}</td>
                <td>{{ item.category['name'] }}</td>
                <td>{{ item.group['name'] }}</td>
                <td>{{ item.supplier['name'] }}</td>
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
</template>

<style lang="scss" scoped></style>
