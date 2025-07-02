

export const HTTP_CONFIG = {
    baseUrl: process.env.NEXT_PUBLIC_API_API || "http://localhost",
    port: Number(process.env.NEXT_PUBLIC_API_PORT) || 5000,
    paths: {
        workShift: {
            findAll: "/work-shift",
            findById: "/work-shift/by-id",
            create: "/work-shift",
        },
        retailOutlet: {
            findAll: "/retail-outlet",
            findById: "/retail-outlet/by-id",
            create: "/retail-outlet",
        },
        seo: {
            findAll: "/seo",
            findById: "/seo/by-id",
            create: "/seo",
        }
    }
} as const;