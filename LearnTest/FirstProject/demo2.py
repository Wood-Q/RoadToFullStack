from selenium import webdriver


class TestCase:
    def __init__(self):
        self.driver = webdriver.Chrome()
        self.driver.get("https://www.baidu.com")
        self.driver.maximize_window()

    def test_prop(self):
        # 获取浏览器名称
        print(self.driver.name)
        # 获取当前url
        print(self.driver.current_url)
        # 获取页面源码
        print(self.driver.page_source)
        # 获取页面加载策略
        print(self.driver.page_load_strategy)
        # 获取当前窗口句柄
        print(self.driver.current_window_handle)
        # 刷新页面
        self.driver.refresh()
        # 后退
        self.driver.back()
        # 前进
        self.driver.forward()
        # 关闭当前窗口
        self.driver.close()
        # 关闭浏览器
        self.driver.quit()


if __name__ == "__main__":
    case = TestCase()
    case.test_prop()