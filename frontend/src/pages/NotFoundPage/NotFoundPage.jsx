import React, {useContext} from 'react';
import {AUTH_ROUTE} from "../../utils/consts";
import {useNavigate} from "react-router-dom";
import {observer} from "mobx-react-lite";
import {StoreContext} from "../../index";

const NotFoundPage = observer(() => {
        const {user} = useContext(StoreContext)
        const navigate = useNavigate()
        React.useEffect(() => {
            if (user.isAuth) {
                navigate(AUTH_ROUTE)
            }
        }, [])

        return (
            <div>
                    <div>PAGE NOT FOUND</div>
            </div>
        );
    }
)

export default NotFoundPage;