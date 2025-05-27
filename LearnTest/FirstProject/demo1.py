from selenium import webdriver
from selenium.webdriver.chrome.options import Options
from time import sleep
from selenium.webdriver.common.by import By

class TestCase:
    def __init__(self):
        self.driver = webdriver.Chrome()

    def test_baidu(self):
        self.driver.maximize_window()
        sleep(1)
        self.driver.get("https://www.baidu.com")
        sleep(1)
        self.driver.find_element(By.NAME,'wd').send_keys("selenium")
        sleep(1)
        self.driver.find_element(By.ID,'su').click()
        sleep(1)
        self.driver.find_element(By.LINK_TEXT,"百度首页").click()
        self.driver.quit()

if __name__ == "__main__":
    case = TestCase()
    case.test_baidu()