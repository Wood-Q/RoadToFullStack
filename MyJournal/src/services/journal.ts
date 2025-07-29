import { invoke } from "@tauri-apps/api/core";

export async function createJournal(newTitle: string, newBody: string) {
  try {
    const successMessage: string = await invoke("create_journal", {
      newTitle,
      newBody,
    });
    return `成功: ${successMessage}`;
  } catch (error) {
    return `失败: ${error}`;
  }
}
