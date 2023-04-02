<script setup>
import { computed, ref, reactive } from "vue";
import { useRoute, useRouter } from "vue-router";

import Card from '@/components/Card.vue'

import { URL } from '@/utils/constants';
import { formatDate, useFetch } from '@/utils/common'

const route = useRoute();
const router = useRouter()
const id = computed(() => route.params.id);

const items = reactive({})
getDoc()

function getDoc() {
    useFetch(URL, id.value)
        .then((res) => {
            let data = {
                value: res.data
            }
            items['value'] = data.value
        });
}

function deleteItem() {
    useFetch(`${URL}delete/`, id.value, { method: 'DELETE' })
        .then((res) => {
            items['value'] = res;
        });

    router.push({ name: 'board', replace: true })
}

function editItem() {

}

</script>

<template>
    <section v-if="items.value">
        <div class="col-sm-6">
            <button type="button" class="btn btn-primary" @click="editItem">Редактировать</button>
            <button type="button" class="btn btn-primary mx-2" @click="deleteItem">Удалить</button>
        </div>
        <Card>
            <template #header>
                Государственный контракт № {{ items.value.number }} от {{ formatDate(items.date) }} г.
            </template>
            <template #text>
                <ul ul class="list-unstyled">
                    <li class="mb-2">
                        <div class="fw-bold d-inline">Наименование: </div>
                        <div class="fst-italic d-inline">{{ items.value.title }}</div>
                    </li>
                    <li class="mb-2">
                        <div class="fw-bold d-inline">Цена контракта: </div>
                        <div class="fst-italic d-inline">{{ items.value.price }} руб.</div>
                    </li>
                    <li class="mb-2">
                        <div class="fw-bold d-inline">Срок действия </div>
                        <div class="fst-italic d-inline">с: {{ formatDate(items.value.start_date) }} г. по: {{
                            formatDate(items.value.end_date) }} г.</div>
                    </li>
                    <li class="mb-2">
                        <div class="fw-bold d-inline">Поставщик услуг: </div>
                        <div class="fst-italic d-inline">{{ items.value.distributor }}</div>
                    </li>
                    <li class="mb-2">
                        <div class="fw-bold d-inline">Способ размещения: </div>
                        <div class="fst-italic d-inline">{{ items.value.category }}</div>
                    </li>
                    <li class="mb-2">
                        <div class="fw-bold d-inline">Статус контракта: </div>
                        <div class="fst-italic d-inline" v-if="items.value.status">Действует</div>
                        <div class="fst-italic d-inline" v-else>Не действует</div>
                    </li>
                    <li class="mb-2">
                        <div class="fw-bold d-inline">Примечание: </div>
                        <div class="fst-italic d-inline">{{ items.value.descriptions }}</div>
                    </li>
                </ul>

                <ul class="list-unstyled" v-if="items.value.files">
                    <li>
                        <div class="fw-bold d-inline">Прикрепленные файлы: </div>
                        <sup><small class="badge rounded-pill bg-info me-1">{{ items.value.files.attr.length
                        }}</small></sup>
                    </li>
                    <li v-for="item in items.value.files.attr" :key="item.name">
                        <div>файл:
                            <i class="bi bi-file-earmark-arrow-down-fill" style="color: cornflowerblue;"></i>
                        </div>
                    </li>
                </ul>
            </template>
            <template #footer>
                <div class="d-inline">дата создания: {{ formatDate(items.value.created_at) }}</div>
                <div class="mx-2 d-inline" v-if="!formatDate(items.value.update_at) === '01.01.1'">обновлено: {{
                    formatDate(items.value.update_at) }}</div>
            </template>
        </Card>
    </section>
</template>