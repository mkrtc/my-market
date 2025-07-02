"use client"
import { RetailOutletEntity } from '@/entities';
import {useEffect, useMemo, useState, type FC} from 'react';
import { RetailOutletService } from './services/retail-outlet.service';
import { Table, TableRow } from '@/shared';
import styles from "./styles/retail-outlet.module.css";
import { AddRetailOutletView } from './views/AddRetailOutletView';

export const RetailOutletComponent:FC = ({}) => {
    const service = useMemo(() => new RetailOutletService, []);
    const [retailOutlets, setRetailOutlets] = useState<RetailOutletEntity[]>([]);
    const [error, setError] = useState<string>("");

    useEffect(() => {
        (async () => {
            const response = await service.getRetailOutlets();
            if(response instanceof Error) {
                setError(() => response.message);
                return;
            }
            setRetailOutlets(() => response);
        })()
    }, [service])


    return (
        <div className={styles.ro}>
            {error && <span>{error}</span>}
            <AddRetailOutletView roState={setRetailOutlets} service={service}/>
            <Table cols={["ID", "Наименование", "Адресс", "Дата открытия", "Дата закрытия", "Организация"]}>
                {retailOutlets.map(ro => (
                    <TableRow key={ro.id} values={[
                        {key: ro.id, value: ro.id},
                        {key: ro.fullName, value: ro.fullName},
                        {key: ro.address, value: ro.address},
                        {key: ro.openedDate.getTime(), value: ro.openedDate.toLocaleDateString("fr-CH")},
                        {key: ro.closedDate?.getTime() || 0, value: ro.closedDate?.toLocaleDateString("fr-CH") || "-"},
                        {key: `seo-${ro.seoId}`, value: ro.seo?.orgName || "-"},
                        
                    ]}/>
                ))}
            </Table>
        </div>
    );
}