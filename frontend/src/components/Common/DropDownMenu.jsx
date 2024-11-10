import React, { useState, useRef, useEffect } from 'react';
import PropTypes from 'prop-types';
import { Link } from 'react-router-dom';
import { createPortal } from 'react-dom';

const DropdownMenu = ({ label, items, icon }) => {
    const [isOpen, setIsOpen] = useState(false);
    const [dropdownPosition, setDropdownPosition] = useState({ top: 0, left: 0 });
    const menuRef = useRef(null);

    const toggleMenu = () => {
        if (menuRef.current) {
            const rect = menuRef.current.getBoundingClientRect();
            setDropdownPosition({
                top: rect.top + window.scrollY,
                left: rect.right + window.scrollX,
            });
            setIsOpen(!isOpen);
        }
    };

    const closeDropdown = (event) => {
        if (menuRef.current && !menuRef.current.contains(event.target)) {
            setIsOpen(false);
        }
    };

    useEffect(() => {
        if (isOpen) {
            document.addEventListener('mousedown', closeDropdown);
        } else {
            document.removeEventListener('mousedown', closeDropdown);
        }

        return () => {
            document.removeEventListener('mousedown', closeDropdown);
        };
    }, [isOpen]);

    return (
        <div ref={menuRef} className="relative">
            <div className="flex items-center p-2 hover:bg-gray-800 rounded cursor-pointer" onClick={toggleMenu}>
                {icon && <i className={`fas fa-${icon} text-xl mr-4`}></i>}
                <span className="text-lg">{label}</span>
            </div>
            {isOpen && createPortal(
                <div
                    className="absolute bg-gray-800 text-white rounded shadow-lg z-50"
                    style={{
                        top: dropdownPosition.top,
                        left: dropdownPosition.left,
                        minWidth: '12rem',
                    }}
                >
                    {items.map((item, index) => (
                        <Link
                            key={index}
                            to={item.link}
                            className="block px-4 py-2 text-white hover:bg-gray-700"
                            onClick={() => setIsOpen(false)}  // Close the dropdown on click
                        >
                            {item.label}
                        </Link>
                    ))}
                </div>,
                document.body
            )}
        </div>
    );
};

DropdownMenu.propTypes = {
    label: PropTypes.string.isRequired,
    items: PropTypes.arrayOf(
        PropTypes.shape({
            label: PropTypes.string.isRequired,
            link: PropTypes.string.isRequired,
        })
    ).isRequired,
    icon: PropTypes.string,  // Optional icon prop
};

export default DropdownMenu;
