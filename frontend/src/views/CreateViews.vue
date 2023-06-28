<script setup>
import { reactive } from 'vue';
import { storeToRefs } from 'pinia';
import { useRouter } from 'vue-router';

import { useDocsStore } from '@/stores';
import { useAuthStore } from '@/stores';

const { optionsSelect, error } = storeToRefs(useDocsStore());
const { create, getInitials, getOptionSelect } = useDocsStore();

const authStore = useAuthStore();

const init = getInitials()
const attrs = reactive({ ...init });
let selectedFiles = reactive([]);

const router = useRouter();

function onChangeCategory(e) {
    attrs.category = e.target.value;
}

function onChangeSupplier(e) {
    attrs.supplier = e.target.value;
}

function onChangeGroup(e) {
    attrs.group = e.target.value;
}

function onChangeFile(e) {

    if (e.target.files) {
        for (let i = 0; i < e.target.files.length; i++) {
            selectedFiles.push(e.target.files[i])
        }
        // TODO добавляется какое то поле value. Решено удалить.
        delete selectedFiles['value']
    }
}

function createForm() {

    try {
        if (attrs.date > attrs.start_date || attrs.start_date > attrs.end_date) {
            throw new Error('Не верные даты');
        }

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

        formData.set('author', authStore.user.id);

        create(formData);
    } catch (err) {
        error.value = err;
        console.log(err);
    }
}

function closeForm() {
    router.replace({ name: 'board' });
}

getOptionSelect();

</script>

<template>
    <div v-if="authStore.user.isActive" class="col-md-6">
        <h4 class="mb-2">Create</h4>
        <form method="POST" class="was-validated" enctype="multipart/form-data" @submit.prevent="createForm">
            <div class="row g-3">
                <div class="col-12">
                    <input type="text" class="form-control" maxlength="160" v-model.trim="attrs.title"
                        placeholder="Наименование" required>
                    <div class="invalid-feedback">
                        Valid title is required.
                    </div>
                </div>
                <div class="col-sm-6">
                    <input type="text" class="form-control" maxlength="50" v-model.trim="attrs.number" placeholder="Номер"
                        required>
                    <div class="invalid-feedback">
                        Valid number is required.
                    </div>
                </div>
                <div class="col-sm-6">
                    <input type="number" class="form-control" v-model.trim="attrs.price" placeholder="Цена контракта"
                        required>
                    <div class="invalid-feedback">
                        Valid price is required.
                    </div>
                </div>

                <div class="col-sm-4">
                    <label>Дата заключения</label>
                    <input type="date" class="form-control" v-model="attrs.date" required>
                    <div class="invalid-feedback">
                        Valid date is required.
                    </div>
                </div>
                <div class="col-sm-4">
                    <label>Дата начала действия</label>
                    <input type="date" class="form-control" v-model="attrs.start_date" required>
                    <div class="invalid-feedback">
                        Valid date is required.
                    </div>
                </div>
                <div class="col-sm-4">
                    <label>Дата окончания действия</label>
                    <input type="date" class="form-control" v-model="attrs.end_date" required>
                    <div class="invalid-feedback">
                        Valid date is required.
                    </div>
                </div>
                <div class="col-sm-6">
                    <select class="form-floating form-select" @change="onChangeCategory($event)" required>
                        <option disabled selected value="">Способ заключения</option>
                        <option v-for="item in optionsSelect.categories" :key=item.id :value="item.id">
                            {{ item.name }}
                        </option>
                    </select>
                </div>
                <div class="col-sm-6">
                    <select class="form-floating form-select" @change="onChangeGroup($event)" required>
                        <option disabled selected value="">Группа</option>
                        <option v-for="item in optionsSelect.groups" :key=item.id :value="item.id">{{ item.name }}
                        </option>
                    </select>
                </div>
                <div class="col-sm-12">
                    <select class="form-floating form-select" @change="onChangeSupplier($event)" required>
                        <option disabled selected value="">Поставщик</option>
                        <option v-for="item in optionsSelect.suppliers" :key=item.id :value="item.id">{{ item.name }}
                        </option>
                    </select>
                </div>
                <div class="col-sm-12">
                    <textarea class="form-control" maxlength="150" v-model.trim="attrs.description" row="3"
                        placeholder="Примечание"></textarea>
                </div>
                <div class="col-sm-12 mb-2">
                    <input type="file" @change="onChangeFile($event)" ref="file" class="form-control" multiple>
                </div>
            </div>
            <div class="d-grid gap-2 d-md-flex justify-content-md-end">
                <button type="submit" class="btn btn-primary">OK</button>
                <button type="button" class="btn btn-primary" @click="closeForm">Close</button>
            </div>
        </form>
        <div class=" mt-2 alert alert-primary" role="alert" v-show="error">
            <div>
                {{ error }}
            </div>
        </div>
    </div>
    <div v-else class=" col-md-6 mt-2 alert alert-primary" role="alert">
        <div>
            <span>Error: Пользователь не авторизирован</span>
        </div>
    </div>
</template>

<style lang="scss" scoped></style>

