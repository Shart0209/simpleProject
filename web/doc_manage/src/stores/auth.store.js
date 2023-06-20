import { computed, ref, reactive } from 'vue';
import { useRouter } from 'vue-router';
import { defineStore } from 'pinia';
import { SHA256 } from 'crypto-js';

import { Buffer } from 'buffer';

const baseURL = `${import.meta.env.VITE_API_URL}/apiV2/auth`;

export const useAuthStore = defineStore('auth', () => {
  const router = useRouter();
  const error = ref(null);

  const access_token = ref(localStorage.getItem('access_token'));
  const token_decode = computed(() => (access_token.value != null && access_token.value !='' ? parseJwt(access_token.value) : null));
  const user = reactive({
    id: computed(()=> token_decode.value ? token_decode.value.sub : null),
    name: computed(()=> token_decode.value ? token_decode.value.name : null),
    role: computed(()=> token_decode.value ? token_decode.value.role : null),
    login: computed(()=> token_decode.value ? token_decode.value.login : null),
    isActive: computed(() => (access_token.value != null && access_token.value !='' ? true : false)),
  });
  
  function parseJwt(token) {
    var base64Payload = token.split('.')[1];
    var payload = Buffer.from(base64Payload, 'base64');
    return JSON.parse(payload.toString());
  }

  async function login(username, password) {
    error.value = null;

    let formData = new FormData();
    formData.append('username', username.value);
    formData.append('password', SHA256(password.value).toString());

    try {
      const response = await fetch(`${baseURL}/login`, {
        method: 'POST',
        body: formData,
      });

      let result = await response.json();

      if (!response.ok) {
        const message = `${result.errors} ${response.status}`;
        throw new Error(message);
      }

      let tk = null;
      tk = JSON.stringify(result);
      tk = JSON.parse(tk);

      access_token.value = tk.access_token;
      
      localStorage.setItem('access_token', tk.access_token);

      // redirect to home page
      router.push({ name: 'home', replace: true });
    } catch (err) {
      error.value = err;
      console.log(err);
    }
  }

  function logout() {
    access_token.value = null;
    localStorage.removeItem('access_token');
    router.push({ name: 'home', replace: true });
  }

  function authHeader() {
    if (user.isActive) {
      return { Authorization: `Bearer ${access_token.value}` };
    } else {
      return {};
    }
  }

  async function refreshToken() {
    error.value = null;
    try {
      let response = await fetch(`${baseURL}/refresh`, {
        method: 'POST',
        headers: authHeader(),
      });

      let result = await response.json();

      if (!response.ok) {
        const message = `${result.errors} Status: ${response.status}`;
        throw new Error(message);
      }

      user.value = result.access_token;

      localStorage.setItem('access_token', result.access_token);
    } catch (err) {
      error.value = err;
      console.log(err);
    }
  }

  return { user, error, login, logout, authHeader, refreshToken };
});
