import ChatsPage from 'pages/ChatsPage/ui/ChatsPage';
import { MainPage } from 'pages/MainPage';
import NotFoundPage from 'pages/NotFoundPage/ui/NotFoundPage';
import { ProfilePage } from 'pages/ProfilePage/ui/ProfilePage';
import SignInPage from 'pages/SignInPage/ui/SignInPage';
import SignUpPage from 'pages/SignUpPage/ui/SignUpPage';
import { Navigate, RouteProps } from 'react-router-dom';

export enum AppRoutes {
    ROOT = 'root',
    FEED = 'feed',
    SIGN_UP = 'sign_up',
    SIGN_IN = 'sign_in',
    PROFILE = 'profile',
    CHATS = 'chats',

    NOT_FOUND = 'not_found',
}

export type AppRoutesProps = RouteProps & {
    authOnly?: boolean;
};

export const RoutesPaths: Record<AppRoutes, string> = {
    [AppRoutes.ROOT]: '/',
    [AppRoutes.FEED]: '/feed',

    [AppRoutes.SIGN_IN]: '/sign_in',
    [AppRoutes.SIGN_UP]: '/sign_up',
    [AppRoutes.PROFILE]: '/profile/:username',
    [AppRoutes.CHATS]: '/chats',

    [AppRoutes.NOT_FOUND]: '*',
};

export const routeConfig: Record<AppRoutes, AppRoutesProps> = {
    [AppRoutes.ROOT]: {
        path: RoutesPaths.root,
        element: <Navigate to={RoutesPaths[AppRoutes.FEED]} />,
        authOnly: true,
    },
    [AppRoutes.FEED]: {
        path: RoutesPaths.feed,
        element: <MainPage />,
        authOnly: true,
    },
    [AppRoutes.SIGN_IN]: {
        path: RoutesPaths.sign_in,
        element: <SignInPage />,
    },
    [AppRoutes.SIGN_UP]: {
        path: RoutesPaths.sign_up,
        element: <SignUpPage />,
    },
    [AppRoutes.PROFILE]: {
        path: RoutesPaths.profile,
        element: <ProfilePage />,
        authOnly: true,
    },
    [AppRoutes.CHATS]: {
        path: RoutesPaths.chats,
        element: <ChatsPage />,
        authOnly: true,
    },
    [AppRoutes.NOT_FOUND]: {
        path: RoutesPaths.not_found,
        element: <NotFoundPage />,
        authOnly: true,
    },
};
