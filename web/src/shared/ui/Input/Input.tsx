import { classNames } from "shared/lib/classNames/classNames";
import cls from "./Input.module.scss";
import { InputHTMLAttributes, useState } from "react";
import { Button } from "../Button/Button";
import show from "shared/assets/images/show.png";
import hide from "shared/assets/images/hide.png";
import showWhite from "shared/assets/images/show-white.png";
import hideWhite from "shared/assets/images/hide-white.png";
import { useTheme } from "app/providers/ThemeProvider";
import { Theme } from "app/providers/ThemeProvider/lib/ThemeContext";

type HTMLInputProps = Omit<InputHTMLAttributes<HTMLInputElement>, 'value' | 'onChange' | 'readOnly'>

interface InputProps extends HTMLInputProps {
    className?: string;
    value?: string | number;
    onChange?: (value: string) => void;
    togglePasswordVisibility?: boolean;
}

export const Input = (props : InputProps) => {
    const {
        className,
        type = "text",
        togglePasswordVisibility = false,
        onChange,
        placeholder,
        value,
        ...othetProps
    } = props;

    const {theme} = useTheme();

    const showIcon = theme === Theme.DARK ? showWhite : show;
    const hideIcon = theme === Theme.DARK ? hideWhite : hide;

    const [visible, setVisible] = useState(false);

    const onChangeHandler = (e: React.ChangeEvent<HTMLInputElement>) => {
        onChange?.(e.target.value);
    }


    const toggleVisibility = (e : any) => {
        e.preventDefault();
        
        setVisible(prev => !prev);
    }

    const getType = () : string => {
        if(togglePasswordVisibility) {
            return visible ? "text" : "password";
        }
        
        return type;
    }

    return (
        <div
            className={classNames("", {}, [className, cls.wrapper])}
        >
            <input
                className={cls.Input}
                type={getType()}
                value={value}
                placeholder={placeholder}
                onChange={onChangeHandler}
                {...othetProps}
            />
            {togglePasswordVisibility
                ?
                    <Button
                        className={cls.pass}
                        onClick={toggleVisibility}
                    >
                        <img src={visible ? showIcon : hideIcon} alt="show"/>
                    </Button>
                :
                    null
            }
        </div>

    )
}