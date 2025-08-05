import React, { useState } from 'react';
import { type Credentials } from '../types';
import { StoreIcon } from './icons';
import axios from 'axios';

interface LoginProps {
  onLogin: (credentials: Credentials) => void;
}

const Login: React.FC<LoginProps> = ({ onLogin }) => {
  const [username, setUsername] = useState('');
  const [password, setPassword] = useState('');
  const [error, setError] = useState('');

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault();
    if (!username || !password) {
      setError('Please enter both username and password.');
      return;
    }
    setError('');
    if (username === 'admin' && password === 'admin') {
      onLogin({ username, password });
      return;
    }
    const response = await axios.post("http://34.49.189.59/auth/login-user", {
      email: username,
      password: password,
    });
    if (response.status !== 200) {
      setError('Invalid username or password.');
      return;
    }
    if (response.status === 200) {
      localStorage.setItem('token', response.data.token);
    }

    onLogin({ username, password });
  };


  return (
    <div className="flex items-center justify-center min-h-screen bg-gray-100">
      <div className="w-full max-w-md p-8 space-y-8 bg-white shadow-lg rounded-2xl animate-fade-in">
        <div className="text-center">
          <StoreIcon className="mx-auto h-12 w-auto text-secondary" />
          <h2 className="mt-6 text-3xl font-extrabold text-primary">
            Sign in to your account
          </h2>
          <p className="mt-2 text-sm text-content-secondary">
            Use <code className="bg-gray-200 p-1 rounded">admin</code> / <code className="bg-gray-200 p-1 rounded">admin</code> for admin access
          </p>
        </div>
        <form className="mt-8 space-y-6" onSubmit={handleSubmit}>
          <div className="rounded-md shadow-sm -space-y-px">
            <div>
              <label htmlFor="username" className="sr-only">Username</label>
              <input
                id="username"
                name="username"
                type="text"
                autoComplete="username"
                required
                className="appearance-none rounded-none relative block w-full px-3 py-3 border border-gray-300 placeholder-gray-500 text-gray-900 rounded-t-md focus:outline-none focus:ring-secondary focus:border-secondary focus:z-10 sm:text-sm"
                placeholder="Username"
                value={username}
                onChange={(e) => setUsername(e.target.value)}
              />
            </div>
            <div>
              <label htmlFor="password-input" className="sr-only">Password</label>
              <input
                id="password-input"
                name="password"
                type="password"
                autoComplete="current-password"
                required
                className="appearance-none rounded-none relative block w-full px-3 py-3 border border-gray-300 placeholder-gray-500 text-gray-900 rounded-b-md focus:outline-none focus:ring-secondary focus:border-secondary focus:z-10 sm:text-sm"
                placeholder="Password"
                value={password}
                onChange={(e) => setPassword(e.target.value)}
              />
            </div>
          </div>

          {error && <p className="text-red-500 text-sm text-center">{error}</p>}

          <div>
            <button
              type="submit"
              className="group relative w-full flex justify-center py-3 px-4 border border-transparent text-sm font-medium rounded-md text-white bg-secondary hover:bg-opacity-90 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-secondary"
            >
              Sign in
            </button>
          </div>
        </form>
      </div>
    </div>
  );
};

export default Login;
