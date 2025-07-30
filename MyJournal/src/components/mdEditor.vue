<template>
  <div class="editor-layout">
    <div class="editor-header">
      <div class="editor-tabs">
        <button 
          class="tab-btn" 
          :class="{ active: activeTab === 'write' }" 
          @click="activeTab = 'write'"
        >
          <svg xmlns="http://www.w3.org/2000/svg" width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
            <path d="M11 4H4a2 2 0 0 0-2 2v14a2 2 0 0 0 2 2h14a2 2 0 0 0 2-2v-7"></path>
            <path d="M18.5 2.5a2.121 2.121 0 0 1 3 3L12 15l-4 1 1-4 9.5-9.5z"></path>
          </svg>
          编辑
        </button>
        <button 
          class="tab-btn" 
          :class="{ active: activeTab === 'preview' }" 
          @click="activeTab = 'preview'"
        >
          <svg xmlns="http://www.w3.org/2000/svg" width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
            <path d="M1 12s4-8 11-8 11 8 11 8-4 8-11 8-11-8-11-8z"></path>
            <circle cx="12" cy="12" r="3"></circle>
          </svg>
          预览
        </button>
        <button 
          class="tab-btn" 
          :class="{ active: activeTab === 'split' }" 
          @click="activeTab = 'split'"
        >
          <svg xmlns="http://www.w3.org/2000/svg" width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
            <rect x="3" y="3" width="18" height="18" rx="2" ry="2"></rect>
            <line x1="12" y1="3" x2="12" y2="21"></line>
          </svg>
          分屏
        </button>
      </div>
      
      <div class="editor-status">
        <span class="word-count">{{ stats.words }} 字</span>
        <div class="sync-status">
          <svg xmlns="http://www.w3.org/2000/svg" width="14" height="14" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
            <polyline points="23,4 23,10 17,10"></polyline>
            <polyline points="1,20 1,14 7,14"></polyline>
            <path d="M20.49 9A9 9 0 0 0 5.64 5.64L1 10m22 4l-4.64 4.36A9 9 0 0 1 3.51 15"></path>
          </svg>
          <span>已同步</span>
        </div>
      </div>
    </div>

    <div class="markdown-editor" :class="`mode-${activeTab}`">
      <div class="editor-pane" v-show="activeTab === 'write' || activeTab === 'split'">
        <div class="editor-toolbar">
          <div class="toolbar-group">
            <button class="tool-btn" title="粗体 (Ctrl+B)" @click="insertFormat('**', '**')">
              <svg xmlns="http://www.w3.org/2000/svg" width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
                <path d="M6 4h8a4 4 0 0 1 4 4 4 4 0 0 1-4 4H6z"></path>
                <path d="M6 12h9a4 4 0 0 1 4 4 4 4 0 0 1-4 4H6z"></path>
              </svg>
            </button>
            <button class="tool-btn" title="斜体 (Ctrl+I)" @click="insertFormat('*', '*')">
              <svg xmlns="http://www.w3.org/2000/svg" width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
                <line x1="19" y1="4" x2="10" y2="4"></line>
                <line x1="14" y1="20" x2="5" y2="20"></line>
                <line x1="15" y1="4" x2="9" y2="20"></line>
              </svg>
            </button>
            <button class="tool-btn" title="链接 (Ctrl+K)" @click="insertFormat('[', '](url)')">
              <svg xmlns="http://www.w3.org/2000/svg" width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
                <path d="M10 13a5 5 0 0 0 7.54.54l3-3a5 5 0 0 0-7.07-7.07l-1.72 1.71"></path>
                <path d="M14 11a5 5 0 0 0-7.54-.54l-3 3a5 5 0 0 0 7.07 7.07l1.71-1.71"></path>
              </svg>
            </button>
          </div>
          
          <div class="toolbar-group">
            <button class="tool-btn" title="标题" @click="insertFormat('## ', '')">
              <svg xmlns="http://www.w3.org/2000/svg" width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
                <path d="M4 12h8m-8-6v12m8-12v12m4-7h6m-6-2v4m6-4v4"></path>
              </svg>
            </button>
            <button class="tool-btn" title="列表" @click="insertFormat('- ', '')">
              <svg xmlns="http://www.w3.org/2000/svg" width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
                <line x1="8" y1="6" x2="21" y2="6"></line>
                <line x1="8" y1="12" x2="21" y2="12"></line>
                <line x1="8" y1="18" x2="21" y2="18"></line>
                <line x1="3" y1="6" x2="3.01" y2="6"></line>
                <line x1="3" y1="12" x2="3.01" y2="12"></line>
                <line x1="3" y1="18" x2="3.01" y2="18"></line>
              </svg>
            </button>
            <button class="tool-btn" title="引用" @click="insertFormat('> ', '')">
              <svg xmlns="http://www.w3.org/2000/svg" width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
                <path d="M3 21c3 0 7-1 7-8V5c0-1.25-.756-2.017-2-2H4c-1.25 0-2 .75-2 1.972V11c0 1.25.75 2 2 2 1 0 1 0 1 1v1c0 1-1 2-2 2s-1 .008-1 1.031V20c0 1 0 1 1 1z"></path>
              </svg>
            </button>
          </div>
        </div>
        
        <textarea
          ref="editorTextarea"
          class="editor"
          :value="mdContentStore.mdText"
          @input="updateContent"
          @keydown="handleKeydown"
          placeholder="开始写你的日记..."
        ></textarea>
      </div>
      
      <div class="preview-pane" v-show="activeTab === 'preview' || activeTab === 'split'">
        <div class="preview-header">
          <h3 class="preview-title">预览</h3>
          <button class="export-btn" title="导出">
            <svg xmlns="http://www.w3.org/2000/svg" width="14" height="14" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
              <path d="M21 15v4a2 2 0 0 1-2 2H5a2 2 0 0 1-2-2v-4"></path>
              <polyline points="7,10 12,15 17,10"></polyline>
              <line x1="12" y1="15" x2="12" y2="3"></line>
            </svg>
          </button>
        </div>
        <div class="preview" v-html="renderedHtml"></div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { computed, ref } from "vue";
