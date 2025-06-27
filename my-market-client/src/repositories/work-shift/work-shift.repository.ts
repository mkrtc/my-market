import { IWorkShiftEntity, WorkShiftEntity } from "@/entities";
import { HTTP_CONFIG, HttpProvider } from "@/providers";


export class WorkShiftRepository{

    constructor(
        private readonly httpProvider: HttpProvider
    ){}

    public async find(): Promise<WorkShiftEntity> {
        const ws = await this.httpProvider.get<IWorkShiftEntity>(HTTP_CONFIG.paths.workShift.findAll);
        return new WorkShiftEntity(ws);
    }

    public async findById(id: number): Promise<WorkShiftEntity> {
        const ws = await this.httpProvider.get<IWorkShiftEntity>(HTTP_CONFIG.paths.workShift.findById, {query: {id}});
        return new WorkShiftEntity(ws);
    }

    public async create(dto: any): Promise<WorkShiftEntity> {
        const ws = await this.httpProvider.post<IWorkShiftEntity>(HTTP_CONFIG.paths.workShift.create, {body: dto});
        return new WorkShiftEntity(ws);
    }
}