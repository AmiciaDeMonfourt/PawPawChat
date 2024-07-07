import { classNames } from 'shared/lib/classNames/classNames';
import cls from './ProfilePanel.module.scss';
import { Text, TextSize } from 'shared/ui/Text/Text';
import { useTranslation } from 'react-i18next';
import { DefaultAvatar } from 'shared/ui/DefaultAvatar/DefaultAvatar';
import { Button, ButtonTheme } from 'shared/ui/Button/Button';
import { AppLink } from 'shared/ui/AppLink/AppLink';
import { IconLabel } from 'shared/ui/IconLabel/IconLabel';
import Friends from 'shared/assets/icons/friends.svg';
import Groups from 'shared/assets/icons/groups.svg';
import Subscriptions from 'shared/assets/icons/Subscriptions.svg';
import Subscribers from 'shared/assets/icons/subscribers.svg';

interface ProfilePanelProps {
    className?: string;
}

export const ProfilePanel = ({ className }: ProfilePanelProps) => {
    const { t } = useTranslation();

    return (
        <div className={classNames(cls.ProfilePanel, {}, [className])}>
            <div className={cls.person}>
                <DefaultAvatar name="Viktor" />
                <Text textSize={TextSize.LARGE}>Name</Text>
            </div>
            <div className={cls.info}>
                <div className={cls.infos}>
                    <Text>{`${t('Username')}: ${'viktor'}`}</Text>
                    <Text>{`${t('Phone')}: ${'8 800 555 35 35'}`}</Text>
                    <Text>{`${t('Birthday')}: ${'22.04.2004'}`}</Text>
                    <Button theme={ButtonTheme.COLORED}>{t('More')}</Button>
                </div>
                <div className={cls.groups}>
                    <Button theme={ButtonTheme.COLORED}>
                        <IconLabel
                            icon={Friends}
                            text={`${t('Friends')}: ${100}`}
                        />
                    </Button>{' '}
                    <Button theme={ButtonTheme.COLORED}>
                        <IconLabel
                            icon={Subscribers}
                            text={`${t('Subscribers')}: ${65}`}
                        />
                    </Button>
                    <Button theme={ButtonTheme.COLORED}>
                        <IconLabel
                            icon={Subscriptions}
                            text={`${t('Subscriptions')}: ${65}`}
                        />
                    </Button>
                    <Button theme={ButtonTheme.COLORED}>
                        <IconLabel
                            icon={Groups}
                            text={`${t('Groups')}: ${15}`}
                        />
                    </Button>
                </div>
                <div>
                    <Button theme={ButtonTheme.COLORED}>
                        {t('Edit profile')}
                    </Button>
                </div>
            </div>
        </div>
    );
};

