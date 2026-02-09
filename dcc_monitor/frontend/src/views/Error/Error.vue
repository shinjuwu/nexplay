<template>
  <div class="mt-4 text-center text-6xl font-bold sm:text-8xl">404</div>
  <div class="mt-2 text-center text-2xl sm:text-3xl">抱歉，找不到该页面</div>
  <div class="mt-2 text-center text-xl sm:text-2xl">
    <router-link v-slot="{ navigate }" to="/" custom>
      <button type="button" class="btn btn-primary" @click="navigate">返回</button>
    </router-link>
  </div>
  <div class="fixed inset-0 -z-10 bg-gray-100"></div>
</template>
