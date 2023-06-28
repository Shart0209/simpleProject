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
}

function isValidForm() {
    error.value = null;
    try {

        if (updItem.data.date > updItem.data.start_date || updItem.data.start_date > updItem.data.end_date) {
            throw new Error('Не верные даты');
        }

        for (let key in updItem.data) {
            if (key === 'files' && updItem.data[key] === undefined) {
                continue;
            } else if (updItem.data[key] != item.value.data[key]) {
                return;
            } else {
                continue;
            }
        }

        throw new Error('Форма не изменена');

    } catch (err) {
        console.log(err);
        error.value = err;
    }
}

async function updateForm() {

    isValidForm()
    if (error.value) {
        return;
    }

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
                <div v-show="authStore.user.isActive" class="col-sm-6 mb-3">
                    <div class="btn-group" role="group">
                        <button type="button" class="btn btn-primary me-1" @click="updateItem">
                            <i class="bi bi-pencil-fill"></i>
                        </button>
                        <button type="button" class="btn btn-primary me-1" @click="refreshForm">
                            <i class="bi bi-arrow-clockwise"></i>
                        </button>
                        <button type="button" class="btn btn-primary" @click="deleteItem">
                            <i class="bi bi-trash3"></i>
                        </button>
                    </div>
                </div>
                <div v-show="!authStore.user.isActive" class="mb-3">
                    <button type="button" class="btn btn-primary me-1" @click="refreshForm">
                        <i class="bi bi-arrow-clockwise"></i>
                    </button>
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
                            <li class="mb-2" v-show="item.data.description">
                                <div class="fw-bold d-inline">Примечание: </div>
                                <div class="fst-italic d-inline">{{ item.data.description }}</div>
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
                            <div class="col col-sm-8">
                                <div class="d-inline">Дата создания: {{ formatDate(item.data.created_at, 'ru') }} г. </div>
                                <div class="d-inline">Автор: {{ item.data.author }}</div>
                            </div>
                            <div class="ms-auto" v-if="item.data.updated_at > item.data.created_at">Обновлено:
                                {{ formatDate(item.data.updated_at, 'ru') }} г.
                            </div>
                        </div>
                    </template>
                </Card>
                <div class=" mt-2 alert alert-primary" role="alert" v-show="error">
                    <div>
                        {{ error }}
                    </div>
                </div>
            </div>
        </div>
        <div class=" mt-2 alert alert-primary" role="alert" v-else>
            <div>
                Ооопс данные не найдены.
            </div>
        </div>
    </div>
    <UpdForm v-show="hidden" @close="hidden = false" @ok="updateForm" :updItem="updItem.data" />
</template> 

<style lang="scss" scoped></style>