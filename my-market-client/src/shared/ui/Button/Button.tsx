import { FC, MouseEvent, ReactNode } from "react";
import styles from "./styles/button.module.css";

interface ButtonProps{
    children?: ReactNode;
    onClick?: (event: MouseEvent<HTMLButtonElement>) => void;
    style?: "danger" | "success" | "primary" | "disabled";
    disabled?: boolean;
    type?: "button" | "submit";
    form?: string;
}

export const Button:FC<ButtonProps> = ({children, disabled, type, style, form, onClick}) => {

    return <button form={form} type={type || "button"} disabled={disabled} onClick={disabled ? undefined : onClick} className={`${styles.button} ${styles[`button_${style || "primary"}`]}`}>{children}</button>
}