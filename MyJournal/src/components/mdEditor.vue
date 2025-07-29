<template>
  <div class="editor-layout">
    <div class="markdown-editor">
      <textarea
        class="editor"
        :value="mdContentStore.mdText"
        @input="updateContent"
      ></textarea>
      <div class="preview" v-html="renderedHtml"></div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { computed } from "vue";
import { marked } from "marked";
import { useMdContentStore } from "../store/mdContent";

const mdContentStore = useMdContentStore();

const renderedHtml = computed(() => {
  return marked(mdContentStore.mdText);
});

const updateContent = (event: Event) => {
  const target = event.target as HTMLTextAreaElement;
  mdContentStore.setContent(target.value);
};
</script>

<style scoped>
.editor-layout {
  grid-area: editor;
  display: flex;
  flex-direction: column;
  height: 100%;
  background-color: #1e1e1e;
  color: #e0e0e0;
  padding: 1.2em;
  box-sizing: border-box;
}

.editor-controls {
  display: flex;
  align-items: center;
  margin-bottom: 1.2em;
  gap: 1.2em;
}

.editor-controls button {
  background-color: #c46a21;
  color: #fff;
  border: none;
  padding: 0.6em 1.2em;
  border-radius: 0.5em;
  cursor: pointer;
  transition: background-color 0.2s ease;
  font-size: 0.9rem;
  font-weight: 600;
}

.editor-controls button:hover {
  background-color: #a45a11;
}

.create-result {
  color: #a0a0a0;
  font-size: 0.9rem;
}

.markdown-editor {
  flex-grow: 1;
  display: flex;
  gap: 1.2em;
  height: calc(100% - 3.5em); /* Adjust based on controls height */
}

.editor,
.preview {
  flex: 1;
  padding: 1.5em;
  border: 1px solid #333;
  border-radius: 0.5em;
  overflow-y: auto;
  background-color: #2a2a2a;
  color: #e0e0e0;
  font-family: -apple-system, BlinkMacSystemFont, "Segoe UI", Roboto, Helvetica,
    Arial, sans-serif;
  font-size: 1rem;
  line-height: 1.6;
}

.editor {
  font-family: monospace;
  resize: none;
}

.preview :deep(h1),
.preview :deep(h2),
.preview :deep(h3) {
  color: #c46a21;
  border-bottom: 1px solid #444;
  padding-bottom: 0.3em;
}

.preview :deep(p) {
  margin: 1em 0;
}

.preview :deep(ul),
.preview :deep(ol) {
  padding-left: 2em;
}

.preview :deep(li) {
  margin-bottom: 0.5em;
}

.preview :deep(blockquote) {
  border-left: 0.25em solid #c46a21;
  padding-left: 1em;
  margin-left: 0;
  color: #a0a0a0;
}

.preview :deep(code) {
  background-color: #3c3c3c;
  padding: 0.2em 0.4em;
  border-radius: 0.25em;
  font-family: monospace;
}

.preview :deep(pre) {
  background-color: #3c3c3c;
  padding: 1em;
  border-radius: 0.5em;
  overflow-x: auto;
}

.preview :deep(pre code) {
  padding: 0;
  background-color: transparent;
}

/* Custom Scrollbar */
.editor::-webkit-scrollbar,
.preview::-webkit-scrollbar {
  width: 0.5em;
}

.editor::-webkit-scrollbar-track,
.preview::-webkit-scrollbar-track {
  background: #2a2a2a;
}

.editor::-webkit-scrollbar-thumb,
.preview::-webkit-scrollbar-thumb {
  background-color: #555;
  border-radius: 0.25em;
  border: 0.125em solid #2a2a2a;
}

.editor::-webkit-scrollbar-thumb:hover,
.preview::-webkit-scrollbar-thumb:hover {
  background-color: #777;
}
</style>
