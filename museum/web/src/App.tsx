import { createResource, type Component } from 'solid-js';

import logo from './logo.svg';
import styles from './App.module.css';

import { eventsApi } from './infra/api'

const fetchEvents = async () => {
  const res = await eventsApi.listSpecialEvents();
  return res.data;
}

const App: Component = () => {
  const [events] = createResource(fetchEvents);

  return (
    <div class={styles.App}>
      <header class={styles.header}>
        <img src={logo} class={styles.logo} alt="logo" />
        <p>
          Edit <code>src/App.tsx</code> and save to reload.
        </p>
        <a
          class={styles.link}
          href="https://github.com/solidjs/solid"
          target="_blank"
          rel="noopener noreferrer"
        >
          Learn Solid
        </a>
      </header>
    </div>
  );
};

export default App;
