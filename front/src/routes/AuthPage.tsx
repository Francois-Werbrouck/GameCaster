import { Route, Router } from '@solidjs/router';
import { Component } from 'solid-js';
import LoginForm from '~/components/LoginForm';
import SignupForm from '~/components/SignupForm';
import { Tabs, TabsContent, TabsList, TabsTrigger } from "~/components/ui/tabs"
import PasswordReset from './auth/PasswordReset';




const AuthPage: Component = () => {


  return (
  <>

    <div class='flex align-middle justify-center mt-5'>
    <Tabs defaultValue="account" class="w-[400px]">
      <TabsList class="grid w-full grid-cols-2">
        <TabsTrigger value="login">Login</TabsTrigger>
        <TabsTrigger value="create an account">Create an account</TabsTrigger>
      </TabsList>
      <TabsContent value="login">
        <LoginForm />
      </TabsContent>
      <TabsContent value="create an account">
        <SignupForm />
      </TabsContent>
    </Tabs>
    </div>
    </>
  );
};

export default AuthPage;
