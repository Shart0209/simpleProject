<script setup>
import { ref } from 'vue'

import { useFetch } from '@/utils/const.js'
import { requestOptions } from '@/utils/const.js'


const title = ref('')

const distributor = ref('Поставщик')
const distributors = ref([
    { id: 1, name: 'ПАО МТС' },
    { id: 2, name: 'ПАО Ростелеком' }
])

const category = ref('Способ заключения')
const categories = ref([
    { id: 1, name: 'Аукцион' },
    { id: 2, name: 'Единственный поставщик' }])


function onSubmit() {

    requestOptions.method = 'POST'
    requestOptions.body = JSON.stringify({ title: title.value })

    const { data, error } = useFetch('http://localhost:8080/documents/add', requestOptions)
    console.log(error)
}

</script>

<template>
    <div class="box p-3 rounded-3 d-flex justify-content-between">
        <div class="container px-5">
            <form method="POST" enctype="multipart/form-data" @submit.prevent="onSubmit">
                <div class="form-floating mb-3">
                    <input type="text" required class="form-control" id="title" v-model="title"
                        placeholder="Наименование ГК">
                    <label for="title">Наименование</label>
                </div>
                <div class="row">
                    <div class="col">
                        <div class="form-floating mb-3">
                            <input type="text" class="form-control" id="number" v-model="number" placeholder="Номер ГК">
                            <label for="number">Номер</label>
                        </div>
                    </div>
                    <div class="col">
                        <div class="form-floating mb-3">
                            <input type="date" class="form-control" id="date" v-model="dates"
                                placeholder="Дата заключения ГК">
                            <label for="date">Дата заключения</label>
                        </div>
                    </div>
                </div>
                <div class="row">
                    <div class="col">
                        <div class="form-floating mb-3">
                            <input type="number" class="form-control" id="price" v-model="price"
                                placeholder="Цена контракта ГК">
                            <label for="price">Цена контракта</label>
                        </div>
                    </div>
                    <div class="col">
                        <div class="form-floating mb-3">
                            <input type="date" class="form-control" id="start_date" v-model="start_date"
                                placeholder="Дата начала действия ГК">
                            <label for="start_date">Дата начала действия</label>
                        </div>
                    </div>
                    <div class="col">
                        <div class="form-floating mb-3">
                            <input type="date" class="form-control" id="end_date" v-model="end_date"
                                placeholder="Дата окончания действия ГК">
                            <label for="end_date">Дата окончания действия</label>
                        </div>
                    </div>
                </div>
                <div class="row">
                    <div class="col">
                        <select class="form-floating required  mb-3 form-select" id="category" v-model="category">
                            <option disabled selected>Способ заключения</option>
                            <option v-for="item in categories" value="{{ item.id }}" :key=item.id>{{ item.name }}</option>
                        </select>
                    </div>
                    <div class="col">
                        <select class="form-floating required mb-3 form-select" id="distributor" v-model="distributor">
                            <option disabled selected>Поставщик</option>
                            <option v-for="item in distributors" value="{{ item.id }}" :key=item.id>{{ item.name }}</option>
                        </select>
                    </div>
                </div>
                <div class="form-floating mb-3">
                    <textarea class="form-control" id="description" v-model="description" rows="2"
                        placeholder="Примечание"></textarea>
                    <label for="description">Примечание</label>
                </div>
                <div class="input-group mb-3">
                    <input type="file" class="form-control" multiple id="files" v-on:change="files" aria-describedby="files"
                        aria-label="Upload">
                </div>
                <button type="submit" class="btn btn-primary">Submit</button>
            </form>
        </div>
    </div>
</template>

