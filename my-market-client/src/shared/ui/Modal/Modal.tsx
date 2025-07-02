import { FC, MouseEvent, ReactNode } from "react";
import styles from "./styles/modal.module.css";

interface ModalProps {
    children?: ReactNode;
    open?: boolean;
    onClose?: () => void | Promise<void> | Promise<() => void> | Promise<() => void> | Promise<void> | void;
    footer?: ReactNode;
    header?: ReactNode;
}

export const Modal: FC<ModalProps> = ({ children, open, footer, header, onClose }) => {

    const onModalClose = (event: MouseEvent<HTMLDivElement>) => {
        if ((event.target as HTMLElement).id !== "modal") return;
        onClose?.();
    }

    return (
       <div id="modal" onClick={onModalClose} className={`${styles.modal} ${open ? styles.modal_open : styles.modal_close}`}>
            <div className={styles.modal_body}>
                {header && open &&
                    <div className={styles.modal_header}>
                        {header}
                    </div>
                }
                {open && children}
                {footer && open &&
                    <div className={styles.modal_footer}>
                        {footer}
                    </div>
                }
            </div>
        </div>
    )
}
