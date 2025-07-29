<template>
  <div class="app-grid-container">
    <TopNavbar />
    <Sidebar />
    <MdEditor />
    <RightToolbar />
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
.app-grid-container {
  display: grid;
  grid-template-columns: 18% 1fr 18%; /* Sidebar, Editor, Toolbar */
  grid-template-rows: 60px 1fr; /* TopNav, Main Content */
  grid-template-areas:
    "sidebar top-nav top-nav"
    "sidebar editor right-toolbar";
  height: 100vh;
  width: 100vw;
  background-color: #1e1e1e;
}

/* Assign components to grid areas */
/* Note: These selectors will work if the child components have a single root element. 
   We will need to assign the grid-area property directly in the child components' styles. 
   This is a better practice for component-based architectures. 
   I will add the grid-area property to each component in the next steps. */
</style>
