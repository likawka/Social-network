import React from "react";
import Anchor from "../Common/Anchor";

const Footer = () => {
    return (
        <footer className="bg-gray-800 p-4">
            <div className="container mx-auto text-center text-white">
                <div className="flex justify-center space-x-6 mb-4">
                    <Anchor to="/" className="text-white">
                    Home
                    </Anchor>
                    <Anchor
                    to="https://github.com/01-edu/public/tree/master/subjects/social-network/audit"
                    className='text-white'
                    external={true}
                    >
                        Audit
                    </Anchor>
                    <Anchor to="http://localhost:8080/swagger/" className='text-white' external={true}>
                        API
                    </Anchor>
                    </div>
                    <div className="text-sm">
                        &copy; {new Date().getFullYear()} Anime Network. All rights reserved.
                    </div>
                </div>
            </footer>
    );
};

export default Footer;