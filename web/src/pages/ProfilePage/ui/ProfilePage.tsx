import { classNames } from "shared/lib/classNames/classNames";
import cls from "./ProfilePage.module.scss";
import { ContentWrapper } from "shared/ui/ContentWrapper/ContentWrapper";
import { ProfilePanel } from "widgets/ProfilePanel/ui/ProfilePanel";
import { StatsPanel } from "widgets/StatsPanel";

interface ProfilePageProps {
    className?: string
}

export const ProfilePage = ({className} : ProfilePageProps) => {
    return (
            <ContentWrapper>
                <ProfilePanel className={cls.ProfilePanel}/>
            </ContentWrapper>
    )
}