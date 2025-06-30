

export interface Path{
    path: string;
    title: string;
    key: string;
    isActive: boolean;
}

export class LayoutService{
    private static _paths: Path[] = [
        {
            path: "/",
            title: "Главная",
            key: "/",
            isActive: false
        },
        {
            path: "/retail-outlets",
            title: "Тороговые точки",
            key: "/retail-outlets",
            isActive: false
        },
        {
            path: "/work-shift",
            title: "Смены",
            key: "/work-shift",
            isActive: false
        },
    ]

    public static getPaths(pathname: string){
        return this._paths.map(path => ({...path, isActive: path.key === pathname}));
    }
}