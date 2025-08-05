import React, { useState, useEffect } from 'react';
import { type User } from '../types';
import { EditIcon, XIcon } from './icons';

interface CustomerManagementProps {
  users: User[];
  onUpdateUser: (user: User) => void;
}

const CustomerManagement: React.FC<CustomerManagementProps> = ({ users, onUpdateUser }) => {
  const [isModalOpen, setIsModalOpen] = useState(false);
  const [editingUser, setEditingUser] = useState<User | null>(null);

  const handleOpenModal = (user: User) => {
    setEditingUser(user);
    setIsModalOpen(true);
  };

  const handleCloseModal = () => {
    setEditingUser(null);
    setIsModalOpen(false);
  };

  const handleSave = (updatedUser: User) => {
    onUpdateUser(updatedUser);
    handleCloseModal();
  }


  return (
    <div className="bg-surface rounded-lg shadow-md overflow-hidden">
      <div className="overflow-x-auto">
        <table className="min-w-full divide-y divide-gray-200">
          <thead className="bg-gray-50">
            <tr>
              <th scope="col" className="px-6 py-3 text-left text-xs font-medium text-content-secondary uppercase tracking-wider">Customer</th>
              <th scope="col" className="px-6 py-3 text-left text-xs font-medium text-content-secondary uppercase tracking-wider">Username</th>
              <th scope="col" className="px-6 py-3 text-left text-xs font-medium text-content-secondary uppercase tracking-wider">Address</th>
              <th scope="col" className="px-6 py-3 text-left text-xs font-medium text-content-secondary uppercase tracking-wider">Joined</th>
              <th scope="col" className="relative px-6 py-3"><span className="sr-only">Edit</span></th>
            </tr>
          </thead>
          <tbody className="bg-white divide-y divide-gray-200">
            {users.map((user) => (
              <tr key={user.id}>
                <td className="px-6 py-4 whitespace-nowrap font-medium text-primary">{user.fullName}</td>
                <td className="px-6 py-4 whitespace-nowrap text-content-secondary">{user.username}</td>
                <td className="px-6 py-4 whitespace-nowrap text-content-secondary">{user.address}</td>
                <td className="px-6 py-4 whitespace-nowrap text-content-secondary">{user.joinDate.toLocaleDateString()}</td>
                <td className="px-6 py-4 whitespace-nowrap text-right">
                  <button onClick={() => handleOpenModal(user)} className="text-secondary hover:text-opacity-80 p-1">
                    <EditIcon className="h-5 w-5" />
                  </button>
                </td>
              </tr>
            ))}
          </tbody>
        </table>
      </div>
      {isModalOpen && editingUser && <EditUserModal user={editingUser} onSave={handleSave} onClose={handleCloseModal} />}
    </div>
  );
};

interface EditUserModalProps {
  user: User;
  onSave: (user: User) => void;
  onClose: () => void;
}

const EditUserModal: React.FC<EditUserModalProps> = ({ user, onSave, onClose }) => {
  const [formData, setFormData] = useState({ fullName: user.fullName, address: user.address });

  useEffect(() => {
    setFormData({ fullName: user.fullName, address: user.address });
  }, [user]);

  const handleChange = (e: React.ChangeEvent<HTMLInputElement | HTMLTextAreaElement>) => {
    setFormData(prev => ({ ...prev, [e.target.name]: e.target.value }));
  };

  const handleSubmit = (e: React.FormEvent) => {
    e.preventDefault();
    onSave({ ...user, ...formData });
  };

  return (
    <div className="fixed inset-0 bg-black bg-opacity-50 flex items-center justify-center z-30 animate-fade-in p-4">
      <div className="bg-surface rounded-lg shadow-2xl w-full max-w-md animate-slide-up relative">
        <button onClick={onClose} className="absolute top-4 right-4 text-gray-400 hover:text-gray-600">
          <XIcon className="h-6 w-6" />
        </button>
        <form onSubmit={handleSubmit} className="p-8">
          <h2 className="text-2xl font-bold text-primary mb-6">Edit Customer Details</h2>
          <div className="space-y-4">
            <div>
              <label htmlFor="fullName" className="block text-sm font-medium text-content-primary">Full Name</label>
              <input type="text" name="fullName" id="fullName" value={formData.fullName} onChange={handleChange} className="mt-1 w-full p-2 border border-gray-300 rounded-md shadow-sm focus:ring-secondary focus:border-secondary" />
            </div>
            <div>
              <label htmlFor="address" className="block text-sm font-medium text-content-primary">Address</label>
              <textarea name="address" id="address" value={formData.address} onChange={handleChange} rows={3} className="mt-1 w-full p-2 border border-gray-300 rounded-md shadow-sm focus:ring-secondary focus:border-secondary"></textarea>
            </div>
          </div>
          <div className="mt-8 flex justify-end gap-4">
            <button type="button" onClick={onClose} className="py-2 px-4 border border-gray-300 rounded-md shadow-sm text-sm font-medium text-gray-700 bg-white hover:bg-gray-50">Cancel</button>
            <button type="submit" className="py-2 px-4 border border-transparent rounded-md shadow-sm text-sm font-medium text-white bg-accent hover:bg-opacity-90">Save Changes</button>
          </div>
        </form>
      </div>
    </div>
  );
}


export default CustomerManagement;
