const serverUrl = import.meta.env.VITE_SERVER_URL;

const APIRoot = `http://${serverUrl}:8080/api`;

export default {
  API_URL: APIRoot,
  API_VERSION: 'v1',
};
