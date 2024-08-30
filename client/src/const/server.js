const serverUrl = import.meta.env.VITE_SERVER_URL;

const APIRoot = `${serverUrl}/api`;

export default {
  API_URL: APIRoot,
  API_VERSION: 'v1',
};
