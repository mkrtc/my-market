import { IWorkShiftEntity, WorkShiftEntity } from "../work-shift";


export interface IExpenseEntity{
    readonly id: number;
    article: string;
    debit: number;
    credit: number;
    payed: boolean;
    work_shift_id: number;
    work_shift: IWorkShiftEntity | null;
}

export class ExpenseEntity{
    private readonly _id: number;
    private _article: string;
    private _debit: number;
    private _credit: number;
    private _payed: boolean;
    private _workShiftId: number;
    private _workShift: WorkShiftEntity | null;

    constructor(entity: IExpenseEntity){
        this._id = entity.id;
        this._article = entity.article;
        this._debit = entity.debit;
        this._credit = entity.credit;
        this._payed = entity.payed;
        this._workShiftId = entity.work_shift_id;
        this._workShift = entity.work_shift ? new WorkShiftEntity(entity.work_shift) : null;
    }

    public get id(){
        return this._id;
    }

    public get article(){
        return this._article;
    }

    public get debit(){
        return this._debit;
    }

    public get credit(){
        return this._credit;
    }

    public get payed(){
        return this._payed;
    }

    public get workShiftId(){
        return this._workShiftId;
    }

    public get workShift(){
        return this._workShift;
    }
}