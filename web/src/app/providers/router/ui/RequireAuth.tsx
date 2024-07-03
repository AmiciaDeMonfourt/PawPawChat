import { getIsAuth } from "entities/User/model/selectors/getIsAuth";
import { ReactNode } from "react";
import { useSelector } from "react-redux";
import { Navigate, Route } from "react-router-dom";
import { AppRoutes, AppRoutesProps, RoutesPaths } from "shared/config/routeConfig/routeConfig"

interface RequireAuthProps {
    defaultRoute?: AppRoutes,
    children: ReactNode,
}

export const RequireAuth = (props : RequireAuthProps) => {

    const isAuth = useSelector(getIsAuth);

    const {
        defaultRoute = AppRoutes.SIGN_IN,
        children
    } = props;

    return (
        isAuth
            ?
                children
            :
                <Navigate to={RoutesPaths[defaultRoute]}/>
    )
}