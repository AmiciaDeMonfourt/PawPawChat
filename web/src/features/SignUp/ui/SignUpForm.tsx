import { classNames } from 'shared/lib/classNames/classNames';
import cls from './SignUpForm.module.scss';
import { useDispatch, useSelector } from 'react-redux';
import { AppDispatch } from 'app/providers/StoreProvider/config/store';
import { SignUp } from '../model/service/SignUp';
import { Button, ButtonTheme } from 'shared/ui/Button/Button';
import { AppRoutes, RoutesPaths } from 'shared/config/routeConfig/routeConfig';
import { AppLink } from 'shared/ui/AppLink/AppLink';
import { Input } from 'shared/ui/Input/Input';
import { getSignUp } from '../model/selectors/getSignUp/getSignUp';
import { useCallback } from 'react';
import { signUpActions } from '../model/slice/signUpSlice';
import { getIsAuth } from 'entities/User/model/selectors/getIsAuth';
import { Navigate } from 'react-router-dom';

interface SignUpFormProps {
    className?: string;
}

export const SignUpForm = ({ className }: SignUpFormProps) => {
    const dispatch = useDispatch<AppDispatch>();
    const data = useSelector(getSignUp);
    const isAuth = useSelector(getIsAuth);

    const { Username, Email, Password, errors } = data;

    const onUsernameChange = useCallback(
        (value: string) => {
            dispatch(signUpActions.setUsername(value));
        },
        [dispatch],
    );

    const onEmailChange = useCallback(
        (value: string) => {
            dispatch(signUpActions.setEmail(value));
        },
        [dispatch],
    );

    const onPasswordChange = useCallback(
        (value: string) => {
            dispatch(signUpActions.setPassword(value));
        },
        [dispatch],
    );

    const onSignUp = useCallback(
        (e: any) => {
            e.preventDefault();
            dispatch(SignUp());
        },
        [dispatch],
    );

    if (isAuth) return <Navigate to={RoutesPaths[AppRoutes.HOME]} />;

    return (
        <div className={classNames(cls.SignUpForm, {}, [className])}>
            <h1 className={cls.title}>Регистрация</h1>
            <form className={cls.form}>
                <Input
                    type="text"
                    placeholder="username"
                    className={cls.username}
                    value={Username}
                    onChange={onUsernameChange}
                />
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
                    togglePasswordVisibility={true}
                    value={Password}
                    onChange={onPasswordChange}
                />
                <Button onClick={onSignUp} theme={ButtonTheme.CLASSIC}>
                    Зарегистрироваться
                </Button>
            </form>
            <AppLink className={cls.noAcc} to={RoutesPaths[AppRoutes.SIGN_IN]}>
                У меня уже есть аккаунт
            </AppLink>
        </div>
    );
};
