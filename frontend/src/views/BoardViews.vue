<script setup>
import { ref } from 'vue';
import { useRouter } from 'vue-router';
import { storeToRefs } from 'pinia';

import Board from '@/components/Board.vue';

import { useDocsStore } from '@/stores';
import { useAuthStore } from '@/stores';

const authStore = useAuthStore();
const { getAllDocs } = useDocsStore();
const { items, optionsSelect } = storeToRefs(useDocsStore());

const router = useRouter();

const filterSP = ref(0);
const filterGR = ref(0);
const FilterKeyStatus = ref(undefined);

async function createForm() {
    // redirect to create page
    router.push({ name: 'add', replace: true });
}

async function refreshForm() {
    getAllDocs();
}

function selectReset() {
    filterSP.value = 0;
    filterGR.value = 0;
    FilterKeyStatus.value = undefined;
}

getAllDocs();

</script>

<template>
    <section>
        <!-- Component Board -->
        <div class="col-md-10 offset-md-1">
            <div class="row mb-3 gx-1">
                <div class="col">
                    <div class="btn-group" role="group">
                        <button v-show="authStore.user.isActive" @click="createForm" type="button"
                            class="btn btn-primary me-1"><i class="bi bi-file-earmark-plus-fill"></i></button>
                        <button type="button" class="btn btn-primary" @click="refreshForm">
                            <i class="bi bi-arrow-clockwise"></i>
                        </button>
                    </div>
                </div>
                <div class="col-md-auto">
                    <button type="button" class="btn btn-primary" @click="selectReset"><i
                            class="bi bi-x-circle"></i></button>
                </div>
                <div class="col-sm-3">
                    <select class="form-floating form-select" v-model="filterSP">
                        <option selected :value=0>-- Способ заключения --</option>
                        <option v-for="item in optionsSelect.categories" :key=item.id :value="item.id">
                            {{ item.name }}
                        </option>
                    </select>
                </div>
                <div class="col-sm-2">
                    <select class="form-floating form-select" v-model="filterGR">
                        <option selected :value=0>-- Услуга --</option>
                        <option v-for="item in optionsSelect.groups" :key=item.id :value="item.id">{{ item.name }}
                        </option>
                    </select>
                </div>
                <div class="col-sm-2">
                    <select class="form-floating form-select" v-model="FilterKeyStatus">
                        <option selected :value=undefined>-- Статус --</option>
                        <option :value=true>Действителен</option>
                        <option :value=false>Не действителен</option>
                    </select>
                </div>
            </div>
            <Board :items="items.data" :filterKeySP=filterSP :filterKeyGR=filterGR :FilterKeyST=FilterKeyStatus />

            <div v-show="items.error" class="mt-2 alert alert-primary" role="alert">
                Ооопля, пусто однако!
            </div>
        </div>
    </section>
</template>


<style lang="scss" scoped></style>