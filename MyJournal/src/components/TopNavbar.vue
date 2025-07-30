<template>
  <nav class="top-navbar">
    <div class="navbar-left">
      <div class="logo">
        <svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
          <path d="M20.24 12.24a6 6 0 0 0-8.49-8.49L5 10.5V19h8.5z"></path>
          <line x1="16" y1="8" x2="22" y2="2"></line>
          <line x1="12" y1="12" x2="18" y2="6"></line>
        </svg>
        <span>一叶日记</span>
      </div>
      <div class="breadcrumb">
        <span class="current-file">新建日记</span>
      </div>
    </div>
    
    <div class="navbar-center">
      <div class="search-box">
        <svg xmlns="http://www.w3.org/2000/svg" width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
          <circle cx="11" cy="11" r="8"></circle>
          <path d="m21 21-4.35-4.35"></path>
        </svg>
        <input type="text" placeholder="搜索日记..." />
      </div>
    </div>
    
    <div class="navbar-right">
      <div class="toolbar-actions">
        <button class="icon-btn" title="撤销">
          <svg xmlns="http://www.w3.org/2000/svg" width="18" height="18" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
            <path d="M3 7v6h6"></path>
            <path d="M21 17a9 9 0 0 0-9-9 9 9 0 0 0-6 2.3L3 13"></path>
          </svg>
        </button>
        <button class="icon-btn" title="重做">
          <svg xmlns="http://www.w3.org/2000/svg" width="18" height="18" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
            <path d="M21 7v6h-6"></path>
            <path d="M3 17a9 9 0 0 1 9-9 9 9 0 0 1 6 2.3L21 13"></path>
          </svg>
        </button>
        <div class="divider"></div>
        <button class="primary-btn" @click="handleCreateJournal">
          <svg xmlns="http://www.w3.org/2000/svg" width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
            <path d="M19 21H5a2 2 0 0 1-2-2V5a2 2 0 0 1 2-2h11l5 5v11a2 2 0 0 1-2 2z"></path>
            <polyline points="17,21 17,13 7,13 7,21"></polyline>
            <polyline points="7,3 7,8 15,8"></polyline>
          </svg>
          保存日记
        </button>
      </div>
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
  display: flex;
  align-items: center;
  justify-content: space-between;
  height: 64px;
  padding: 0 24px;
  background: rgba(42, 42, 42, 0.95);
  backdrop-filter: blur(10px);
  border-bottom: 1px solid #333;
  color: #e0e0e0;
  position: relative;
  z-index: 100;
}

.navbar-left {
  display: flex;
  align-items: center;
  gap: 24px;
  flex: 1;
}

.logo {
  display: flex;
  align-items: center;
  gap: 8px;
  font-weight: 600;
  font-size: 1.1rem;
  color: #f0a500;
}

.logo svg {
  color: #f0a500;
}

.breadcrumb {
  display: flex;
  align-items: center;
  font-size: 0.9rem;
  color: #a0a0a0;
}

.current-file {
  padding: 4px 12px;
  background: rgba(240, 165, 0, 0.1);
  border-radius: 16px;
  color: #f0a500;
  font-weight: 500;
}

.navbar-center {
  flex: 1;
  display: flex;
  justify-content: center;
  max-width: 400px;
}

.search-box {
  display: flex;
  align-items: center;
  gap: 8px;
  background: rgba(255, 255, 255, 0.05);
  border: 1px solid #444;
  border-radius: 20px;
  padding: 8px 16px;
  width: 100%;
  transition: all 0.2s ease;
}

.search-box:focus-within {
  border-color: #f0a500;
  background: rgba(240, 165, 0, 0.05);
}

.search-box svg {
  color: #a0a0a0;
  flex-shrink: 0;
}

.search-box input {
  background: transparent;
  border: none;
  color: #e0e0e0;
  font-size: 0.9rem;
  width: 100%;
  outline: none;
}

.search-box input::placeholder {
  color: #666;
}

.navbar-right {
  flex: 1;
  display: flex;
  justify-content: flex-end;
}

.toolbar-actions {
  display: flex;
  align-items: center;
  gap: 8px;
}

.icon-btn {
  display: flex;
  align-items: center;
  justify-content: center;
  width: 36px;
  height: 36px;
  background: transparent;
  border: none;
  border-radius: 8px;
  color: #a0a0a0;
  cursor: pointer;
  transition: all 0.2s ease;
}

.icon-btn:hover {
  background: rgba(255, 255, 255, 0.1);
  color: #e0e0e0;
}

.divider {
  width: 1px;
  height: 24px;
  background: #444;
  margin: 0 8px;
}

.primary-btn {
  display: flex;
  align-items: center;
  gap: 8px;
  background: linear-gradient(135deg, #f0a500 0%, #e09400 100%);
  color: #fff;
  border: none;
  padding: 10px 20px;
  border-radius: 8px;
  cursor: pointer;
  font-weight: 600;
  font-size: 0.9rem;
  transition: all 0.2s ease;
  box-shadow: 0 2px 8px rgba(240, 165, 0, 0.3);
}

.primary-btn:hover {
  background: linear-gradient(135deg, #e09400 0%, #d08800 100%);
  box-shadow: 0 4px 12px rgba(240, 165, 0, 0.4);
  transform: translateY(-1px);
}

.primary-btn:active {
  transform: translateY(0);
  box-shadow: 0 2px 6px rgba(240, 165, 0, 0.3);
}

/* 响应式设计 */
@media (max-width: 992px) {
  .navbar-center {
    display: none;
  }
  
  .breadcrumb {
    display: none;
  }
}

@media (max-width: 768px) {
  .top-navbar {
    padding: 0 16px;
  }
  
  .toolbar-actions .icon-btn {
    display: none;
  }
  
  .divider {
    display: none;
  }
  
  .primary-btn {
    padding: 8px 16px;
    font-size: 0.8rem;
  }
}
</style>
