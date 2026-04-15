<script setup lang="ts">
import Button from "@/components/ui/button/Button.vue";
import { Input } from "@/components/ui/input";
import { Label } from "@/components/ui/label";
import api from "@/lib/api";
import { AxiosError } from "axios";
import { reactive } from "vue";
import { toast } from "vue-sonner";
const form = reactive({
  username: "",
  password: " ",
});
const login = async () => {
  const loading = toast.loading("Loading");
  try {
    const respon = await api.post("/login", form);
  } catch (error) {
    const err = error as AxiosError;
    if (err.status == 422) {
      toast.error("Validation Error", {
        description: (err.response?.data as any).errors,
      });
    } else {
      console.log("Suksess");
    }
  } finally {
    toast.dismiss(loading);
  }
};
</script>

<template>
  <form action="" @submit.prevent="login" method="POST">
    <div>
      <div class="mb-3 flex flex-col space-y-1">
        <Label for="username" class="text-muted-foreground">Username</Label>
        <Input v-model="form.username" name="username" />
      </div>
      <div class="mb-3 flex flex-col space-y-1">
        <Label for="password" class="text-muted-foreground">Kata Sandi</Label>
        <Input v-model="form.password" name="password" />
      </div>
      <div>
        <Button>Login</Button>
      </div>
    </div>
  </form>
</template>
