from colorama import Fore, init

init(autoreset=True)

class Logger:
  
  def log(message: str):
    print(f"[LOG] {message}")
    
  def error(message: str):
    print(f"{Fore.RED}[ERROR] {message}")
  
  def database(message: str, error = False):
    COLOR = Fore.RED if error else Fore.BLUE
    print(f"{COLOR}[DATABASE] {message}")
    