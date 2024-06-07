
import { Component, Match, Show, Switch, createResource, createSignal } from 'solid-js';
import Alert, { AlertSeverity } from './components/Alert';



const HostStatus: Component = () => {

  const [hostStatus, setHostStatus] = createSignal<string>("");

  const [status] = createResource(`/api/status`, async (url) => {
    const response = await fetch(url);

    return response.json()
  })

  const [authStatus, setAuthStatus] = createSignal<string>("");

  const [auth] = createResource(`/api/authenticated/test`, async (url) => {
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

      <div class='bg-cyan-50'>
        <Show when={auth.loading}>
          <p>Loading...</p>
        </Show>
        <Switch>
          <Match when={auth.error}>
            <span>Error: {status.error}</span>
          </Match>
          <Match when={auth.state == "ready"}>
            <div>{JSON.stringify(auth())}</div>
            <Alert message='Something is here' severity={AlertSeverity.low} />
          </Match>
        </Switch>
      </div>
    </>
  );
};

export default HostStatus;
