import { Button, Form, Input, Modal } from '@/shared';
import { Dispatch, SetStateAction, useState, type FC } from 'react';
import { SeoService } from '../services/seo.service';
import { Create } from '../../../repositories/seo/seo.types';
import { SeoEntity } from '@/entities';

interface AddSeoViewProps {
    service: SeoService;
    seoState: Dispatch<SetStateAction<SeoEntity[]>>;
}
export const AddSeoView: FC<AddSeoViewProps> = ({ service, seoState }) => {
    const [modalOpened, setModalOpen] = useState<boolean>(false);
    const [error, setError] = useState<string>("");

    const create = async (data: Create) => {
        const seo = await service.createSeo(data);
        if (seo instanceof Error) {
            setError(() => seo.message);
            return;
        }
        setError(() => "");
        seoState(prev => [...prev, seo]);
        setModalOpen(() => false);
    }

    return (
        <>
            <div>
                <Button onClick={() => setModalOpen(() => true)}>Добавить SEO</Button>
            </div>
            <Modal open={modalOpened} onClose={() => setModalOpen(() => false)} footer={<Button type='submit' form='add-seo-form'>Сохранить</Button>} header={error && "Добавить владельца"}>
                <Form id='add-seo-form' action={create}>
                    <div>
                        <Input id='shortName' name='shortName' label="Наименование" />
                        <Input id='fullName' name='fullName' label="Полное наименование" />
                        <Input id='orgName' name='orgName' label="Организация" />
                    </div>
                </Form>
            </Modal>
        </>
    );
}