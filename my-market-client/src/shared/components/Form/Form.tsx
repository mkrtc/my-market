import type { FC, ReactNode } from 'react';

interface FormProps {
    children?: ReactNode;
    // eslint-disable-next-line @typescript-eslint/no-explicit-any
    action?: (data: any) => void | Promise<void>;
    id?: string;
}
export const Form: FC<FormProps> = ({ action, id, children }) => {

    const onAction = (formData: FormData) => {
        const data = Object.fromEntries(formData);
        action?.(data);
    }

    return (
        <form id={id} action={onAction}>
            {children}
        </form>
    );
}