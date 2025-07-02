import { IWorkShiftEntity, WorkShiftEntity } from "@/entities";
import { HTTP_CONFIG, HttpProvider } from "@/providers";
import { Create } from "./work-shift.types";


export class WorkShiftRepository{

    constructor(
        private readonly httpProvider: HttpProvider
    ){}

    public async find(): Promise<WorkShiftEntity[]> {
        const ws = await this.httpProvider.get<IWorkShiftEntity[]>(HTTP_CONFIG.paths.workShift.findAll);
        return ws.map(w => new WorkShiftEntity(w));
    }

    public async findById(id: number): Promise<WorkShiftEntity> {
        const ws = await this.httpProvider.get<IWorkShiftEntity>(HTTP_CONFIG.paths.workShift.findById, {query: {id}});
        return new WorkShiftEntity(ws);
    }

    public async create(dto: Create): Promise<WorkShiftEntity> {
        const ws = await this.httpProvider.post<IWorkShiftEntity>(HTTP_CONFIG.paths.workShift.create, {body: dto});
        return new WorkShiftEntity(ws);
    }
}