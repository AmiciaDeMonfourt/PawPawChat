import { classNames } from "shared/lib/classNames/classNames";
import cls from "./ProfilePanel.module.scss";
import { Text } from "shared/ui/Text/Text";

interface ProfilePanelProps {
    className?: string
}

export const ProfilePanel = ({className} : ProfilePanelProps) => {
    return (
        <div className={classNames(cls.ProfilePanel, {}, [className])}>
            <div className={cls.person}>
                <div className={cls.avatar}></div>
                <Text>Имя Фамилия</Text>
            </div>
            <div className={cls.info}>
                <ul className={cls.infoList}>
                    <li>
                        <Text>Возраст: 10</Text>
                    </li>
                    <li>
                        <Text>Город: Москва</Text>
                    </li>
                    <li>
                        <Text>Страна: Россия</Text>
                    </li>
                </ul>
            </div>
            <div className={cls.friends}>
                
            </div>
        </div>
    )
}