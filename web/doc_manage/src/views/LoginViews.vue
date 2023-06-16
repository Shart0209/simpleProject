<script setup>
import { ref } from 'vue'
import { storeToRefs } from 'pinia'
import Card from '@/components/Card.vue';
import { useAuthStore } from '@/stores';

const username = ref('')
const password = ref('')

const { error } = storeToRefs(useAuthStore());
const { login } = useAuthStore();

async function onSubmitLogin() {
    await login(username, password)
}

</script>

<template>
    <div class="col-sm-4">
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
                <button type="button" class="btn btn-primary mx-2" @click="onSubmitLogin">Войти</button>
            </template>
        </Card>
        <div>
            <p v-if="error">{{ error }}</p>
        </div>
    </div>
</template>