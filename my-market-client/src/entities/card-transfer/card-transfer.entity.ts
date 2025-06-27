import { IWorkShiftEntity, WorkShiftEntity } from "../work-shift";


export interface ICardTransferEntity{
    readonly id: number;
    sum: number;
    work_shift_id: number;
    work_shift: IWorkShiftEntity | null;
}

export class CardTransferEntity{
    private readonly _id: number;
    private _sum: number;
    private _workShiftId: number;
    private _workShift: WorkShiftEntity | null;

    constructor(entity: ICardTransferEntity){
        this._id = entity.id;
        this._sum = entity.sum;
        this._workShiftId = entity.work_shift_id;
        this._workShift = entity.work_shift ? new WorkShiftEntity(entity.work_shift) : null;
    }

    public get id(){
        return this._id;
    }

    public get sum(){
        return this._sum;
    }

    public get workShiftId(){
        return this._workShiftId
    }

    public get workShift(){
        return this._workShift;
    }
}