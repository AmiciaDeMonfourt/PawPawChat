import { SignInForm } from 'features/SignIn';
import {
    ContentWrapper,
    PageAlign,
} from 'shared/ui/ContentWrapper/ContentWrapper';

const SignInPage = () => {
    return (
        <ContentWrapper fullscreen={true} align={PageAlign.CENTER}>
            <SignInForm />
        </ContentWrapper>
    );
};

export default SignInPage;
