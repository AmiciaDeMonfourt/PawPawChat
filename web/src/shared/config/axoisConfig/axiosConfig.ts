import axios from "axios";
import { LOCAL_STORAGE_TOKEN_KEY } from "shared/consts/localstorage";

export const $api = axios.create({
    baseURL: 'http://localhost:8080/',
});

$api.interceptors.request.use((config) => {
    if (config.headers && localStorage.getItem(LOCAL_STORAGE_TOKEN_KEY)) {
        config.headers.Authorization = `Bearer ${localStorage.getItem(LOCAL_STORAGE_TOKEN_KEY)}`;
    }
    return config;
});