import { ref } from 'vue';
import { useRouter } from 'vue-router';
import { defineStore } from 'pinia';

const baseURL = `${import.meta.env.VITE_API_URL}/apiV2/auth`;

export const useAuthStore = defineStore('auth', () => {
  const router = useRouter();
  const error = ref();
  const user = ref(localStorage.getItem('access_token'));

  async function login(username, password) {
    error.value = null;

    let formData = new FormData();
    formData.append('username', username.value);
    formData.append('password', password.value);

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

      user.value = tk.access_token;

      localStorage.setItem('access_token', tk.access_token);

      // redirect to home page
      router.push({ name: 'home', replace: true });
    } catch (err) {
      error.value = err;
      console.log(err);
    }
  }

  function logout() {
    user.value = null;
    localStorage.removeItem('access_token');
    router.push({ name: 'home', replace: true });
  }

  function authHeader() {
    if (user.value != undefined) {
      return { Authorization: `Bearer ${user.value}` };
    } else {
      return {}
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
      error.value = err
      console.log(err);
    }
  }

  return { user, error, login, logout, authHeader, refreshToken };
});
