import React, { useState, useEffect } from 'react';
import { type Product } from '../types';
import ProductCard from './ProductCard';
import axios from 'axios';

interface ProductGridProps {
  products: [];
  setProducts;
  onAddToCart: (product: Product) => void;
}

const ProductGrid: React.FC<ProductGridProps> = ({ setProducts, products, onAddToCart }) => {
  const fetchProducts = async () => {
    const token = localStorage.getItem("token")
    console.log(token)
    const response = await axios.get("http://34.49.189.59/product/getproducts", {
      headers: {
        Authorization: `Bearer ${token}`,
      },
    });
    if (response.status === 200) {
      setProducts(response.data);
    } else {
      alert("Error while fetching products....")
    }
  }
  useEffect(() => {
    fetchProducts()
  }, [])
  return (
    <div className="animate-fade-in">
      <h2 className="text-3xl font-bold text-primary mb-2">Our Products</h2>
      <p className="text-lg text-content-secondary mb-8">Browse our collection of curated items.</p>
      <div className="grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-3 xl:grid-cols-4 gap-6">
        {products.map((product) => (
          <ProductCard key={product.id} product={product} onAddToCart={onAddToCart} />
        ))}
      </div>
    </div>
  );
};

export default ProductGrid;
