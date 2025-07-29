<template>
  <aside class="sidebar">
    <div class="sidebar-header">
      <div class="user-profile">
        <div class="avatar">
          <svg xmlns="http://www.w3.org/2000/svg" width="40" height="40" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="1.5" stroke-linecap="round" stroke-linejoin="round" class="feather feather-leaf"><path d="M20.24 12.24a6 6 0 0 0-8.49-8.49L5 10.5V19h8.5z"></path><line x1="16" y1="8" x2="22" y2="2"></line><line x1="12" y1="12" x2="18" y2="6"></line></svg>
        </div>
        <div class="user-info">
          <h1 class="username">一叶日记</h1>
          <p class="motto">琐碎即是生活</p>
        </div>
      </div>
    </div>

  </aside>
</template>

<script setup lang="ts">
// No script logic needed for this static component
</script>

<style scoped>
.sidebar {
  grid-area: sidebar;
  width: 100%;
  height: 100vh;
  background-color: #1e1e1e;
  color: #e0e0e0;
  display: flex;
  flex-direction: column;
  padding: 1.2em;
  box-sizing: border-box;
  font-family: -apple-system, BlinkMacSystemFont, "Segoe UI", Roboto, Helvetica, Arial, sans-serif;
}

.sidebar-header {
  margin-bottom: 1.2em;
}

.user-profile {
  display: flex;
  align-items: center;
  margin-bottom: 1.2em;
}

.avatar {
  width: 3em;
  height: 3em;
  border-radius: 50%;
  background-color: #c46a21;
  display: flex;
  align-items: center;
  justify-content: center;
  margin-right: 1em;
}

.avatar svg {
  color: #f5d5b1;
}

.user-info .username {
  font-size: 1.2rem;
  font-weight: 600;
  margin: 0;
}

.user-info .motto {
  font-size: 0.8rem;
  color: #a0a0a0;
  margin: 0.3em 0 0 0;
}

</style>