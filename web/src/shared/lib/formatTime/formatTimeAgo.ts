export function formatTimeAgo(timeInMilliseconds: number) {
    const targetDate = new Date(timeInMilliseconds);
    const currentDate = new Date();

    const deltaTimeAtMinutes =
        (currentDate.getTime() - targetDate.getTime()) / 1000 / 60;
    const deltaTimeAtHours = deltaTimeAtMinutes / 60;
    const deltaTimeAtDays = deltaTimeAtHours / 24;

    if (deltaTimeAtMinutes < 1) {
        return 'just now';
    }

    if (deltaTimeAtMinutes >= 1 && deltaTimeAtMinutes < 60) {
        return `${Math.floor(deltaTimeAtMinutes)} minute${Math.floor(deltaTimeAtMinutes) === 1 ? '' : 's'} ago`;
    }

    if (deltaTimeAtHours > 1 && deltaTimeAtHours <= 3) {
        if (Math.floor(deltaTimeAtHours) === 1)
            return `${parseInt(deltaTimeAtHours.toString())} hour ago`;

        return `${parseInt(deltaTimeAtHours.toString())} hours ago`;
    }

    if (deltaTimeAtHours > 3 && deltaTimeAtHours < 24) {
        const timeStr = `${targetDate.getHours().toString().padStart(2, '0')}
                        :${targetDate.getMinutes().toString().padStart(2, '0')}`;
        if (currentDate.getHours() - deltaTimeAtHours > 0) {
            return `today at ${timeStr}`;
        }

        return `yesterday at ${timeStr}`;
    }

    if (deltaTimeAtDays >= 1 && deltaTimeAtDays < 30) {
        return `${Math.floor(deltaTimeAtDays)} day${Math.floor(deltaTimeAtDays) === 1 ? '' : 's'} ago`;
    }

    return 'bbb';
}
