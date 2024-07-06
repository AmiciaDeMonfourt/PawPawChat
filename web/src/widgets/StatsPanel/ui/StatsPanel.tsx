import { classNames } from 'shared/lib/classNames/classNames';
import cls from './StatsPanel.module.scss';

interface StatsPanelProps {
    className?: string;
}

export const StatsPanel = ({ className }: StatsPanelProps) => {
    return <div className={classNames(cls.StatsPanel, {}, [className])}></div>;
};
