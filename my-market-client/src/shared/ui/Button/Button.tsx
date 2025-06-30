import { FC, MouseEvent, ReactNode } from "react";
import styles from "./styles/button.module.css";

interface ButtonProps{
    children?: ReactNode;
    onClick?: (event: MouseEvent<HTMLButtonElement>) => void;
    style?: "danger" | "success" | "primary" | "disabled";
    disabled?: boolean;
    type?: "button" | "submit";
}

export const Button:FC<ButtonProps> = ({children, disabled, type, style, onClick}) => {

    return <button type={type || "button"} disabled={disabled} onClick={disabled ? undefined : onClick} className={`${styles.button} ${styles[`button_${style || "primary"}`]}`}>{children}</button>
}