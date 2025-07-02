"use client"
import { usePathname } from "next/navigation";
import { FC, useEffect, useState } from "react";
import { LayoutService, Path } from "./services/layout.service";
import Link from "next/link";
import styles from "./styles/layout.module.css";


interface LayoutComponentProps{
    children: React.ReactNode
}

export const LayoutComponent:FC<LayoutComponentProps> = ({children}) => {
    const pathname = usePathname();
    const [paths, setPaths] = useState<Path[]>([]);
    
    useEffect(() => {
        setPaths(() => LayoutService.getPaths(pathname));
    }, [pathname])
    
    return (
        <>
            <header className={styles.header}>
                {
                    paths.map(path => (
                        <nav key={path.key}>
                            <Link className={`${styles.header_nav} ${path.isActive && styles["header_nav-active"]}`} href={path.path}>{path.title}</Link>
                        </nav>
                    ))
                }
            </header>
            <main>{children}</main>
        </>
    )
}