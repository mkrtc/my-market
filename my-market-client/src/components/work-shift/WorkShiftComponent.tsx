"use client"
import { useEffect, useMemo, useState, type FC } from 'react';
import styles from "./styles/work-shift.module.css";
import { WorkShiftService } from './services/work-sift.service';
import { WorkShiftEntity } from '@/entities';
import { numberToCurrency, Table, TableRow } from '@/shared';
import { AddWorkShiftView } from './views/AddWorkShiftView';

export const WorkShiftComponent: FC = ({ }) => {
    const service = useMemo(() => new WorkShiftService, []);
    const [workShifts, setWorkShifts] = useState<WorkShiftEntity[]>([]);
    const [error, setError] = useState<string>("");

    useEffect(() => {
        (async () => {
            const response = await service.find();
            if (response instanceof Error) {
                setError(() => response.message);
                return;
            }
            setWorkShifts(() => response);
        })()
    }, [service])


    return (
        <div className={styles.ws}>
            {error && <span>{error}</span>}
            <AddWorkShiftView service={service} wsState={setWorkShifts}/>
            <Table cols={["ID", "Дата", "Нал", "Б/Н", "Переводы", "Касса", "Торговая точка", "Выручка", "Расходы", "Касса(рс)", "Остаток"]}>
                {workShifts.map(ws => (
                    <TableRow key={ws.id} values={[
                        { key: ws.id, value: ws.id },
                        { key: 'ws-date', value: ws.date.toLocaleDateString("fr-CH") },
                        { key: 'ws-cash', value: numberToCurrency(ws.cash) },
                        { key: 'ws-cash-less', value: numberToCurrency(ws.cashLess) },
                        { key: 'ws-card-transfers', value: numberToCurrency(ws.cardTransfers.reduce((c, t) => c += t.sum, 0)) },
                        { key: 'ws-cash-register', value: numberToCurrency(ws.cashRegister) },
                        { key: 'ws-full-name', value: ws.retailOutlet?.fullName },
                        { key: `ws-profit`, value: numberToCurrency((ws.cash + ws.cashLess) + ws.cardTransfers.reduce((c, t) => c += t.sum, 0) )},
                        { key: 'ws-expenses', value: numberToCurrency(ws.expenses.reduce((c, e) => c += e.debit, 0)) },
                        { key: 'ws-cash-register-rc', value: numberToCurrency(ws.cash - ws.expenses.reduce((c, e) => c += e.debit, 0)) },
                        { key: 'ws-remainder', value: numberToCurrency(ws.cashRegister + (ws.cash + ws.cashLess) + ws.cardTransfers.reduce((c, t) => c += t.sum, 0) - ws.expenses.reduce((c, e) => c += e.debit, 0))}
                    ]} />
                ))}
            </Table>
        </div>
    );
}