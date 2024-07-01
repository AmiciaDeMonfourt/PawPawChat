import { classNames } from "shared/lib/classNames/classNames";
import cls from "./StoreProvider.module.scss";
import { ReactNode } from "react";
import { Provider } from "react-redux";
import { createReduxStore } from "../config/store";
import { StateSchema } from "../config/StateSchema";

interface StoreProviderProps {
    children?: ReactNode
    initialState?: StateSchema
}

export const StoreProvider = ({children, initialState} : StoreProviderProps) => {
    return (
        <Provider store={createReduxStore(initialState)}>
            {children}
        </Provider>
    )
}