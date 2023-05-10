import React, {useContext} from 'react';
import './HomePage.scss'
import {useNavigate} from "react-router-dom";
import {StoreContext} from "../../index";
import Button from "../../components/Button/Button";
import {buttonProps} from "../../components/Button/ButtonProps";

function HomePage(props) {
    const {user} = useContext(StoreContext)
    const navigate = useNavigate()

    const logOut = () => {
        localStorage.clear()
        user.clearCache()
    }


    return (<div className={'home-page'}>
        <div className={'page-title'}>
            HOME PAGE
        </div>

        <Button text={'Выйти из аккаунта'}
                size={buttonProps.size.small}
                color={buttonProps.color.light}
                bgColor={buttonProps.background_color.dark_v1}
                onClck={() => logOut()}
        />
    </div>);
}

export default HomePage;