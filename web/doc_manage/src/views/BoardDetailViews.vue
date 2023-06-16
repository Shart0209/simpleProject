<script setup>

import { computed, ref, reactive } from 'vue';
import { useRoute } from 'vue-router';
import { storeToRefs } from 'pinia'
import { useDocsStore } from '@/stores'
import { useAuthStore } from '@/stores';

import Card from '@/components/Card.vue';
import UpdForm from '@/components/UpdForm.vue';

const authStore = useAuthStore();
const { item, error } = storeToRefs(useDocsStore());
const { getByIDDoc, deleteByID, downloadFile, formatDate, update } = useDocsStore();

const route = useRoute();
const id = computed(() => route.params.id);

const hidden = ref(false);
const updItem = reactive({
    data: {}
});

getByIDDoc(id.value);

async function deleteItem() {
    deleteByID(id.value);
}

async function downloadItem(fileID, fileName) {
    downloadFile(id.value, fileID, fileName);
}

async function updateForm() {

    hidden.value = !hidden.value;

    updItem.data = {};

    // Copy obj
    Object.assign(updItem.data, item.value.data);

    updItem.data.files = [];
    delete updItem.data.created_at;
    delete updItem.data.update_at;
}

async function updateItem() {

    let newItem = {};
    for (let key in updItem.data) {

        if (updItem.data[key] != item.value.data[key]) {
            if (key === 'files' && updItem.data.files.length === 0) {
                continue;
            } else if (key === 'category' || key === 'supplier' || key === 'group') {
                newItem[key] = updItem.data[key].id;
            } else {
                newItem[key] = updItem.data[key];
            }
        }
    }

    update(newItem, id.value);
}

</script>

<template>
    <div class="col-md-8">
        <div v-if="!item.error">
            <div v-show="!hidden">
                <div v-show="authStore.user" class="col-sm-6 mb-3">
                    <button type="button" class="btn btn-primary" @click="updateForm">Редактировать</button>
                    <button type="button" class="btn btn-primary mx-2" @click="deleteItem">Удалить</button>
                </div>
                <Card>
                    <template #header>
                        Государственный контракт № {{ item.data.number }} от {{ formatDate(item.data.date, 'ru') }} г.
                    </template>
                    <template #body>
                        <ul ul class="list-unstyled">
                            <li class="mb-2">
                                <div class="fw-bold d-inline">Наименование: </div>
                                <div class="fst-italic d-inline">{{ item.data.title }}</div>
                            </li>
                            <li class="mb-2">
                                <div class="fw-bold d-inline">Цена контракта: </div>
                                <div class="fst-italic d-inline">{{ item.data.price }} руб.</div>
                            </li>
                            <li class="mb-2">
                                <div class="fw-bold d-inline">Срок действия </div>
                                <div class="fst-italic d-inline">с: {{ formatDate(item.data.start_date, 'ru') }} г. по: {{
                                    formatDate(item.data.end_date, 'ru') }}
                                    г.</div>
                            </li>
                            <li class="mb-2">
                                <div class="fw-bold d-inline">Поставщик услуг: </div>
                                <div class="fst-italic d-inline">{{ item.data.supplier?.name }}</div>
                            </li>
                            <li class="mb-2">
                                <div class="fw-bold d-inline">Способ заключения: </div>
                                <div class="fst-italic d-inline">{{ item.data.category?.name }}</div>
                            </li>
                            <li class="mb-2">
                                <div class="fw-bold d-inline">Группа услуг: </div>
                                <div class="fst-italic d-inline">{{ item.data.group?.name }}</div>
                            </li>
                            <li class="mb-2">
                                <div class="fw-bold d-inline">Статус контракта: </div>
                                <div class="fst-italic d-inline" v-if="item.data.status">Действует</div>
                                <div class="fst-italic d-inline" v-else>Не действует</div>
                            </li>
                            <li class="mb-2">
                                <div class="fw-bold d-inline">Примечание: </div>
                                <div class="fst-italic d-inline">"{{ item.data.description }}"</div>
                            </li>
                        </ul>

                        <ul class="list-unstyled mb-0" v-if="item.data.files">
                            <li>
                                <div class="fw-bold d-inline">Прикрепленные файлы: </div>
                                <sup><small class="badge rounded-pill bg-info me-1">{{ item.data.files.attr.length
                                }}</small></sup>
                            </li>
                            <li v-for="fl in item.data.files.attr" :key="item.data.name">
                                <ul>
                                    <li class="d-flex justify-content-start">
                                        {{ fl.name }}
                                        {{ (fl.size / 1024).toFixed(1) }} КБ
                                        <a @click="downloadItem(fl.id, fl.name)">
                                            <i class="bi bi-file-earmark-arrow-down-fill"
                                                style="color: cornflowerblue;"></i>
                                        </a>
                                    </li>
                                </ul>
                            </li>
                        </ul>
                    </template>
                    <template #footer>
                        <div class="d-inline">Дата создания: {{ formatDate(item.data.created_at, 'ru') }} г.</div>
                        <div class="mx-2 d-inline" v-if="!item.data.hidden_at === '01.01.1'">обновлено: {{
                            formatDate(item.data.hidden_at, 'ru') }}</div>
                    </template>
                </Card>
                <p v-if="error">{{ error }}</p>
            </div>
        </div>
        <div v-else>
            <p>Ооопс данные не найдены.</p>
        </div>
    </div>
    <UpdForm v-show="hidden" @close="hidden = false" @ok="updateItem" :updItem="updItem.data" />
</template>