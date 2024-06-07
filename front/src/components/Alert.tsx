import { Component, Match, Show, Switch, createSignal, onCleanup } from 'solid-js';
import { Alert as AlertComponent, AlertDescription, AlertTitle } from '~/components/ui/alert';

export enum AlertSeverity {
  low,
  medium,
  high,
}

const Alert: Component<{message: string, severity: AlertSeverity}> = (props) => {

  const { message, severity } = props;

  const [visible, setVisible] = createSignal(true);

  // Hide the alert after 5 seconds
  const timeout = setTimeout(() => setVisible(false), 5000);


  // Cleanup the timeout if the component is destroyed before it triggers
  onCleanup(() => clearTimeout(timeout));


  return (
    <Show when={visible()}>
      <div class='alert absolute top-2 right-2'>
        <AlertComponent>
          <Switch>
            <Match when={severity === AlertSeverity.low}>
              Attention
            </Match >
            <Match when={severity === AlertSeverity.medium}>
              Warning
            </Match >
            <Match when={severity === AlertSeverity.high}>
              Error
            </Match >
          </Switch>
          <AlertTitle></AlertTitle>
          <AlertDescription>
            {message}
          </AlertDescription>
        </AlertComponent>
      </div>
    </Show>
  );
};

export default Alert;
