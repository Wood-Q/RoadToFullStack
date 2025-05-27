from selenium import webdriver
from selenium.webdriver.chrome.options import Options
from time import sleep
from selenium.webdriver.common.by import By

# 先到系统里找到chromedriver
# 然后示例化  webdriver.Chrome()
options = Options()
driver = webdriver.Chrome(options=options)

# 调用get方法，访问百度
driver.get("https://www.baidu.com")

# 通过id找到输入框，输入selenium
driver.find_element(By.ID,'kw').send_keys("selenium")

# 通过id找到搜索按钮，点击
driver.find_element(By.ID,'su').click()

sleep(10)

driver.quit()