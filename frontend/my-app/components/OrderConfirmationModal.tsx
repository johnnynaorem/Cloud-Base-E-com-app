import React from 'react';
import { type Order } from '../types';
import { XIcon } from './icons';

interface OrderConfirmationModalProps {
    order: Order;
    onClose: () => void;
}

const OrderConfirmationModal: React.FC<OrderConfirmationModalProps> = ({ order, onClose }) => {
    console.log("Order Confirmation Modal:", order);
    return (
        <div className="fixed inset-0 bg-black bg-opacity-60 flex items-center justify-center z-50 animate-fade-in p-4">
            <div className="bg-surface rounded-lg shadow-2xl w-full max-w-lg animate-slide-up relative max-h-[90vh] flex flex-col">
                <div className="p-6 border-b">
                    <h2 className="text-2xl font-bold text-primary">Order Successful!</h2>
                    <p className="text-content-secondary">Thank you for your purchase.</p>
                    <button onClick={onClose} className="absolute top-4 right-4 text-gray-400 hover:text-gray-600">
                        <XIcon className="h-6 w-6" />
                        <span className="sr-only">Close confirmation</span>
                    </button>
                </div>

                <div className="p-6 overflow-y-auto">
                    <div className="space-y-2 mb-4">
                        <p><span className="font-semibold">Order ID:</span> {order.id}</p>
                        <p><span className="font-semibold">Date:</span> {order.date.toLocaleDateString()}</p>
                    </div>

                    <h3 className="font-semibold text-lg mb-2">Order Summary:</h3>
                    <ul className="divide-y divide-gray-200 border-b border-t">
                        {order.items.map(item => (
                            <li key={item.id} className="flex py-3 items-center">
                                <img src={item.product_image} alt={item.product_name} className="h-16 w-16 rounded-md object-cover mr-4" />
                                <div className="flex-1">
                                    <p className="font-medium">{item.product_name}</p>
                                    <p className="text-sm text-content-secondary">Qty: {item.quantity}</p>
                                </div>
                                <p className="font-medium">${(item.product_price * item.quantity).toFixed(2)}</p>
                            </li>
                        ))}
                    </ul>
                </div>

                <div className="p-6 bg-gray-50 rounded-b-lg">
                    <div className="flex justify-between items-center text-xl font-bold">
                        <span>Total</span>
                        <span>${order.totalPrice.toFixed(2)}</span>
                    </div>
                    <button
                        onClick={onClose}
                        className="mt-6 w-full py-2 px-4 bg-secondary text-white font-semibold rounded-lg shadow-md hover:bg-opacity-90 transition-colors"
                    >
                        Continue Shopping
                    </button>
                </div>
            </div>
        </div>
    );
};

export default OrderConfirmationModal;
