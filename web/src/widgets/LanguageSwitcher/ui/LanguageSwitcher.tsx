import { Button } from 'shared/ui/Button/Button';
import i18n from 'shared/config/i18n/i18n';
import { useTranslation } from 'react-i18next';

interface LanguageSwitcherProps {
    className?: string;
}

export const LanguageSwitcher = ({ className }: LanguageSwitcherProps) => {
    const { t } = useTranslation();
    const onToggleLanguge = () => {
        i18n.changeLanguage(i18n.language === 'ru' ? 'en' : 'ru');
    };

    return (
        <Button className={className} onClick={onToggleLanguge}>
            {i18n.language}
        </Button>
    );
};
