import { IRetailOutletEntity, RetailOutletEntity } from "../retail-outlet";


export interface ISeoEntity{
    readonly id: number;
    full_name: string;
    short_name: string;
    org_name: string;
    retail_outlets: IRetailOutletEntity[] | null;
}

export class SeoEntity{
    private readonly _id: number;
    private _fullName: string;
    private _shortName: string;
    private _orgName: string;
    private _retailOutlets: RetailOutletEntity[] | null;
    
    constructor(entity: ISeoEntity){
        this._id = entity.id;
        this._fullName = entity.full_name;
        this._shortName = entity.short_name;
        this._orgName = entity.org_name;
        this._retailOutlets = entity.retail_outlets?.map(ro => new RetailOutletEntity(ro)) || null
    }

    public get id(){
        return this._id;
    }

    public get fullName(){
        return this._fullName;
    }

    public get shortName(){
        return this._shortName;
    }

    public get orgName(){
        return this._orgName;
    }
}