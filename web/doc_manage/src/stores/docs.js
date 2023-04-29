import { ref, reactive } from "vue";
import { defineStore } from "pinia";

export const useDocsStore = defineStore("docs", () => {
  const items = reactive({
    data: [],
  });
  const error = ref(null);
  const optionsSelect = reactive({});

  function reset() {
    items.data = [];
    error.value = null;
  }

  async function getOptionSelect(url) {
    try {
      let response = await fetch(url);
      let result = await response.json();
      if (!response.ok) {
        const message = `An error has occured: ${result.errors} error: ${response.status}`;
        throw new Error(message);
      }

      optionsSelect.categories = result.data.Category;
      optionsSelect.distributors = result.data.Distributor;
    } catch (error) {
      console.log(error);
    }
  }

  function formatDateISO(v) {
    return new Date(v).toLocaleDateString("sv");
  }
  function formatDateRU(v) {
    return new Date(v).toLocaleDateString('ru-RU');
  }
  async function useFetchDocs(url, ext = "", options = {}) {
    reset();

    try {
      let response = await fetch(`${url}${ext}`, options);
      let result = await response.json();
      if (!response.ok) {
        const message = `An error has occured: ${result.errors} error: ${response.status}`;
        throw new Error(message);
      }

      if (result.data.title) {
        const item = result.data;
        items.data.push(item);

        items.data[0].date = formatDateISO(result.data.date);
        items.data[0].start_date = formatDateISO(result.data.start_date);
        items.data[0].end_date = formatDateISO(result.data.end_date);
        console.log(items.data);
        console.log('well done');
      } else {
        for (const item in result.data) {
          items.data.push(result.data[item]);
        }
      }
    } catch (err) {
      error.value = err;
    }
  }

  

  return { items, error, optionsSelect, useFetchDocs, getOptionSelect };
});
