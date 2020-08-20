export interface Import {
    id: number;
    supplier: string;
    user: string;
    importTime: Date | string;
    status: string;
}

export interface ImportDetail {
    id: number;
    importId: number;
    product: string;
    qty: number;
}