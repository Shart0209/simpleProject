<script setup>
import { useAuthStore } from '@/stores';

const authStore = useAuthStore();
</script>

<template>
    <div class="container">
        <div class="col-md-10 offset-1">
            <nav class="navbar navbar-expand-lg border-bottom py-1 mb-4">
                <div class="collapse navbar-collapse" id="navbarNav">
                    <ul class="navbar-nav fs-5 me-auto">
                        <li class="nav-item">
                            <router-link class="nav-link" to="/">
                                <i class="bi bi-house-fill"></i>
                            </router-link>
                        </li>
                        <li class="nav-item text-uppercase">
                            <router-link class="nav-link" to="/board">board</router-link>
                        </li>
                    </ul>
                    <ul class="navbar-nav fs-4">

                        
                        <li v-if="authStore.user.isActive" class="nav-item ">
                            <a @click="authStore.logout()" class="nav-link" href="#">
                                <i class="bi bi-door-open-fill"></i>
                            </a>
                        </li>
                        <li v-else class="nav-item">
                            <router-link class="nav-link" to="/login">
                                <i class="bi bi-person-circle"></i>
                            </router-link>
                        </li>
                    </ul>
                </div>
            </nav>
        </div>
    </div>
</template>

<style lang="scss" scoped></style>