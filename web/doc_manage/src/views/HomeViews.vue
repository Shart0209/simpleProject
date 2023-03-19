<script setup>
import {ref, reactive} from "vue";

import Table from "@/components/DocList.vue";
import CreateForm from "@/components/CreateForm.vue";
import Modal from "@/components/Modal.vue";


// import { useFetch } from '@/utils/fetch.js'
// import { requestOptions } from '@/utils/const.js'
// const { data, error } = useFetch('http://localhost:8080/documents/', requestOptions)

const currentDate = new Date().toJSON().slice(0, 10)

let docsList = ref ([
  {
    id: 1,
    title: 'Оказание слуг по связи - телефония',
    number: 'ЧИТ01-12-243/24',
    date: currentDate,
    start_date: currentDate,
    end_date: currentDate,
    distributors: 'ПАО МТС',
    description: '',
    category: 'Аукцион',
    price: 15000,
    status: true
  },
  {
    id: 2,
    title: 'Оказание слуг по связи - интернет',
    number: 'ЧИТ01-12-232/24',
    date: currentDate,
    start_date: currentDate,
    end_date: currentDate,
    distributors: 'ПАО МТС',
    description: '',
    category: 'Единственный поставщик',
    price: 150000,
    status: false
  },
  {
    id: 3,
    title: 'Оказание слуг по связи - мобильная связь',
    number: 'ЧИТ01-12-454/24',
    date: currentDate,
    start_date: currentDate,
    end_date: currentDate,
    distributors: 'ПАО МТС',
    description: '',
    category: 'Аукцион',
    price: 115200,
    status: true
  },
  {
    id: 4,
    title: 'Оказание слуг по связи - телефония',
    number: 'ЧИТ01-12-683/24',
    date: currentDate,
    start_date: currentDate,
    end_date: currentDate,
    distributors: 'ПАО МТС',
    description: '',
    category: 'Единственный поставщик',
    price: 10000,
    status: true
  },
  {
    id: 5,
    title: 'Оказание слуг по связи - телефония',
    number: 'ЧИТ01-12-686/24',
    date: currentDate,
    start_date: currentDate,
    end_date: currentDate,
    distributors: 'ПАО МТС',
    description: '',
    category: 'Единственный поставщик',
    price: 256000,
    status: true
  },
])

const error = ref('')

let attr = ref({
    title: '',
    number: '',
    date: '',
    start_date: '',
    end_date: '',
    category: 'Способ заключения',
    distributor: 'Поставщик',
    description: '',
    price: 0,
    status: true
})

const docsArr = {
  categories: [
    {id: 1, name: 'Аукцион'},
    {id: 2, name: 'Единственный поставщик'}
  ],
  distributors: [
    {id: 1, name: 'ПАО МТС'},
    {id: 2, name: 'ПАО Ростелеком'}
  ]
}

let thisModal= ref(null);

function showModal(){
  thisModal.value.show()
}

function createDoc(payload) {

  console.log(payload);
  docsList.push(payload)
}

</script>

<template>
  <main class="container flex-shrink-0 mb-4">
    <div class="col-md-7 py-2 col-lg-8">
      <button type="button" class="btn btn-primary" @click="showModal">Create</button>
    </div>
    <Modal
        :attr="attr"
        :docsArr="docsArr"
        @create="createDoc"
        ref="thisModal"
    />

    <Table
        :data="docsList"
        :error="error"
    />


  </main>
</template>