import { marked } from "marked";
import { useMdContentStore } from "../store/mdContent";

const mdContentStore = useMdContentStore();
const activeTab = ref<'write' | 'preview' | 'split'>('split');
const editorTextarea = ref<HTMLTextAreaElement>();

const renderedHtml = computed(() => {
  return marked(mdContentStore.mdText);
});

const stats = computed(() => {
  const text = mdContentStore.mdText;
  const words = text.trim() ? text.trim().split(/\s+/).length : 0;
  return { words };
});

const updateContent = (event: Event) => {
  const target = event.target as HTMLTextAreaElement;
  mdContentStore.setContent(target.value);
};

const insertFormat = (before: string, after: string) => {
  if (!editorTextarea.value) return;
  
  const textarea = editorTextarea.value;
  const start = textarea.selectionStart;
  const end = textarea.selectionEnd;
  const selectedText = textarea.value.substring(start, end) || '文本';
  
  const newText = 
    textarea.value.substring(0, start) + 
    before + selectedText + after + 
    textarea.value.substring(end);
  
  mdContentStore.setContent(newText);
  
  // 重新聚焦并设置光标位置
  setTimeout(() => {
    textarea.focus();
    const newCursorPos = start + before.length + selectedText.length + after.length;
    textarea.setSelectionRange(newCursorPos, newCursorPos);
  }, 0);
};

const handleKeydown = (event: KeyboardEvent) => {
  // 快捷键支持
  if (event.ctrlKey || event.metaKey) {
    switch (event.key) {
      case 'b':
        event.preventDefault();
        insertFormat('**', '**');
        break;
      case 'i':
        event.preventDefault();
        insertFormat('*', '*');
        break;
      case 'k':
        event.preventDefault();
        insertFormat('[', '](url)');
        break;
    }
  }
  
  // Tab 键支持
  if (event.key === 'Tab') {
    event.preventDefault();
    const textarea = event.target as HTMLTextAreaElement;
    const start = textarea.selectionStart;
    const end = textarea.selectionEnd;
    
    const newText = 
      textarea.value.substring(0, start) + 
      '  ' + 
      textarea.value.substring(end);
    
    mdContentStore.setContent(newText);
    
    setTimeout(() => {
      textarea.setSelectionRange(start + 2, start + 2);
    }, 0);
  }
};
</script>

