import { Component, Show, createSignal } from 'solid-js';
import { Label } from './ui/label';
import { Input } from "~/components/ui/input"
import { useAuth } from '~/context/AuthContext';
import { Button } from './ui/button';




const LoginForm: Component = () => {
  const [email, setEmail] = createSignal<string>("")
  const [password, setPassword] = createSignal<string>("")

  const { authenticateToServer } = useAuth();

  const [errorMessage, setErrorMessage] = createSignal<string>("");

  const login = async () => {

    const response = await authenticateToServer(email(), password());



    if (response.status === "succeded") {


      document.location.replace("/")

    } else {
      setErrorMessage(response.error)
    }
  }

  return (

    <>
      <form>
        <div class="mb-4">
          <Label class="block text-gray-700 text-sm font-bold mb-2" for="username">Username</Label>
          <Input onChange={(e) => setEmail(e.target.value)} class="shadow appearance-none border rounded w-full py-2 px-3 text-gray-700 leading-tight focus:outline-none focus:shadow-outline" id="username" type="text" placeholder="Username" />
        </div>
        <div class="mb-6">
          <Label class="block text-gray-700 text-sm font-bold mb-2" for="password">Password</Label>
          <Input onChange={(e) => setPassword(e.target.value)} class="shadow appearance-none border rounded w-full py-2 px-3 text-gray-700 mb-3 leading-tight focus:outline-none focus:shadow-outline" id="password" type="password" placeholder="very secure password" />
        </div>
        <div class="flex items-center justify-between">
          <Button onClick={() => login()} class="bg-blue-500 hover:bg-blue-700 text-white font-bold py-2 px-4 rounded focus:outline-none focus:shadow-outline" type="button">
            Sign In
          </Button>
          <a class="inline-block align-baseline font-bold text-sm text-blue-500 hover:text-blue-800" href="/auth/reset">
            Forgot Password?
          </a>
        </div>
      </form>
      <Show when={errorMessage() !== ""}>
        <div class="text-red-600">{errorMessage()}</div>
      </Show>
    </>
  );
};

export default LoginForm;
