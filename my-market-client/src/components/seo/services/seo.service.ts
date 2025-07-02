import { SeoEntity } from "@/entities";
import { HttpProvider } from "@/providers";
import { SeoRepository } from "@/repositories";
import { Create } from "../../../repositories/seo/seo.types";


export class SeoService{
    private readonly seoRepo: SeoRepository;

    constructor(){
        const httpProvider = new HttpProvider;
        this.seoRepo = new SeoRepository(httpProvider);
    }

    public async getSeo(): Promise<SeoEntity[] | Error>{
        try{
            return await this.seoRepo.find();
        }catch(e){
            return e as Error;
        }
    }

    public async createSeo(dto: Create){
        try{
            return await this.seoRepo.create(dto);
        }catch(e){
            return e as Error;
        }
    }
}