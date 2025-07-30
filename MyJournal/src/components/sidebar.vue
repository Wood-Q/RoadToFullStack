<template>
  <aside class="sidebar">
    <div class="sidebar-header">
      <div class="user-profile">
        <div class="avatar">
          <svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="1.5" stroke-linecap="round" stroke-linejoin="round">
            <path d="M20.24 12.24a6 6 0 0 0-8.49-8.49L5 10.5V19h8.5z"></path>
            <line x1="16" y1="8" x2="22" y2="2"></line>
            <line x1="12" y1="12" x2="18" y2="6"></line>
          </svg>
        </div>
        <div class="user-info">
          <h1 class="username">一叶日记</h1>
          <p class="motto">琐碎即是生活</p>
        </div>
      </div>
    </div>

    <div class="sidebar-nav">
      <div class="nav-section">
        <h3 class="section-title">快速访问</h3>
        <ul class="nav-list">
          <li class="nav-item active">
            <svg xmlns="http://www.w3.org/2000/svg" width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
              <path d="M14 2H6a2 2 0 0 0-2 2v16a2 2 0 0 0 2 2h12a2 2 0 0 0 2-2V8z"></path>
              <polyline points="14,2 14,8 20,8"></polyline>
              <line x1="16" y1="13" x2="8" y2="13"></line>
              <line x1="16" y1="17" x2="8" y2="17"></line>
              <polyline points="10,9 9,9 8,9"></polyline>
            </svg>
            <span>今日日记</span>
          </li>
          <li class="nav-item">
            <svg xmlns="http://www.w3.org/2000/svg" width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
              <path d="M19 21H5a2 2 0 0 1-2-2V5a2 2 0 0 1 2-2h11l5 5v11a2 2 0 0 1-2 2z"></path>
              <polyline points="17,21 17,13 7,13 7,21"></polyline>
              <polyline points="7,3 7,8 15,8"></polyline>
            </svg>
            <span>所有日记</span>
            <span class="count">12</span>
          </li>
          <li class="nav-item">
            <svg xmlns="http://www.w3.org/2000/svg" width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
              <polygon points="12,2 15.09,8.26 22,9.27 17,14.14 18.18,21.02 12,17.77 5.82,21.02 7,14.14 2,9.27 8.91,8.26"></polygon>
            </svg>
            <span>收藏夹</span>
            <span class="count">3</span>
          </li>
          <li class="nav-item">
            <svg xmlns="http://www.w3.org/2000/svg" width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
              <path d="M3 3h18v18H3zM9 9h6v6H9z"></path>
            </svg>
            <span>草稿箱</span>
            <span class="count">2</span>
          </li>
        </ul>
      </div>

      <div class="nav-section">
        <div class="section-header">
          <h3 class="section-title">最近日记</h3>
          <button class="add-btn" title="新建日记">
            <svg xmlns="http://www.w3.org/2000/svg" width="14" height="14" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
              <line x1="12" y1="5" x2="12" y2="19"></line>
              <line x1="5" y1="12" x2="19" y2="12"></line>
            </svg>
          </button>
        </div>
        <ul class="journal-list">
          <li class="journal-item">
            <div class="journal-info">
              <h4 class="journal-title">今天的思考</h4>
              <p class="journal-meta">2024年1月15日</p>
            </div>
            <div class="journal-actions">
              <button class="action-btn" title="编辑">
                <svg xmlns="http://www.w3.org/2000/svg" width="12" height="12" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
                  <path d="M11 4H4a2 2 0 0 0-2 2v14a2 2 0 0 0 2 2h14a2 2 0 0 0 2-2v-7"></path>
                  <path d="M18.5 2.5a2.121 2.121 0 0 1 3 3L12 15l-4 1 1-4 9.5-9.5z"></path>
                </svg>
              </button>
            </div>
          </li>
          <li class="journal-item">
            <div class="journal-info">
              <h4 class="journal-title">周末的计划</h4>
              <p class="journal-meta">2024年1月14日</p>
            </div>
            <div class="journal-actions">
              <button class="action-btn" title="编辑">
                <svg xmlns="http://www.w3.org/2000/svg" width="12" height="12" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
                  <path d="M11 4H4a2 2 0 0 0-2 2v14a2 2 0 0 0 2 2h14a2 2 0 0 0 2-2v-7"></path>
                  <path d="M18.5 2.5a2.121 2.121 0 0 1 3 3L12 15l-4 1 1-4 9.5-9.5z"></path>
                </svg>
              </button>
            </div>
          </li>
          <li class="journal-item">
            <div class="journal-info">
              <h4 class="journal-title">读书笔记：《人生的智慧》</h4>
              <p class="journal-meta">2024年1月13日</p>
            </div>
            <div class="journal-actions">
              <button class="action-btn" title="编辑">
                <svg xmlns="http://www.w3.org/2000/svg" width="12" height="12" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
                  <path d="M11 4H4a2 2 0 0 0-2 2v14a2 2 0 0 0 2 2h14a2 2 0 0 0 2-2v-7"></path>
                  <path d="M18.5 2.5a2.121 2.121 0 0 1 3 3L12 15l-4 1 1-4 9.5-9.5z"></path>
                </svg>
              </button>
            </div>
          </li>
        </ul>
      </div>
    </div>

    <div class="sidebar-footer">
      <button class="settings-btn">
        <svg xmlns="http://www.w3.org/2000/svg" width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
          <circle cx="12" cy="12" r="3"></circle>
          <path d="M12 1v6m0 6v6m11-7h-6m-6 0H1m18-5a2 2 0 0 1-2 2H5a2 2 0 0 1-2-2V7a2 2 0 0 1 2-2h14a2 2 0 0 1 2 2v4z"></path>
        </svg>
        <span>设置</span>
      </button>
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
  background: linear-gradient(180deg, #2a2a2a 0%, #242424 100%);
  color: #e0e0e0;
  display: flex;
  flex-direction: column;
  border-right: 1px solid #333;
  overflow-y: auto;
  font-family: -apple-system, BlinkMacSystemFont, "Segoe UI", Roboto, Helvetica, Arial, sans-serif;
}

