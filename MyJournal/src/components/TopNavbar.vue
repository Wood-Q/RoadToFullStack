<template>
  <nav class="top-navbar">
    <div class="logo">一叶日记</div>
    <div class="actions">
      <button @click="handleCreateJournal">保存日记</button>
    </div>
  </nav>
</template>

<script setup lang="ts">
import { createJournal } from "../services/journal";
import { useMdContentStore } from "../store/mdContent";
import { useToastStore } from "../store/toast";

const mdContentStore = useMdContentStore();
const toastStore = useToastStore();

async function handleCreateJournal() {
  try {
    const result = await createJournal("新的日记", mdContentStore.mdText);
    toastStore.showToast(result, "success");
  } catch (error) {
    toastStore.showToast(String(error), "error");
  }
}
</script>

<style scoped>
.top-navbar {
  margin-left: 1.2em;
  grid-area: top-nav;
  display: flex;
  justify-content: space-between;
  align-items: center;
  padding: 0 1.2em;
  background-color: #2a2a2a; /* Slightly different from main background */
  border-bottom: 1px solid #333;
  color: #e0e0e0;
  border-radius: 0.5em;
}

.logo {
  font-weight: 600;
  font-size: 1.2rem;
}

.actions button {
  background-color: transparent;
  color: #ef8e2c;
  border: 1px solid #555;
  padding: 0.5em 1em;
  border-radius: 0.5em;
  cursor: pointer;
  margin-left: 0.6em;
  transition: background-color 0.2s ease, border-color 0.2s ease;
}

.actions button:hover {
  background-color: #3c3c3c;
  border-color: #777;
}
</style>
