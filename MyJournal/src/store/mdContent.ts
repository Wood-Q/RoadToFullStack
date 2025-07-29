import { defineStore } from "pinia";

export interface MdContent {
  mdText: string;
}

export const useMdContentStore = defineStore("mdContent", {
  state: () => {
    return {
      mdText: "# 123",
    };
  },
  actions: {
    setContent(newContent: string) {
      this.mdText = newContent;
    },
  },
});
