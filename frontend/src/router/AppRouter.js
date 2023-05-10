import React, {useContext} from 'react';
import {Navigate, Route, Routes} from "react-router-dom";
import {AUTH_ROUTE, HOME_ROUTE, NOTFOUND_ROUTE} from "../utils/consts";
import {authRoutes, publicRoutes} from "./routes";
import {StoreContext} from "../index";
import {observer} from "mobx-react-lite";

const AppRouter = observer(() => {
    const {user} = useContext(StoreContext)

    return (<Routes>
        {user.isAuth && authRoutes.map(({path, Component}) =>
            <Route key={path} path={path} element={Component} exact/>)}
        {publicRoutes.map(({path, Component}) =>
            <Route key={path} path={path} element={Component} exact/>)}

        {user.isAuth && <Route path="/" element={<Navigate to={HOME_ROUTE}/>}/>}

        {!user.isAuth && <Route path="*" element={<Navigate to={AUTH_ROUTE}/>}/>}
        <Route path="*" element={<Navigate to={NOTFOUND_ROUTE} replace/>}/>
    </Routes>);
})

export default AppRouter;