import { HttpProvider } from "@/providers";
import { RetailOutletRepository } from "@/repositories";
import { Create } from "../../../repositories/retail-outlet/retail-outlet.types";

export type CreateRoDto = {
    openedDate: string;
    closedDate: string
} & Create;

export class RetailOutletService{
    private readonly retailOutletRepo: RetailOutletRepository;

    constructor(){
        const httpProvider = new HttpProvider;
        this.retailOutletRepo = new RetailOutletRepository(httpProvider);
    }

    public getRetailOutlets(){
        try{
            return this.retailOutletRepo.find();
        }catch(e){
            return e as Error;
        }
    }

    public createRetailOutlet(dto: CreateRoDto){
        const data: Create = {
            ...dto,
            seoId: +dto.seoId,
            openedDate: new Date(dto.openedDate).getTime(),
            closedDate: dto.closedDate ? new Date(dto.closedDate).getTime() : 0
        }
        try{
            return this.retailOutletRepo.create(data);
        }catch(e){
            return e as Error;
        }
    }
}