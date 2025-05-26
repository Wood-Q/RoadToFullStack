from selenium import webdriver
from selenium.webdriver.common.by import By
from selenium.webdriver.support.ui import WebDriverWait
from selenium.webdriver.support import expected_conditions as EC
import time
driver = webdriver.Chrome()

driver.get("http://sahitest.com/demo/")

driver.find_element(By.LINK_TEXT,"Form Test").click()

time.sleep(10)



