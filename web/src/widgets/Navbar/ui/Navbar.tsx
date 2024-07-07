import { classNames } from 'shared/lib/classNames/classNames';
import cls from './Navbar.module.scss';
import { Link } from 'react-router-dom';
import { AppRoutes, RoutesPaths } from 'shared/config/routeConfig/routeConfig';
import { ThemeSwitcher } from 'widgets/ThemeSwitcher';
import { LanguageSwitcher } from 'widgets/LanguageSwitcher';

interface NavbarProps {
    className?: string;
}

export const Navbar = ({ className }: NavbarProps) => {
    return (
        <div className={classNames(cls.Navbar, {}, [className])}>
            <Link to={RoutesPaths[AppRoutes.FEED]} className={cls.logo}>
                PawPawChat🐾
            </Link>
            <div className={cls.leftPanel}>
                <LanguageSwitcher className={cls.Lang} />
                <ThemeSwitcher />
            </div>
        </div>
    );
};

