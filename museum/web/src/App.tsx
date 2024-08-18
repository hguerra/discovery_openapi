import { createResource, For, type Component } from 'solid-js';

import logo from './logo.svg';
import styles from './App.module.css';

import { eventsApi } from './infra/api';

const fetchEvents = async () => {
  const res = await eventsApi.listSpecialEvents();
  return res.data;
};

const App: Component = () => {
  const [events] = createResource(fetchEvents);

  return (
    <div class={styles.App}>
      <header class={styles.header}>
        <img src={logo} class={styles.logo} alt='logo' />
        <p>
          Edit <code>src/App.tsx</code> and save to reload.
        </p>
        <a
          class={styles.link}
          href='https://github.com/solidjs/solid'
          target='_blank'
          rel='noopener noreferrer'
        >
          Learn Solid
        </a>
      </header>
      <main class={styles.main}>
        <For each={events()}>
          {(event, i) => (
            <li>
              <a
                target='_blank'
                href={`https://www.youtube.com/watch?v=${event.eventId}`}
              >
                {i() + 1}: {event.name}
              </a>
            </li>
          )}
        </For>
      </main>
    </div>
  );
};

export default App;
