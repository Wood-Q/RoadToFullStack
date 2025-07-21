from rich.traceback import install
from src.agents.agent_runner import run_agent
from src.agents.researcher import research
from dotenv import load_dotenv

# 运行这一行，rich 就会自动接管后续所有的报错信息
install()

load_dotenv()


def main():
    run_agent(agent=research, message="明日方舟缪尔赛思的种族是什么")


if __name__ == "__main__":
    main()
