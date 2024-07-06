import { classNames } from 'shared/lib/classNames/classNames';
import cls from './Sidebar.module.scss';
import {
    defaultSidebarWith,
    maxSidebarWith,
    minSidebarWith,
} from 'widgets/Sidebar/config/config';
import { useResize } from 'shared/lib/resize/UseResize';
import { SidebarContent } from '../SidebarContent/SidebarContent';
import { Text } from 'shared/ui/Text/Text';
import { useDispatch, useSelector } from 'react-redux';
import { getUser } from 'entities/User/model/selectors/getUser';

interface SidebarProps {
    className?: string;
}

export const Sidebar = ({ className }: SidebarProps) => {
    const { ref, width, onMouseDown } = useResize({
        defaultWidth: defaultSidebarWith,
        minWidth: minSidebarWith,
        maxWidth: maxSidebarWith,
    });

    return (
        <div
            className={classNames(cls.Sidebar, {}, [className])}
            ref={ref}
            style={{ width: `${width}px` }}
        >
            <SidebarContent />
            <div className={cls.resizer} onMouseDown={onMouseDown} />
        </div>
    );
};
