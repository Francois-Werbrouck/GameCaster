
import { Component, Match, Show, Switch, createResource, createSignal } from 'solid-js';
import { BASE_URL } from '.';
import Alert, { AlertSeverity } from './components/Alert';



const HostStatus: Component = () => {

  const [hostStatus, setHostStatus] = createSignal<string>("");

  const [status] = createResource(`${BASE_URL}/status`, async (url) => {
    const response = await fetch(url);

    return response.json()
  })



  return (
    <>
      <div class='bg-cyan-50'>
        <Show when={status.loading}>
          <p>Loading...</p>
        </Show>
        <Switch>
          <Match when={status.error}>
            <span>Error: {status.error}</span>
          </Match>
          <Match when={status.state == "ready"}>
            <div>{JSON.stringify(status())}</div>
            <Alert message='Something is here' severity={AlertSeverity.low} />
          </Match>
        </Switch>
      </div>
    </>
  );
};

export default HostStatus;
