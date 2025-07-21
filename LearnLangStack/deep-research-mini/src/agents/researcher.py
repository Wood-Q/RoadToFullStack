from langgraph.prebuilt import create_react_agent
from src.models.chat_model import chat_model
from src.tools.web_crawl import web_crawl
from src.tools.web_search import web_search

research=create_react_agent(
    chat_model,
    tools=[web_search, web_crawl],
    name="researcher",
)