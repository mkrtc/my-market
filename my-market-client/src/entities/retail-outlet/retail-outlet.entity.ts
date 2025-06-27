import { ISeoEntity, SeoEntity } from "../seo";


export interface IRetailOutletEntity{
    readonly id: number;
    full_name: string;
    address: string;
    opened_date: string;
    closed_date: string;
    seo_id: number;
    seo: ISeoEntity | null;
}

export class RetailOutletEntity{
    private readonly _id: number;
    private _fullName: string;
    private _address: string;
    private _openedDate: Date;
    private _closedDate: Date;
    private _seoId: number;
    private _seo: SeoEntity | null;

    constructor(entity: IRetailOutletEntity){
        this._id = entity.id;
        this._fullName = entity.full_name;
        this._address = entity.address;
        this._openedDate = new Date(entity.opened_date);
        this._closedDate = new Date(entity.closed_date)
        this._seoId = entity.seo_id;
        this._seo = entity.seo ? new SeoEntity(entity.seo) : null;
    }

    public get id(){
        return this._id;
    }

    public get fullName(){
        return this._fullName;
    }

    public get address(){
        return this._address;
    }

    public get openedDate(){
        return this._openedDate;
    }

    public get closedDate(){
        return this._closedDate;
    }

    public get seoId(){
        return this._seoId;
    }

    public get seo(){
        return this._seo;
    }

}