.sidebar-header {
  padding: 20px;
  border-bottom: 1px solid #333;
}

.user-profile {
  display: flex;
  align-items: center;
}

.avatar {
  width: 48px;
  height: 48px;
  border-radius: 12px;
  background: linear-gradient(135deg, #f0a500 0%, #e09400 100%);
  display: flex;
  align-items: center;
  justify-content: center;
  margin-right: 12px;
  box-shadow: 0 4px 12px rgba(240, 165, 0, 0.3);
}

.avatar svg {
  color: #fff;
}

.user-info .username {
  font-size: 1.1rem;
  font-weight: 600;
  margin: 0 0 4px 0;
  color: #f0a500;
}

.user-info .motto {
  font-size: 0.8rem;
  color: #a0a0a0;
  margin: 0;
}

.sidebar-nav {
  flex: 1;
  padding: 0;
  overflow-y: auto;
}

.nav-section {
  margin-bottom: 8px;
}

.section-header {
  display: flex;
  align-items: center;
  justify-content: space-between;
  padding: 16px 20px 8px;
}

.section-title {
  font-size: 0.75rem;
  font-weight: 600;
  color: #666;
  text-transform: uppercase;
  letter-spacing: 0.5px;
  margin: 0;
  padding: 16px 20px 8px;
}

.section-header .section-title {
  padding: 0;
}

.add-btn {
  width: 24px;
  height: 24px;
  border-radius: 6px;
  background: rgba(240, 165, 0, 0.1);
  border: none;
  color: #f0a500;
  cursor: pointer;
  display: flex;
  align-items: center;
  justify-content: center;
  transition: all 0.2s ease;
}

.add-btn:hover {
  background: rgba(240, 165, 0, 0.2);
}

.nav-list {
  list-style: none;
  padding: 0;
  margin: 0;
}

.nav-item {
  display: flex;
  align-items: center;
  padding: 12px 20px;
  cursor: pointer;
  transition: all 0.2s ease;
  color: #a0a0a0;
  position: relative;
}

.nav-item:hover {
  background: rgba(255, 255, 255, 0.05);
  color: #e0e0e0;
}

.nav-item.active {
  background: rgba(240, 165, 0, 0.1);
  color: #f0a500;
  border-right: 3px solid #f0a500;
}

.nav-item svg {
  margin-right: 12px;
  flex-shrink: 0;
}

.nav-item span {
  flex: 1;
  font-size: 0.9rem;
  font-weight: 500;
}

.count {
  background: rgba(255, 255, 255, 0.1);
  color: #a0a0a0;
  font-size: 0.75rem;
  padding: 2px 8px;
  border-radius: 12px;
  margin-left: 8px;
}

.journal-list {
  list-style: none;
  padding: 0;
  margin: 0;
}

.journal-item {
  display: flex;
  align-items: center;
  padding: 12px 20px;
  cursor: pointer;
  transition: all 0.2s ease;
  border-left: 3px solid transparent;
}

.journal-item:hover {
  background: rgba(255, 255, 255, 0.05);
  border-left-color: #f0a500;
}

.journal-info {
  flex: 1;
  min-width: 0;
}

.journal-title {
  font-size: 0.85rem;
  font-weight: 500;
  margin: 0 0 4px 0;
  color: #e0e0e0;
  white-space: nowrap;
  overflow: hidden;
  text-overflow: ellipsis;
}

.journal-meta {
  font-size: 0.75rem;
  color: #666;
  margin: 0;
}

.journal-actions {
  opacity: 0;
  transition: opacity 0.2s ease;
}

.journal-item:hover .journal-actions {
  opacity: 1;
}

.action-btn {
  width: 24px;
  height: 24px;
  border-radius: 4px;
  background: transparent;
  border: none;
  color: #a0a0a0;
  cursor: pointer;
  display: flex;
  align-items: center;
  justify-content: center;
  transition: all 0.2s ease;
}

.action-btn:hover {
  background: rgba(255, 255, 255, 0.1);
  color: #e0e0e0;
}

.sidebar-footer {
  padding: 16px 20px;
  border-top: 1px solid #333;
}

.settings-btn {
  display: flex;
  align-items: center;
  width: 100%;
  padding: 12px;
  background: transparent;
  border: none;
  color: #a0a0a0;
  cursor: pointer;
  border-radius: 8px;
  transition: all 0.2s ease;
  font-size: 0.9rem;
}

.settings-btn:hover {
  background: rgba(255, 255, 255, 0.05);
  color: #e0e0e0;
}

.settings-btn svg {
  margin-right: 8px;
}

/* 自定义滚动条 */
.sidebar::-webkit-scrollbar,
.sidebar-nav::-webkit-scrollbar {
  width: 4px;
}

.sidebar::-webkit-scrollbar-track,
.sidebar-nav::-webkit-scrollbar-track {
  background: transparent;
}

.sidebar::-webkit-scrollbar-thumb,
.sidebar-nav::-webkit-scrollbar-thumb {
  background-color: #444;
  border-radius: 2px;
}

.sidebar::-webkit-scrollbar-thumb:hover,
.sidebar-nav::-webkit-scrollbar-thumb:hover {
  background-color: #555;
}
</style>