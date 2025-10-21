import { Component, createSignal, Show } from 'solid-js';
import { useNavigate } from '@solidjs/router';
import { BASE_URL } from '~/index';
import Alert, { AlertSeverity } from '~/components/Alert';

const SessionSelection: Component = () => {
  const [selectedRole, setSelectedRole] = createSignal<string>('');
  const [isLoading, setIsLoading] = createSignal<boolean>(false);
  const [error, setError] = createSignal<string>('');
  const navigate = useNavigate();

  const handleRoleSelection = async (role: string) => {
    setIsLoading(true);
    setError('');

    try {
      const response = await fetch(`${BASE_URL}/api/authenticated/user/role`, {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json',
        },
        credentials: 'include',
        body: JSON.stringify({ role }),
      });

      if (!response.ok) {
        throw new Error('Failed to set role');
      }

      setSelectedRole(role);
      // Navigate to the main page or role-specific page
      navigate('/');
    } catch (err) {
      setError('Failed to select role. Please try again.');
      console.error(err);
    } finally {
      setIsLoading(false);
    }
  };

  return (
    <div class="flex flex-col items-center justify-center min-h-screen bg-gradient-to-br from-slate-900 via-purple-900 to-slate-900">
      <div class="bg-white/10 backdrop-blur-lg rounded-2xl shadow-2xl p-8 max-w-2xl w-full mx-4">
        <h1 class="text-4xl font-bold text-white text-center mb-2">
          Welcome to GameCaster
        </h1>
        <p class="text-gray-300 text-center mb-8">
          Choose your role to begin your adventure
        </p>

        <Show when={error()}>
          <div class="mb-6">
            <Alert message={error()} severity={AlertSeverity.high} />
          </div>
        </Show>

        <div class="grid grid-cols-1 md:grid-cols-3 gap-6">
          {/* DM (Dungeon Master) Card */}
          <button
            onClick={() => handleRoleSelection('DM')}
            disabled={isLoading()}
            class="group relative overflow-hidden rounded-xl bg-gradient-to-br from-red-500 to-red-700 p-6 text-white transition-all duration-300 hover:scale-105 hover:shadow-2xl disabled:opacity-50 disabled:cursor-not-allowed"
          >
            <div class="relative z-10">
              <div class="text-5xl mb-4">🎲</div>
              <h2 class="text-2xl font-bold mb-2">DM</h2>
              <p class="text-sm text-red-100">
                Lead the adventure, control the world, and guide your players through epic quests
              </p>
            </div>
            <div class="absolute inset-0 bg-gradient-to-t from-black/20 to-transparent opacity-0 group-hover:opacity-100 transition-opacity" />
          </button>

          {/* Player Card */}
          <button
            onClick={() => handleRoleSelection('Player')}
            disabled={isLoading()}
            class="group relative overflow-hidden rounded-xl bg-gradient-to-br from-blue-500 to-blue-700 p-6 text-white transition-all duration-300 hover:scale-105 hover:shadow-2xl disabled:opacity-50 disabled:cursor-not-allowed"
          >
            <div class="relative z-10">
              <div class="text-5xl mb-4">⚔️</div>
              <h2 class="text-2xl font-bold mb-2">Player</h2>
              <p class="text-sm text-blue-100">
                Create your character, join the party, and embark on thrilling adventures
              </p>
            </div>
            <div class="absolute inset-0 bg-gradient-to-t from-black/20 to-transparent opacity-0 group-hover:opacity-100 transition-opacity" />
          </button>

          {/* Casting Card */}
          <button
            onClick={() => handleRoleSelection('Casting')}
            disabled={isLoading()}
            class="group relative overflow-hidden rounded-xl bg-gradient-to-br from-purple-500 to-purple-700 p-6 text-white transition-all duration-300 hover:scale-105 hover:shadow-2xl disabled:opacity-50 disabled:cursor-not-allowed"
          >
            <div class="relative z-10">
              <div class="text-5xl mb-4">📺</div>
              <h2 class="text-2xl font-bold mb-2">Casting</h2>
              <p class="text-sm text-purple-100">
                View and spectate games in real-time, perfect for streaming and watching
              </p>
            </div>
            <div class="absolute inset-0 bg-gradient-to-t from-black/20 to-transparent opacity-0 group-hover:opacity-100 transition-opacity" />
          </button>
        </div>

        <Show when={isLoading()}>
          <div class="mt-6 text-center">
            <div class="inline-block animate-spin rounded-full h-8 w-8 border-4 border-white border-t-transparent" />
            <p class="text-white mt-2">Setting up your session...</p>
          </div>
        </Show>
      </div>
    </div>
  );
};

export default SessionSelection;
