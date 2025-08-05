import React from 'react';
import { type PaymentStatus } from '../types';
import { ClockIcon, CheckCircleIcon, XCircleIcon } from './icons';

interface PaymentStatusToggleProps {
  status: PaymentStatus;
  onUpdateStatus: (newStatus: PaymentStatus) => void;
}

const statusOptions: {
  value: PaymentStatus;
  label: string;
  icon: React.ElementType;
  classes: string;
}[] = [
  { value: 'Pending', label: 'Pending', icon: ClockIcon, classes: 'bg-yellow-100 text-yellow-800 hover:bg-yellow-200' },
  { value: 'Paid', label: 'Paid', icon: CheckCircleIcon, classes: 'bg-green-100 text-green-800 hover:bg-green-200' },
  { value: 'Failed', label: 'Failed', icon: XCircleIcon, classes: 'bg-red-100 text-red-800 hover:bg-red-200' },
];

const PaymentStatusToggle: React.FC<PaymentStatusToggleProps> = ({ status, onUpdateStatus }) => {
  return (
    <div className="flex items-center gap-2">
      {statusOptions.map(({ value, label, icon: Icon, classes }) => {
        const isActive = status === value;
        return (
          <button
            key={value}
            onClick={() => onUpdateStatus(value)}
            className={`flex items-center gap-2 py-2 px-3 rounded-full text-sm font-semibold transition-all duration-200 ease-in-out transform ${
              isActive
                ? `${classes} ring-2 ring-offset-1 ring-blue-500 scale-105`
                : 'bg-gray-100 text-gray-600 hover:bg-gray-200'
            }`}
          >
            <Icon className="h-5 w-5" />
            {label}
          </button>
        );
      })}
    </div>
  );
};

export default PaymentStatusToggle;
