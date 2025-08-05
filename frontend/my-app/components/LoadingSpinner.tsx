import React from 'react';

interface LoadingSpinnerProps {
  text?: string;
}

const LoadingSpinner: React.FC<LoadingSpinnerProps> = ({ text = "Loading..." }) => {
  return (
    <div className="flex flex-col items-center justify-center gap-4 animate-fade-in p-4">
      <div className="w-12 h-12 border-4 border-dashed rounded-full animate-spin border-secondary"></div>
      <p className="text-md font-semibold text-content-secondary">{text}</p>
    </div>
  );
};

export default LoadingSpinner;
