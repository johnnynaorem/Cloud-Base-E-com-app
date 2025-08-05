
import React, { useMemo } from 'react';
import { type CartItem, type Product, type OrderItem } from '../types';
import { XIcon, Trash2Icon } from './icons';

interface ShoppingCartProps {
  cartItems: any[];
  products: any[];
  onClose: () => void;
  onUpdateQuantity: (productId: string, quantity: number) => void;
  onRemoveItem: (productId: string) => void;
  onCheckout: () => void;
}

const ShoppingCart: React.FC<ShoppingCartProps> = ({ cartItems, products, onClose, onUpdateQuantity, onRemoveItem, onCheckout }) => {
  const enrichedCartItems = useMemo((): any => {
    return cartItems
      .map(item => {
        const product = products.find(p => p.ID === item.productId);
        if (!product) {
          return null; // Product may have been deleted
        }
        return { ...product, quantity: item.quantity };
      })
      .filter((item): item is OrderItem => item !== null);
  }, [cartItems, products]);

  const totalPrice = useMemo(() => {
    return enrichedCartItems.reduce((total, item) => total + item.product_price * item.quantity, 0);
  }, [enrichedCartItems]);

  console.log("Enriched cart items:", enrichedCartItems);

  return (
    <div className="fixed inset-0 bg-black bg-opacity-60 z-40 animate-fade-in" aria-modal="true" role="dialog">
      <div className="fixed inset-y-0 right-0 w-full max-w-md bg-surface shadow-xl flex flex-col">
        <div className="flex items-center justify-between p-6 border-b border-gray-200">
          <h2 className="text-2xl font-bold text-primary">Shopping Cart</h2>
          <button onClick={onClose} className="text-gray-500 hover:text-gray-800">
            <XIcon className="h-6 w-6" />
            <span className="sr-only">Close cart</span>
          </button>
        </div>

        {enrichedCartItems.length === 0 ? (
          <div className="flex-1 flex flex-col items-center justify-center text-center p-6">
            <p className="text-lg text-content-secondary">Your cart is empty.</p>
            <button onClick={onClose} className="mt-4 py-2 px-4 bg-secondary text-white font-semibold rounded-lg shadow-md hover:bg-opacity-90 transition-colors">
              Continue Shopping
            </button>
          </div>
        ) : (
          <>
            <div className="flex-1 overflow-y-auto p-6">
              <ul className="divide-y divide-gray-200">
                {enrichedCartItems.map(item => (
                  <li key={item.id} className="flex py-4">
                    <div className="h-24 w-24 flex-shrink-0 overflow-hidden rounded-md border border-gray-200">
                      <img src={item.product_image} alt={item.product_name} className="h-full w-full object-cover object-center" />
                    </div>
                    <div className="ml-4 flex flex-1 flex-col">
                      <div>
                        <div className="flex justify-between text-base font-medium text-primary">
                          <h3>{item.product_name}</h3>
                          <p className="ml-4">${(item.product_price * item.quantity).toFixed(2)}</p>
                        </div>
                        <p className="mt-1 text-sm text-content-secondary">${item.product_price.toFixed(2)} each</p>
                      </div>
                      <div className="flex flex-1 items-end justify-between text-sm">
                        <div className="flex items-center border border-gray-300 rounded-md">
                          <button onClick={() => onUpdateQuantity(item.ID, item.quantity - 1)} disabled={item.quantity <= 1} className="px-2 py-1 text-lg font-bold text-gray-600 disabled:opacity-50">-</button>
                          <input
                            type="number"
                            value={item.quantity}
                            onChange={(e) => onUpdateQuantity(item.id, Math.max(1, parseInt(e.target.value) || 1))}
                            className="w-12 text-center border-l border-r border-gray-300"
                          />
                          <button onClick={() => onUpdateQuantity(item.ID, item.quantity + 1)} disabled={item.quantity >= item.inventory} className="px-2 py-1 text-lg font-bold text-gray-600 disabled:opacity-50">+</button>
                        </div>
                        <button onClick={() => onRemoveItem(item.ID)} className="font-medium text-red-600 hover:text-red-800 flex items-center gap-1">
                          <Trash2Icon className="h-4 w-4" /> Remove
                        </button>
                      </div>
                    </div>
                  </li>
                ))}
              </ul>
            </div>

            <div className="border-t border-gray-200 p-6 bg-gray-50">
              <div className="flex justify-between text-lg font-bold text-primary">
                <p>Subtotal</p>
                <p>${totalPrice.toFixed(2)}</p>
              </div>
              <p className="mt-1 text-sm text-content-secondary">Shipping and taxes calculated at checkout.</p>
              <div className="mt-6">
                <button
                  onClick={onCheckout}
                  className="w-full flex items-center justify-center rounded-md border border-transparent bg-accent px-6 py-3 text-base font-medium text-white shadow-sm hover:bg-opacity-90"
                >
                  Place Order
                </button>
              </div>
            </div>
          </>
        )}
      </div>
    </div>
  );
};

export default ShoppingCart;
