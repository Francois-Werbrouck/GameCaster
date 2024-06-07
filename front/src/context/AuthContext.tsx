import { log } from "console";
import { JSXElement, createSignal, useContext } from "solid-js";
import { createContext } from "solid-js";


interface WithChildren {
  children?: JSXElement;
}

export type AuthenticationServerResponse = SuccessResponse | FailResponse

interface SuccessResponse {
  status: 'succeded';
  token: never
}

interface FailResponse {
  status: 'failed';
  error: string;
  token: never; // Explicitly state that token should not be present
}


export interface AuthContextActions {
  isLogged: () => boolean;
  authenticateToServer: (email: string, password: string) => Promise<AuthenticationServerResponse>;
  registerToServer: (email: string, password: string) => Promise<AuthenticationServerResponse>;
  logout: () => void;
  id: () => number;
  email: () => string;

}

const AuthContext = createContext<AuthContextActions>();



export const AuthProvider = (props: WithChildren) => {

  const [isLogged, setIsLogged] = createSignal<boolean>(false);
  const [id, setId] = createSignal<number>(-1);
  const [email, setEmail] = createSignal<string>("");



  const login = async (email: string, password: string) => {


    const request = await fetch(`/api/auth/login`, {
      method: "POST",
      headers: {
        "Content-Type": "application/json"
      },
      body: JSON.stringify({
        email: email,
        password: password
      })
    })

    const response: AuthenticationServerResponse = await request.json();

    console.log(response)

    if (response.status === "succeded") {
      setIsLogged(true);


      setEmail(email);
    }

    return response;


  }

  const register = async (email: string, password: string) => {

    console.log(JSON.stringify({
      email: email,
      password: password
    })
    )
    const request = await fetch(`/api/auth/register`, {
      method: "POST",
      headers: {
        "Content-Type": "application/json"
      },
      body: JSON.stringify({
        email: email,
        password: password
      })
    })

    try {

      const response: AuthenticationServerResponse = await request.json();

      if (response.status === "succeded") {
        setIsLogged(true);


        setEmail(email);
      }

      return response;
    } catch (e: any) {

      return JSON.stringify({
        status: "failed",
        error: "unable to parse json from server"
      })

    }



  }

  const values: AuthContextActions = {
    isLogged: isLogged,
    email: email,
    id: id,
    logout: () => setIsLogged(false),
    authenticateToServer: login,
    registerToServer: register,
  };
  return (
    <AuthContext.Provider value={values}>
      {props.children}
    </AuthContext.Provider>
  );
}



export function useAuth(): AuthContextActions {
  const contextValue = useContext(AuthContext);

  if (!contextValue) {
    throw new Error("useAuth must be used within an AuthProvider");
  }

  return contextValue;
}
