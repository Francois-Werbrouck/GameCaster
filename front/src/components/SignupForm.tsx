import { Component, Show, createSignal } from "solid-js";
import { Input } from "./ui/input";
import { Label } from "./ui/label";
import { Button } from "~/components/ui/button"
import { useAuth } from "~/context/AuthContext";
import Alert, { AlertSeverity } from "./Alert";




const SignupForm: Component = () => {

  const [email, setEmail] = createSignal<string>("")
  const [password, setPassword] = createSignal<string>("")
  const [passwordConfirmation, setPasswordConfirmation] = createSignal<string>("")

  const { registerToServer } = useAuth();

  const [errorMessage, setErrorMessage] = createSignal<string>("");

  const signup = async () => {

    let error_messages: string[] = [];

    if (!email().includes("@")) {
      error_messages.push("Email adress does not contain @")
    }

    if (!(password() === passwordConfirmation())) {
      error_messages.push("Confirmation password is not identical to first password")
    }
    if (password().length < 8) {
      error_messages.push("Password contains less than 8 characters")
    }

    if (error_messages.length < 1) {

      const response = await registerToServer(email(), password());

      if (response.status === "failed") {
        setErrorMessage(response.error)
      }

      if (response.status === "succeded") {
        window.location.replace("/")

      }

      console.log(response);
    } else {
      setErrorMessage(error_messages.join(" and "));
    }

  }




  return (
    <>
      <form>
        <div class="mb-4">
          <Label class="block text-gray-700 text-sm font-bold mb-2" for="Email">Email</Label>
          <Input onChange={(e) => setEmail(e.target.value)} class="shadow appearance-none border rounded w-full py-2 px-3 text-gray-700 leading-tight focus:outline-none focus:shadow-outline" id="Email" type="email" placeholder="Username" />
        </div>
        <div class="mb-6">
          <Label class="block text-gray-700 text-sm font-bold mb-2" for="new-password">Password</Label>
          <Input onChange={(e) => setPassword(e.target.value)} class="shadow appearance-none border rounded w-full py-2 px-3 text-gray-700 mb-3 leading-tight focus:outline-none focus:shadow-outline" id="password" type="password" name="new-password" placeholder="very secure password" />
        </div>

        <div class="mb-6">
          <Label class="block text-gray-700 text-sm font-bold mb-2" for="password-confirmation">Confirm Password</Label>
          <Input onChange={(e) => setPasswordConfirmation(e.target.value)} class="shadow appearance-none border rounded w-full py-2 px-3 text-gray-700 mb-3 leading-tight focus:outline-none focus:shadow-outline" id="password-confirmation" type="password" name="new-password-confirmation" placeholder="very secure password" />
        </div>
        <div class="flex items-center justify-between">
          <Button class=" hover:bg-blue-700 text-white font-bold py-2 px-4 rounded focus:outline-none focus:shadow-outline" type="button" onClick={() => { signup() }}>
            Register
          </Button>
          <a class="inline-block align-baseline font-bold text-sm text-blue-500 hover:text-blue-800" href="/auth/reset">
            Forgot Password?
          </a>
        </div>
      </form>
      <Show when={errorMessage() !== ""}>
        <Alert message="This is very serious" severity={AlertSeverity.high}/>
        <div class="text-red-600">{errorMessage()}</div>
      </Show>
    </>
  );
}

export default SignupForm
