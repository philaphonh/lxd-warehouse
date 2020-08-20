export interface Export {
    id: number;
    distributor: string;
    user: string;
    exportTime: Date | string;
}

export interface ExportDetail {
    id: number;
    exportId: number;
    product: string;
    qty: number;
}