<template>
  <aside class="right-toolbar">
    <div class="toolbar-section">
      <p>字体调整</p>
      <!-- Add font adjustment tools here -->
    </div>
    <div class="toolbar-section">
      <p>媒体插入</p>
      <!-- Add media insertion tools here -->
    </div>
  </aside>
</template>

<script setup lang="ts">
</script>

<style scoped>
.right-toolbar {
  grid-area: right-toolbar;
  background-color: #2a2a2a;
  padding: 1.2em;
  color: #e0e0e0;
  border-left: 1px solid #333;
}

.toolbar-section {
  margin-bottom: 1.2em;
}

.toolbar-section p {
  font-weight: 600;
  margin-bottom: 0.6em;
}
</style>