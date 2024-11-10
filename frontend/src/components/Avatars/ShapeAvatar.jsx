import React from 'react';
import PropTypes from 'prop-types';

const ShapeAvatar = ({ shape, color, isSelected, onClick }) => {
  const renderShape = () => {
    switch (shape) {
      case 'circle':
        return <div className={`w-10 h-10 rounded-full ${color}`}></div>;
      case 'square':
        return <div className={`w-10 h-10 ${color}`}></div>;
      case 'triangle':
        return (
          <div className="w-10 h-10 flex items-center justify-center">
            <div className="w-0 h-0 border-l-[20px] border-r-[20px] border-b-[34px] border-b-blue-500 border-l-transparent border-r-transparent"></div>
          </div>
        );
      case 'hexagon':
        return (
          <div className="w-10 h-10 flex items-center justify-center">
            <div className="relative w-10 h-10">
              <div className="absolute top-0 left-0 w-0 h-0 border-l-[18px] border-r-[18px] border-b-[10px] bg-purple-500 border-b-transparent border-l-transparent border-r-transparent"></div>
              <div className={`w-[36px] h-[20px] ${color}`}></div>
              <div className="absolute bottom-0 left-0 w-0 h-0 border-l-[18px] border-r-[18px] border-t-[10px] bg-purple-500 border-t-transparent border-l-transparent border-r-transparent"></div>
            </div>
          </div>
        );
      case 'pentagon':
        return (
          <div className="w-10 h-10 flex items-center justify-center">
            <div
              className="w-0 h-0 border-l-[15px] border-r-[15px] border-b-[25px] border-transparent border-b-orange-500"
              style={{ transform: 'rotate(-35deg)' }}
            ></div>
            <div
              className="w-[20px] h-[10px] bg-orange-500"
              style={{ transform: 'translateY(-5px)' }}
            ></div>
          </div>
        );
      default:
        return null;
    }
  };

  return (
    <div
      onClick={onClick}
      className={`cursor-pointer p-2 ${isSelected ? 'ring-4 ring-blue-500' : ''}`}
    >
      {renderShape()}
    </div>
  );
};

ShapeAvatar.propTypes = {
  shape: PropTypes.string.isRequired,
  color: PropTypes.string.isRequired,
  isSelected: PropTypes.bool,
  onClick: PropTypes.func.isRequired,
};

export default ShapeAvatar;
