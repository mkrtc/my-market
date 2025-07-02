

export interface Create{
    cash: number;
    cashLess: number;
    cashRegister: number;
    retailOutletId: number;
    cardTransfers: number[];
    expenses: CreateExpense[];
    date: number;
}

interface CreateExpense{
    article: string;
    debit: number;
    credit: number;
    payed: boolean;
}