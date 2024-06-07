import { A } from "@solidjs/router";
import { Component, JSXElement } from "solid-js";

interface NavProp {
  children?: JSXElement;
}


const Nav: Component<NavProp> = (props: { children?: JSXElement }) => {


  return (
    <>
      <header class="bg-yellow-600 h-12 flex items-center w-full fixed top-0 z-10">
        <div class="container mx-auto">
          <h1 class="text-white">Logo here</h1>
          <a href="/"> home</a>
          <a href="/auth"> Login</a>
          <a href="/auth/reset"> Forgotten password?</a>
        </div>
      </header>

      <main class="flex-grow mt-12">
        <div class="container ">
          <h2 class="text-2xl ml-4 font-bold mb-4">Main Content</h2>

          {props.children}
        </div>
      </main>

      <footer class="bg-yellow-600 h-12 flex items-center w-full mt-auto">
        <div class="container mx-auto">
          <p class="text-white">Bottom Bar</p>
        </div>
      </footer>
    </>
  );
}

export default Nav;
