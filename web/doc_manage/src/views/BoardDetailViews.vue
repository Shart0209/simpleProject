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

async function updateItem() {
    hidden.value = !hidden.value;

    updItem.data = {};

    // Copy obj
    Object.assign(updItem.data, item.value.data);

    updItem.data.files = undefined;
    delete updItem.data.created_at;
    delete updItem.data.update_at;
}

async function updateForm() {
    let formData = new FormData();

    for (let key in updItem.data) {
        if (key === 'category' || key === 'group' || key === 'supplier') {
            formData.set(key, updItem.data[key].id);
        } else if (key === 'files') { 
            for (let key in updItem.data.files) {
            formData.append("files", updItem.data.files[key], updItem.data.files[key].name);
        }
        } else {
            formData.append(key, updItem.data[key]);
        }
    }
    if (updItem.data.files === undefined) {
        formData.delete('files');
    }
    formData.delete('author');
    formData.delete('update_at');
    
    update(formData, id.value);
}

async function refreshForm() {
    getByIDDoc(id.value);
}

</script>

<template>
    <div class="col-md-8">
        <div v-if="!item.error">
            <div v-show="!hidden">
                <div v-show="authStore.user" class="col-sm-6 mb-3">
                    <div class="btn-group" role="group">
                        <button type="button" class="btn btn-primary me-1" @click="updateItem">Edit</button>
                        <button type="button" class="btn btn-primary me-1" @click="refreshForm">Refresh</button>
                        <button type="button" class="btn btn-primary" @click="deleteItem">Delete</button>
                    </div>
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
                                <sup><small class="badge rounded-pill bg-info me-1">{{ item.data.files.length
                                }}</small></sup>
                            </li>
                            <li v-for="fl in item.data.files" :key="item.data.files.name">
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
                        <div class="d-grid gap-2 d-md-flex">
                            <div class="d-inline">Дата создания: {{ formatDate(item.data.created_at, 'ru') }} г.</div>
                            <div class="d-inline" v-if="!item.data.hidden_at === '01.01.1'">обновлено: {{
                                formatDate(item.data.hidden_at, 'ru') }}
                            </div>
                            <div class="d-inline">Автор: {{ item.data.author }}</div>
                        </div>
                    </template>
                </Card>
                <p v-if="error">{{ error }}</p>
            </div>
        </div>
        <div v-else>
            <span>Ооопс данные не найдены.</span>
        </div>
    </div>
    <UpdForm v-show="hidden" @close="hidden = false" @ok="updateForm" :updItem="updItem.data" />
</template> 