export const DateParser = (date) => {
    const year = date.slice(0, 4);
    const month = date.slice(4,7);
    const day = date.slice(6);

    const dateFormatted = `${month}/${day}/${year}`
    return dateFormatted
}