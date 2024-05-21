import { Component, Match, Show, Switch, createResource, createSignal } from 'solid-js';
import { Alert as AlertComponent, AlertDescription, AlertTitle } from '~/components/ui/alert';

export enum AlertSeverity {
  low,
  medium,
  high,
}

const Alert = (props: { message: string, severity: AlertSeverity }) => {

  const { message, severity } = props;




  return (
    <div class='absolute top-2 right-2'>
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
  );
};

export default Alert;
