import {makeAutoObservable} from "mobx";
import {localStorageParams} from "../utils/consts";

export default class UserStore {
    get userToken() {
        return this._userToken;
    }

    setUserToken(value) {
        localStorage.setItem(localStorageParams.user_token, value)
        this._userToken = value;
    }
    constructor() {
        this._isAuth = false
        this._userToken = ''
        this._userRole = ''
        this._user = {}
        makeAutoObservable(this)
    }

    get isAuth() {
        return this._isAuth;
    }

    setIsAuth(value) {
        this._isAuth = value;
    }

    get userRole() {
        return this._userRole;
    }

     setUserRole(value) {
        localStorage.setItem(localStorageParams.user_role, value)
        this._userRole = value;
    }

    clearCache() {
        this._isAuth = false
        this._userToken = ''
        this._userRole = ''
        this._user = {}
    }
}