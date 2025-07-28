<template>
  <div class="markdown-editor">
    <textarea class="editor" v-model="markdownText"></textarea>
    <div class="preview" v-html="renderedHtml"></div>
  </div>
  <button @click="createJournal">Create Journal</button>
  <div class="create-result">处理结果为：{{ createResult }}</div>
</template>

<script setup lang="ts">
import { ref, computed } from "vue";
import { marked } from "marked";
import { invoke } from "@tauri-apps/api/core";

const createResult = ref("");
// 1. 使用 ref 创建一个响应式变量来存储用户输入的 Markdown
const markdownText = ref(
  "# Hello, Vue!\n\n- 这是一个列表\n- 使用 `marked` 渲染"
);

// 2. 使用 computed 创建一个计算属性
//    它会依赖 markdownText，当 markdownText 变化时，它会自动重新计算
const renderedHtml = computed(() => {
  // 3. 调用 marked 函数将 Markdown 文本转换为 HTML 字符串
  return marked(markdownText.value);
});


async function createJournal() {
  try {
    const successMessage: string = await invoke("create_journal", {
      // ✅ 将 new_title 改为 newTitle
      newTitle: "hello",
      // ✅ 将 new_body 改为 newBody
      newBody: markdownText.value,
    });
    createResult.value = `成功: ${successMessage}`;
  } catch (error) {
    createResult.value = `失败: ${error}`;
  }
}

// const greetMsg = ref("");
// const name = ref("");

// async function greet() {
//   // Learn more about Tauri commands at https://tauri.app/develop/calling-rust/
//   greetMsg.value = await invoke("greet", { name: name.value });
// }


</script>

<style scoped>
.markdown-editor {
  display: flex;
  height: 90vh;
  gap: 1rem; /* 增加一点间距 */
}

.editor,
.preview {
  flex: 1;
  padding: 1em;
  border: 1px solid #ccc;
  border-radius: 4px;
  overflow-y: auto; /* 内容超出时滚动 */
  font-family: monospace;
}

.preview {
  font-family: sans-serif;
  /* 为预览区域添加一些基本的样式 */
  background-color: #f9f9f9;
}

.create-result {
  color: green;
}
</style>
