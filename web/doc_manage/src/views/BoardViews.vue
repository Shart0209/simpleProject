<script setup>
import { reactive, onMounted } from "vue";

import { Modal } from "bootstrap";

import Board from "@/components/Board.vue";
import ModalForm from "@/components/Modal.vue";
import CreateForm from '@/components/CreateForm.vue';

import { initialAttrs, URL, optionsSelect } from '@/utils/constants';
import { useFetch } from '@/utils/common';

const state = reactive({
    modal_demo: null
})

const attrs = reactive({ ...initialAttrs });
let selectedFiles = reactive([])

const items = reactive({
    res: {},
    error: {},
})


// function getDoc() {
//     useFetch(URL)
//         .then((res) => {
//             items.res = res.res;
//             items.error = res.error;
//         });
// }




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
    getDoc()

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
    getDoc()
}

onMounted(() => {
    state.modal_demo = new Modal('#modal_demo', {
        keyboard: false
    })
    getDoc()
})

</script>

<template>
    {{ items }}
    <section>
        <!-- Component Board -->
        <!-- <Board :items="items.value.res">
                <template #menu>
                    <button type="button" class="btn btn-primary" @click="showModal">Create</button>
                </template>
            </Board> -->
    </section>
    <section>
        <!-- {{ items.value.error }} -->
    </section>

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


