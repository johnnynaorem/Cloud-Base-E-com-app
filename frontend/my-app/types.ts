
import { ROLES, TRACKING_STEPS } from './constants';

export interface Product {
  id: string;
  name: string;
  description: string;
  price: number;
  imageUrl: string;
  imagePrompt: string;
  inventory: number;
}

export type Role = typeof ROLES[keyof typeof ROLES];

/**
 * Represents an item in the live shopping cart. Only stores reference ID and quantity.
 */
export interface CartItem {
  productId: string;
  quantity: number;
}

/**
 * Represents a snapshot of a product at the time of purchase. Stored inside an Order.
 */
export interface OrderItem extends Product {
  quantity: number;
}

export type PaymentStatus = 'Pending' | 'Paid' | 'Failed';
export type TrackingStatus = (typeof TRACKING_STEPS)[number];

export interface Order {
  id: string;
  userId: string;
  customerName: string; // denormalized for easy display
  items: OrderItem[];
  totalPrice: number;
  date: Date;
  paymentStatus: PaymentStatus;
  trackingStatus: TrackingStatus;
}

export interface Credentials {
  username: string;
  password: string;
}

export interface User extends Credentials {
  id: string;
  role: Role;
  fullName: string;
  address: string;
  joinDate: Date;
}
