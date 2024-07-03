import { MainPage } from "pages/MainPage"
import NotFoundPage from "pages/NotFoundPage/ui/NotFoundPage"
import { ProfilePage } from "pages/ProfilePage/ui/ProfilePage"
import SignInPage from "pages/SignInPage/ui/SignInPage"
import SignUpPage from "pages/SignUpPage/ui/SignUpPage"
import { RouteProps } from "react-router-dom"

export enum AppRoutes {
    HOME = "home",
    SIGN_UP = "sign_up",
    SIGN_IN = "sign_in",
    NOT_FOUND = "not_found",
    PROFILE = "profile",
}

export type AppRoutesProps = RouteProps & {
    authOnly?: boolean,
}

export const RoutesPaths : Record<AppRoutes, string> = {
    [AppRoutes.HOME]: "/",
    [AppRoutes.SIGN_IN]: "/sign_in",
    [AppRoutes.SIGN_UP]: "/sign_up",
    [AppRoutes.PROFILE]: "/profile",
    [AppRoutes.NOT_FOUND]: "*",
}

export const  routeConfig : Record<AppRoutes, AppRoutesProps> = {
    [AppRoutes.HOME]: {
        path: RoutesPaths.home,
        element: <MainPage/>,
        authOnly: true,
    },
    [AppRoutes.SIGN_IN]: {
        path: RoutesPaths.sign_in,
        element: <SignInPage/>
    },
    [AppRoutes.SIGN_UP]: {
        path: RoutesPaths.sign_up,
        element: <SignUpPage/>
    },
    [AppRoutes.PROFILE]: {
        path: RoutesPaths.profile,
        element: <ProfilePage/>,
        authOnly: true,
    },
    [AppRoutes.NOT_FOUND]: {
        path: RoutesPaths.not_found,
        element: <NotFoundPage/>,
        authOnly: true
    }
}