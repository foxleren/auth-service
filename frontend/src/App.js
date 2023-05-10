import React from 'react';
import AppRouter from "./router/AppRouter";
import {BrowserRouter} from "react-router-dom";
import Header from "./components/Header/Header";
import {localStorageParams} from "./utils/consts";
import {useContext} from "react";
import {observer} from "mobx-react-lite";
import {StoreContext} from "./index";

const App = observer(() => {
    const {user} = useContext(StoreContext)
    const [isLoading, setIsLoading] = React.useState(true)

    React.useEffect(() => {
        const user_token = localStorage.getItem(localStorageParams.user_token)
        if (user_token) {
            user.setUserToken(user_token)
            user.setIsAuth(true)
            const user_role = localStorage.getItem(localStorageParams.user_role)
            user.setUserRole(user_role)

        }
        setIsLoading(false)
    }, [])

    if (isLoading) {
        return (
            <div>
                is loading
            </div>
        )
    }

    return (<div className="App">
        <BrowserRouter>
            <Header/>
            <AppRouter/>
        </BrowserRouter>
    </div>);
})

export default App;
