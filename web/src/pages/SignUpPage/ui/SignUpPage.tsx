import { SignUpForm } from "features/SignUp";
import { ContentWrapper, PageAlign } from "shared/ui/ContentWrapper/ContentWrapper";



const SignUpPage = () => {
    return (
        <ContentWrapper fullscreen={true} align={PageAlign.CENTER}>
            <SignUpForm/>
        </ContentWrapper>
    )
}

export default SignUpPage;