import { formatTimeAgo } from './formatTimeAgo';

describe('formatTime', () => {
    const currentDate = new Date();
    const milliSecondsInMinute = 1000 * 60;

    test('now', () => {
        const now = currentDate.getTime();
        expect(formatTimeAgo(now)).toBe('just now');
    });
    test('1 minute', () => {
        const oneMinuteAgo = currentDate.getTime() - milliSecondsInMinute;
        expect(formatTimeAgo(oneMinuteAgo)).toBe('1 minute ago');
    });
    test('exactly 5 minutes ago', () => {
        const fiveMinutesAgo = currentDate.getTime() - milliSecondsInMinute * 5;

        expect(formatTimeAgo(fiveMinutesAgo)).toBe('5 minutes ago');
    });
    test('not exactly 5 minutes ago', () => {
        const fiveMinutesAgo =
            currentDate.getTime() - milliSecondsInMinute * 5.5;

        expect(formatTimeAgo(fiveMinutesAgo)).toBe('5 minutes ago');
    });
});
