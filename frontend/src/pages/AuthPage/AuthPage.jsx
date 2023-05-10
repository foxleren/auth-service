import React, {useContext} from 'react';
import './AuthPage.scss'
import AuthForm from "../../components/AuthForm/AuthForm";
import {useNavigate} from "react-router-dom";
import {HOME_ROUTE} from "../../utils/consts";
import {StoreContext} from "../../index";
import {observer} from "mobx-react-lite";

const AuthPage = observer(() => {
    const {user} = useContext(StoreContext)
    const navigate = useNavigate()
    React.useEffect(() => {
        if (user.isAuth) {
            navigate(HOME_ROUTE)
        }
    }, [user.isAuth])

    return (
        <div className={'auth-page'}>
            <AuthForm/>
        </div>
    );
})

export default AuthPage;