import { MainPage } from "pages/MainPage"
import SignInPage from "pages/SignInPage/ui/SignInPage"
import SignUpPage from "pages/SignUpPage/ui/SignUpPage"
import { RouteProps } from "react-router-dom"

export enum AppRoutes {
    HOME = "home",
    SIGN_UP = "sign_up",
    SIGN_IN = "sign_in",
}

export const RoutesPaths : Record<AppRoutes, string> = {
    [AppRoutes.HOME]: "/",
    [AppRoutes.SIGN_IN]: "/sign_in",
    [AppRoutes.SIGN_UP]: "/sign_up",
}

export const  routeConfig : Record<AppRoutes, RouteProps> = {
    [AppRoutes.HOME]: {
        path: RoutesPaths.home,
        element: <MainPage/>
    },
    [AppRoutes.SIGN_IN]: {
        path: RoutesPaths.sign_in,
        element: <SignInPage/>
    },
    [AppRoutes.SIGN_UP]: {
        path: RoutesPaths.sign_up,
        element: <SignUpPage/>
    }
}