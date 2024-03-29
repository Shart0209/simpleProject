import { ref, reactive } from 'vue';
import { useRouter } from 'vue-router';
import { defineStore } from 'pinia';

import { useAuthStore } from '@/stores';

const baseURL = 'http://localhost:8888/apiV1/docs';

export const useDocsStore = defineStore('docs', () => {
  const router = useRouter();
  let items = reactive({
    data: [],
  });
  let item = reactive({
    data: {},
  });
  const error = ref({});
  const optionsSelect = reactive({});

  const { authHeader, refreshToken } = useAuthStore();

  function reset() {
    items.data = [];
    items.error = null;
    item.data = [];
    item.error = null;

    error.value = null;
  }

  function getInitials() {
    let initialAttrs;
    return (initialAttrs = {
      title: '',
      number: '',
      price: '',
      date: '',
      start_date: '',
      end_date: '',
      description: '',
      category: '',
      supplier: '',
      status: true,
      group: '',
      author: '',
    });
  }

  function formatDate(d, locale) {
    switch (locale) {
      case 'ru':
        return new Date(d).toLocaleDateString('ru-RU');
      case 'iso':
        return new Date(d).toLocaleDateString('sv');
    }
  }

  async function getOptionSelect() {
    error.value = null;

    if (optionsSelect.categories != undefined) {
      return;
    }

    try {
      let response = await fetch(`${baseURL}/sps`);
      let result = await response.json();

      if (!response.ok) {
        const message = `${result.errors} ${response.status}`;
        throw new Error(message);
      }

      optionsSelect.categories = result.data.Category;
      optionsSelect.groups = result.data.Groups;
      optionsSelect.suppliers = result.data.Suppliers;
    } catch (err) {
      error.value = err;
      console.log(err);
    }
  }

  async function getAllDocs() {
    reset();
    await getOptionSelect();

    try {
      const response = await fetch(`${baseURL}/`);
      let result = await response.json();

      if (!response.ok) {
        const message = `${result.errors} Status: ${response.status}`;
        throw new Error(message);
      }

      if (result.data === null) {
        const message = 'response is empty';
        throw new Error(message);
      }

      for (const item in result.data) {
        result.data[item]['date'] = formatDate(result.data[item]['date'], 'ru');

        result.data[item].supplier = optionsSelect.suppliers.find(
          (el) => el.id == result.data[item].supplier
        );
        result.data[item].category = optionsSelect.categories.find(
          (el) => el.id == result.data[item].category
        );
        result.data[item].group = optionsSelect.groups.find(
          (el) => el.id == result.data[item].group
        );
      }

      items.data = result.data;
    } catch (err) {
      items.error = err;
      console.log(err);
    }
  }

  async function getByIDDoc(id) {
    reset();
    await getOptionSelect();

    try {
      const response = await fetch(`${baseURL}/${id}`);
      let result = await response.json();

      if (!response.ok) {
        const message = `${result.errors} Status: ${response.status}`;
        throw new Error(message);
      }

      if (result.data === null) {
        const message = 'response is empty';
        throw new Error(message);
      }

      result.data['date'] = formatDate(result.data['date'], 'iso');
      result.data['start_date'] = formatDate(result.data['start_date'], 'iso');
      result.data['end_date'] = formatDate(result.data['end_date'], 'iso');

      let tmp = JSON.stringify(result);
      tmp = JSON.parse(tmp);

      tmp.data.supplier = optionsSelect.suppliers.find(
        (el) => el.id == tmp.data.supplier
      );
      tmp.data.category = optionsSelect.categories.find(
        (el) => el.id == tmp.data.category
      );
      tmp.data.group = optionsSelect.groups.find(
        (el) => el.id == tmp.data.group
      );

      item.data = tmp.data;
    } catch (err) {
      item.error = err;
      console.log(err);
    }
  }

  async function create(form) {
    error.value = null;

    try {
      await fetch(`${baseURL}/add`, {
        method: 'POST',
        headers: authHeader(),
        body: form,
      });
    } catch (err) {
      error.value = err;
      console.log(err);
    }

    router.replace({ name: 'board' });
  }

  function viewFormData(form) {
    for (let [name, value] of form) {
      console.log(`${name} = ${value}`);
    }
  }

  async function update(form, id) {
    error.value = null;

    try {
      let response = await fetch(`${baseURL}/update/${id}`, {
        method: 'POST',
        headers: authHeader(),
        body: form,
      });
      if (!response.ok) {
        const message = `${result.errors} ${response.status}`;
        throw new Error(message);
      }

      router.replace({ name: 'board' });
    } catch (err) {
      error.value = err;
      console.log(err);
    }
  }

  async function deleteByID(id) {
    error.value = null;
    try {
      await fetch(`${baseURL}/delete/${id}`, {
        method: 'DELETE',
        headers: authHeader(),
      });
    } catch (err) {
      error.value = err;
      console.log(err);
    }

    router.replace({ name: 'board' });
  }

  async function downloadFile(docID, fileID, fileName) {
    error.value = null;

    try {
      let response = await fetch(`${baseURL}/download/${fileID}`, {
        method: 'POST',
        headers: authHeader(),
        body: JSON.stringify(docID),
      });

      if (!response.ok) {
        let result = await response.json();

        if (result.errors === 'token has invalid claims: token is expired') {
          refreshToken();
        } else if (result.errors === 'token is empty') {
          throw new Error('Вы не авторизованы');
        }

        const message = `${result.errors} Status: ${response.status}`;
        throw new Error(message);
      }

      let result = await response.blob();

      let a = document.createElement('a');
      a.href = window.URL.createObjectURL(result);
      a.download = fileName;
      a.click();
    } catch (err) {
      error.value = err;
      console.log(err);
    }
  }

  return {
    items,
    item,
    error,
    optionsSelect,
    getAllDocs,
    getByIDDoc,
    getOptionSelect,
    getInitials,
    create,
    deleteByID,
    update,
    formatDate,
    downloadFile,
  };
});
