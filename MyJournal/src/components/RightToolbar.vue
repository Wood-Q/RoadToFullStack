<template>
  <aside class="right-toolbar">
    <div class="toolbar-header">
      <h3 class="toolbar-title">工具栏</h3>
    </div>

    <div class="toolbar-section">
      <h4 class="section-title">文本格式</h4>
      <div class="format-buttons">
        <button class="format-btn" title="粗体" @click="insertFormat('**', '**')">
          <svg xmlns="http://www.w3.org/2000/svg" width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
            <path d="M6 4h8a4 4 0 0 1 4 4 4 4 0 0 1-4 4H6z"></path>
            <path d="M6 12h9a4 4 0 0 1 4 4 4 4 0 0 1-4 4H6z"></path>
          </svg>
        </button>
        <button class="format-btn" title="斜体" @click="insertFormat('*', '*')">
          <svg xmlns="http://www.w3.org/2000/svg" width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
            <line x1="19" y1="4" x2="10" y2="4"></line>
            <line x1="14" y1="20" x2="5" y2="20"></line>
            <line x1="15" y1="4" x2="9" y2="20"></line>
          </svg>
        </button>
        <button class="format-btn" title="删除线" @click="insertFormat('~~', '~~')">
          <svg xmlns="http://www.w3.org/2000/svg" width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
            <path d="M16 4H9a3 3 0 0 0-2.83 4"></path>
            <path d="M14 12a4 4 0 0 1 0 8H6"></path>
            <line x1="4" y1="12" x2="20" y2="12"></line>
          </svg>
        </button>
        <button class="format-btn" title="代码" @click="insertFormat('`', '`')">
          <svg xmlns="http://www.w3.org/2000/svg" width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
            <polyline points="16,18 22,12 16,6"></polyline>
            <polyline points="8,6 2,12 8,18"></polyline>
          </svg>
        </button>
      </div>
      
      <div class="format-buttons">
        <button class="format-btn wide" title="标题 1" @click="insertFormat('# ', '')">
          H1
        </button>
        <button class="format-btn wide" title="标题 2" @click="insertFormat('## ', '')">
          H2
        </button>
        <button class="format-btn wide" title="标题 3" @click="insertFormat('### ', '')">
          H3
        </button>
      </div>

      <div class="format-buttons">
        <button class="format-btn" title="无序列表" @click="insertFormat('- ', '')">
          <svg xmlns="http://www.w3.org/2000/svg" width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
            <line x1="8" y1="6" x2="21" y2="6"></line>
            <line x1="8" y1="12" x2="21" y2="12"></line>
            <line x1="8" y1="18" x2="21" y2="18"></line>
            <line x1="3" y1="6" x2="3.01" y2="6"></line>
            <line x1="3" y1="12" x2="3.01" y2="12"></line>
            <line x1="3" y1="18" x2="3.01" y2="18"></line>
          </svg>
        </button>
        <button class="format-btn" title="有序列表" @click="insertFormat('1. ', '')">
          <svg xmlns="http://www.w3.org/2000/svg" width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
            <line x1="10" y1="6" x2="21" y2="6"></line>
            <line x1="10" y1="12" x2="21" y2="12"></line>
            <line x1="10" y1="18" x2="21" y2="18"></line>
            <path d="M4 6h1v4"></path>
            <path d="M4 10h2"></path>
            <path d="M6 18H4c0-1 2-2 2-3s-1-1.5-2-1"></path>
          </svg>
        </button>
        <button class="format-btn" title="引用" @click="insertFormat('> ', '')">
          <svg xmlns="http://www.w3.org/2000/svg" width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
            <path d="M3 21c3 0 7-1 7-8V5c0-1.25-.756-2.017-2-2H4c-1.25 0-2 .75-2 1.972V11c0 1.25.75 2 2 2 1 0 1 0 1 1v1c0 1-1 2-2 2s-1 .008-1 1.031V20c0 1 0 1 1 1z"></path>
            <path d="M15 21c3 0 7-1 7-8V5c0-1.25-.757-2.017-2-2h-4c-1.25 0-2 .75-2 1.972V11c0 1.25.75 2 2 2h.75c0 2.25.25 4-2.75 4v3c0 1 0 1 1 1z"></path>
          </svg>
        </button>
        <button class="format-btn" title="链接" @click="insertFormat('[', '](url)')">
          <svg xmlns="http://www.w3.org/2000/svg" width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
            <path d="M10 13a5 5 0 0 0 7.54.54l3-3a5 5 0 0 0-7.07-7.07l-1.72 1.71"></path>
            <path d="M14 11a5 5 0 0 0-7.54-.54l-3 3a5 5 0 0 0 7.07 7.07l1.71-1.71"></path>
          </svg>
        </button>
      </div>
    </div>

    <div class="toolbar-section">
      <h4 class="section-title">快速插入</h4>
      <div class="quick-insert">
        <button class="insert-btn" @click="insertText('---\n')" title="分割线">
          <svg xmlns="http://www.w3.org/2000/svg" width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
            <line x1="3" y1="12" x2="21" y2="12"></line>
          </svg>
          <span>分割线</span>
        </button>
        <button class="insert-btn" @click="insertText('```\n\n```\n')" title="代码块">
          <svg xmlns="http://www.w3.org/2000/svg" width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
            <polyline points="16,18 22,12 16,6"></polyline>
            <polyline points="8,6 2,12 8,18"></polyline>
          </svg>
          <span>代码块</span>
        </button>
        <button class="insert-btn" @click="insertText('| 列1 | 列2 | 列3 |\n|-----|-----|-----|\n| 内容 | 内容 | 内容 |\n')" title="表格">
          <svg xmlns="http://www.w3.org/2000/svg" width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
            <path d="M9 3H5a2 2 0 0 0-2 2v4m6-6h10a2 2 0 0 1 2 2v4M9 3v18m0 0h10a2 2 0 0 0 2-2V9M9 21H5a2 2 0 0 1-2-2V9m0 0h18"></path>
          </svg>
          <span>表格</span>
        </button>
      </div>
    </div>

    <div class="toolbar-section">
      <h4 class="section-title">文档统计</h4>
      <div class="stats">
        <div class="stat-item">
          <span class="stat-label">字符数</span>
          <span class="stat-value">{{ stats.characters }}</span>
        </div>
        <div class="stat-item">
          <span class="stat-label">单词数</span>
          <span class="stat-value">{{ stats.words }}</span>
        </div>
        <div class="stat-item">
          <span class="stat-label">行数</span>
          <span class="stat-value">{{ stats.lines }}</span>
        </div>
        <div class="stat-item">
          <span class="stat-label">预计阅读</span>
          <span class="stat-value">{{ stats.readTime }}分钟</span>
        </div>
      </div>
    </div>

    <div class="toolbar-section">
      <h4 class="section-title">主题设置</h4>
      <div class="theme-options">
        <button class="theme-btn active" title="深色主题">
          <div class="theme-preview dark"></div>
          <span>深色</span>
        </button>
        <button class="theme-btn" title="浅色主题">
          <div class="theme-preview light"></div>
          <span>浅色</span>
        </button>
      </div>
    </div>
  </aside>
