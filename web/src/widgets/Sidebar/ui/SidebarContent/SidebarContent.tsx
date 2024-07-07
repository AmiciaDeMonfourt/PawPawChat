import cls from './SidebarContent.module.scss';
import { AppLink, AppLinkTheme } from 'shared/ui/AppLink/AppLink';

import Profile from 'shared/assets/icons/account_circle.svg';
import Settings from 'shared/assets/icons/settings.svg';
import Chats from 'shared/assets/icons/3p.svg';
import Exit from 'shared/assets/icons/exit_to_app.svg';
import Feed from 'shared/assets/icons/view_agenda.svg';

import { IconLabel } from 'shared/ui/IconLabel/IconLabel';
import { useDispatch, useSelector } from 'react-redux';
import { getUser } from 'entities/User/model/selectors/getUser';
import { Button, ButtonTheme } from 'shared/ui/Button/Button';
import { AppDispatch } from 'app/providers/StoreProvider/config/store';
import { userActions } from 'entities/User';
import { useTranslation } from 'react-i18next';

export const SidebarContent = () => {
    const user = useSelector(getUser);
    const dispatch = useDispatch<AppDispatch>();
    const { t } = useTranslation();
    const onLogout = () => dispatch(userActions.logout());

    return (
        <div className={cls.SidebarContent}>
            <div className={cls.TopPanel}>
                <ul className={cls.Options}>
                    <li className={cls.Option}>
                        <AppLink
                            className={cls.option}
                            theme={AppLinkTheme.CLEAR}
                            to={'/feed'}
                        >
                            <IconLabel icon={Feed} text={t('Feed')} />
                        </AppLink>
                    </li>
                    <li className={cls.Option}>
                        <AppLink theme={AppLinkTheme.CLEAR} to={'/chats'}>
                            <IconLabel icon={Chats} text={t('Chats')} />
                        </AppLink>
                    </li>
                    <li className={cls.Option}>
                        <AppLink
                            theme={AppLinkTheme.CLEAR}
                            to={`/profile/${user.username}`}
                        >
                            <IconLabel icon={Profile} text={t('Profile')} />
                        </AppLink>
                    </li>
                    <li className={cls.Option}>
                        <AppLink theme={AppLinkTheme.CLEAR} to={'/settings'}>
                            <IconLabel icon={Settings} text={t('Settings')} />
                        </AppLink>
                    </li>
                </ul>
            </div>
            <div className={cls.BottomPanel}>
                <Button theme={ButtonTheme.CLEAR} onClick={onLogout}>
                    <IconLabel icon={Exit} text={t('Exit')} />
                </Button>
            </div>
        </div>
    );
};

