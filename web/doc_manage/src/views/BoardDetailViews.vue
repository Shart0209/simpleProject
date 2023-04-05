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
import { formatDate } from '@/utils/common'


const { items, error, optionsSelect } = storeToRefs(useDocsStore())
const { useFetchDocs, getOptionSelect } = useDocsStore()

let selectedFiles = reactive([])

const state = reactive({
    modal_demo: null
})

const route = useRoute();
const router = useRouter()
const id = computed(() => route.params.id);

useFetchDocs(URL, id.value);

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

// TODO Доделать резет формы, а именно поля selectedFiles
function resetForm() {
    // selectedFiles.value = null
    // Object.assign(attrs, initialAttrs);
}

function showModal() {
    resetForm()
    state.modal_demo.show()
}

function closeModal() {
    state.modal_demo.hide()
    resetForm()
}
onMounted(() => {
    state.modal_demo = new Modal('#modal_demo', {
        keyboard: false
    })

})
</script>

<template>
    <section v-if="!error">
        <div class="col-sm-6">
            <button type="button" class="btn btn-primary" @click="showModal">Редактировать</button>
            <button type="button" class="btn btn-primary mx-2" @click="deleteItem">Удалить</button>
        </div>
        <Card>
            <template #header>
                Государственный контракт № {{ items.data.number }} от {{ formatDate(items.data.date) }} г.
            </template>
            <template #text>
                <ul ul class="list-unstyled">
                    <li class="mb-2">
                        <div class="fw-bold d-inline">Наименование: </div>
                        <div class="fst-italic d-inline">{{ items.data.title }}</div>
                    </li>
                    <li class="mb-2">
                        <div class="fw-bold d-inline">Цена контракта: </div>
                        <div class="fst-italic d-inline">{{ items.data.price }} руб.</div>
                    </li>
                    <li class="mb-2">
                        <div class="fw-bold d-inline">Срок действия </div>
                        <div class="fst-italic d-inline">с: {{ formatDate(items.data.start_date) }} г. по: {{
                            formatDate(items.data.end_date) }} г.</div>
                    </li>
                    <li class="mb-2">
                        <div class="fw-bold d-inline">Поставщик услуг: </div>
                        <div class="fst-italic d-inline">{{ items.data.distributor }}</div>
                    </li>
                    <li class="mb-2">
                        <div class="fw-bold d-inline">Способ размещения: </div>
                        <div class="fst-italic d-inline">{{ items.data.category }}</div>
                    </li>
                    <li class="mb-2">
                        <div class="fw-bold d-inline">Статус контракта: </div>
                        <div class="fst-italic d-inline" v-if="items.data.status">Действует</div>
                        <div class="fst-italic d-inline" v-else>Не действует</div>
                    </li>
                    <li class="mb-2">
                        <div class="fw-bold d-inline">Примечание: </div>
                        <div class="fst-italic d-inline">{{ items.data.descriptions }}</div>
                    </li>
                </ul>

                <ul class="list-unstyled" v-if="items.data.files">
                    <li>
                        <div class="fw-bold d-inline">Прикрепленные файлы: </div>
                        <sup><small class="badge rounded-pill bg-info me-1">{{ items.data.files.attr.length
                        }}</small></sup>
                    </li>
                    <li v-for="item in items.data.files.attr" :key="item.name">
                        <div>файл:
                            <i class="bi bi-file-earmark-arrow-down-fill" style="color: cornflowerblue;"></i>
                        </div>
                    </li>
                </ul>
            </template>
            <template #footer>
                <div class="d-inline">дата создания: {{ formatDate(items.data.created_at) }}</div>
                <div class="mx-2 d-inline" v-if="!formatDate(items.data.update_at) === '01.01.1'">обновлено: {{
                    formatDate(items.data.update_at) }}</div>
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
                <!-- Component Create Form -->
                <UpdForm :attrs="items.data" :selectedFiles="selectedFiles" :optionsSelect="optionsSelect" />
            </template>

            <template #footer>
                <button type="button" class="btn btn-secondary" @click="closeModal">Close</button>
                <button type="button" class="btn btn-primary" @click="updDoc">Save</button>
            </template>
        </ModalForm>
    </section>
</template>