</template>

<script setup lang="ts">
import { computed } from 'vue';
import { useMdContentStore } from '../store/mdContent';

const mdContentStore = useMdContentStore();

// 文档统计
const stats = computed(() => {
  const text = mdContentStore.mdText;
  const characters = text.length;
  const words = text.trim() ? text.trim().split(/\s+/).length : 0;
  const lines = text.split('\n').length;
  const readTime = Math.ceil(words / 200); // 假设每分钟阅读200字

  return {
    characters,
    words,
    lines,
    readTime: readTime || 1
  };
});

// 插入格式化文本
const insertFormat = (before: string, after: string) => {
  const currentText = mdContentStore.mdText;
  const newText = currentText + before + '文本' + after;
  mdContentStore.setContent(newText);
};

// 插入文本
const insertText = (text: string) => {
  const currentText = mdContentStore.mdText;
  const newText = currentText + text;
  mdContentStore.setContent(newText);
};
</script>

<style scoped>
.right-toolbar {
  grid-area: right-toolbar;
  background: linear-gradient(180deg, #2a2a2a 0%, #242424 100%);
  border-left: 1px solid #333;
  color: #e0e0e0;
  height: 100vh;
  overflow-y: auto;
  font-family: -apple-system, BlinkMacSystemFont, "Segoe UI", Roboto, Helvetica, Arial, sans-serif;
}

.toolbar-header {
  padding: 20px;
  border-bottom: 1px solid #333;
}

.toolbar-title {
  font-size: 1.1rem;
  font-weight: 600;
  margin: 0;
  color: #f0a500;
}

.toolbar-section {
  padding: 20px;
  border-bottom: 1px solid #333;
}

.toolbar-section:last-child {
  border-bottom: none;
}

.section-title {
  font-size: 0.85rem;
  font-weight: 600;
  color: #a0a0a0;
  margin: 0 0 16px 0;
  text-transform: uppercase;
  letter-spacing: 0.5px;
}

.format-buttons {
  display: grid;
  grid-template-columns: repeat(4, 1fr);
  gap: 8px;
  margin-bottom: 12px;
}

.format-buttons:last-child {
  margin-bottom: 0;
}

.format-btn {
  display: flex;
  align-items: center;
  justify-content: center;
  width: 36px;
  height: 36px;
  background: rgba(255, 255, 255, 0.05);
  border: 1px solid #444;
  border-radius: 6px;
  color: #a0a0a0;
  cursor: pointer;
  transition: all 0.2s ease;
  font-size: 0.75rem;
  font-weight: 600;
}

.format-btn.wide {
  grid-column: span 1;
  font-size: 0.7rem;
}

.format-btn:hover {
  background: rgba(240, 165, 0, 0.1);
  border-color: #f0a500;
  color: #f0a500;
}

.quick-insert {
  display: flex;
  flex-direction: column;
  gap: 8px;
}

.insert-btn {
  display: flex;
  align-items: center;
  gap: 8px;
  padding: 10px 12px;
  background: rgba(255, 255, 255, 0.05);
  border: 1px solid #444;
  border-radius: 6px;
  color: #a0a0a0;
  cursor: pointer;
  transition: all 0.2s ease;
  font-size: 0.8rem;
  text-align: left;
}

.insert-btn:hover {
  background: rgba(240, 165, 0, 0.1);
  border-color: #f0a500;
  color: #f0a500;
}

.stats {
  display: flex;
  flex-direction: column;
  gap: 12px;
}

.stat-item {
  display: flex;
  justify-content: space-between;
  align-items: center;
  padding: 8px 12px;
  background: rgba(255, 255, 255, 0.03);
  border-radius: 6px;
}

.stat-label {
  font-size: 0.8rem;
  color: #a0a0a0;
}

.stat-value {
  font-size: 0.85rem;
  font-weight: 600;
  color: #f0a500;
}

.theme-options {
  display: flex;
  gap: 8px;
}

.theme-btn {
  display: flex;
  flex-direction: column;
  align-items: center;
  gap: 6px;
  padding: 12px;
  background: rgba(255, 255, 255, 0.05);
  border: 1px solid #444;
  border-radius: 8px;
  color: #a0a0a0;
  cursor: pointer;
  transition: all 0.2s ease;
  font-size: 0.75rem;
  flex: 1;
}

.theme-btn.active {
  background: rgba(240, 165, 0, 0.1);
  border-color: #f0a500;
  color: #f0a500;
}

.theme-btn:hover:not(.active) {
  background: rgba(255, 255, 255, 0.08);
  border-color: #555;
}

.theme-preview {
  width: 24px;
  height: 16px;
  border-radius: 4px;
  border: 1px solid #444;
}

.theme-preview.dark {
  background: linear-gradient(135deg, #1a1a1a 0%, #2d2d2d 100%);
}

.theme-preview.light {
  background: linear-gradient(135deg, #f5f5f5 0%, #ffffff 100%);
}

/* 自定义滚动条 */
.right-toolbar::-webkit-scrollbar {
  width: 4px;
}

.right-toolbar::-webkit-scrollbar-track {
  background: transparent;
}

.right-toolbar::-webkit-scrollbar-thumb {
  background-color: #444;
  border-radius: 2px;
}

.right-toolbar::-webkit-scrollbar-thumb:hover {
  background-color: #555;
}
</style>