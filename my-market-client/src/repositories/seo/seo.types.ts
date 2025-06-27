

export interface FindOptions{
    limit?: number;
    order?: "asc" | "desc";
}

export interface Create{
    fullName: string;
    shortName: string;
    orgName: string;
}