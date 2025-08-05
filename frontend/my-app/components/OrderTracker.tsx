import React from 'react';
import { type TrackingStatus } from '../types';
import { TRACKING_STEPS } from '../constants';

interface OrderTrackerProps {
  currentStatus: TrackingStatus;
  onUpdateStatus: (newStatus: TrackingStatus) => void;
}

const OrderTracker: React.FC<OrderTrackerProps> = ({ currentStatus, onUpdateStatus }) => {
  const currentIndex = TRACKING_STEPS.indexOf(currentStatus);

  const handleNext = () => {
    if (currentIndex < TRACKING_STEPS.length - 1) {
      onUpdateStatus(TRACKING_STEPS[currentIndex + 1]);
    }
  };

  const handlePrev = () => {
    if (currentIndex > 0) {
      onUpdateStatus(TRACKING_STEPS[currentIndex - 1]);
    }
  };

  return (
    <div>
      <div className="flex items-center justify-between">
        {TRACKING_STEPS.map((step, index) => {
          const isCompleted = index < currentIndex;
          const isCurrent = index === currentIndex;
          return (
            <React.Fragment key={step}>
              <div className="flex flex-col items-center text-center">
                <div
                  className={`h-8 w-8 rounded-full flex items-center justify-center transition-colors duration-300 ${
                    isCurrent ? 'bg-secondary text-white ring-4 ring-blue-200' : isCompleted ? 'bg-accent text-white' : 'bg-gray-300 text-gray-500'
                  }`}
                >
                  {isCompleted ? '✓' : index + 1}
                </div>
                <p className={`mt-2 text-xs font-semibold ${isCurrent || isCompleted ? 'text-primary' : 'text-content-secondary'}`}>{step}</p>
              </div>
              {index < TRACKING_STEPS.length - 1 && (
                <div className={`flex-1 h-1 mx-2 transition-colors duration-300 ${isCompleted ? 'bg-accent' : 'bg-gray-300'}`}></div>
              )}
            </React.Fragment>
          );
        })}
      </div>
      <div className="flex justify-end gap-3 mt-4">
        <button onClick={handlePrev} disabled={currentIndex === 0} className="py-1 px-3 text-sm font-medium rounded-md bg-gray-200 hover:bg-gray-300 disabled:opacity-50 disabled:cursor-not-allowed">
          Previous Step
        </button>
        <button onClick={handleNext} disabled={currentIndex === TRACKING_STEPS.length - 1} className="py-1 px-3 text-sm font-medium rounded-md bg-secondary text-white hover:bg-opacity-90 disabled:opacity-50 disabled:cursor-not-allowed">
          Next Step
        </button>
      </div>
    </div>
  );
};

export default OrderTracker;
