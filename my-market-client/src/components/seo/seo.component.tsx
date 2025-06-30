"use client"
import { FC, useEffect, useMemo, useState } from "react";
import { SeoService } from "./services/seo.service";
import { SeoEntity } from "@/entities";
import styles from "./styles/seo.module.css";
import { Button, Input, Modal, Table, TableRow } from "@/shared";

export const SeoComponent:FC = () => {
    const service = useMemo(() => new SeoService, []);
    const [seo, setSeo] = useState<SeoEntity[]>([]);
    const [error, setError] = useState<string>("");
    useEffect(() => {
        (async () => {
            const s = await service.getSeo();
            if(s instanceof Error){
                setError(() => s.message);
            }else{
                setSeo(() => s);
            }
        })()
    }, [])
    return (
        <div className={styles.seo}>
            {error && <span>{error}</span>}
            <div>
                <Button>Добавить SEO</Button>
            </div>
            <Modal open>
                <form>
                    <Input />
                </form>
            </Modal>
            <Table cols={["ID", "Наименование", "Полное наименование", "Организация"]}>
                {seo.map(s => (
                    <TableRow key={s.id} values={[
                        {key: s.id, value: s.id},
                        {key: s.fullName, value: s.fullName},
                        {key: s.shortName, value: s.shortName},
                        {key: s.orgName, value: s.orgName},
                    ]}/>
                ))}
            </Table>
        </div>
    )
}