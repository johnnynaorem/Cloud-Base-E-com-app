import React from 'react';
import { type Product } from '../types';
import { ShoppingCartIcon } from './icons';

interface ProductCardProps {
  product: Product;
  onAddToCart: (product: Product) => void;
}

const ProductCard: React.FC<ProductCardProps> = ({ product, onAddToCart }) => {
  const isSoldOut = product.product_stock <= 0;

  return (
    <div className="group bg-surface rounded-lg shadow-md overflow-hidden flex flex-col justify-between transform transition-all duration-300 hover:shadow-xl hover:-translate-y-1">
      <div>
        <div className="relative w-full h-48 sm:h-56 bg-gray-200 overflow-hidden">
          <img
            src={product.product_image}
            alt={product.product_name}
            className={`w-full h-full object-cover transition-transform duration-300 group-hover:scale-105 ${isSoldOut ? 'grayscale' : ''}`}
          />
          {isSoldOut && (
            <div className="absolute top-2 right-2 bg-red-600 text-white text-xs font-bold px-2 py-1 rounded-full">
              SOLD OUT
            </div>
          )}
        </div>
        <div className="p-5">
          <h3 className="text-lg font-bold text-primary truncate">{product.product_name}</h3>
          <p className="text-sm text-content-secondary mt-1 h-10 overflow-hidden">{product.product_desc}</p>
          <div className="mt-4 text-xl font-extrabold text-primary">
            ${product.product_price.toFixed(2)}
          </div>
        </div>
      </div>
      <div className="p-5 pt-0">
        <button
          onClick={() => onAddToCart(product)}
          disabled={isSoldOut}
          className="w-full flex items-center justify-center gap-2 py-2 px-4 border border-transparent rounded-md shadow-sm text-sm font-medium text-white bg-secondary hover:bg-opacity-90 disabled:bg-gray-400 disabled:cursor-not-allowed transition-colors"
        >
          <ShoppingCartIcon className="h-5 w-5" />
          Add to Cart
        </button>
      </div>
    </div>
  );
};

export default ProductCard;
