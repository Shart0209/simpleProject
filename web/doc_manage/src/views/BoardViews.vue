<script setup>
import { reactive, onMounted } from "vue";
import { storeToRefs } from 'pinia'

import { Modal } from "bootstrap";

import Board from "@/components/Board.vue";
import ModalForm from "@/components/Modal.vue";
import CreateForm from '@/components/CreateForm.vue';

import { initialAttrs, URL } from '@/utils/constants';
import { useDocsStore } from '@/stores/docs'


const { items, error, optionsSelect } = storeToRefs(useDocsStore())
const { useFetchDocs, getOptionSelect } = useDocsStore()

const state = reactive({
    modal_demo: null
})

const attrs = reactive({ ...initialAttrs });
let selectedFiles = reactive([])

async function createDoc() {

    let formData = new FormData();

    for (let key in attrs) {
        formData.append(key, attrs[key]);
    }

    if (selectedFiles) {
        delete selectedFiles['value']
        for (let key in selectedFiles) {
            formData.append("files", selectedFiles[key], selectedFiles[key].name);
        }
    }

    try {
        await fetch(`${URL}add`, {
            method: 'POST',
            body: formData
        });
    } catch (error) {
        console.error('Ошибка:', error);
    }


    state.modal_demo.hide()
    resetForm()
    useFetchDocs(URL)
}

// TODO Доделать резет формы, а именно поля selectedFiles
function resetForm() {
    selectedFiles.value = null
    Object.assign(attrs, initialAttrs);
}

function showModal() {
    resetForm()
    state.modal_demo.show()
}

function closeModal() {
    state.modal_demo.hide()
    resetForm()
    useFetchDocs(URL)
}

onMounted(() => {
    state.modal_demo = new Modal('#modal_demo', {
        keyboard: false
    })
    useFetchDocs(URL)
    getOptionSelect(`${URL}sps`)
})

</script>

<template>
    <section>
        <!-- Component Board -->
        <Board :items="items.data">
            <template #menu>
                <button type="button" class="btn btn-primary" @click="showModal">Create</button>
            </template>
        </Board>
    </section>
    <p v-if="error">Ооопля, пусто однако!</p>

    <section>
        <!-- Component Modal -->
        <ModalForm id="modal_demo" :close_func="closeModal">
            <template #title>
                Создание
            </template>

            <template #body>
                <!-- Component Create Form -->
                <CreateForm :attrs="attrs" :selectedFiles="selectedFiles" :optionsSelect="optionsSelect" />
            </template>

            <template #footer>
                <button type="button" class="btn btn-secondary" @click="closeModal">Close</button>
                <button type="button" class="btn btn-primary" @click="createDoc">Save</button>
            </template>
        </ModalForm>
    </section>
</template>


