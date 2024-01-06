import { useState } from "react";
import * as z from "zod";
import { zodResolver } from "@hookform/resolvers/zod";
import { useForm } from "react-hook-form";
import { Button } from "@/components/ui/button";
import {
  Card,
  CardContent,
  CardDescription,
  CardFooter,
  CardHeader,
  CardTitle,
} from "@/components/ui/card";
import { Input } from "@/components/ui/input";
import {
  Form,
  FormControl,
  FormField,
  FormItem,
  FormLabel,
  FormMessage,
} from "@/components/ui/form";
import { toast } from "@/components/ui/use-toast";
// import { paths } from "@/lib/api/v1";
import createClient from "openapi-fetch";

const registerFormSchema = z.object({
  username: z
    .string({ required_error: "This field may not be blank." })
    .min(3, { message: "Username must be at least 3 characters." })
    .max(30, { message: "Username must not be longer than 30 characters." }),
  email: z
    .string({ required_error: "This field may not be blank." })
    .email({ message: "Please enter a valid email." }),
  password: z
    .string({ required_error: "This field may not be blank." })
    .min(8, { message: "Password must be at least 8 characters." })
    .max(30, { message: "Password must not be longer than 30 characters." }),
});

type RegisterFormValues = z.infer<typeof registerFormSchema>;

interface RegisterFormProps {
  onLoginURLClick: () => void;
  onRegister: () => void;
}

// const client = createClient<paths>({ baseUrl: "/" });
// const { POST } = client;

export default function RegisterForm({
  onRegister,
  onLoginURLClick,
}: RegisterFormProps) {
  const [loading, setLoading] = useState<boolean>(false);
  const form = useForm<RegisterFormValues>({
    resolver: zodResolver(registerFormSchema),
    defaultValues: { username: "", email: "", password: "" },
  });

  async function onSubmit(formData: RegisterFormValues) {
    // try {
    //   setLoading(true);
    //   const { error, response } = await POST("/api/users/register/", {
    //     body: {
    //       username: formData.username,
    //       email: formData.email,
    //       password: formData.password,
    //     },
    //   });
    //   if (response.ok) {
    //     onRegister();
    //     return toast({
    //       title: "Account Created",
    //     });
    //   }
    //   if (error?.username)
    //     form.setError("username", { message: error.username[0] });
    //   if (error?.email) form.setError("email", { message: error.email[0] });
    //   if (error?.password)
    //     form.setError("password", { message: error.password[0] });
    //   if (error?.detail) toast({ title: error.detail });
    // } catch (error) {
    //   console.log(error);
    //   return toast({
    //     title: "Something went wrong!",
    //   });
    // } finally {
    //   setLoading(false);
    // }
  }

  return (
    <Card className="w-full md:w-[400px]">
      <CardHeader className="space-y-1">
        <CardTitle className="text-2xl">Create an account</CardTitle>
        <CardDescription>
          Enter your email below to create your account
        </CardDescription>
      </CardHeader>
      <CardContent className="grid gap-4">
        <Form {...form}>
          <form onSubmit={form.handleSubmit(onSubmit)} className="space-y-4">
            <FormField
              control={form.control}
              name="username"
              render={({ field }) => (
                <FormItem className="space-y-1">
                  <FormLabel>Username</FormLabel>
                  <FormControl>
                    <Input
                      autoComplete="username"
                      placeholder="johnsmith"
                      {...field}
                    />
                  </FormControl>
                  <FormMessage />
                </FormItem>
              )}
            />
            <FormField
              control={form.control}
              name="email"
              render={({ field }) => (
                <FormItem className="space-y-1">
                  <FormLabel>Email</FormLabel>
                  <FormControl>
                    <Input
                      autoComplete="email"
                      placeholder="user@example.com"
                      {...field}
                    />
                  </FormControl>
                  <FormMessage />
                </FormItem>
              )}
            />
            <FormField
              control={form.control}
              name="password"
              render={({ field }) => (
                <FormItem className="space-y-1">
                  <FormLabel>Password</FormLabel>
                  <FormControl>
                    <Input
                      autoComplete="new-password"
                      type="password"
                      {...field}
                    />
                  </FormControl>
                  <FormMessage />
                </FormItem>
              )}
            />
            <Button type="submit" className="w-full" disabled={loading}>
              {loading ? "Loading..." : "Create Account"}
            </Button>
          </form>
        </Form>
      </CardContent>
      <CardFooter>
        <a
          className="text-sm text-muted-foreground hover:text-primary underline underline-offset-4 cursor-pointer text-center w-full"
          onClick={() => onLoginURLClick()}
        >
          Already have an account? Sign In
        </a>
      </CardFooter>
    </Card>
  );
}
