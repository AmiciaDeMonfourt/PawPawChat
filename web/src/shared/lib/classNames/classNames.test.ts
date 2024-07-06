import { describe } from 'node:test';
import { classNames } from './classNames';

describe('classNames', () => {
    test('single class', () => {
        expect(classNames('red')).toBe('red');
    });
    test('additional classes', () => {
        expect(classNames('red', {}, ['active', 'dotted'])).toBe(
            'red active dotted',
        );
    });
    test('mods', () => {
        expect(
            classNames('header', {
                collapsed: true,
                place: 'gg',
                hovered: false,
            }),
        ).toBe('header collapsed place');
    });
    test('all params', () => {
        expect(
            classNames(
                'sidebar',
                {
                    collapsed: true,
                    resize: false,
                },
                ['red', 'left'],
            ),
        ).toBe('sidebar red left collapsed');
    });
});
