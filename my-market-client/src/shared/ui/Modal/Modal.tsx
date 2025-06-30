import { FC, ReactNode } from "react";
import styles from "./styles/modal.module.css";

interface ModalProps{
    children?: ReactNode;
    open?: boolean;
}

export const Modal:FC<ModalProps> = ({children, open}) => {

    return (
        <div className={`${styles.modal} ${open ? styles.modal_open : styles.modal_close}`}>
            <div className={styles.modal_body}>
        
            </div>
        </div>
    )
}
