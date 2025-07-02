import { Button, Form, Input, Modal } from '@/shared';
import { Dispatch, SetStateAction, useState, type FC } from 'react';
import { RetailOutletEntity } from '@/entities';
import { CreateRoDto, RetailOutletService } from '../services/retail-outlet.service';

interface AddRetailOutletViewViewProps {
    service: RetailOutletService;
    roState: Dispatch<SetStateAction<RetailOutletEntity[]>>;
}
export const AddRetailOutletView: FC<AddRetailOutletViewViewProps> = ({ service, roState }) => {
    const [modalOpened, setModalOpen] = useState<boolean>(false);
    const [error, setError] = useState<string>("");

    const create = async (data: CreateRoDto) => {
        const ro = await service.createRetailOutlet(data);
        if (ro instanceof Error) {
            setError(() => ro.message);
            return;
        }
        setError(() => "");
        roState(prev => [...prev, ro]);
        setModalOpen(() => false);
    }

    return (
        <>
            <div>
                <Button onClick={() => setModalOpen(() => true)}>Добавить торговую точку</Button>
            </div>
            <Modal open={modalOpened} onClose={() => setModalOpen(() => false)} footer={<Button type='submit' form='add-seo-form'>Сохранить</Button>} header={error && "Добавить торгову. точку"}>
                <Form id='add-seo-form' action={create}>
                    <div>
                        <Input id='fullName' name='fullName' label="Нименование" />
                        <Input id='address' name='address' label="Адрес" />
                        <Input type='date' id='openedDate' name='openedDate' label="Дата открытия" />
                        <Input type='date' id='closedDate' name='closedDate' label="Дата закрытия" />
                        <Input id='seoId' name='seoId' label="SEO" />
                    </div>
                </Form>
            </Modal>
        </>
    );
}