<style scoped>
.editor-layout {
  grid-area: editor;
  display: flex;
  flex-direction: column;
  height: 100vh;
  background: linear-gradient(180deg, #2a2a2a 0%, #242424 100%);
  color: #e0e0e0;
  font-family: -apple-system, BlinkMacSystemFont, "Segoe UI", Roboto, Helvetica, Arial, sans-serif;
}

.editor-header {
  display: flex;
  align-items: center;
  justify-content: space-between;
  padding: 16px 20px;
  border-bottom: 1px solid #333;
  background: rgba(42, 42, 42, 0.95);
  backdrop-filter: blur(10px);
}

.editor-tabs {
  display: flex;
  gap: 4px;
}

.tab-btn {
  display: flex;
  align-items: center;
  gap: 6px;
  padding: 8px 16px;
  background: transparent;
  border: 1px solid #444;
  border-radius: 6px;
  color: #a0a0a0;
  cursor: pointer;
  transition: all 0.2s ease;
  font-size: 0.85rem;
  font-weight: 500;
}

.tab-btn.active {
  background: rgba(240, 165, 0, 0.1);
  border-color: #f0a500;
  color: #f0a500;
}

.tab-btn:hover:not(.active) {
  background: rgba(255, 255, 255, 0.05);
  border-color: #555;
  color: #e0e0e0;
}

.editor-status {
  display: flex;
  align-items: center;
  gap: 16px;
  font-size: 0.8rem;
  color: #a0a0a0;
}

.word-count {
  padding: 4px 8px;
  background: rgba(255, 255, 255, 0.05);
  border-radius: 12px;
  font-weight: 500;
}

.sync-status {
  display: flex;
  align-items: center;
  gap: 4px;
  color: #4ade80;
}

.markdown-editor {
  flex: 1;
  display: flex;
  overflow: hidden;
}

.mode-write .preview-pane {
  display: none !important;
}

.mode-preview .editor-pane {
  display: none !important;
}

.mode-split .editor-pane,
.mode-split .preview-pane {
  flex: 1;
}

.mode-write .editor-pane,
.mode-preview .preview-pane {
  flex: 1;
}

.editor-pane {
  display: flex;
  flex-direction: column;
  border-right: 1px solid #333;
}

.editor-toolbar {
  display: flex;
  align-items: center;
  gap: 16px;
  padding: 12px 16px;
  border-bottom: 1px solid #333;
  background: rgba(255, 255, 255, 0.02);
}

.toolbar-group {
  display: flex;
  gap: 4px;
}

.tool-btn {
  display: flex;
  align-items: center;
  justify-content: center;
  width: 32px;
  height: 32px;
  background: transparent;
  border: 1px solid #444;
  border-radius: 4px;
  color: #a0a0a0;
  cursor: pointer;
  transition: all 0.2s ease;
}

.tool-btn:hover {
  background: rgba(240, 165, 0, 0.1);
  border-color: #f0a500;
  color: #f0a500;
}

.editor {
  flex: 1;
  padding: 20px;
  border: none;
  background: transparent;
  color: #e0e0e0;
  font-family: 'JetBrains Mono', 'Fira Code', Consolas, monospace;
  font-size: 14px;
  line-height: 1.6;
  resize: none;
  outline: none;
}

.editor::placeholder {
  color: #666;
  font-style: italic;
}

.preview-pane {
  display: flex;
  flex-direction: column;
  background: rgba(255, 255, 255, 0.02);
}

.preview-header {
  display: flex;
  align-items: center;
  justify-content: space-between;
  padding: 12px 16px;
  border-bottom: 1px solid #333;
}

.preview-title {
  font-size: 0.9rem;
  font-weight: 600;
  margin: 0;
  color: #a0a0a0;
}

.export-btn {
  display: flex;
  align-items: center;
  justify-content: center;
  width: 28px;
  height: 28px;
  background: transparent;
  border: 1px solid #444;
  border-radius: 4px;
  color: #a0a0a0;
  cursor: pointer;
  transition: all 0.2s ease;
}

.export-btn:hover {
  background: rgba(240, 165, 0, 0.1);
  border-color: #f0a500;
  color: #f0a500;
}

.preview {
  flex: 1;
  padding: 20px;
  overflow-y: auto;
  font-family: -apple-system, BlinkMacSystemFont, "Segoe UI", Roboto, Helvetica, Arial, sans-serif;
  font-size: 14px;
  line-height: 1.7;
}

.preview :deep(h1),
.preview :deep(h2),
.preview :deep(h3),
.preview :deep(h4),
.preview :deep(h5),
.preview :deep(h6) {
  color: #f0a500;
  margin: 24px 0 16px 0;
  font-weight: 600;
}

.preview :deep(h1) {
  font-size: 2rem;
  border-bottom: 2px solid #333;
  padding-bottom: 8px;
}

.preview :deep(h2) {
  font-size: 1.5rem;
  border-bottom: 1px solid #333;
  padding-bottom: 4px;
}

.preview :deep(h3) {
  font-size: 1.25rem;
}

.preview :deep(p) {
  margin: 16px 0;
  color: #e0e0e0;
}

.preview :deep(ul),
.preview :deep(ol) {
  margin: 16px 0;
  padding-left: 24px;
}

.preview :deep(li) {
  margin-bottom: 8px;
  color: #e0e0e0;
}

.preview :deep(blockquote) {
  border-left: 4px solid #f0a500;
  padding: 16px 20px;
  margin: 24px 0;
  background: rgba(240, 165, 0, 0.05);
  color: #d0d0d0;
  font-style: italic;
}

.preview :deep(code) {
  background: rgba(255, 255, 255, 0.1);
  padding: 2px 6px;
  border-radius: 4px;
  font-family: 'JetBrains Mono', 'Fira Code', Consolas, monospace;
  font-size: 0.9em;
  color: #fbbf24;
}

.preview :deep(pre) {
  background: rgba(0, 0, 0, 0.3);
  padding: 16px;
  border-radius: 8px;
  overflow-x: auto;
  margin: 16px 0;
  border: 1px solid #333;
}

.preview :deep(pre code) {
  background: transparent;
  padding: 0;
  color: #e0e0e0;
}

.preview :deep(table) {
  width: 100%;
  border-collapse: collapse;
  margin: 16px 0;
}

.preview :deep(th),
.preview :deep(td) {
  border: 1px solid #444;
  padding: 8px 12px;
  text-align: left;
}

.preview :deep(th) {
  background: rgba(240, 165, 0, 0.1);
  color: #f0a500;
  font-weight: 600;
}

.preview :deep(hr) {
  border: none;
  height: 2px;
  background: linear-gradient(90deg, transparent, #333, transparent);
  margin: 32px 0;
}

/* 自定义滚动条 */
.editor::-webkit-scrollbar,
.preview::-webkit-scrollbar {
  width: 8px;
}

.editor::-webkit-scrollbar-track,
.preview::-webkit-scrollbar-track {
  background: rgba(255, 255, 255, 0.05);
  border-radius: 4px;
}

.editor::-webkit-scrollbar-thumb,
.preview::-webkit-scrollbar-thumb {
  background: #444;
  border-radius: 4px;
}

.editor::-webkit-scrollbar-thumb:hover,
.preview::-webkit-scrollbar-thumb:hover {
  background: #555;
}

/* 响应式设计 */
@media (max-width: 768px) {
  .editor-header {
    padding: 12px 16px;
  }
  
  .editor-tabs {
    gap: 2px;
  }
  
  .tab-btn {
    padding: 6px 12px;
    font-size: 0.8rem;
  }
  
  .editor-status {
    gap: 12px;
    font-size: 0.75rem;
  }
  
  .markdown-editor.mode-split {
    flex-direction: column;
  }
  
  .mode-split .editor-pane {
    border-right: none;
    border-bottom: 1px solid #333;
  }
}
</style>
