<template>
  <div class="max-w-full min-h-screen flex items-center justify-center">
    <div class="max-w-sm min-w-sm mx-auto">
      <div class="shadow p-7 rounded-lg border">
        <div class="text-sm">
          <div class="text-sm font-semibold">Welcome Back!</div>
          <p class="text-muted-foreground">Use your username and password to log in.</p>
        </div>
        <div class="mt-6">
          <LoginForm />
        </div>
      </div>
    </div>
  </div>
</template>
<script setup lang="ts">
import LoginForm from "@/features/auth/components/login-form.vue";
</script>
