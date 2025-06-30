import { FC, Key, MouseEvent, ReactNode } from "react";
import styles from "./styles/table.module.css";

export interface TableRowValue{
    value?: ReactNode;
    key?: Key;
}

interface TableRowProps{
    values: TableRowValue[],
    onClick?: (values: TableRowValue[], event: MouseEvent<HTMLTableRowElement>) => void
}

export const TableRow:FC<TableRowProps> = ({values, onClick}) => {

    return (
        <tr className={styles.table_row} onClick={e => onClick?.(values, e)}>
            {values.map(value => (
                <td key={value.key} className={styles.table_row_data}>{value.value}</td>
            ))}
        </tr>
    )
}