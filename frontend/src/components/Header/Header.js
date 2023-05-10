import React, {useContext} from 'react';
import './Header.scss'
import {observer} from 'mobx-react-lite'
import {StoreContext} from "../../index";

const Header = observer(() => {
    return (<header>
       <div className={`header-container auth`}>
            <div className={'header-logo'}>
                LOGO
            </div>
        </div>
    </header>);
})

export default Header;