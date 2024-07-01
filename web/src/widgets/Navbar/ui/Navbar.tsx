import { classNames } from "shared/lib/classNames/classNames";
import cls from "./Navbar.module.scss";
import { Link } from "react-router-dom";
import { AppRoutes, RoutesPaths } from "shared/config/routeConfig/routeConfig";
import { ThemeSwitcher } from "widgets/ThemeSwitcher";

interface NavbarProps {
    className?: string
}

export const Navbar = ({className} : NavbarProps) => {
    return (
        <div className={classNames(cls.Navbar, {}, [className])}>
            <Link to={RoutesPaths[AppRoutes.HOME]} className={cls.logo}>PawPawChat🐾</Link>
            <div className={cls.leftPanel}>
                <ThemeSwitcher/>
            </div>
        </div>
    )
}