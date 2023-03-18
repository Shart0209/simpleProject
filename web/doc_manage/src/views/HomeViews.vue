<template>
    <div class="buttons p-2 d-flex justify-content-between">
    <button type="button" class="btn btn-primary" @click="">Create</button>
    <button type="button" class="btn btn-primary" @click="">Delete</button>
    </div>
    
    <div class="box p-3 rounded-3 d-flex justify-content-between">       
        <div class="container-fluid py-1">
            <div v-if="error">Oops! Error encountered: {{ error.message }}</div>
            <table v-else-if="data" class="table table-hover table-striped table-sm">
                <thead>
                    <tr>
                        <th scope="col">#</th>
                        <th scope="col">Номер ГК</th>
                        <th scope="col">Дата</th>
                        <th scope="col">Категория</th>
                        <th scope="col">Цена</th>
                    </tr>
                </thead>
                <tbody>
                    <tr v-for="item in data.data" :key="item.id">
                        <th scope="row">{{ item.id }}</th>
                        <td>{{ item.number }}</td>
                        <td>{{ item.date }}</td>
                        <td>{{ item.category }}</td>
                        <td>{{ item.price }}</td>
                    </tr>
                </tbody>
            </table>
            <div v-else>Loading...</div>
        </div>
    </div>  
</template>


<script setup>

import { useFetch } from '@/utils/fetch.js'
import { requestOptions } from '@/utils/const.js'


const { data, error } = useFetch('http://localhost:8080/documents/', requestOptions)

</script>

