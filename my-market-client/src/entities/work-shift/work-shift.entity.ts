import { IRetailOutletEntity, RetailOutletEntity } from "../retail-outlet";


export interface IWorkShiftEntity{
    readonly id: number;
    cash: number;
    cash_less: number;
    cash_register: number;
    retail_outlet_id: number;
    retail_outlet: IRetailOutletEntity | null;
    createdAt: string;
	// CardTransfers  []CardTransferModel   `gorm:"foreignKey:WorkShiftId" json:"card_transfers"`
	// Expenses       []ExpenseModel        `gorm:"foreignKey:WorkShiftId" json:"expenses"`
}

export class WorkShiftEntity{
    private readonly _id: number;
    private _cash: number;
    private _cashLess: number;
    private _cashRegister: number;
    private _retailOutletId: number;
    private _retailOutlet: RetailOutletEntity | null;
    private _createdAt: Date;

    constructor(entity: IWorkShiftEntity){
        this._id = entity.id;
        this._cash = entity.cash;
        this._cashLess = entity.cash_less;
        this._cashRegister = entity.cash_register;
        this._retailOutletId = entity.retail_outlet_id;
        this._retailOutlet = entity.retail_outlet ? new RetailOutletEntity(entity.retail_outlet) : null;
        this._createdAt = new Date(entity.createdAt);
    }

    public get id(){
        return this._id;
    }
    public get cash(){
        return this._cash;
    }
    public get cashLess(){
        return this._cashLess;
    }
    public get cashRegister(){
        return this._cashRegister;
    }
    public get retailOutletId(){
        return this._retailOutletId;
    }
    public get retailOutlet(){
        return this._retailOutlet;
    }
    public get createdAt(){
        return this._createdAt;
    }
}