/* eslint-disable @typescript-eslint/no-unused-vars */
import { HttpProvider } from "@/providers";
import { WorkShiftRepository } from "@/repositories";
import { Create } from "../../../repositories/work-shift/work-shift.types";


export interface CreateWsDto{
    cash: string;
    cashLess: string;
    cashRegister: string;
    retailOutletId: string;
    date: string;
    [key: string]: string;
}

export class WorkShiftService{
    private readonly workShiftRepo: WorkShiftRepository;

    constructor(){
        const httpProvider = new HttpProvider;
        this.workShiftRepo = new WorkShiftRepository(httpProvider);
    }

    public find(){
        try{
            return this.workShiftRepo.find();
        }catch(e){
            return e as Error;
        }
    }

    public async create(dto: CreateWsDto){
        const data: Create = {
            cardTransfers: [],
            expenses: [],
            date: new Date(dto.date).getTime(),
            cash: Number(dto.cash),
            cashLess: Number(dto.cashLess),
            cashRegister: Number(dto.cashRegister),
            retailOutletId: Number(dto.retailOutletId)
        }

        // eslint-disable-next-line @typescript-eslint/no-explicit-any
        const expenses: any[] = [];

        for (const [key, value] of Object.entries(dto)){
            if(key.startsWith("cardTransfer")){
                data.cardTransfers.push(Number(value))
            }

            if(key.startsWith("expense")){
                const [_, i, name] = key.split("-");
                let v: string | number | boolean = value;
                if(name === "payed"){
                    v = value === "on" ? true : false;
                }

                if(name === "credit" || name == "debit"){
                    v = +value;
                }
                if(!expenses[Number(i)]){
                    expenses[Number(i)] = {[name]: v}
                }else{
                    expenses[Number(i)][name] = v;
                }
            }
        }

        data.expenses = expenses;

        try{
            return this.workShiftRepo.create(data);
        }catch(e){
            return e as Error;
        }
    }
}