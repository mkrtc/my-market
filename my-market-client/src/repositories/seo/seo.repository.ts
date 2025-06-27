import { ISeoEntity, SeoEntity } from "@/entities";
import { HTTP_CONFIG, HttpProvider } from "@/providers";
import { Create, FindOptions } from "./seo.types";


export class SeoRepository{

    constructor(
        private readonly httpProvider: HttpProvider
    ){}

    public async findById(id: number): Promise<SeoEntity> {
        const seo = await this.httpProvider.get<ISeoEntity>(HTTP_CONFIG.paths.seo.findById, {query: {id}});
        return new SeoEntity(seo);
    }

    public async find(options?: FindOptions): Promise<SeoEntity[]> {
        const seo = await this.httpProvider.get<ISeoEntity[]>(HTTP_CONFIG.paths.seo.findAll, {query: options});
        return seo.map(s => new SeoEntity(s));
    }

    public async create(dto: Create): Promise<SeoEntity> {
        const seo = await this.httpProvider.post<ISeoEntity>(HTTP_CONFIG.paths.seo.create, {body: dto});
        return new SeoEntity(seo);
    }
}