
import React, { useState, useEffect, useCallback, useMemo } from 'react';
import { type Product, type CartItem, type Order, type Credentials, type PaymentStatus, type User, type TrackingStatus, type OrderItem } from './types';
import { INITIAL_PRODUCTS, ROLES, INITIAL_USERS } from './constants';
import AdminDashboard from './components/AdminDashboard';
import ProductGrid from './components/ProductGrid';
import ShoppingCart from './components/ShoppingCart';
import OrderConfirmationModal from './components/OrderConfirmationModal';
import Login from './components/Login';
import { StoreIcon, UserIcon, ShieldCheckIcon, ShoppingCartIcon, LogoutIcon } from './components/icons';
import axios from 'axios';

export default function App(): React.ReactNode {
  const [users, setUsers] = useState<User[]>(INITIAL_USERS);
  const [currentUser, setCurrentUser] = useState<User | null>(null);
  const [products, setProducts] = useState([]);
  const [cart, setCart] = useState([]);
  const [orders, setOrders] = useState<Order[]>([]);
  const [isProductModalOpen, setIsProductModalOpen] = useState(false);
  const [editingProduct, setEditingProduct] = useState<Product | null>(null);
  const [isCartOpen, setIsCartOpen] = useState(false);
  const [lastOrder, setLastOrder] = useState<Order | null>(null);

  const handleLogin = useCallback((credentials: Credentials) => {
    const foundUser = users.find(
      (user) => user.username === credentials.username && user.password === credentials.password
    );
    if (foundUser) {
      setCurrentUser(foundUser);
    } else {
      const guestUser: User = {
        id: 'guest-' + Date.now(),
        username: credentials.username,
        password: '',
        role: ROLES.CUSTOMER,
        fullName: credentials.username,
        address: 'N/A',
        joinDate: new Date(),
      };
      setCurrentUser(guestUser);
    }
  }, [users]);

  const handleLogout = useCallback(() => {
    // localStorage.removeItem("token")
    setCurrentUser(null);
    setCart([]);
    setLastOrder(null);
    setIsCartOpen(false);
  }, []);

  const handleUpdateUser = useCallback((updatedUser: User) => {
    setUsers(prevUsers => prevUsers.map(user => user.id === updatedUser.id ? updatedUser : user));
    setOrders(prevOrders => prevOrders.map(order =>
      order.userId === updatedUser.id ? { ...order, customerName: updatedUser.fullName } : order
    ));
  }, []);

  const handleOpenProductModal = useCallback((product: Product | null = null) => {
    setEditingProduct(product);
    setIsProductModalOpen(true);
  }, []);

  const handleCloseProductModal = useCallback(() => {
    setIsProductModalOpen(false);
    setEditingProduct(null);
  }, []);

  const handleSaveProduct = useCallback((product: any) => {
    setProducts(prev => {
      const existing = prev.find(p => p.id === product.ID);
      if (existing) {
        return prev.map(p => p.id === product.id ? product : p);
      }
      return [...prev, { ...product, id: Date.now().toString() }];
    });
    handleCloseProductModal();
  }, [handleCloseProductModal]);

  const handleDeleteProduct = useCallback((productId: string) => {
    setProducts(prev => prev.filter(p => p.id !== productId));
    setCart(prevCart => prevCart.filter(item => item.productId !== productId));
  }, []);

  const handleAddToCart = useCallback((product: any) => {
    console.log("Adding to cart:", product);
    console.log("Current cart:", cart);
    if (product.product_stock <= 0) return;

    setCart(prevCart => {
      const existingItem = prevCart.find(item => item.productId === product.ID);
      if (existingItem) {
        if (existingItem.quantity < product.product_stock) {
          return prevCart.map(item =>
            item.productId === product.ID ? { ...item, quantity: item.quantity + 1 } : item
          );
        }
        return prevCart;
      }
      return [...prevCart, { productId: product.ID, quantity: 1 }];
    });
  }, []);

  const handleUpdateCartQuantity = useCallback((productId: string, quantity: number) => {
    setCart(prevCart => {
      const product = products.find(p => p.ID === productId);
      if (!product) { // Product might have been deleted
        return prevCart.filter(item => item.productId !== productId);
      }

      const cappedQuantity = Math.min(quantity, product.product_stock);

      if (cappedQuantity > 0) {
        return prevCart.map(item =>
          item.productId === productId ? { ...item, quantity: cappedQuantity } : item
        );
      }
      return prevCart.filter(item => item.productId !== productId);
    });
  }, [products]);

  const handleRemoveFromCart = useCallback((productId: string) => {
    setCart(prevCart => prevCart.filter(item => item.productId !== productId));
  }, []);

  const placeOrder = async (customerId,
    productId,
    quantity,
    totalAmount) => {
    const token = localStorage.getItem("token")
    const response = await axios.post("http://34.49.189.59/order/create", {
      customerId,
      productId,
      quantity,
      totalAmount,
    }, {
      headers: {
        Authorization: `Bearer ${token}`,
      },
    });
    if (response.status === 201) {
      console.log("Order placed successfully:", response.data);
      return response.data.result;
    } else {
      console.error("Failed to place order:", response);
      throw new Error("Failed to place order");
    }
  }

  const handlePlaceOrder = useCallback(async () => {
    if (!currentUser || cart.length === 0) return;

    const orderItems: any = cart.map(cartItem => {
      const product = products.find(p => p.ID === cartItem.productId);
      if (!product || product.inventory < cartItem.quantity) return null;
      return { ...product, quantity: cartItem.quantity };
    }).filter((item): item is OrderItem => item !== null);

    console.log("Placing order with items:", orderItems);

    if (orderItems.length !== cart.length) {
      alert("Some items in your cart are no longer available or have insufficient stock. Your cart has been updated.");
      const validProductIds = orderItems.map(item => item.id);
      setCart(prevCart => prevCart.filter(item => validProductIds.includes(item.productId)));
      return;
    }

    setProducts(prevProducts => {
      const updatedProducts = [...prevProducts];
      orderItems.forEach(orderItem => {
        const productIndex = updatedProducts.findIndex(p => p.ID === orderItem.id);
        if (productIndex !== -1) {
          updatedProducts[productIndex].product_stock -= orderItem.quantity;
        }
      });
      return updatedProducts;
    });
    console.log(orderItems)

    const orderResponse = await placeOrder("1", orderItems[0].ID.toString(), orderItems[0].quantity, orderItems.reduce((total, item) => total + item.product_price * item.quantity, 0));
    const newOrder: Order = {
      id: orderResponse.orderId,
      userId: currentUser.id,
      customerName: currentUser.fullName,
      items: orderItems,
      totalPrice: orderItems.reduce((total, item) => total + item.product_price * item.quantity, 0),
      date: new Date(),
      paymentStatus: 'Pending',
      trackingStatus: 'Processing',
    };
    setOrders(prevOrders => [newOrder, ...prevOrders]);
    setLastOrder(newOrder);

    setCart([]);
    setIsCartOpen(false);
  }, [cart, currentUser, products]);

  const handleUpdatePaymentStatus = useCallback((orderId: string, status: PaymentStatus) => {
    setOrders(prevOrders =>
      prevOrders.map(order =>
        order.id === orderId ? { ...order, paymentStatus: status } : order
      )
    );
  }, []);

  const handleUpdateTrackingStatus = useCallback((orderId: string, status: TrackingStatus) => {
    setOrders(prevOrders =>
      prevOrders.map(order =>
        order.id === orderId ? { ...order, trackingStatus: status } : order
      )
    );
  }, []);


  const cartItemCount = useMemo(() => {
    return cart.reduce((count, item) => count + item.quantity, 0);
  }, [cart]);

  if (!currentUser) {
    return <Login onLogin={handleLogin} />;
  }

  return (
    <div className="min-h-screen bg-background text-content-primary font-sans">
      <header className="bg-surface shadow-md sticky top-0 z-20">
        <nav className="container mx-auto px-4 sm:px-6 lg:px-8">
          <div className="flex items-center justify-between h-20">
            <div className="flex items-center gap-3">
              <StoreIcon className="h-9 w-9 text-secondary" />
              <h1 className="text-3xl font-bold text-primary">Johnny's Store</h1>
            </div>
            <div className="flex items-center gap-6">
              <div className="flex items-center gap-2 text-sm font-medium">
                {currentUser.role === ROLES.ADMIN ? (
                  <ShieldCheckIcon className="h-5 w-5 text-accent" />
                ) : (
                  <UserIcon className="h-5 w-5 text-secondary" />
                )}
                <span className="text-content-secondary">
                  {currentUser.role === ROLES.ADMIN ? 'Admin Mode' : `Welcome, ${currentUser.username}!`}
                </span>
              </div>

              {currentUser.role === ROLES.CUSTOMER && (
                <button onClick={() => setIsCartOpen(true)} className="relative text-gray-500 hover:text-secondary transition-colors" aria-label={`Open shopping cart with ${cartItemCount} items`}>
                  <ShoppingCartIcon className="h-7 w-7" />
                  {cartItemCount > 0 && (
                    <span className="absolute -top-2 -right-2 flex items-center justify-center h-5 w-5 bg-accent text-white text-xs font-bold rounded-full">
                      {cartItemCount}
                    </span>
                  )}
                </button>
              )}
              <button onClick={handleLogout} className="flex items-center gap-2 text-gray-500 hover:text-red-600 transition-colors" aria-label="Logout">
                <LogoutIcon className="h-6 w-6" />
                <span className="text-sm font-medium hidden sm:block">Logout</span>
              </button>
            </div>
          </div>
        </nav>
      </header>

      <main className="container mx-auto px-4 sm:px-6 lg:px-8 py-10">
        {currentUser.role === ROLES.CUSTOMER && <ProductGrid setProducts={setProducts} products={products} onAddToCart={handleAddToCart} />}
        {currentUser.role === ROLES.ADMIN && (
          <AdminDashboard
            setProducts={setProducts}
            products={products}
            orders={orders}
            users={users}
            onAddProduct={() => handleOpenProductModal(null)}
            onEditProduct={handleOpenProductModal}
            onDeleteProduct={handleDeleteProduct}
            onUpdatePaymentStatus={handleUpdatePaymentStatus}
            onUpdateTrackingStatus={handleUpdateTrackingStatus}
            onUpdateUser={handleUpdateUser}
            isModalOpen={isProductModalOpen}
            editingProduct={editingProduct}
            onCloseModal={handleCloseProductModal}
            onSaveProduct={handleSaveProduct}
          />
        )}
      </main>

      {isCartOpen && (
        <ShoppingCart
          cartItems={cart}
          products={products}
          onClose={() => setIsCartOpen(false)}
          onUpdateQuantity={handleUpdateCartQuantity}
          onRemoveItem={handleRemoveFromCart}
          onCheckout={handlePlaceOrder}
        />
      )}

      {lastOrder && (
        <OrderConfirmationModal
          order={lastOrder}
          onClose={() => setLastOrder(null)}
        />
      )}

      <footer className="w-full text-center py-6 mt-10 border-t border-gray-200">
        <p className="text-sm text-content-secondary">Copyright @Johnny Naorem 2025. </p>
      </footer>
    </div>
  );
}
