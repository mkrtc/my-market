import { Button, Form, Input, Modal } from '@/shared';
import { Dispatch, SetStateAction, useState, type FC } from 'react';
import { WorkShiftEntity } from '@/entities';
import { CreateWsDto, WorkShiftService } from '../services/work-sift.service';
import styles from "../styles/work-shift.module.css";

interface AddWorkShiftViewViewProps {
    service: WorkShiftService;
    wsState: Dispatch<SetStateAction<WorkShiftEntity[]>>;
}
export const AddWorkShiftView: FC<AddWorkShiftViewViewProps> = ({ service, wsState }) => {
    const [modalOpened, setModalOpen] = useState<boolean>(false);
    const [error, setError] = useState<string>("");
    const [cardTransfers, setCardTransfer] = useState<number[]>([]);
    const [expenses, setExpense] = useState<number[]>([]);

    const create = async (data: CreateWsDto) => {
        const ws = await service.create(data);
        if (ws instanceof Error) {
            setError(() => ws.message);
            return;
        }
        setError(() => "");
        wsState(prev => [...prev, ws]);
        setModalOpen(() => false);
    }

    const AddCardTransfer: FC<{ n: number }> = ({ n }) => (
        <>
            <span className={styles.label}>{`Перевод №${n + 1}`}</span>
            <Input id={`cardTransfer-${n}`} name={`cardTransfer-${n}`} label='Сумма' />
        </>
    )

    const AddExpense: FC<{ n: number }> = ({ n }) => (
        <>
            <span className={styles.label}>{`Расход №${n + 1}`}</span>
            <Input id={`expense-${n}-article`} name={`expense-${n}-article`} label='Статья' />
            <Input id={`expense-${n}-debit`} name={`expense-${n}-debit`} label='Дебет' />
            <Input id={`expense-${n}-credit`} name={`expense-${n}-credit`} label='Кредит' />
            <Input id={`expense-${n}-payed`} type='checkbox' name={`expense-${n}-payed`} label='Оплачено' />
        </>
    )


    return (
        <>
            <div>
                <Button onClick={() => setModalOpen(() => true)}>Добавить смену</Button>
            </div>
            <Modal open={modalOpened} onClose={() => setModalOpen(() => false)} footer={<Button type='submit' form='add-seo-form'>Сохранить</Button>} header={error || "Добавить смену"}>
                <Form id='add-seo-form' action={create}>
                    <div className={styles.ws_modal}>
                        <Input id='date' name='date' label="Дата" type='date' value={new Date().toDateString()}/>
                        <Input id='cash' name='cash' label="Наличные" />
                        <Input id='cashLess' name='cashLess' label="Безналичные" />
                        <Input id='cashRegister' name='cashRegister' label="Касса" />
                        <Input id='retailOutletId' name='retailOutletId' label="Торговая точка(ID)" />
                        <div className={styles.ws_modal_btns}>
                            <Button onClick={() => setExpense(prev => [...prev, prev.length])}>Добавить расход</Button>
                            <Button onClick={() => setCardTransfer(prev => [...prev, prev.length])}>Добавить Перевод</Button>
                        </div>
                        <span className={styles.label}>Расходы: {expenses.length}</span>
                        <span className={styles.label}>Переводы: {cardTransfers.length}</span>
                        {expenses.map(n => <AddExpense key={n} n={n} />)}
                        <span className={styles.label}></span>
                        {cardTransfers.map(n => <AddCardTransfer key={n} n={n} />)}
                    </div>
                </Form>
            </Modal>
        </>
    );
}