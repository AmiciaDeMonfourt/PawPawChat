import { classNames } from "shared/lib/classNames/classNames";
import "app/styles/index.scss";
import { AppRouter } from "./providers/router/ui/AppRouter";
import { ThemeSwitcher } from "widgets/ThemeSwitcher";
import { useTheme } from "./providers/ThemeProvider";
import { Navbar } from "widgets/Navbar";
import { Sidebar } from "widgets/Sidebar/ui/Sidebar/Sidebar";
import { ContentWrapper } from "shared/ui/ContentWrapper/ContentWrapper";

export const App = () => {

    const {theme} = useTheme();

    return (
        <div className={classNames("app", {}, [theme])}>
            <Navbar/>
            <AppRouter/>
        </div>
    )
}