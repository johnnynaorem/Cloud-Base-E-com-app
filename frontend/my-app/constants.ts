import { type Product, type User, type Role } from './types';

export const ROLES = {
  CUSTOMER: 'customer',
  ADMIN: 'admin',
} as const;

export const TRACKING_STEPS = ['Processing', 'Shipped', 'In Transit', 'Delivered'] as const;

export const INITIAL_PRODUCTS: Product[] = [
  {
    id: '1',
    name: 'Ergonomic Desk Chair',
    description: 'A comfortable and supportive chair designed for long hours of work. Features adjustable lumbar support and breathable mesh back.',
    price: 299.99,
    imageUrl: 'https://placehold.co/600x450/e2e8f0/475569?text=Desk+Chair',
    imagePrompt: 'A high-quality ergonomic office chair with a sleek black design, studio lighting, on a plain white background.',
    inventory: 10,
  },
  {
    id: '2',
    name: 'Wireless Mechanical Keyboard',
    description: 'A compact 75% layout keyboard with tactile switches, RGB backlighting, and bluetooth connectivity.',
    price: 129.99,
    imageUrl: 'https://placehold.co/600x450/e2e8f0/475569?text=Keyboard',
    imagePrompt: 'A stylish wireless mechanical keyboard with RGB backlighting, on a dark wooden desk next to a mouse.',
    inventory: 25,
  },
  {
    id: '3',
    name: '4K Ultra-Wide Monitor',
    description: 'A 34-inch curved monitor with stunning 4K resolution, perfect for immersive gaming and productivity.',
    price: 799.00,
    imageUrl: 'https://placehold.co/600x450/e2e8f0/475569?text=Monitor',
    imagePrompt: 'A 34-inch curved 4K ultra-wide monitor displaying a beautiful mountain landscape, on a clean white desk.',
    inventory: 8,
  },
    {
    id: '4',
    name: 'Noise-Cancelling Headphones',
    description: 'Immerse yourself in sound with these premium over-ear headphones featuring adaptive noise cancellation and 30-hour battery life.',
    price: 349.00,
    imageUrl: 'https://placehold.co/600x450/e2e8f0/475569?text=Headphones',
    imagePrompt: 'Sleek, silver noise-cancelling headphones resting on a minimalist marble surface, soft ambient light.',
    inventory: 0,
  },
];


export const INITIAL_USERS: User[] = [
    {
        id: 'user-admin',
        username: 'admin',
        password: 'admin',
        role: ROLES.ADMIN,
        fullName: 'Store Administrator',
        address: '123 Admin Way, Management City, 10101',
        joinDate: new Date('2023-01-01'),
    },
    {
        id: 'user-1',
        username: 'johnny_nao',
        password: 'password123',
        role: ROLES.CUSTOMER,
        fullName: 'Johnny Naorem',
        address: '456 Customer Ave, Shopper Town, 20202',
        joinDate: new Date('2025-08-5'),
    }
];
