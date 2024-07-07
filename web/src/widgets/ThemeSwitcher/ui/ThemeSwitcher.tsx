import { useTheme } from 'app/providers/ThemeProvider';
import { Theme } from 'app/providers/ThemeProvider/lib/ThemeContext';
import { classNames } from 'shared/lib/classNames/classNames';
import { Button, ButtonTheme } from 'shared/ui/Button/Button';

import cls from './ThemeSwitcher.module.scss';

interface ThemeSwitcherProps {
    className?: string;
}

export const ThemeSwitcher = ({ className }: ThemeSwitcherProps) => {
    const { theme, toggleTheme } = useTheme();
    return (
        <Button
            theme={ButtonTheme.CLEAR}
            onClick={toggleTheme}
            className={classNames(cls.ThemeSwitcher, {}, [className])}
        >
            {`${theme === Theme.DARK ? '🌙' : '☀️'}`}
        </Button>
    );
};

