
export type RequestMethod = "GET" | "POST" | "PATCH" | "PUT" | "DELETE"

export interface RequestOptions{
    query?: object;
    params?: object;
    body?: object;
}