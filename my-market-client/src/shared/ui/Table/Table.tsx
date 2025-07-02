import { FC } from "react";
import styles from "./styles/table.module.css";

interface TableProps{
    cols: string[];
    children: React.ReactNode;
}

export const Table:FC<TableProps> = ({cols, children}) => {

    return (
        <table className={styles.table}>
            <thead className={styles.table_header}>
                <tr className={styles.table_row}>
                    {cols.map(col => (
                        <th className={styles.table_row_data} key={col}>{col}</th>
                    ))}
                </tr>
            </thead>
            <tbody className={styles.table_body}>
                {children}
            </tbody>
        </table>
    )
}