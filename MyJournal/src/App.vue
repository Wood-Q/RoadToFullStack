<template>
  <div class="app-container">
    <TopNavbar />
    <div class="main-content">
      <Sidebar />
      <MdEditor />
      <RightToolbar />
    </div>
    <Toast />
  </div>
</template>

<script setup>
import MdEditor from "./components/mdEditor.vue";
import Sidebar from "./components/sidebar.vue";
import TopNavbar from "./components/TopNavbar.vue";
import RightToolbar from "./components/RightToolbar.vue";
import Toast from "./components/Toast.vue";
</script>

<style scoped>
.app-container {
  display: flex;
  flex-direction: column;
  height: 100vh;
  width: 100vw;
  background: linear-gradient(135deg, #1a1a1a 0%, #2d2d2d 100%);
  color: #e0e0e0;
  font-family: -apple-system, BlinkMacSystemFont, "Segoe UI", Roboto, Helvetica,
    Arial, sans-serif;
}

.main-content {
  display: grid;
  grid-template-columns: 280px 1fr 320px;
  grid-template-areas: "sidebar editor toolbar";
  flex: 1;
  gap: 1px;
  background-color: #333;
  padding: 1px;
  overflow: hidden;
}

/* 响应式设计 */
@media (max-width: 1200px) {
  .main-content {
    grid-template-columns: 240px 1fr 280px;
  }
}

@media (max-width: 992px) {
  .main-content {
    grid-template-columns: 200px 1fr 240px;
  }
}

@media (max-width: 768px) {
  .main-content {
    grid-template-columns: 1fr;
    grid-template-areas: "editor";
  }

  .main-content :deep(.sidebar),
  .main-content :deep(.right-toolbar) {
    display: none;
  }
}
</style>
