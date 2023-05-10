import React, {useContext} from 'react';
import './AuthForm.scss'
import {useNavigate} from "react-router-dom";
import {AUTH_ROUTE, formTypes, localStorageParams, RESET_PASSWORD_ROUTE} from "../../utils/consts";
import {resetPassword, signIn} from "../../http/authApi";
import {HTTP_STATUS_CODES} from "../../http/HttpStatus";
import {StoreContext} from "../../index";
import {observer} from "mobx-react-lite";
import SubLine from "../SubLine/SubLine";

const AuthForm = observer(({formType = formTypes.loginForm}) => {

    const {user} = useContext(StoreContext)

    const [login, setLogin] = React.useState('')
    const [password, setPassword] = React.useState('')
    const [isPasswordVisible, setIsPasswordVisible] = React.useState(false)
    const [isDataInvalid, setIsDataInvalid] = React.useState(false)
    const [isRequestFailed, setIsRequestFailed] = React.useState(false)
    const [isPasswordSent, setIsPasswordSent] = React.useState(false)

    const navigate = useNavigate()

    const completeForm = async (e) => {
        e.preventDefault()
        let responseCode = HTTP_STATUS_CODES.NETWORK_ERR

        if (formType === formTypes.loginForm) {
            responseCode = await signIn(login, password)
        }
        if (formType === formTypes.resetPasswordForm) {
            responseCode = await resetPassword(login)
        }

        if (responseCode === HTTP_STATUS_CODES.STATUS_OK) {
            setIsDataInvalid(false)
            setIsRequestFailed(false)

            if (formType === formTypes.loginForm) {
                user.setUserToken(localStorage.getItem(localStorageParams.user_token))
                user.setUserRole(localStorage.getItem(localStorageParams.user_role))
                user.setIsAuth(true)
            }

            if (formType === formTypes.resetPasswordForm) {
                setIsPasswordSent(true)
            }
        }
        if (responseCode === HTTP_STATUS_CODES.INTERNAL_SERVER_ERR || responseCode === HTTP_STATUS_CODES.BAD_REQUEST_ERR) {
            setIsRequestFailed(false)
            setIsDataInvalid(true)
        }
        if (responseCode === HTTP_STATUS_CODES.NETWORK_ERR) {
            setIsRequestFailed(true)
            setIsDataInvalid(false)
        }
    }

    React.useEffect(() => {
        setIsPasswordSent(false)
    }, [])

    return (
        <div className={'auth-form'}>
            {formType === formTypes.loginForm && <div className={'auth-form-title'}>Вход в систему</div>}
            {formType === formTypes.resetPasswordForm && <div className={'auth-form-title'}>Сброс пароля</div>}

            <form className={'auth-form-container'} onSubmit={(e) => completeForm(e)}>
                <SubLine/>
                <div className={'auth-form-field'}>
                    <div className={'auth-form-field-input-container'}>
                        <input className={'auth-form-field-input'}
                               required
                               name={'user_login'}
                               type={'email'}
                               placeholder={'Введите логин'}
                               value={login}
                               onChange={(e) => setLogin(e.target.value)}/>
                    </div>
                </div>
                {formType === formTypes.loginForm && <div className={'auth-form-field'}>
                    <input className={'auth-form-field-input'}
                           required
                           name={'user_password'}
                           type={isPasswordVisible ? 'text' : 'password'}
                           placeholder={'Введите пароль'}
                           value={password}
                           onChange={(e) => setPassword(e.target.value)}/>

                    {isPasswordVisible ? (
                        <svg onClick={() => setIsPasswordVisible(!isPasswordVisible)} width="24" viewBox="0 0 24 24"
                             height="24"
                             fill="none" xmlns="http://www.w3.org/2000/svg">
                            <path xmlns="http://www.w3.org/2000/svg" clipRule="evenodd"
                                  d="m12 6.00018c-6.54545 0-9 6.00002-9 6.00002s2.45455 6 9 6c4.9091 0 9-6 9-6s-4.0909-6.00002-9-6.00002zm-3 6.00002c0 1.6568 1.3431 3 3 3s3-1.3432 3-3c0-1.6569-1.3431-3.00002-3-3.00002s-3 1.34312-3 3.00002z"
                                  fill="rgba(76, 76, 76, 0.5)" fillRule="evenodd"></path>
                        </svg>) : (
                        <svg onClick={() => setIsPasswordVisible(!isPasswordVisible)} width="24" viewBox="0 0 24 24"
                             height="24"
                             fill="none" xmlns="http://www.w3.org/2000/svg">
                            <g xmlns="http://www.w3.org/2000/svg" fill="rgba(76, 76, 76, 0.5)">
                                <path clipRule="evenodd"
                                      d="m11.0957 17.9609c.2924.0257.5938.0392.9045.0392 4.9091 0 9-6 9-6s-.6696-.9821-1.7921-2.15162z"
                                      fillRule="evenodd"></path>
                                <path clipRule="evenodd"
                                      d="m14.5051 6.49498c-.7975-.30791-1.6386-.49486-2.5051-.49486-6.54545 0-9 5.99998-9 5.99998s.75006 1.8335 2.52661 3.4734l3.47339-3.4734c0-1.6568 1.3431-2.99998 3-2.99998z"
                                      fillRule="evenodd"></path>
                                <path d="m5.1001 18.4352h19v2h-19z" opacity=".3"
                                      transform="matrix(.70710678 -.70710678 .70710678 .70710678 -11.54187 9.00586)"></path>
                            </g>
                        </svg>)}
                </div>}

                {formType === formTypes.loginForm && isDataInvalid &&
                    <div className={'auth-form-error-message'}>
                        Неправильный логин или пароль
                    </div>}

                {formType === formTypes.resetPasswordForm && isDataInvalid &&
                    <div className={'auth-form-error-message'}>
                        Неправильный логин
                    </div>}

                {isRequestFailed && <div className={'auth-form-error-message'}>
                    Не удалось выполнить запрос
                </div>}

                {formType === formTypes.loginForm && <>
                    <button className={'auth-form-button'} type={'submit'}>Войти</button>
                    <div className={'reset-password'} onClick={() => navigate(RESET_PASSWORD_ROUTE)}>Сбросить
                        пароль
                    </div>
                </>}

                {formType === formTypes.resetPasswordForm && <>
                    <button className={'auth-form-button'} type={'submit'}>Сбросить
                        пароль
                    </button>

                    <div className={'reset-password'} onClick={() => navigate(AUTH_ROUTE)}>
                        Войти
                    </div>
                </>}

                {formType === formTypes.resetPasswordForm && isPasswordSent && <div className={'auth-form-message'}>
                Пароль выслан
            </div>}
            </form>
        </div>
    );
})

export default AuthForm;