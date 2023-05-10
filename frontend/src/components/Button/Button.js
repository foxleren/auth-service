import React from 'react';
import {buttonProps} from "./ButtonProps";
import './Button.scss'

function Button({
                    text = 'Button',
                    onClck = () => {

                    },
                    size = buttonProps.size.small,
                    color = buttonProps.color.dark,
                    bgColor= buttonProps.background_color.grey,
                    isActive = false,
                    type = 'button'
                }) {
    return (
        <button className={`button ${size} ${color} ${bgColor} ${isActive ? 'active' : ''}`} onClick={() => onClck()}
        type={type}
        >
            {text}
        </button>
    );
}

export default Button;