<script setup>
import { ref } from 'vue';
import { storeToRefs } from 'pinia';
import { useRouter } from 'vue-router';
import { useAuthStore } from '@/stores';

import Card from '@/components/Card.vue';

const username = ref('');
const password = ref('');

const { error } = storeToRefs(useAuthStore());
const { login } = useAuthStore();

const router = useRouter();

async function onSubmitLogin() {
    login(username, password);
}

async function onClose() {
    router.push({ name: 'board', replace: true });
}

</script>

<template>
    <div class="col-sm-4 pt-5">
        <Card>
            <template #header>
                Аутентификация
            </template>
            <template #body>
                <form method="POST" class="was-validated" @submit.prevent>
                    <div class="row g-3">
                        <div class="col-12 mb-2">
                            <input type="text" class="form-control" v-model="username" placeholder="Логин" required>
                            <div class="invalid-feedback">
                                Введите логин.
                            </div>
                        </div>
                        <div class="col-12 mb-2">
                            <input type="password" class="form-control" @keyup.enter="onSubmitLogin" v-model="password"
                                placeholder="Пароль" required>
                            <div class="invalid-feedback">
                                Введите пароль.
                            </div>
                        </div>
                    </div>
                </form>
            </template>
            <template #footer>
                <div class="d-grid gap-2 d-md-flex justify-content-md-end mt-2">
                    <button type="submit" class="btn btn-primary" @click="onSubmitLogin">OK</button>
                    <button type="button" class="btn btn-primary" @click="onClose">Close</button>
                </div>
            </template>
        </Card>

        <div class=" mt-2 alert alert-primary" role="alert" v-show="error">
            <div>
                {{ error }}
            </div>
        </div>
    </div>
</template>

<style lang="scss" scoped></style>