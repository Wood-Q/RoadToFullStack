from selenium import webdriver
from selenium.webdriver.common.by import By

class TestCase:
    def __init__(self):
        self.driver = webdriver.Chrome()
        self.driver.get("http://sahitest.com/demo/")
        self.driver.maximize_window()

    def test_get_title(self):
        self.driver.find_element(By.LINK_TEXT,'Link Test').click()
        e=self.driver.find_element(By.ID,'linkById')
        # 获取元素文本
        print(e.text)
        # 获取元素标签名
        print(e.tag_name)
        # 获取元素大小
        print(e.size)
        # 获取元素位置
        print(e.location)
        # 获取元素href属性
        print(e.get_attribute("href"))
        # 获取元素id属性
        print(e.get_attribute("id"))
        # 获取元素class属性
        print(e.get_attribute("class"))
        # 获取元素style属性
        print(e.get_attribute("style"))
        # 获取元素title属性
        print(e.get_attribute("title"))

if __name__ == "__main__":
    case = TestCase()
    case.test_get_title()
        