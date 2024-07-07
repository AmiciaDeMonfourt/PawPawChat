import { classNames } from 'shared/lib/classNames/classNames';
import cls from './SignInForm.module.scss';
import { useDispatch, useSelector } from 'react-redux';
import { getSignIn } from 'features/SignIn/model/selectors/getSignIn/getSignIn';
import { AppDispatch } from 'app/providers/StoreProvider/config/store';
import { SignIn } from '../../model/service/SignIn';
import { Button, ButtonFont, ButtonTheme } from 'shared/ui/Button/Button';
import { AppRoutes, RoutesPaths } from 'shared/config/routeConfig/routeConfig';
import { AppLink } from 'shared/ui/AppLink/AppLink';
import { Input } from 'shared/ui/Input/Input';
import { useCallback } from 'react';
import { signInActions } from 'features/SignIn/model/slice/signInSlice';
import { getIsAuth } from 'entities/User/model/selectors/getIsAuth';
import { Navigate } from 'react-router-dom';
import { useTranslation } from 'react-i18next';

interface SignInFormProps {
    className?: string;
}

export const SignInForm = ({ className }: SignInFormProps) => {
    const dispatch = useDispatch<AppDispatch>();
    const data = useSelector(getSignIn);
    const isAuth = useSelector(getIsAuth);
    const { t } = useTranslation();
    const { Email, Password, errors } = data;

    const onEmailChange = useCallback(
        (value: string) => {
            dispatch(signInActions.setEmail(value));
        },
        [dispatch],
    );

    const onPasswordChange = useCallback(
        (value: string) => {
            dispatch(signInActions.setPassword(value));
        },
        [dispatch],
    );

    const onSignIn = useCallback(
        (e: any) => {
            e.preventDefault();
            dispatch(SignIn());
        },
        [dispatch],
    );

    if (isAuth) return <Navigate to={RoutesPaths[AppRoutes.FEED]} />;

    return (
        <div className={classNames(cls.SignInForm, {}, [className])}>
            <h1 className={cls.title}>{t('Sign in')}</h1>
            <form className={cls.form}>
                <Input
                    type="text"
                    placeholder="email"
                    className={cls.email}
                    value={Email}
                    onChange={onEmailChange}
                />
                <Input
                    type="password"
                    placeholder="password"
                    className={cls.pass}
                    value={Password}
                    onChange={onPasswordChange}
                    togglePasswordVisibility={true}
                />
                <Button onClick={onSignIn} theme={ButtonTheme.COLORED}>
                    {t('Sign in btn')}
                </Button>
            </form>
            <AppLink className={cls.noAcc} to={RoutesPaths[AppRoutes.SIGN_UP]}>
                {t("Don't have an account yet?")}
            </AppLink>
        </div>
    );
};

