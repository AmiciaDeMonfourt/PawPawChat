import { classNames } from 'shared/lib/classNames/classNames';
import { Button, ButtonTheme } from 'shared/ui/Button/Button';
import {
    ContentWrapper,
    PageAlign,
} from 'shared/ui/ContentWrapper/ContentWrapper';
import cls from './NotFoundPage.module.scss';
import { useNavigate } from 'react-router-dom';

const NotFoundPage = () => {
    const navigate = useNavigate();

    const goBack = () => {
        navigate(-1);
    };

    return (
        <ContentWrapper align={PageAlign.CENTER}>
            <h1 className={classNames('title', {}, [cls.text])}>
                404: Страница не найдена
            </h1>
            <Button theme={ButtonTheme.CLASSIC} onClick={goBack}>
                Назад
            </Button>
        </ContentWrapper>
    );
};

export default NotFoundPage;
