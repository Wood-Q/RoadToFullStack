<template>
  <transition name="toast-fade">
    <div v-if="toast.isVisible" :class="['toast-container', toast.type]">
      <p>{{ toast.message }}</p>
    </div>
  </transition>
</template>

<script setup lang="ts">
import { useToastStore } from '../store/toast';

const toast = useToastStore();
</script>

<style scoped>
.toast-container {
  position: fixed;
  top: 1.2em;
  left: 50%;
  transform: translateX(-50%);
  padding: 1em 1.5em;
  border-radius: 0.5em;
  color: #fff;
  font-weight: 600;
  z-index: 9999;
  box-shadow: 0 0.25em 0.75em rgba(0, 0, 0, 0.15);
  display: flex;
  align-items: center;
  justify-content: center;
}

.toast-container.success {
  background-color: #4CAF50; /* Green */
}

.toast-container.error {
  background-color: #F44336; /* Red */
}

/* Transition effects */
.toast-fade-enter-active,
.toast-fade-leave-active {
  transition: opacity 0.5s, transform 0.5s;
}

.toast-fade-enter-from,
.toast-fade-leave-to {
  opacity: 0;
  transform: translate(-50%, -1.2em);
}
</style>