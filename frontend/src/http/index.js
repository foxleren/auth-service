import axios from "axios";
import {localStorageParams} from "../utils/consts";

const $host = axios.create({
    baseURL: process.env.REACT_APP_API_URL
})

const $authHost = axios.create({
    baseURL: process.env.REACT_APP_API_URL
})

const authInterceptor = (config) => {
    config.headers.authorization = `Bearer ${localStorage.getItem(localStorageParams.user_token)}`
    return config
}

$authHost.interceptors.request.use(authInterceptor)

export {
    $host,
    $authHost
}