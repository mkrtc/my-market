import { CardTransferEntity, ICardTransferEntity } from "../card-transfer";
import { ExpenseEntity, IExpenseEntity } from "../expense";
import { IRetailOutletEntity, RetailOutletEntity } from "../retail-outlet";


export interface IWorkShiftEntity {
    readonly id: number;
    cash: number;
    cash_less: number;
    cash_register: number;
    retail_outlet_id: number;
    retail_outlet: IRetailOutletEntity | null;
    date: string;
    createdAt: string;
    card_transfers: ICardTransferEntity[]
    expenses: IExpenseEntity[]
}

export class WorkShiftEntity {
    private readonly _id: number;
    private _cash: number;
    private _cashLess: number;
    private _cashRegister: number;
    private _retailOutletId: number;
    private _retailOutlet: RetailOutletEntity | null;
    private _date: Date;
    private _cardTransfers: CardTransferEntity[];
    private _expenses: ExpenseEntity[];
    private _createdAt: Date;

    constructor(entity: IWorkShiftEntity) {
        this._id = entity.id;
        this._cash = entity.cash;
        this._cashLess = entity.cash_less;
        this._cashRegister = entity.cash_register;
        this._retailOutletId = entity.retail_outlet_id;
        this._retailOutlet = entity.retail_outlet ? new RetailOutletEntity(entity.retail_outlet) : null;
        this._cardTransfers = entity.card_transfers.map((ct) => new CardTransferEntity(ct));
        this._expenses = entity.expenses.map((e) => new ExpenseEntity(e));
        this._createdAt = new Date(entity.createdAt);
        this._date = new Date(entity.date);
    }

    public get id() {
        return this._id;
    }
    public get cash() {
        return this._cash;
    }
    public get cashLess() {
        return this._cashLess;
    }
    public get cashRegister() {
        return this._cashRegister;
    }
    public get retailOutletId() {
        return this._retailOutletId;
    }
    public get retailOutlet() {
        return this._retailOutlet;
    }
    public get cardTransfers() {
        return this._cardTransfers;
    }
    public get expenses() {
        return this._expenses;
    }
    public get createdAt() {
        return this._createdAt;
    }
    public get date() {
        return this._date;
    }
}