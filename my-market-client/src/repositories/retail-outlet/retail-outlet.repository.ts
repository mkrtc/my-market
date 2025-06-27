import { IRetailOutletEntity, RetailOutletEntity } from "@/entities";
import { HTTP_CONFIG, HttpProvider } from "@/providers";
import { Create } from "./retail-outlet.types";


export class RetailOutletRepository{

    constructor(
        private readonly httpProvider: HttpProvider
    ){}

    public async find(): Promise<RetailOutletEntity> {
        const ro = await this.httpProvider.get<IRetailOutletEntity>(HTTP_CONFIG.paths.retailOutlet.findAll);
        return new RetailOutletEntity(ro);
    }

    public async findById(id: number){
        const ro = await this.httpProvider.get<IRetailOutletEntity>(HTTP_CONFIG.paths.retailOutlet.findById, {query: {id}});
        return new RetailOutletEntity(ro);
    }

    public async create(dto: Create){
        const ro = await this.httpProvider.post<IRetailOutletEntity>(HTTP_CONFIG.paths.retailOutlet.create, {body: dto});
        return new RetailOutletEntity(ro);
    }
}