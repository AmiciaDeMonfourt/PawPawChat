import { MutableRefObject, useEffect, useRef, useState } from "react";


interface useResizeProps {
    defaultWidth: number;
    minWidth: number;
    maxWidth: number;
}

interface useResizeResult {
    ref: MutableRefObject<any>;
    width: number;
    onMouseDown: (e: any) => void;
}

export function useResize(props : useResizeProps) : useResizeResult {

    const {
        defaultWidth,
        minWidth,
        maxWidth,
    } = props

    const [width, setWidth] = useState(defaultWidth);

    const ref = useRef(null);

    const onMouseDown = () => {
        document.addEventListener('mousemove', onMouseMove);
        document.addEventListener('mouseup', onMouseUp);
        document.body.classList.add("no-select");
    }

    const onMouseMove = (e : any) => {
        const newWidth = e.clientX - ref.current.getBoundingClientRect().left;
        const newCorrectWidth = Math.min(Math.max(minWidth, newWidth),maxWidth);

        setWidth(newCorrectWidth);
    }

    const onMouseUp = () => {
        document.removeEventListener('mousemove', onMouseMove);
        document.removeEventListener('mouseup', onMouseUp);
        document.body.classList.remove("no-select");
    }

    useEffect(() => {
        return () => {
            onMouseUp();
        }
    }, [])

    return {
        ref,
        width,
        onMouseDown
    }
}