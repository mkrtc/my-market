import { ChangeEvent, FC, HTMLInputTypeAttribute } from "react";
import styles from "./styles/input.module.css";

interface InputProps{
    type?: HTMLInputTypeAttribute;
    value?: string | number;
    onChange?: (event: ChangeEvent<HTMLInputElement>) => void;
    id?: string;
    name?: string;
}

export const Input:FC<InputProps> = ({type, value, id, name, onChange}) => {
    return <input id={id} name={name} className={styles.input} type={type || "text"} defaultValue={value} onChange={onChange}/>
}