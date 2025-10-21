
import { Component, createResource, Show } from 'solid-js';
import { useNavigate } from '@solidjs/router';

import HostStatus from '~/HostStatus';
import { BASE_URL } from '~/index';



const MainPage: Component = () => {
  const navigate = useNavigate();

  const [user] = createResource(async () => {
    try {
      const response = await fetch(`${BASE_URL}/api/authenticated/user`, {
        credentials: 'include',
      });
      
      if (!response.ok) {
        return null;
      }
      
      return await response.json();
    } catch (error) {
      console.error('Failed to fetch user:', error);
      return null;
    }
  });

  // Redirect to session selection if user is authenticated but has no role
  const checkUserRole = () => {
    if (user.state === 'ready' && user() && !user().role) {
      navigate('/session');
    }
  };

  checkUserRole();

  return (
    <Show 
      when={!user.loading && user()}
      fallback={<HostStatus />}
    >
      <Show 
        when={user()?.role}
        fallback={<div class="flex items-center justify-center min-h-screen">
          <div class="text-center">
            <div class="animate-spin rounded-full h-12 w-12 border-4 border-purple-500 border-t-transparent mx-auto mb-4" />
            <p class="text-gray-600">Setting up your session...</p>
          </div>
        </div>}
      >
        <div class="p-6">
          <div class="bg-white rounded-lg shadow-md p-6 mb-6">
            <h2 class="text-2xl font-bold mb-2">Welcome, {user()?.name || 'User'}!</h2>
            <p class="text-gray-600">You are logged in as: <span class="font-semibold text-purple-600">{user()?.role}</span></p>
          </div>
          <HostStatus />
        </div>
      </Show>
    </Show>
  );
};

export default MainPage;
