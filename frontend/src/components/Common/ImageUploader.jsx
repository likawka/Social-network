import React, { useState } from "react";
import PropTypes from "prop-types";

const ImageUploader = ({ image, setImage, label = 'Upload Image', maxSizeMB, width, height}) => {
    const [error, setError] = useState('');
    const [imagePreview, setImagePreview] = useState('');

    // Convert
    const toBase64 = (file) => {
        return new Promise((resolve, reject) => {
            const reader = new FileReader();
            reader.readAsDataURL(file);
            reader.onload = () => resolve(reader.result); // Base64 encoded image
            reader.onerror = (error) => reject(error);
        });
    };

    // Validate image
    const validateImage = (file) => {
        // File size
        if (file.size > maxSizeMB * 1024 * 1024) {
            return `File size should be less than ${maxSizeMB}MB`;
        }

        // Image dimensions
        if (width && height) {
            const img = new Image();
            img.src = URL.createObjectURL(file);
            return new Promise((resolve) => {
                img.onload = () => {
                    if (img.width !== width || img.height !== height) {
                        resolve(`Image dimensions should be ${width}x${height} pixels`);
                    } else {
                        resolve('');
                    }
                };
            });
        }
        return Promise.resolve('');
    };

    // Handle image change
    const handleChangeImage = async (e) => {
        const file = e.target.files[0];
        if (file) {
            const validationError = await validateImage(file);
            if (validationError) {
                setError(validationError);
                setImage(null);
                setImagePreview('');
            } else {
                setError('');
                const base64Img = await toBase64(file);
                setImage(base64Img);
                setImagePreview(URL.createObjectURL(file));
            }
        }
    };
    
    return (
        <div className="mb-4">
          <label className="block text-sm font-medium text-gray-700">{label}</label>
          <input
            type="file"
            accept="image/*"
            onChange={handleChangeImage}
            className="mt-1 block w-full text-sm text-grey-500 file:mr-4 file:py-2 file:px-4 file:rounded file:border-0 file:text-sm
            file:font-semibold file:bg-blue-50 file:text-blue-700 hover:file:bg-blue-700 hover:file:text-white"
          />
          {error && <p className="text-red-500 text-xs mt-2">{error}</p>}
          {/* Show a preview of the image */}
          {imagePreview && <img src={imagePreview} alt="Preview" className="mt-2 max-w-xs" />}
        </div>
      );
};

ImageUploader.propTypes = {
    image: PropTypes.string, // Base64 string
    setImage: PropTypes.func.isRequired,
    label: PropTypes.string,
    maxSizeMB: PropTypes.number.isRequired,
    width: PropTypes.number,
    height: PropTypes.number,
  };

export default ImageUploader;