<script setup>
import { computed } from "vue";
import { useRouter } from "vue-router";
import { routes } from "@/router";

const siteUrl = import.meta.env.VITE_BUILD_ADDRESS;
const router = useRouter();
const activeRoute = computed(() => router.currentRoute.value.path);
const isActive = (path) => path === activeRoute.value;

</script>

<template>
    <nav class="navbar navbar-expand-lg bg-body-tertiary">
        <div class="container">
            <div class="collapse navbar-collapse" id="navbarNav">
                <ul class="navbar-nav">
                    <li class="nav-item text-uppercase" v-for="route in routes" :key="route.path">
                        <router-link class="nav-link" :to="route.path"
                            v-if="route.children[0].name !== '404' && route.children[0].name !== 'board_detail_id'"
                            :title="route.children[0].name" :class="{ active: isActive(route.path) }">
                            {{ route.children[0].name }}
                        </router-link>
                    </li>
                </ul>
            </div>
        </div>
    </nav>
</template>

<style lang="scss" scoped></style>