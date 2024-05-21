import type { Component } from 'solid-js';

import logo from './logo.svg';
import styles from './App.module.css';
import HostStatus from './HostStatus';

const App: Component = () => {
  return (
    <div class='bg-blue-900 flex-col'>
      <header class='justify-items-center' >
        <img src={logo} class='w-1/2 h-1/2 ' alt="logo" />
        <p>
          Edit <code>src/*.tsx</code> and save to reload.
        </p>
        <a
          class="bg-red-50"
          href="https://github.com/solidjs/solid"
          target="_blank"
          rel="noopener noreferrer"
        >
          Learn Solid
        </a>
      </header>
      <div class='align-bottom'>
        <HostStatus />
      </div>
    </div >
  );
};

export default App;
