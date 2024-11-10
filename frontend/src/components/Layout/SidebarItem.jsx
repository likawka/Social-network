import React from "react";
import PropTypes from "prop-types";
import { Link } from "react-router-dom";
import DropdownMenu from "../Common/DropDownMenu";

const SidebarItem = ({ icon, label, link, items }) => {
    if (items && items.length > 0) {
        return (
            <DropdownMenu
                label={label}
                items={items}
                icon={icon}
            />
        );
    }

    return (
        <Link to={link} className="flex items-center p-2 hover:bg-gray-800 rounded">
            <i className={`fas fa-${icon} text-xl mr-4`}></i>
            <span className="text-lg">{label}</span>
        </Link>
    );
};

SidebarItem.propTypes = {
    icon: PropTypes.string.isRequired,
    label: PropTypes.string.isRequired,
    link: PropTypes.string,
    items: PropTypes.arrayOf(
        PropTypes.shape({
            label: PropTypes.string.isRequired,
            link: PropTypes.string.isRequired,
        })
    ),
};

export default SidebarItem;
