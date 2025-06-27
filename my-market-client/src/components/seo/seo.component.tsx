"use client"
import { FC, useEffect, useMemo, useState } from "react";
import { SeoService } from "./services/seo.service";
import { SeoEntity } from "@/entities";
import styles from "./styles/seo.module.css";

export const SeoComponent:FC = () => {
    const service = useMemo(() => new SeoService, []);
    const [seo, setSeo] = useState<SeoEntity[]>([]);
    const [error, setError] = useState<string>("");
    useEffect(() => {
        (async () => {
            const s = await service.getSeo();
            console.log("ue", s)
            if(s instanceof Error){
                setError(() => s.message);
            }else{
                setSeo(() => s);
            }
        })()
    }, [])

    console.log(seo)
    return (
        <div className={styles.seo}>
            {error && <span>{error}</span>}
            <table className={styles.seo_table}>
                <thead>
                    <tr>
                        <th>ID</th>
                        <th>Наименование</th>
                        <th>Полное наименование</th>
                        <th>Организация</th>
                    </tr>
                </thead>
                <tbody>
                    {seo.map(s => (
                        <tr key={s.id}>
                            <td>{s.id}</td>
                            <td>{s.fullName}</td>
                            <td>{s.shortName}</td>
                            <td>{s.orgName}</td>
                        </tr>
                    ))}
                </tbody>
            </table>
        </div>
    )
}