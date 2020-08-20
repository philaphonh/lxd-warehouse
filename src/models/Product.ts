export default interface Product {
    id: number;
    name: string;
    brand: string;
    type: string;
    description: string;
    price: number;
    qty: number;
    imageUrl: string;
}

export interface ImportedProduct {
    id: number;
    name: string;
    brand: string;
    type: string;
    description: string;
    price: number;
    qty: number;
    quantity: number;
    imageUrl: string; 
}

export interface ExportedProduct {
    id: number;
    name: string;
    brand: string;
    type: string;
    description: string;
    price: number;
    qty: number;
    quantity: number;
    imageUrl: string;  
}