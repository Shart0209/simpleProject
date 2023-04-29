<script setup>
import { computed, reactive, onMounted } from "vue";
import { useRoute, useRouter } from "vue-router";
import { storeToRefs } from 'pinia'
import { useDocsStore } from '@/stores/docs'

import { Modal } from "bootstrap";

import ModalForm from "@/components/Modal.vue";
import UpdForm from '@/components/UpdForm.vue';
import Card from '@/components/Card.vue'

import { URL } from '@/utils/constants';
import { formatDateRU } from '@/utils/common'

const { items, error, optionsSelect } = storeToRefs(useDocsStore())
const { useFetchDocs, getOptionSelect } = useDocsStore()

const route = useRoute();
const router = useRouter();
const id = computed(() => route.params.id);

let selectedFiles = reactive([])

const state = reactive({
    modal_demo: null
})

async function deleteItem() {
    try {
        await fetch(`${URL}delete/${id.value}`, { method: 'DELETE' })
    } catch (error) {
        console.log(error);
    }

    router.push({ name: 'board', replace: true })
}

function updDoc() {

}

async function downloadFile(name) {
    try {
        let response = await fetch(`${URL}download/${name}`)
        let result = await response.blob()
        if (!response.ok) {
            const message = `An error has occured: ${result.errors} error: ${response.status}`;
            throw new Error(message);
        }

        let a = document.createElement("a");
        a.href = window.URL.createObjectURL(result);
        a.download = 'file' + '.' + name.split('.')[1]
        a.click();

    } catch (err) {
        console.log(err);
    }
}

function showModal() {
    state.modal_demo.show()
}

function closeModal() {
    state.modal_demo.hide()
}

onMounted(() => {
    state.modal_demo = new Modal('#modal_demo', {
        keyboard: false
    })
    useFetchDocs(URL, id.value);
})

</script>

<template>
    <section v-if="!error">
        <div class="col-sm-6">
            <button type="button" class="btn btn-primary" @click="showModal">Редактировать</button>
            <button type="button" class="btn btn-primary mx-2" @click="deleteItem">Удалить</button>
        </div>
        <Card v-for="data in items.data" :key="data.id">
            <template #header>
                Государственный контракт № {{ data.number }} от {{ formatDateRU(data.date) }} г.
            </template>
            <template #text>
                <ul ul class="list-unstyled" >
                    <li class="mb-2">
                        <div class="fw-bold d-inline">Наименование: </div>
                        <div class="fst-italic d-inline">{{ data.title }}</div>
                    </li>
                    <li class="mb-2">
                        <div class="fw-bold d-inline">Цена контракта: </div>
                        <div class="fst-italic d-inline">{{ data.price }} руб.</div>
                    </li>
                    <li class="mb-2">
                        <div class="fw-bold d-inline">Срок действия </div>
                        <div class="fst-italic d-inline">с: {{ formatDateRU(data.start_date) }} г. по: {{
                            formatDateRU(data.end_date) }} г.</div>
                    </li>
                    <li class="mb-2">
                        <div class="fw-bold d-inline">Поставщик услуг: </div>
                        <div class="fst-italic d-inline">{{ data.distributor }}</div>
                    </li>
                    <li class="mb-2">
                        <div class="fw-bold d-inline">Способ размещения: </div>
                        <div class="fst-italic d-inline">{{ data.category }}</div>
                    </li>
                    <li class="mb-2">
                        <div class="fw-bold d-inline">Статус контракта: </div>
                        <div class="fst-italic d-inline" v-if="data.status">Действует</div>
                        <div class="fst-italic d-inline" v-else>Не действует</div>
                    </li>
                    <li class="mb-2">
                        <div class="fw-bold d-inline">Примечание: </div>
                        <div class="fst-italic d-inline">{{ data.descriptions }}</div>
                    </li>
                </ul>

                <ul class="list-unstyled" v-if="data.files">
                    <li>
                        <div class="fw-bold d-inline">Прикрепленные файлы: </div>
                        <sup><small class="badge rounded-pill bg-info me-1">{{ data.files.attr.length
                        }}</small></sup>
                    </li>
                    <li v-for="item in data.files.attr" :key="item.name">
                        <div>
                            {{ item.name }}
                            <a @click="downloadFile(item.name)">
                                <i class="bi bi-file-earmark-arrow-down-fill" style="color: cornflowerblue;"></i>
                            </a>
                        </div>
                    </li>
                </ul>
            </template>
            <template #footer>
                <div class="d-inline">дата создания: {{ formatDateRU(data.created_at) }}</div>
                <div class="mx-2 d-inline" v-if="!formatDateRU(data.update_at) === '01.01.1'">обновлено: {{
                    formatDateRU(data.update_at) }}</div>
            </template>
        </Card>
    </section>
    <p v-else>{{ error.message }}</p>

    <section>
        <!-- Component Modal -->
        <ModalForm id="modal_demo" :close_func="closeModal">
            <template #title>
                Создание
            </template>

            <template #body>
                <!-- Component Update Form -->
                <UpdForm :attrs="items.data" :selectedFiles="selectedFiles" :optionsSelect="optionsSelect" />
            </template>

            <template #footer>
                <button type="button" class="btn btn-secondary" @click="closeModal">Close</button>
                <button type="button" class="btn btn-primary" @click="updDoc">Save</button>
            </template>
        </ModalForm>
    </section>
</template>