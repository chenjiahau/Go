import axios from "axios";

import server from "@/const/server";
import message from "@/const/message";

const axiosInstance = axios.create({
  baseURL: server.API_URL,
});

class APIHandler {
  constructor() {
    this.axios = axiosInstance;
    this.token = null;
  }

  setToken(token) {
    this.token = token;
    this.axios.defaults.headers.common["Authorization"] = `Bearer ${token}`;
  }

  get(url, queryString = null, config) {
    if (queryString) {
      url += queryString;
    }

    return this.axios.get(url, config);
  }

  post(url, data, config) {
    return this.axios.post(url, data, config);
  }

  put(url, data, config) {
    return this.axios.put(url, data, config);
  }

  delete(url, config) {
    return this.axios.delete(url, config);
  }

  patch(url, data, config) {
    return this.axios.patch(url, data, config);
  }

  extractSuccessMessage(data) {
    return data.data?.data?.message || message.success.DEFAULT;
  }

  extractErrorMessage(data) {
    return data.response?.data?.error?.message || message.error.UNKNOWN;
  }
}

export default new APIHandler();
