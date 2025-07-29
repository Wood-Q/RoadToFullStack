import { defineStore } from 'pinia';

export const useToastStore = defineStore('toast', {
  state: () => ({
    isVisible: false,
    message: '',
    type: 'success' as 'success' | 'error',
  }),
  actions: {
    showToast(message: string, type: 'success' | 'error' = 'success') {
      this.message = message;
      this.type = type;
      this.isVisible = true;

      setTimeout(() => {
        this.hideToast();
      }, 3000); // Auto-hide after 3 seconds
    },
    hideToast() {
      this.isVisible = false;
    },
  },
});