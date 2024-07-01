import { classNames } from "shared/lib/classNames/classNames";
import cls from "./Sidebar.module.scss";
import { useRef, useState } from "react";
import { defaultSidebarWith, maxSidebarWith, minSidebarWith } from "widgets/Sidebar/config/config";
import { ThemeSwitcher } from "widgets/ThemeSwitcher";

interface SidebarProps {
    className?: string
}

export const Sidebar = ({className} : SidebarProps) => {

    const [sidebarWidth, setSidebarWidth] = useState(defaultSidebarWith);

    const sidebarRef = useRef(null);

    const onMouseDown = () => {
        document.addEventListener('mousemove', onMouseMove);
        document.addEventListener('mouseup', onMouseUp);
        document.body.classList.add("no-select");
    }
    const onMouseMove = (e : any) => {
        const newWidth = e.clientX - sidebarRef.current.getBoundingClientRect().left;
        const newCorrectWidth = Math.min(Math.max(minSidebarWith, newWidth), maxSidebarWith);

        setSidebarWidth(newCorrectWidth);
    }
    const onMouseUp = () => {
        document.removeEventListener('mousemove', onMouseMove);
        document.removeEventListener('mouseup', onMouseUp);
        document.body.classList.remove("no-select");
    }

    return (
        <div 
            className={classNames(cls.Sidebar, {}, [className])}
            ref={sidebarRef}
            style={{width: `${sidebarWidth}px`}}
        >
            <div className={cls.SidebarContent}>

            </div>
            <div
                className={cls.resizer}
                onMouseDown={onMouseDown}
            />
            <div className={cls.switchers}>
            </div>
        </div>
    )
}