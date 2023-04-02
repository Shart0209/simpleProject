export function formatDate(v) {
  return new Date(v).toLocaleDateString("ru-RU");
}

export async function useFetch(url, ext = "", options = {}) {
  let error = {
    message: "",
    status: null,
  };

  let res = null;

  try {
    let response = await fetch(`${url}${ext}`, options);
    let data = await response.json();
    if (!response.ok) {
      error.message = data.errors;
      error.status = response.status;

      return { res, error };
    }
    res = data.data;
    error = null;
    
    return { res, error };
  } catch (err) {
    console.log(err);
  }
}
