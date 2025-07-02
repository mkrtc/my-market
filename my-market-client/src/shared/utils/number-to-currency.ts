

export const numberToCurrency = (num: number) => {
    return Intl.NumberFormat("ru", {
        currency: "RUB",
        style: "currency"
    }).format(num)
}