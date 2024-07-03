import { classNames } from "shared/lib/classNames/classNames";
import "app/styles/index.scss";
import { AppRouter } from "./providers/router/ui/AppRouter";
import { useTheme } from "./providers/ThemeProvider";
import { Navbar } from "widgets/Navbar";
import { useEffect } from "react";
import { useDispatch, useSelector } from "react-redux";
import { AppDispatch } from "./providers/StoreProvider/config/store";
import { initUser } from "entities/User/model/service/initUser";
import { getIsInit } from "entities/User/model/selectors/getIsInit";

export const App = () => {

    const {theme} = useTheme();
    const dispatch = useDispatch<AppDispatch>();
    const isInit = useSelector(getIsInit);

    useEffect(() => {
        dispatch(initUser());
    }, [dispatch]);

    return (
        <div className={classNames("app", {}, [theme])}>
            <Navbar/>
            {!isInit || <AppRouter/>}
        </div>
    )
}