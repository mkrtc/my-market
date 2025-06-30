"use client"
import { usePathname } from "next/navigation";
import { FC, useEffect, useMemo, useState } from "react";
import { LayoutService, Path } from "./services/layout.service";
import Link from "next/link";
import styles from "./styles/layout.module.css";


interface LayoutComponentProps{
    children: React.ReactNode
}

export const LayoutComponent:FC<LayoutComponentProps> = ({children}) => {
    console.log(1)
    const pathname = usePathname();
    const [paths, setPaths] = useState<Path[]>([]);
    console.log(2, paths)
    
    useEffect(() => {
        setPaths(() => LayoutService.getPaths(pathname));
    }, [pathname])
    
    console.log(3, paths)
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