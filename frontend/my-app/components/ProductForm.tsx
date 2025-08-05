import React, { useState, useEffect } from 'react';
import { type Product } from '../types';
import { generateProductImage } from '../services/geminiService';
import LoadingSpinner from './LoadingSpinner';
import { XIcon, ImageIcon } from './icons';
import axios from 'axios';

interface ProductFormProps {
  product: Product | null;
  onSave: (product: Product) => void;
  onClose: () => void;
}

const ProductForm: React.FC<ProductFormProps> = ({ product, onSave, onClose }) => {
  console.log(product)
  const [formData, setFormData] = useState({});
  const [isGenerating, setIsGenerating] = useState(false);
  const [error, setError] = useState('');

  useEffect(() => {
    if (product) {
      setFormData({ name: product.product_name, description: product.product_desc, price: product.product_price, imageUrl: product.product_image, inventory: product.product_stock });
    } else {
      setFormData({ name: '', description: '', price: 0, imageUrl: '', imagePrompt: '', inventory: 0 });
    }
  }, [product]);

  const handleChange = async (e: React.ChangeEvent<HTMLInputElement | HTMLTextAreaElement>) => {
    const { name, value } = e.target;
    let parsedValue: string | number = value;
    if (name === 'price' || name === 'inventory') {
      parsedValue = name === 'price' ? parseFloat(value) || 0 : parseInt(value, 10) || 0;
      if (parsedValue < 0) parsedValue = 0;
    }
    setFormData(prev => ({ ...prev, [name]: parsedValue }));
  }

  const handleGenerateImage = async () => {
    if (!formData.imagePrompt) {
      setError('Please enter a prompt for the image.');
      return;
    }
    setIsGenerating(true);
    setError('');
    try {
      const imageUrl = await generateProductImage(formData.imagePrompt);
      setFormData(prev => ({ ...prev, imageUrl }));
    } catch (e) {
      setError(e instanceof Error ? e.message : 'An unknown error occurred.');
    } finally {
      setIsGenerating(false);
    }
  };

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault();
    if (!formData.name || formData.price <= 0 || !formData.imageUrl) {
      setError('Please fill out the name, set a valid price, and generate an image.');
      return;
    }
    // onSave({ ...formData, id: product?.id || '' });
    const updateResponse = await axios.put(`http://localhost:8082/product/${product?.ID}`, {
      ProductName: formData.name,
      ProductDesc: formData.description,
      ProductPrice: formData.price,
      ProductStock: formData.inventory,
      ProductCategory: formData.category
    });

    if (updateResponse.status !== 200) {
      alert("Update Competed")
      setFormData({ name: updateResponse.data.product_name, description: updateResponse.data.product_desc, price: updateResponse.data.product_price, imageUrl: updateResponse.data.product_image, inventory: updateResponse.data.product_stock });
    }
  };

  return (
    <div className="fixed inset-0 bg-black bg-opacity-50 flex items-center justify-center z-30 animate-fade-in p-4">
      <div className="bg-surface rounded-lg shadow-2xl w-full max-w-2xl animate-slide-up relative max-h-[90vh] overflow-y-auto">
        <button onClick={onClose} className="absolute top-4 right-4 text-gray-400 hover:text-gray-600">
          <XIcon className="h-6 w-6" />
        </button>
        <form onSubmit={handleSubmit} className="p-8">
          <h2 className="text-2xl font-bold text-primary mb-6">{product ? 'Edit Product' : 'Add New Product'}</h2>

          <div className="grid grid-cols-1 md:grid-cols-2 gap-6">
            <div className="md:col-span-1 space-y-4">
              <div>
                <img src="" alt="" />
              </div>
              <div className="aspect-w-4 aspect-h-3 bg-gray-100 rounded-md flex items-center justify-center overflow-hidden">
                {isGenerating ? <LoadingSpinner text="Creating image..." /> :
                  formData.imageUrl ? <img src={formData.imageUrl} alt="Generated product" className="w-full h-full object-cover" /> : <span className="text-sm text-content-secondary">Image preview</span>
                }
              </div>
            </div>

            <div className="md:col-span-1 space-y-4">
              <div>
                <label htmlFor="name" className="block text-sm font-medium text-content-primary">Product Name</label>
                <input type="text" id="name" name="name" value={formData.name} onChange={handleChange} className="mt-1 w-full p-2 border border-gray-300 rounded-md shadow-sm focus:ring-secondary focus:border-secondary" required />
              </div>
              <div className="grid grid-cols-2 gap-4">
                <div>
                  <label htmlFor="price" className="block text-sm font-medium text-content-primary">Price</label>
                  <input type="number" id="price" name="price" value={formData.price} onChange={handleChange} min="0" step="0.01" className="mt-1 w-full p-2 border border-gray-300 rounded-md shadow-sm focus:ring-secondary focus:border-secondary" required />
                </div>
                <div>
                  <label htmlFor="inventory" className="block text-sm font-medium text-content-primary">Inventory</label>
                  <input type="number" id="inventory" name="inventory" value={formData.inventory} onChange={handleChange} min="0" step="1" className="mt-1 w-full p-2 border border-gray-300 rounded-md shadow-sm focus:ring-secondary focus:border-secondary" required />
                </div>
              </div>
              <div>
                <label htmlFor="description" className="block text-sm font-medium text-content-primary">Description</label>
                <textarea id="description" name="description" value={formData.description} onChange={handleChange} rows={5} className="mt-1 w-full p-2 border border-gray-300 rounded-md shadow-sm focus:ring-secondary focus:border-secondary"></textarea>
              </div>
            </div>
          </div>

          {error && <p className="text-red-600 text-sm mt-4 text-center">{error}</p>}

          <div className="mt-8 flex justify-end gap-4">
            <button type="button" onClick={onClose} className="py-2 px-4 border border-gray-300 rounded-md shadow-sm text-sm font-medium text-gray-700 bg-white hover:bg-gray-50">Cancel</button>
            <button type="submit" className="py-2 px-4 border border-transparent rounded-md shadow-sm text-sm font-medium text-white bg-accent hover:bg-opacity-90">Save Product</button>
          </div>
        </form>
      </div>
    </div>
  );
};

export default ProductForm;
