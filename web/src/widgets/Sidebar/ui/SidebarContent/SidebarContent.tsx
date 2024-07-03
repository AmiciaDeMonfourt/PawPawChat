import cls from "./SidebarContent.module.scss";
import { AppLink, AppLinkTheme } from "shared/ui/AppLink/AppLink";

import Profile from "shared/assets/icons/account_circle.svg";
import Settings from "shared/assets/icons/settings.svg";
import Chats from "shared/assets/icons/3p.svg";
import Exit from "shared/assets/icons/exit_to_app.svg";
import Feed from "shared/assets/icons/view_agenda.svg";

import { IconLabel } from "shared/ui/IconLabel/IconLabel";
import { Text } from "shared/ui/Text/Text";
import { useDispatch, useSelector } from "react-redux";
import { getUser } from "entities/User/model/selectors/getUser";
import { Button, ButtonTheme } from "shared/ui/Button/Button";
import { AppDispatch } from "app/providers/StoreProvider/config/store";
import { userActions } from "entities/User";

export const SidebarContent = () => {

    const user = useSelector(getUser);
    const dispatch = useDispatch<AppDispatch>();

    const onLogout = () => dispatch(userActions.logout());

    return (
        <div className={cls.SidebarContent}>
                <div className={cls.TopPanel}>
                    <Text className={cls.user}>{user?.username}</Text>
                    <ul className={cls.Options}>
                        <li>
                            <AppLink className={cls.option} theme={AppLinkTheme.CLEAR} to={"/profile"}>
                                <IconLabel icon={Feed} text={"Лента"}/>
                            </AppLink>
                        </li>
                        <li>
                            <AppLink theme={AppLinkTheme.CLEAR} to={"/chats"}>
                                <IconLabel icon={Chats} text={"Чаты"}/>
                            </AppLink>
                        </li>
                        <li>
                            <AppLink theme={AppLinkTheme.CLEAR} to={"/profile"}>
                                <IconLabel icon={Profile} text={"Профиль"}/>
                            </AppLink>
                        </li>
                        <li>
                            <AppLink theme={AppLinkTheme.CLEAR} to={"/settings"}>
                                <IconLabel icon={Settings} text={"Настройки"}/>
                            </AppLink>
                        </li>
                    </ul>
                </div>
                <div className={cls.BottomPanel}>
                    <Button
                        theme={ButtonTheme.CLEAR}
                        onClick={onLogout}
                    >
                        <IconLabel icon={Exit} text={"Выйти"}/>
                    </Button>
                </div>
            </div>
    )
}