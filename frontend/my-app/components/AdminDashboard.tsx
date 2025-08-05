
import React, { useState, useMemo, useEffect } from 'react';
import { type Product, type Order, type PaymentStatus, type User, type TrackingStatus } from '../types';
import ProductForm from './ProductForm';
import { PlusCircleIcon, EditIcon, Trash2Icon, StoreIcon, ClipboardListIcon, UsersIcon } from './icons';
import PaymentStatusToggle from './PaymentStatusToggle';
import OrderTracker from './OrderTracker';
import CustomerManagement from './CustomerManagement';
import axios from 'axios';

interface AdminDashboardProps {
  setProducts;
  products: Product[];
  orders: Order[];
  users: User[];
  onAddProduct: () => void;
  onEditProduct: (product: Product) => void;
  onDeleteProduct: (id: string) => void;
  onUpdatePaymentStatus: (orderId: string, status: PaymentStatus) => void;
  onUpdateTrackingStatus: (orderId: string, status: TrackingStatus) => void;
  onUpdateUser: (user: User) => void;
  isModalOpen: boolean;
  editingProduct: Product | null;
  onCloseModal: () => void;
  onSaveProduct: (product: Product) => void;
}

const AdminDashboard: React.FC<AdminDashboardProps> = (props) => {
  const {
    setProducts, products, orders, users, onAddProduct, onEditProduct, onDeleteProduct,
    onUpdatePaymentStatus, onUpdateTrackingStatus, onUpdateUser,
    isModalOpen, editingProduct, onCloseModal, onSaveProduct,
  } = props;
  const [activeView, setActiveView] = useState<'products' | 'sales' | 'customers'>('sales');
  const [selectedOrder, setSelectedOrder] = useState([]);
  const [allOrders, setAllOrders] = useState([]);
  const [orderProduct, setOrderProduct] = useState({});

  const getOrderProduct = async (id: string) => {
    const token = localStorage.getItem("token")
    const response = await axios.get(`http://34.49.189.59/product/${id}`, {
      headers: {
        Authorization: `Bearer ${token}`,
      },
    });
    if (response.status === 200) {
      setOrderProduct(response.data);
    } else {
      alert("Error while fetching products....")
    }
  }

  const fetchProducts = async () => {
    const token = localStorage.getItem("token")
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

  useEffect(() => {
    if (selectedOrder) {
      // Find the latest version of the selected order from the master `orders` list
      const freshOrderData = orders.find(o => o.id === selectedOrder.id);
      setSelectedOrder(freshOrderData || null); // Update state, or deselect if it's gone
    } else if (orders.length > 0 && !selectedOrder) {
      // If no order is selected, default to the first one
      setSelectedOrder(orders[0]);
    }
  }, [orders]); // Rerun this effect whenever the master `orders` array changes

  const getAllOrders = async () => {
    const token = localStorage.getItem("token")
    const response = await axios.get("http://34.49.189.59/order/getorders", {
      headers: {
        Authorization: `Bearer ${token}`
      }
    });
    if (response.status === 200) {
      return setAllOrders(response.data.orders);
    } else if (response.status === 404) {
      return setAllOrders([]);
    } else {
      alert("Failed to fetch orders");
    }
  }

  useEffect(() => {
    getAllOrders();
  }, [])

  const selectedOrderCustomer = useMemo(() => {
    if (!selectedOrder) return null;
    return users.find(u => u.id === selectedOrder.userId) || null;
  }, [selectedOrder, users]);

  console.log("Selected Order:", selectedOrder);

  return (
    <div className="animate-fade-in">
      <div className="mb-8">
        <h2 className="text-3xl font-bold text-primary">Admin Dashboard</h2>
        <p className="text-lg text-content-secondary mt-1">Manage your store's products, sales, and customers.</p>

        <div className="mt-6 border-b border-gray-200">
          <nav className="-mb-px flex space-x-8" aria-label="Tabs">
            <TabButton name="Products" icon={<StoreIcon className="h-5 w-5" />} activeView={activeView} setActiveView={() => setActiveView('products')} />
            <TabButton name="Sales" icon={<ClipboardListIcon className="h-5 w-5" />} activeView={activeView} setActiveView={() => setActiveView('sales')} />
            <TabButton name="Customers" icon={<UsersIcon className="h-5 w-5" />} activeView={activeView} setActiveView={() => setActiveView('customers')} />
          </nav>
        </div>
      </div>

      {activeView === 'products' && <ProductManagementView {...{ products, onAddProduct, onEditProduct, onDeleteProduct }} />}

      {activeView === 'sales' && (
        <div className="grid grid-cols-1 lg:grid-cols-3 gap-8">
          <div className="lg:col-span-1 bg-surface rounded-lg shadow-md overflow-hidden h-[70vh]">
            <h3 className="p-4 text-lg font-semibold border-b">All Orders</h3>
            <ul className="overflow-y-auto h-[calc(70vh-57px)]">
              {allOrders?.map(order => (
                <li key={order.orderId} onClick={() => { setSelectedOrder(order); getOrderProduct(order.productId) }} className={`p-4 border-b cursor-pointer hover:bg-gray-100 ${selectedOrder?.orderId === order.orderId ? 'bg-blue-50' : ''}`}>
                  <p className="font-semibold text-primary">Johnny</p>
                  <p className="text-sm text-content-secondary">{order.orderId}</p>
                  <p className="text-sm text-content-secondary">{order.orderDate}</p>
                </li>
              ))}
              {allOrders.length === 0 && <p className="p-4 text-content-secondary">No orders yet.</p>}
            </ul>
          </div>
          <div className="lg:col-span-2 bg-surface rounded-lg shadow-md p-6 h-[70vh] overflow-y-auto">
            {selectedOrder ? (
              <div>
                <h3 className="text-xl font-bold mb-4">Order Details</h3>
                <div className="space-y-6">
                  {selectedOrderCustomer && (
                    <div>
                      <h4 className="font-semibold text-content-primary mb-2">Customer Information</h4>
                      <div className="text-sm bg-gray-50 p-3 rounded-md">
                        <p><strong>Name:</strong> {selectedOrderCustomer.fullName}</p>
                        <p><strong>Address:</strong> {selectedOrderCustomer.address}</p>
                        <p><strong>Username:</strong> {selectedOrderCustomer.username}</p>
                      </div>
                    </div>
                  )}

                  <div>
                    <h4 className="font-semibold text-content-primary mb-2">Payment Status</h4>
                    <PaymentStatusToggle
                      status={selectedOrder.paymentStatus}
                      onUpdateStatus={(status) => onUpdatePaymentStatus(selectedOrder.id, status)}
                    />
                  </div>

                  <div>
                    <h4 className="font-semibold text-content-primary mb-2">Order Tracking</h4>
                    <OrderTracker
                      currentStatus={selectedOrder.trackingStatus}
                      onUpdateStatus={(status) => onUpdateTrackingStatus(selectedOrder.id, status)}
                    />
                  </div>

                  <div>
                    <h4 className="font-semibold text-content-primary mb-2">Items Ordered</h4>
                    <ul className="divide-y divide-gray-200 border-t border-b">
                      {/* {selectedOrder.map(item => ( */}
                      <li key={selectedOrder.ID} className="flex py-3 items-center">
                        <img src={orderProduct?.product_image} alt={orderProduct?.product_name} className="h-12 w-12 rounded-md object-cover mr-4" />
                        <div className="flex-1">
                          <p className="font-medium">{orderProduct?.product_name}</p>
                          <p className="text-sm text-content-secondary">Qty: {selectedOrder.quantity}</p>
                        </div>
                        <p className="font-medium">${(selectedOrder.totalAmount) / selectedOrder.quantity}</p>
                      </li>
                      {/* ))} */}
                    </ul>
                    <p className="text-right font-bold text-lg mt-2">Total: ${selectedOrder.totalAmount}</p>
                  </div>
                </div>
              </div>
            ) : <p className="text-center text-content-secondary pt-10">Select an order to view details.</p>}
          </div>
        </div>
      )}

      {activeView === 'customers' && (
        <CustomerManagement users={users.filter(u => u.role === 'customer')} onUpdateUser={onUpdateUser} />
      )}

      {isModalOpen && <ProductForm product={editingProduct} onSave={onSaveProduct} onClose={onCloseModal} />}
    </div>
  );
};

const TabButton = ({ name, icon, activeView, setActiveView }: { name: string, icon: React.ReactElement, activeView: string, setActiveView: () => void }) => (
  <button
    onClick={setActiveView}
    className={`whitespace-nowrap py-4 px-1 border-b-2 font-medium text-sm flex items-center gap-2 ${activeView === name.toLowerCase()
      ? 'border-secondary text-secondary'
      : 'border-transparent text-gray-500 hover:text-gray-700 hover:border-gray-300'
      }`}
  >
    {icon}
    {name}
  </button>
);

const ProductManagementView: React.FC<Pick<AdminDashboardProps, 'products' | 'onAddProduct' | 'onEditProduct' | 'onDeleteProduct'>> = ({ products, onAddProduct, onEditProduct, onDeleteProduct }) => (
  <div>
    <div className="flex justify-end items-center mb-4">
      <button onClick={onAddProduct} className="flex items-center gap-2 py-2 px-4 bg-secondary text-white font-semibold rounded-lg shadow-md hover:bg-opacity-90 transition-colors">
        <PlusCircleIcon className="h-5 w-5" />
        Add Product
      </button>
    </div>
    <div className="bg-surface rounded-lg shadow-md overflow-hidden">
      <div className="overflow-x-auto">
        <table className="min-w-full divide-y divide-gray-200">
          <thead className="bg-gray-50">
            <tr>
              <th scope="col" className="px-6 py-3 text-left text-xs font-medium text-content-secondary uppercase tracking-wider">Product</th>
              <th scope="col" className="px-6 py-3 text-left text-xs font-medium text-content-secondary uppercase tracking-wider">Price</th>
              <th scope="col" className="px-6 py-3 text-left text-xs font-medium text-content-secondary uppercase tracking-wider">Stock</th>
              <th scope="col" className="px-6 py-3 text-left text-xs font-medium text-content-secondary uppercase tracking-wider">ID</th>
              <th scope="col" className="relative px-6 py-3"><span className="sr-only">Actions</span></th>
            </tr>
          </thead>
          <tbody className="bg-white divide-y divide-gray-200">
            {products.map((product) => (
              <tr key={product.id}>
                <td className="px-6 py-4 whitespace-nowrap">
                  <div className="flex items-center">
                    <div className="flex-shrink-0 h-10 w-10">
                      <img className="h-10 w-10 rounded-md object-cover" src={product.product_image} alt={product.product_name} />
                    </div>
                    <div className="ml-4">
                      <div className="text-sm font-medium text-primary">{product.product_name}</div>
                    </div>
                  </div>
                </td>
                <td className="px-6 py-4 whitespace-nowrap text-sm text-content-primary">${product.product_price.toFixed(2)}</td>
                <td className="px-6 py-4 whitespace-nowrap text-sm text-content-primary font-medium">{product.product_stock}</td>
                <td className="px-6 py-4 whitespace-nowrap text-sm text-content-secondary">{product.ID}</td>
                <td className="px-6 py-4 whitespace-nowrap text-right text-sm font-medium space-x-2">
                  <button onClick={() => onEditProduct(product)} className="text-secondary hover:text-opacity-80 p-1"><EditIcon className="h-5 w-5" /></button>
                  <button onClick={() => onDeleteProduct(product.ID)} className="text-red-600 hover:text-red-800 p-1"><Trash2Icon className="h-5 w-5" /></button>
                </td>
              </tr>
            ))}
          </tbody>
        </table>
      </div>
    </div>
  </div>
);

export default AdminDashboard;