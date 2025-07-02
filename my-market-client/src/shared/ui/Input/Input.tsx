import { ChangeEvent, FC, HTMLInputTypeAttribute } from "react";
import styles from "./styles/input.module.css";

interface InputProps{
    type?: HTMLInputTypeAttribute;
    value?: string | number;
    onChange?: (event: ChangeEvent<HTMLInputElement>) => void;
    id?: string;
    name?: string;
    placeholder?: string;
    label?: string;
}

export const Input:FC<InputProps> = ({type, value, id, name, placeholder, label, onChange}) => {
    if (!label) {
        return <input id={id} name={name} placeholder={placeholder} className={styles.input} type={type || "text"} defaultValue={value} onChange={onChange}/>
    }else{
        return (
            <div className={styles.inputWrapper}>
                <label htmlFor={id} className={styles.label}>{label}</label>
                <input id={id} name={name} placeholder={placeholder} className={styles.input} type={type || "text"} defaultValue={value} onChange={onChange}/>
            </div>
        )
    }
}