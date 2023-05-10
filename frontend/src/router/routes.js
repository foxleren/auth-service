import {
    AUTH_ROUTE,
    HOME_ROUTE,
    NOTFOUND_ROUTE,
    RESET_PASSWORD_ROUTE,
} from "../utils/consts";
import AuthPage from "../pages/AuthPage/AuthPage";
import NotFoundPage from "../pages/NotFoundPage/NotFoundPage";
import ResetPasswordPage from "../pages/AuthPage/ResetPasswordPage";
import HomePage from "../pages/HomePage/HomePage";

export const authRoutes = [
    {
        path: HOME_ROUTE,
        Component: <HomePage/>
    },
]

export const publicRoutes = [
    {
        path: AUTH_ROUTE,
        Component: <AuthPage/>
    },
    {
        path: RESET_PASSWORD_ROUTE,
        Component: <ResetPasswordPage/>
    },
    {
        path: NOTFOUND_ROUTE,
        Component: <NotFoundPage/>
    },
]