import cls from './ProfilePage.module.scss';
import { ContentWrapper } from 'shared/ui/ContentWrapper/ContentWrapper';
import { ProfilePanel } from 'widgets/ProfilePanel/ui/ProfilePanel';

interface ProfilePageProps {
    className?: string;
}

export const ProfilePage = ({ className }: ProfilePageProps) => {
    return (
        <ContentWrapper>
            <ProfilePanel className={cls.ProfilePanel} />
        </ContentWrapper>
    );
};

