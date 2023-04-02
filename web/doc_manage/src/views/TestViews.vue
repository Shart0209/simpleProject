<script setup>
import { computed, ref, reactive } from "vue";
import { useRoute, useRouter } from "vue-router";

import { URL, reqOptions } from '@/utils/constants';
import { useFetch } from '@/utils/common'


const items = reactive({})


// test 
let r = {}
getDoc()

function getDoc() {
    useFetch(URL)
        .then((res) => {
            r.nested = res
        });
}

const ttt = reactive(r)
console.log(r.nested);

let test = JSON.parse(JSON.stringify(ttt))
console.log(test);

</script>

<template>
    {{ items.value }}

</template>