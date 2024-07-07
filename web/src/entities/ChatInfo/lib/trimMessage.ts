const maxLength = 30;

export function trimMessage(message: string): string {
    if (message.length > maxLength) {
        return message.substring(0, maxLength - 3) + '...';
    }

    return message;
}
