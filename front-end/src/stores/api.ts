import axios from "axios";
import { useAuthStore } from "./auth";

const api = axios.create({
  baseURL: "http://0.0.0.0:8080/api/v1" // replace with your base URL
});

api.interceptors.response.use(
  (response) => response,
  async (error) => {
    const originalRequest = error.config;
    if (error.response.status === 401 && !originalRequest._retry) {
      originalRequest._retry = true;
      await useAuthStore().refresh();
      originalRequest.headers["Authorization"] = "Bearer " + useAuthStore().accessToken;
      return api(originalRequest);
    }
    return Promise.reject(error);
  }
);

export default api;
